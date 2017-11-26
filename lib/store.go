package random

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

const (
	randomArticleURL   = "https://en.wikipedia.org/wiki/Special:Random"
	validHexColorRange = 16777216
)

// NewLinkStore is the constructor for LinkStore. It initializes a data set with
// empty links.
func NewLinkStore(poll time.Duration) *LinkStore {
	data := make([]ColorLink, 1000)
	for i := 0; i < 1000; i++ {
		data[i] = ColorLink{
			Link:  "http://meyerweb.com/eric/thoughts/2014/06/19/rebeccapurple/",
			Color: "rebeccapurple",
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

// LinkStore fetches random links from Wikipedia and stores those links along
// with a random color.
type LinkStore struct {
	mu       sync.Mutex
	data     []ColorLink
	writeIdx int

	pollDuration time.Duration
	client       *http.Client
}

// All returns the entire contents of the store.
func (l *LinkStore) All() []ColorLink {
	l.mu.Lock()
	defer l.mu.Unlock()
	return l.data
}

// Populate will fetch a random article from Wikipedia on the configured
// interval.
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

// NewColorLink creates a color link with the specified URL and a random color.
func NewColorLink(link string) ColorLink {
	randomColor := fmt.Sprintf("#%.6X", rand.Intn(validHexColorRange))
	return ColorLink{
		Link:  link,
		Color: randomColor,
	}
}

// ColorLink holds a URL and a Hex color.
type ColorLink struct {
	Link  string
	Color string
}
