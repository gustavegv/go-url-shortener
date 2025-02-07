// routes_test.go
package routes_test

import (
	"bytes"
	"go-url-shortener/internal/routes"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
)

func TestConcurrentShortenRequests(t *testing.T) {
	// Initialize the Server with a non-concurrent safe map
	s := &routes.Server{
		DataStore: make(map[string]*routes.SavedLinks),
	}

	// Create a test server using your router
	ts := httptest.NewServer(s.NewRouter())
	defer ts.Close()

	// Number of concurrent requests
	const numRequests = 400
	var wg sync.WaitGroup

	for i := 0; i < numRequests; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Simulate a POST request to shorten a URL.
			// All requests use the same long URL for simplicity.
			longURL := "http://example.com"
			resp, err := http.Post(ts.URL+"/api/shorten", "text/plain", bytes.NewBufferString(longURL))
			if err != nil {
				t.Errorf("Goroutine %d: POST request failed: %v", i, err)
				return
			}
			defer resp.Body.Close()
			// You could also check the response body here if needed.
		}(i)
	}

	// Wait for all goroutines to finish
	wg.Wait()

	// Optionally, check the final state of s.DataStore.
	// For example, you might verify that the shortened link exists or that no duplicate entries were made.
	// (This depends on your application's expected behavior.)
}
