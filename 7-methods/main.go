package main

import "fmt"

// Functions start with func and surround the code with { }
func main() {
	//
	fmt.Println("------Basic Method Sections------")
	Basic()

	fmt.Println("------Pointer As Receiver Sections------")
	PointerReceiver()
}

/*
-	Methods can be defined on any file in the package.
- 	You cannot define a method on a type from another package, or on a basic type.

CODE ORGANIZATION TEMPLATE SUGGESTION.

	package models

	// list of packages to import
	import (
		"fmt"
	)

	// list of constants
	const (
		ConstExample = "const before vars"
	)

	// list of variables
	var (
		ExportedVar    = 42
		nonExportedVar = "so say we all"
	)

	// Main type(s) for the file,
	// try to keep the lowest amount of structs per file when possible.
	type User struct {
		FirstName, LastName string
		Location            *UserLocation
	}

	type UserLocation struct {
		City    string
		Country string
	}

	// List of functions
	func NewUser(firstName, lastName string) *User {
		return &User{FirstName: firstName,
			LastName: lastName,
			Location: &UserLocation{
				City:    "Santa Monica",
				Country: "USA",
			},
		}
	}

	// List of methods
	func (u *User) Greeting() string {
		return fmt.Sprintf("Dear %s %s", u.FirstName, u.LastName)
	}

*/
