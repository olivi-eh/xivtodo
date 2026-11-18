package handlers

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"log"
	"net/http"
	"server/models"
	"server/store"
	"server/utils"
	"time"
)

func CharactersHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if utils.ValidateMethodOrFail(w, r, http.MethodGet, http.MethodDelete) != nil {
			return
		}

		userData := models.User{}
		if utils.GetUserWithAuthnOrFail(w, r, &userData) != nil {
			return
		}

		characterID, err := utils.GetCharacterIDOrFail(w, r)
		if err != nil {
			return
		}

		switch r.Method {
		case http.MethodGet:
			addCharacter(w, &userData, characterID)
		case http.MethodDelete:
			removeCharacter(w, &userData, characterID)
		}
	}
	return http.HandlerFunc(fn)
}

func addCharacter(w http.ResponseWriter, userData *models.User, characterID string) {
	if userHasMaxCharacters(userData) && !CharacterInUser(userData, characterID) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println("user already has 8 characters")
		return
	}

	lodestoneProfile, err := utils.GetLodestoneProfile(w, characterID, false)
	if err != nil {
		log.Printf("failed to get Lodestone data: %v", err)
		return
	}

	character := models.Character{
		Updated:       time.Now().Unix(),
		LodestoneData: lodestoneProfile,
		ChecklistData: &models.ChecklistData{
			WeeklyChecklist: []*models.ChecklistItem{},
			DailyChecklist:  []*models.ChecklistItem{},
			AdhocChecklist:  []*models.ChecklistItem{},
		},
	}
	if CharacterProfilePrivate(character.LodestoneData) {
		character.LodestoneData.Character.GuardianDeity = nil
		character.LodestoneData.Character.Town = nil
	}
	if CharacterInUser(userData, characterID) {
		character = *userData.Characters[characterID]
		character.Updated = time.Now().Unix()
		keepCollectedEntities(lodestoneProfile, character.LodestoneData)
		character.LodestoneData = lodestoneProfile
	}

	_, err = store.Client.Collection("users").Doc(userData.DiscordUser.ID).Set(store.Ctx, map[string]interface{}{
		"Characters": map[string]interface{}{
			characterID: character,
		},
	}, firestore.MergeAll)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("failed to store User data: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(character)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("failed to send Character data: %v", err)
		return
	}
}

func removeCharacter(w http.ResponseWriter, userData *models.User, characterID string) {
	if !CharacterInUser(userData, characterID) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println("character is not in user")
		return
	}

	_, err := store.Client.Collection("users").Doc(userData.DiscordUser.ID).Update(store.Ctx, []firestore.Update{
		{
			Path:  "Characters." + characterID,
			Value: firestore.Delete,
		},
	})

	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Printf("Failed to store User data: %v", err)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func CharacterInUser(userData *models.User, id string) bool {
	_, exists := userData.Characters[id]
	return exists
}

func userHasMaxCharacters(userData *models.User) bool {
	return len(userData.Characters) >= 8
}

// keepCollectedEntities carries over stored achievements, minions and mounts when a refresh returns fewer of them.
// None of these can be lost in game, so a shorter list means Lodestone served an incomplete scrape, and storing it
// would wipe the character's progress. Achievements are the exception when the profile has since been made private.
func keepCollectedEntities(fresh, stored *models.LodestoneProfile) {
	if stored == nil {
		return
	}
	achievementsPrivate := fresh.AllAchievementInfo != nil && fresh.AllAchievementInfo.Private
	if !achievementsPrivate && len(fresh.Achievements) < len(stored.Achievements) {
		log.Printf("refresh returned %d achievements, keeping %d stored", len(fresh.Achievements), len(stored.Achievements))
		fresh.Achievements = stored.Achievements
		fresh.AllAchievementInfo = stored.AllAchievementInfo
	}
	if len(fresh.Minions) < len(stored.Minions) {
		log.Printf("refresh returned %d minions, keeping %d stored", len(fresh.Minions), len(stored.Minions))
		fresh.Minions = stored.Minions
	}
	if len(fresh.Mounts) < len(stored.Mounts) {
		log.Printf("refresh returned %d mounts, keeping %d stored", len(fresh.Mounts), len(stored.Mounts))
		fresh.Mounts = stored.Mounts
	}
}

func CharacterProfilePrivate(lodestoneData *models.LodestoneProfile) bool {
	return lodestoneData.Character.ActiveClassJob == nil
}
