package main

import (
	"log"
	"sort"
)

type Users struct {
	FirstName string
	LastName  string
}

func mapsslices() {
	// maps
	myMap := make(map[string]Users)

	me := Users{
		FirstName: "Simon",
		LastName:  "Wardle",
	}

	you := Users{
		FirstName: "Trevor",
		LastName:  "Sawler",
	}

	myMap["me"] = me
	myMap["you"] = you

	log.Println(myMap["me"].FirstName)
	log.Println(myMap["you"].FirstName)

	// slices
	var mySlice []int

	mySlice = append(mySlice, 2)
	mySlice = append(mySlice, 1)
	mySlice = append(mySlice, 3)

	sort.Ints(mySlice)

	log.Println(mySlice)

	// shorthand for slices

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(numbers)

	// print first two elements of slice, starting at first position
	log.Println(numbers[0:2])

	// create a slice of strings
	names := []string{"one", "seven", "fish", "cat"}
	log.Println(names)
}
