package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	random "github.com/enocom/random/lib"
)

func main() {
	rand.Seed(time.Now().Unix())
	log.Println("Starting server on port 8080")

	env, err := readEnvironment()
	if err != nil {
		log.Fatalf("Failed to read environment variable: %v", err)
	}

	store := random.NewStore(env.pollDuration)
	go store.Populate()

	http.Handle("/", random.NewRootHandler(store))
	log.Fatal(http.ListenAndServe(env.port, nil))
}

type env struct {
	pollDuration time.Duration
	port         string
}

func readEnvironment() (env, error) {
	poll := os.Getenv("POLL_DURATION")
	if poll == "" {
		return env{}, errors.New("missing POLL_DURATION (e.g., 1m)")
	}
	pollDuration, err := time.ParseDuration(poll)
	if err != nil {
		return env{}, errors.New("incorrect format of POLL_DURATION (e.g., 1m)")
	}
	port := os.Getenv("PORT")
	if port == "" {
		return env{}, errors.New("missing PORT (e.g., 8080)")
	}

	e := env{
		pollDuration: pollDuration,
		port:         fmt.Sprintf("localhost:%s", port),
	}

	return e, nil
}
