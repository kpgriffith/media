package main

import "testing"

func TestShowAdd(t *testing.T) {
	s := show{
		name:    "Friends",
		season:  1,
		episode: 1,
	}

	s.Add()
}
