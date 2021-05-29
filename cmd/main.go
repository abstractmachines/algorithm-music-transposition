package main

import (
	"fmt"
	// "strconv"
)

/**
- Provide for enharmonic nomenclature and sheet music/visual mappings
- Returns mappings of enharmonic equivalencies (e.g. C is B# is Dbb is ...)
*/
func enharmonic() map[string]string {
	var sharps string = "#"
	var flats string = "U+266"
	fmt.Println(sharps, flats)

	// aMap := {
	// 	"C#": "Db", // ...enumerate
	// }

	aMap := make(map[string]string)

	return aMap
}

func chromatic() []string {
	scale := make([]string, 12)

	scale[0] = "C"
	scale[1] = "C#"
	scale[2] = "D"
	scale[3] = "D#"
	scale[4] = "E"
	scale[5] = "F"
	scale[6] = "F#"
	scale[7] = "G"
	scale[8] = "G#"
	scale[9] = "A"
	scale[10] = "A#"
	scale[11] = "B"

	return scale
}

func main() {
	aScale := chromatic()
	fmt.Printf("The 12 chromatic notes are: %v\n", aScale)

}
