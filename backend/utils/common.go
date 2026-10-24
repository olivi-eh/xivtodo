package utils

import (
	"crypto/subtle"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"server/models"
	"server/store"
	"strconv"
	"strings"
)

// MaxBodyBytes matches Firestore's maximum document size
const MaxBodyBytes = 1 << 20

func ValidateMethodOrFail(w http.ResponseWriter, r *http.Request, methods ...string) error {
	for _, method := range methods {
		if r.Method == method {
			return nil
		}
	}

	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	log.Printf("MethodNotAllowed: %s", r.Method)
	return errors.New("invalid method: " + r.Method)
}

func GetUserWithAuthnOrFail(w http.ResponseWriter, r *http.Request, userData *models.User) error {
	authorization := strings.Split(r.Header.Get("Authorization"), ":")
	if len(authorization) != 2 {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Println("invalid Authorization header")
		return errors.New("invalid Authorization header")
	}

	discordID := authorization[0]
	authorizationCode := authorization[1]

	user, err := store.Client.Collection("users").Doc(discordID).Get(store.Ctx)
	if err != nil || !user.Exists() {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		log.Println(err)
		return errors.New("user not found")
	}

	m := user.Data()
	jsonBody, err := json.Marshal(m)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return err
	}
	if err := json.Unmarshal(jsonBody, userData); err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Println(err)
		return err
	}

	if userData.Settings == nil ||
		subtle.ConstantTimeCompare([]byte(userData.Settings.AuthorizationCode), []byte(authorizationCode)) != 1 {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		log.Println("invalid Authorization code")
		return errors.New("invalid authorization code")
	}

	return nil
}

func GetCharacterIDOrFail(w http.ResponseWriter, r *http.Request) (string, error) {
	characterID := r.URL.Query().Get("id")
	if _, err := strconv.ParseUint(characterID, 10, 32); err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("invalid id query: %q", characterID)
		return "", errors.New("invalid id query")
	}

	return characterID, nil
}
