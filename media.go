// Package media defines an inteface and methods for other types
// to implement.
package main

// media interface defines methods that can be performed on types that
// implement the interface.
type Media interface {
	Add()
	Remove()
	Get()
	print()
}

// Info prints the information for any type implementing the media interface.
func Info(m Media) {
	m.print()
}
