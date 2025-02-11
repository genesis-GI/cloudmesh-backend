package main

import(
	"math/rand"
	"time"
)

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func getRandomToken() string {
	rand.Seed(time.Now().UnixNano())
	b := make([]rune, 16)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}