package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"server/handlers"
	"time"
)

// allowedOrigin derives the frontend origin from its OAuth redirect URI, which handlers require to be set
func allowedOrigin() string {
	redirectURI, err := url.Parse(os.Getenv("DISCORD_REDIRECT_URI"))
	if err != nil || redirectURI.Scheme == "" || redirectURI.Host == "" {
		log.Fatalln("fatal: The DISCORD_REDIRECT_URI env is not an absolute URL")
	}
	return redirectURI.Scheme + "://" + redirectURI.Host
}

func withHeaders(handler http.Handler, origin string) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")

		if r.Method == http.MethodOptions {
			w.Header().Set("Access-Control-Allow-Methods", "GET, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, Cache-Control")
			w.Header().Set("Access-Control-Max-Age", "3600")
			w.WriteHeader(http.StatusNoContent)
			return
		}

		remoteHost := r.Header.Get("X-REAL-IP")
		if remoteHost == "" {
			remoteHost = r.Header.Get("X-FORWARDED-FOR")
		}
		if remoteHost == "" {
			remoteHost = r.RemoteAddr
		}
		log.Println(remoteHost, "-->", r.Method, r.RequestURI)

		handler.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}

func main() {
	origin := allowedOrigin()
	log.Println("allowing CORS requests from", origin)

	mux := http.NewServeMux()
	mux.Handle("/version", withHeaders(handlers.VersionHandler(), origin))
	mux.Handle("/auth", withHeaders(handlers.AuthHandler(), origin))
	mux.Handle("/users", withHeaders(handlers.UsersHandler(), origin))
	mux.Handle("/characters", withHeaders(handlers.CharactersHandler(), origin))
	mux.Handle("/settings", withHeaders(handlers.SettingsHandler(), origin))
	mux.Handle("/encounters", withHeaders(handlers.EncountersHandler(), origin))
	mux.Handle("/checklist", withHeaders(handlers.ChecklistHandler(), origin))

	host := "localhost:8181"
	port := os.Getenv("PORT")
	if port != "" {
		host = ":" + port
	}

	log.Println("listening to", host)
	server := &http.Server{
		Addr:              host,
		Handler:           mux,
		ReadHeaderTimeout: 10 * time.Second,
		ReadTimeout:       30 * time.Second,
		WriteTimeout:      60 * time.Second,
		IdleTimeout:       120 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
