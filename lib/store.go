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
	data := make([]ColorLink, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = ColorLink{
			Link:  "#",
			Color: "white",
		}
	}

	return &LinkStore{
		pollDuration: poll,
		client: &http.Client{
			CheckRedirect: func(*http.Request, []*http.Request) error {
				return http.ErrUseLastResponse
			},
		},
		data:     data,
		writeIdx: 0,
	}
}

type ColorLink struct {
	Link  string
	Color string
}

type LinkStore struct {
	mu       sync.Mutex
	data     []ColorLink
	writeIdx int

	pollDuration time.Duration
	client       *http.Client
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
		l.data[l.writeIdx] = cl
		l.writeIdx++
		if l.writeIdx > 999 {
			l.writeIdx = 0
		}
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
