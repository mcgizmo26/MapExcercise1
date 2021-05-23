package main

import "fmt"

// 1st method of creating a Go Map
// func main() {
// 	// Must have a comma after all key value pairs regardless it it is the last key/value pair.
// 	// Hash map must have the same type for all keys and same type for all values.
// 	colors := map[string]string{
// 		"red":   "#ff0000",
// 		"green": "#00ff00",
// 		"blue":  "#0000ff",
// 	}

// 	fmt.Println(colors)
// }

// 2nd method of creating a Go Map
// func main() {
// 	var colors map[string]string
// 	fmt.Println(colors)
// }

// 3rd method of creating a map
func main() {
	colors := make(map[string]string)

	// Cannot use dot notation for maps
	colors["white"] = "#ffffff"
	colors["red"] = "#ff0000"
	colors["green"] = "#00ff00"
	colors["blue"] = "#0000ff"

	delete(colors, "blue")

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex: ", color, "is", hex)
	}
}
