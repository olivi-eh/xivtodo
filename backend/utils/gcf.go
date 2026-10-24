package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"server/models"
	"strconv"
	"time"
)

var gcfClient = &http.Client{Timeout: 30 * time.Second}

func GetLodestoneProfile(w http.ResponseWriter, characterID string, basic bool) (*models.LodestoneProfile, error) {
	query := url.Values{
		"id":    {characterID},
		"basic": {strconv.FormatBool(basic)},
	}
	resp, err := gcfClient.Get("https://asia-northeast1-xiv-todo.cloudfunctions.net/lodestone-character?" + query.Encode())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode == http.StatusServiceUnavailable {
		http.Error(w, http.StatusText(http.StatusServiceUnavailable), http.StatusServiceUnavailable)
		return nil, errors.New(http.StatusText(http.StatusServiceUnavailable))
	} else if resp.StatusCode == http.StatusNotFound {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return nil, errors.New(http.StatusText(http.StatusNotFound))
	} else if resp.StatusCode != http.StatusOK {
		http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
		return nil, errors.New("unexpected status from lodestone-character: " + resp.Status)
	}

	lodestoneProfile := &models.LodestoneProfile{}
	if err := json.NewDecoder(resp.Body).Decode(lodestoneProfile); err != nil {
		http.Error(w, http.StatusText(http.StatusBadGateway), http.StatusBadGateway)
		return nil, err
	}
	return lodestoneProfile, nil
}
