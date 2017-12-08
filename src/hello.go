package main

import "./greeting"

func main() {
	var s = greeting.Salutation{"Doug", "Hello"}
	greeting.Greet(s, greeting.CreatePrintFunction("??"), true)
}
