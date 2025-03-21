package test

import (
	"io"
	"net/http"
	"strings"
	"testing"
	"time"
)

func TestHomepageContent(t *testing.T) {
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Create a custom request with User-Agent header
	req, err := http.NewRequest("GET", "https://matthias2wym.com/", nil)
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0")

	// Use the client's Do method instead of Get

	resp, err := client.Do(req)
	if err != nil {
		t.Fatalf("Failed to make GET request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	content := string(body)

	// Check for expected content on the homepage
	expectedPhrases := []string{
		"matthias2wym",
		"<html",
		"</html>",
	}

	for _, phrase := range expectedPhrases {
		if !strings.Contains(content, phrase) {
			t.Errorf("Expected phrase '%s' not found in homepage content", phrase)
		}
	}
}
