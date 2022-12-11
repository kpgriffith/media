package main

import (
	"testing"
)

func TestMoviePrint(t *testing.T) {
	m := movie{
		name:   "Godfather",
		genre:  "Drama",
		rating: "R",
	}

	Info(m)
}

func TestShowPrint(t *testing.T) {
	s := show{
		name:    "Friends",
		season:  1,
		episode: 1,
	}

	Info(s)
}
