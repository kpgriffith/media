package main

import (
	"fmt"
)

type movie struct {
	name   string
	genre  string
	rating string
}

func (m movie) print() {
	fmt.Println(m.name, m.genre, m.rating)
}

func (m movie) Add() {
	fmt.Println("adding movie: ", m)
}

func (m movie) Remove() {
	fmt.Println("removing movie: ", m)
}

func (m movie) Get() {
	fmt.Println("getting movie: ", m)
}
