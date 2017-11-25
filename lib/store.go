package random

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

func NewStore(poll time.Duration) *LinkStore {
	return &LinkStore{
		pollDuration: poll,
		client: &http.Client{
			CheckRedirect: func(*http.Request, []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
	}
}

type ColorLink struct {
	Link  string
	Color string
}

type LinkStore struct {
	mu   sync.Mutex
	data []ColorLink

	client       *http.Client
	pollDuration time.Duration
}

func (l *LinkStore) All() []ColorLink {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.data
}

const randomArticleURL = "https://en.wikipedia.org/wiki/Special:Random"

func (l *LinkStore) Populate() {
	for range time.Tick(l.pollDuration) {
		resp, err := l.client.Get(randomArticleURL)
		if err != nil {
			log.Printf("Failed to get random article. Error: %v", err)
			continue
		}
		if resp.StatusCode != http.StatusFound {
			log.Printf("Expected 302 status, got %v", resp.StatusCode)
			continue
		}
		link := resp.Header["Location"]
		if link == nil {
			log.Println("Response had no location header")
			continue
		}
		cl := NewColorLink(link[0])

		l.mu.Lock()
		l.data = append(l.data, cl)
		l.mu.Unlock()
	}
}

const validHexColorRange = 16777216

func NewColorLink(link string) ColorLink {
	randomColor := fmt.Sprintf("#%.6X", rand.Intn(validHexColorRange))
	return ColorLink{
		Link:  link,
		Color: randomColor,
	}
}
