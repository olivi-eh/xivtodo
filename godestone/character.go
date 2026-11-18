package lodestone

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/xivapi/godestone/v2"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Lodestone throttles bursts of requests by answering 200 with pages that don't parse, which godestone reports as
// empty or truncated results rather than errors, so incomplete scrapes are retried after these delays
var retryDelays = []time.Duration{3 * time.Second, 10 * time.Second}

// retryBudget keeps retries well inside the function's 90s timeout and the backend's 120s client timeout
const retryBudget = 60 * time.Second

var errIncomplete = errors.New("incomplete scrape")

type CharacterResponse struct {
	Achievements       []*godestone.AchievementInfo
	AllAchievementInfo *godestone.AllAchievementInfo
	Character          *godestone.Character
	Minions            []*godestone.Minion
	Mounts             []*godestone.Mount
}

func LodestoneCharacter(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if r.Method == http.MethodOptions {
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Access-Control-Max-Age", "3600")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	id64, err := strconv.ParseUint(r.URL.Query().Get("id"), 10, 32)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println(err)
		return
	}
	id := uint32(id64)
	basic := r.URL.Query().Get("basic") == "true"
	deadline := time.Now().Add(retryBudget)

	characterResponse := CharacterResponse{}

	character, err := withRetries(deadline, id, "character", func() (*godestone.Character, error) {
		character, err := scraper.FetchCharacter(id)
		if err == nil && character.Name == "" {
			return nil, errIncomplete
		}
		return character, err
	})
	if err != nil {
		if err.Error() == http.StatusText(http.StatusNotFound) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		} else {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		}
		return
	}
	characterResponse.Character = character

	if !basic {
		type achievementResult struct {
			achievements []*godestone.AchievementInfo
			info         *godestone.AllAchievementInfo
		}
		achievements, err := withRetries(deadline, id, "achievements", func() (achievementResult, error) {
			achievements, info, err := scraper.FetchCharacterAchievements(id)
			if err != nil {
				return achievementResult{}, err
			}
			// Total is parsed from the first page; a public profile always reports it, and every page must be collected
			if !info.Private && (info.TotalAchievements == 0 || uint32(len(achievements)) < info.TotalAchievements) {
				return achievementResult{}, fmt.Errorf("%w: %d of %d achievements", errIncomplete, len(achievements), info.TotalAchievements)
			}
			return achievementResult{achievements, info}, nil
		})
		if err != nil {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		characterResponse.Achievements = achievements.achievements
		characterResponse.AllAchievementInfo = achievements.info

		minions, err := withRetries(deadline, id, "minions", func() ([]*godestone.Minion, error) {
			return scraper.FetchCharacterMinions(id)
		})
		if err != nil {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		characterResponse.Minions = minions

		mounts, err := withRetries(deadline, id, "mounts", func() ([]*godestone.Mount, error) {
			return scraper.FetchCharacterMounts(id)
		})
		if err != nil {
			http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
			return
		}
		characterResponse.Mounts = mounts
	}

	cJSON, err := json.Marshal(characterResponse)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(cJSON)
}

// withRetries calls fetch until it succeeds, retrying failures other than a missing character until the deadline
func withRetries[T any](deadline time.Time, id uint32, what string, fetch func() (T, error)) (T, error) {
	var result T
	var err error
	for attempt := 0; ; attempt++ {
		result, err = fetch()
		if err == nil {
			return result, nil
		}
		log.Printf("character %d: fetching %s failed (attempt %d): %v", id, what, attempt+1, err)
		if err.Error() == http.StatusText(http.StatusNotFound) || attempt >= len(retryDelays) || time.Now().Add(retryDelays[attempt]).After(deadline) {
			return result, err
		}
		time.Sleep(retryDelays[attempt])
	}
}
