package main

import (
	"time"

	"github.com/rl404/fairy"
)

func main() {
	// Cache type.
	t := fairy.RedisCache

	// Init client.
	client, err := fairy.NewCache(t, "localhost:6379", "", time.Minute)
	if err != nil {
		panic(err)
	}

	// Don't forget to close.
	defer client.Close()

	// Sample data. Can be any type.
	key := "key"
	data := []string{"a", "b", "c"}

	// Save to cache. Data will be encoded first.
	if err := client.Set(key, data); err != nil {
		panic(err)
	}

	// Create a new or use existing variable.
	// Data type should be the same as when saving to cache.
	var newData []string

	// Get data from cache. Data will be decoded to inputted
	// variable. Don't forget to use pointer.
	if err := client.Get(key, &newData); err != nil {
		panic(err)
	}

	// Delete data from cache.
	if err := client.Delete(key); err != nil {
		panic(err)
	}
}
