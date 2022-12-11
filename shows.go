package main

import (
	"fmt"
)

type show struct {
	name    string
	season  int
	episode int
}

func (s show) print() {
	fmt.Println(s.name, s.season, s.episode)
}

func (s show) Add() {
	fmt.Println("adding show: ", s)
}

func (s show) Remove() {
	fmt.Println("remove show: ", s)
}

func (s show) Get() {
	fmt.Println("getting show: ", s)
}
