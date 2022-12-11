package main

import "testing"

func TestMovieAdd(t *testing.T) {
	m := movie{
		name:   "Godfather",
		genre:  "Drama",
		rating: "R",
	}

	m.Add()
}
