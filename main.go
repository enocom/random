package main

import (
	"log"
	"math/rand"
	"net/http"
	"time"

	random "github.com/enocom/random/lib"
)

func main() {
	log.Println("Starting server on port 8080")

	rand.Seed(time.Now().Unix())

	store := random.NewStore(10 * time.Second)
	go store.Populate()

	http.Handle("/", random.NewRootHandler(store))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
