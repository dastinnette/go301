package greeting

import "fmt"

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string)

func Greet(salutation []Salutation, do Printer, isFormal bool, times int) {

	for _, s := range salutation {
		message, alternate := CreateMessage(s.Name, s.Greeting)
		if prefix := GetPrefix(s.Name); isFormal {
			do(prefix + message)
		} else {
			do(alternate)
		}
	}
}

func GetPrefix(name string) (prefix string) {
	switch name {
	case "Bob":
		prefix = "Mr. "
	case "Sally", "Doug":
		prefix = "Dr. "
	case "Cindy":
		prefix = "Mrs. "
	default:
		prefix = "Brah "
	}

	return
}

func CreateMessage(name, greeting string) (message string, alternate string) {
	message = greeting + " " + name
	alternate = "HEY!" + name
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func CreatePrintFunction(custom string) Printer {
	return func(s string) {
		fmt.Println(s + custom)
	}
}
