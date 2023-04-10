package main

import (
	"github.com/simonwardle/mydemoprogram/helpers"
	"fmt"
)

func main() {
	var myVar helpers.SomeType
	
	fmt.Println("Examples code for Golang")
	fmt.Println("Enter (A)ll or select and individual option")
	fmt.Println("1. Show Functions")
	fmt.Println("2. Pointers")
	fmt.Println("3. TypeS & Structs")
	fmt.Println("4. Receivers")
	fmt.Println("5. Maps & Slices")
	fmt.Println("6. If & Case")
	fmt.Println("7. Loops")
	fmt.Println("8. Interfaces")
	fmt.Println("9. Packages")
	fmt.Println("10. Channels")
	fmt.Println("11. Using Json")

	var option string
	fmt.Scanln(&option)

	switch option {
	case "1":
		fmt.Println("Running show functions!")
		showFunctions()
	case "2":
		fmt.Println("Running pointers.")
		showPointers()
	case "3":
		fmt.Println("Running types & structs.")
		showTypes()
	case "4":
		fmt.Println("Running receivers.")
		showReceivers()
	case "5":
		fmt.Println("Running Maps & Slices")
		mapsslices()
	case "6":
		fmt.Println("Running If & Case")
		showIf()
	case "7":
		fmt.Println("Running Loops")
		showLoops()
	case "8":
		fmt.Println("Running Interfaces")
		showInterfaces()
	case "9":
		myVar.TypeName = "Some Type from helpers"
		fmt.Println("Running Packages")
		fmt.Println(myVar.TypeName)
	case "10":
		fmt.Println("Running Channels")
		showChannels()
	case "11":
		fmt.Println("11. Using Json")
		showJson()
	default:
		fmt.Println("Running all.")
		fmt.Println("Running Functions")
		showFunctions()
		fmt.Println("Running Pointers")
		showPointers()
		fmt.Println("Running Types")
		showTypes()
		fmt.Println("Running Receivers")
		showReceivers()
		fmt.Println("Running Maps & Slices")
		mapsslices()
		fmt.Println("Running If & Case")
		showIf()
		fmt.Println("Running Loops")
		showLoops()
		fmt.Println("Running Interfaces")
		showInterfaces()
		fmt.Println("Running Packages")
		myVar.TypeName = "Some Type from helpers"
		fmt.Println(myVar.TypeName)
		fmt.Println("Running Channels")
		showChannels()
		fmt.Println("11. Using Json")
		showJson()
	}
}
