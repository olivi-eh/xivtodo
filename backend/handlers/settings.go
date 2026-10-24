package handlers

import (
	"cloud.google.com/go/firestore"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"server/models"
	"server/store"
	"server/utils"
)

func SettingsHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if utils.ValidateMethodOrFail(w, r, http.MethodPut) != nil {
			return
		}

		userData := models.User{}
		if utils.GetUserWithAuthnOrFail(w, r, &userData) != nil {
			return
		}

		settings := models.Settings{}
		err := json.NewDecoder(http.MaxBytesReader(w, r.Body, utils.MaxBodyBytes)).Decode(&settings)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("failed to decode settings: %v", err)
			return
		}

		err = validateSettings(&userData, &settings)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Printf("requested settings invalid: %v", err)
			return
		}

		_, err = store.Client.Collection("users").Doc(userData.DiscordUser.ID).Set(store.Ctx, map[string]interface{}{
			"Settings": settings,
		}, firestore.MergeAll)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("failed to store User data: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(settings)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("failed to send settings data: %v", err)
			return
		}
	}
	return http.HandlerFunc(fn)
}

func validateSettings(userData *models.User, payload *models.Settings) error {
	if payload.AuthorizationCode != userData.Settings.AuthorizationCode {
		return errors.New("AuthorizationCode is immutable")
	}

	if payload.CharacterClaimCode != userData.Settings.CharacterClaimCode {
		return errors.New("CharacterClaimCode is immutable")
	}

	if payload.SpoilersOption < 0 || payload.SpoilersOption > 2 {
		return errors.New("SpoilersOption invalid, must be between [0-2]")
	}

	if payload.PatchNumsOption < 0 || payload.PatchNumsOption > 2 {
		return errors.New("PatchNumsOption invalid, must be between [0-2]")
	}

	if payload.ThemeOption < 0 || payload.ThemeOption > 1 {
		return errors.New("ThemeOption invalid, must be between [0-1]")
	}

	return nil
}
