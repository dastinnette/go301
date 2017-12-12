package main

import "./greeting"

func main() {
	// var s = greeting.Salutation{"Bob", "Hello"}

	slice := []greeting.Salutation{
		{"Bob", "Hello"},
		{"Sally", "Hi"},
		{"Cindy", "What up?"},
	}

	greeting.Greet(slice, greeting.CreatePrintFunction("??"), true, 5)
}
