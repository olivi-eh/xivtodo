package handlers

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/binary"
	"encoding/json"
	"errors"
	"log"
	mathrand "math/rand/v2"
	"net/http"
	"net/url"
	"os"
	"server/models"
	"server/store"
	"server/utils"
	"time"
)

const (
	discordAPIEndpoint = "https://discord.com/api/v10"
)

var (
	discordClientID     string
	discordClientSecret string
	discordRedirectUri  string

	discordClient = &http.Client{Timeout: 10 * time.Second}
)

func init() {
	discordClientID = os.Getenv("DISCORD_CLIENT_ID")
	if discordClientID == "" {
		log.Fatalln("fatal: The DISCORD_CLIENT_ID env is not set")
	}

	discordClientSecret = os.Getenv("DISCORD_CLIENT_SECRET")
	if discordClientSecret == "" {
		log.Fatalln("fatal: The DISCORD_CLIENT_SECRET env is not set")
	}

	discordRedirectUri = os.Getenv("DISCORD_REDIRECT_URI")
	if discordRedirectUri == "" {
		log.Fatalln("fatal: The DISCORD_REDIRECT_URI env is not set")
	}
}

func AuthHandler() http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if utils.ValidateMethodOrFail(w, r, http.MethodGet) != nil {
			return
		}

		code := r.URL.Query().Get("code")

		discordAuthResponse := models.DiscordAuthResponse{}
		err := getDiscordAuthResponse(&discordAuthResponse, code)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			log.Printf("Failed to get DiscordAuthResponse: %v", err)
			return
		}

		discordUser := models.DiscordUser{}
		err = getDiscordUserResponse(&discordUser, &discordAuthResponse)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("Failed to get DiscordUser: %v", err)
			return
		}

		userData := models.User{}
		err = getUserDataOrDefault(&userData, &discordUser)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("Failed to get User data: %v", err)
			return
		}

		_, err = store.Client.Collection("users").Doc(discordUser.ID).Set(store.Ctx, userData)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("Failed to store User data: %v", err)
			return
		}

		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(userData)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			log.Printf("Failed to send User data: %v", err)
			return
		}
	}
	return http.HandlerFunc(fn)
}

func getDiscordAuthResponse(discordAuthResponse *models.DiscordAuthResponse, code string) error {
	values := url.Values{
		"client_id":     {discordClientID},
		"client_secret": {discordClientSecret},
		"grant_type":    {"authorization_code"},
		"code":          {code},
		"redirect_uri":  {discordRedirectUri},
	}

	resp, err := discordClient.Post(discordAPIEndpoint+"/oauth2/token", "application/x-www-form-urlencoded", bytes.NewBufferString(values.Encode()))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// @TODO: Print useful Discord API rate limit data
	//log.Println(resp.Header.Get("X-RateLimit-Global"))
	//log.Println(resp.Header.Get("X-RateLimit-Scope"))
	//log.Println(resp.Header.Get("X-RateLimit-Reset"))
	//log.Println(resp.Header.Get("X-RateLimit-Reset-After"))
	//log.Println(resp.Header.Get("X-RateLimit-Bucket"))
	//rateLimitRemaining := resp.Header.Get("X-RateLimit-Remaining")
	//rateLimitLimit := resp.Header.Get("X-RateLimit-Limit")
	//log.Printf("rate limit for /oauth2/token: %s/%s", rateLimitRemaining, rateLimitLimit)

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(discordAuthResponse)
	if err != nil {
		return err
	}

	return nil
}

func getDiscordUserResponse(discordUser *models.DiscordUser, discordAuthResponse *models.DiscordAuthResponse) error {
	req, err := http.NewRequest(http.MethodGet, discordAPIEndpoint+"/users/@me", nil)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", discordAuthResponse.TokenType+" "+discordAuthResponse.AccessToken)
	resp, err := discordClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	rateLimitRemaining := resp.Header.Get("X-RateLimit-Remaining")
	rateLimitLimit := resp.Header.Get("X-RateLimit-Limit")
	log.Printf("rate limit for /users/@me: %s/%s", rateLimitRemaining, rateLimitLimit)

	if resp.StatusCode != 200 {
		return errors.New(resp.Status)
	}

	err = json.NewDecoder(resp.Body).Decode(discordUser)
	if err != nil {
		return err
	}

	return nil
}

func getUserDataOrDefault(userData *models.User, discordUser *models.DiscordUser) error {
	user, err := store.Client.Collection("users").Doc(discordUser.ID).Get(store.Ctx)
	if err != nil && (user == nil || user.Exists()) {
		return err
	}

	if user.Exists() {
		m := user.Data()
		jsonbody, err := json.Marshal(m)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(jsonbody, userData); err != nil {
			return err
		}
	}

	// Generate default settings if doesn't exist
	userData.DiscordUser = discordUser
	if userData.Settings == nil {
		// The authorization code is a long-lived bearer token, so it must be unguessable
		authorizationCode := rand.Text()

		characterClaimCodeBytes := make([]byte, 8)
		binary.LittleEndian.PutUint64(characterClaimCodeBytes, mathrand.Uint64())
		characterClaimCode := base64.StdEncoding.EncodeToString(characterClaimCodeBytes)

		userData.Settings = &models.Settings{
			AuthorizationCode:  authorizationCode,
			CharacterClaimCode: characterClaimCode,
		}
	} else if isLegacyAuthorizationCode(userData.Settings.AuthorizationCode) {
		// Legacy codes were derived from the Discord ID and a timestamp, so they're guessable
		userData.Settings.AuthorizationCode = rand.Text()
	}

	// Delete characters that have empty Lodestone data
	for id, character := range userData.Characters {
		if character.LodestoneData == nil {
			delete(userData.Characters, id)
		}
	}

	return nil
}

// isLegacyAuthorizationCode reports whether code wasn't generated by rand.Text,
// which always returns 26 characters of the standard base32 alphabet
func isLegacyAuthorizationCode(code string) bool {
	if len(code) != 26 {
		return true
	}
	for _, c := range code {
		if !(c >= 'A' && c <= 'Z') && !(c >= '2' && c <= '7') {
			return true
		}
	}
	return false
}
