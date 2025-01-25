package main

import (
	"belajar-library-golang/packagelib"
	"errors"
	"fmt"
)

func main() {
	firstName := "John"
	lastName := "Doe"
	fmt.Printf("Hello %s %s\n", firstName, lastName)

	// Packge Error
	err := packagelib.CekData("1")
	if err != nil {
		if errors.Is(err, packagelib.ValidationError) {
			fmt.Println("validation error")
		} else if errors.Is(err, packagelib.NotFoundError) {
			fmt.Println("not found error")
		} else {
			fmt.Println("unknown error")
		}
	}

	// Package Os
	packagelib.TestOs()

	// Package Flag
	packagelib.TestFlag()

	// Package Strings
	packagelib.TestStrings()

	// Package Strconv
	packagelib.TestStrconv()
}
