package main

import (
	"errors"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	random "github.com/enocom/random/lib"
)

const version = "1.3.0"

func main() {
	rand.Seed(time.Now().Unix())
	env, err := readEnvironment()
	if err != nil {
		log.Fatalf("Failed to read environment variable: %v", err)
	}
	log.Printf("Starting server on %s", env.addr)

	store := random.NewLinkStore(env.pollDuration)
	go store.Populate()

	http.Handle("/", random.NewRootHandler(store))
	http.Handle("/healthz", random.NewHealthHandler(version))
	log.Fatal(http.ListenAndServe(env.addr, nil))
}

type env struct {
	pollDuration time.Duration
	addr         string
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
	addr := os.Getenv("ADDR")
	if addr == "" {
		return env{}, errors.New("missing ADDR (e.g., 0.0.0.0:80)")
	}

	e := env{
		pollDuration: pollDuration,
		addr:         addr,
	}

	return e, nil
}
