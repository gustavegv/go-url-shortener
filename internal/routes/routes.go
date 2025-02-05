package routes

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
)

type Server struct {
	DataStore map[string]string
}

func (s *Server) NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("/api/data", apiDataHandler)
	mux.HandleFunc("/api/shorten", s.apiShortenLink)

	return mux
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func apiDataHandler(w http.ResponseWriter, r *http.Request) {
	data := "Data..."
	fmt.Fprintln(w, data)
}

func (s *Server) apiShortenLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	shortLink, exists := s.DataStore[string(body)]

	if !exists {
		shortLink = shortenURL()
		s.DataStore[string(body)] = shortLink
	}

	response := fmt.Sprintf("Shortened link: %s", shortLink)

	fmt.Fprintln(w, response)
}

func shortenURL() string {
	return RandString(5)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyz1234567890"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}
