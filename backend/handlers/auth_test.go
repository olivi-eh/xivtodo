package handlers

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"testing"
)

func TestIsLegacyAuthorizationCode(t *testing.T) {
	for range 1000 {
		if c := rand.Text(); isLegacyAuthorizationCode(c) {
			t.Fatalf("new code flagged as legacy: %q", c)
		}
	}
	h := sha1.Sum([]byte("123456789012345678" + "1609459200"))
	if c := base64.URLEncoding.EncodeToString(h[:]); !isLegacyAuthorizationCode(c) {
		t.Fatalf("legacy code not flagged: %q", c)
	}
	if !isLegacyAuthorizationCode("") {
		t.Fatal("empty code not flagged")
	}
}
