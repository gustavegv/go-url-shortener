package routes

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strings"
)

type Server struct {
	DataStore map[string]*SavedLinks
}

type SavedLinks struct {
	LongLink        string
	clickStatistics int
}

func (s *Server) NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/shorten", s.apiShortenLink)
	mux.HandleFunc("/stats/", s.getLinkStatistics)
	mux.HandleFunc("/", s.redirectPage)

	return mux
}

func (s *Server) apiShortenLink(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		return
	}

	var shortURL string

	for shortenedURL, savedLink := range s.DataStore {
		if savedLink.LongLink == string(body) {
			shortURL = shortenedURL
			break
		}
	}

	if shortURL == "" {
		shortURL = shortenURL()
		s.saveShortLink(shortURL, string(body))
	}

	response := map[string]string{"shortened_link": shortURL}
	json.NewEncoder(w).Encode(response)

}

func (s *Server) saveShortLink(short string, long string) {
	s.DataStore[short] = &SavedLinks{
		LongLink:        long,
		clickStatistics: 0,
	}
}

func shortenURL() string {
	return RandString(5)
}

const letterBytes = "bcdfghijklmnpqrstvwxyz1234567890"

func RandString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func (s *Server) redirectPage(w http.ResponseWriter, r *http.Request) {
	found, linkStruct := s.lookupShortLink(r.URL.Path[1:])

	if !found {
		return
	}

	go func() {
		linkStruct.clickStatistics++
	}()

	link := linkStruct.LongLink
	if !strings.HasPrefix(link, "http://") && !strings.HasPrefix(link, "https://") {
		link = "https://" + link // Default to HTTPS
	}

	http.Redirect(w, r, link, http.StatusFound)
}

func (s *Server) getLinkStatistics(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	shortCode := strings.TrimPrefix(path, "/stats/") // Isolate shortend link

	found, linkStruct := s.lookupShortLink(shortCode)

	if !found {
		return
	}

	response := fmt.Sprintf("%s statistics: %d",
		linkStruct.LongLink,
		linkStruct.clickStatistics)

	fmt.Fprintln(w, response)
}

func (s *Server) lookupShortLink(link string) (bool, *SavedLinks) {

	savedLinkStruct, ok := s.DataStore[link]

	if !ok {
		fmt.Println("Invalid short-link")
		return false, nil
	}
	return true, savedLinkStruct
}
