package main

import (
	"net/http"
	"os"
	"strings"
	"testing"

	"mvdan.cc/xurls/v2"
)

func TestURLs(t *testing.T) {
	data, err := os.ReadFile("../README.md")
	if err != nil {
		t.Fatalf("Failed to load markdown file: %v", err)
	}

	rxStrict := xurls.Strict()
	urls := rxStrict.FindAllString(string(data), -1)

	for _, url := range urls {
		// Skip Terraform Registry URLs because of status 200 on non-existing URLs.
		if strings.Contains(url, "registry.terraform.io/providers/") {
			continue
		}

		resp, err := http.Get(url)
		if err != nil {
			t.Errorf("Failed: URL: %s, Error: %v", url, err)
			continue
		}
		resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Failed: URL: %s, Status code: %d", url, resp.StatusCode)
		} else {
			t.Logf("Success: URL: %s, Status code: %d", url, resp.StatusCode)
		}
	}
}
