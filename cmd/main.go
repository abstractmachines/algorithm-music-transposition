package main

import (
	"fmt"
	// "strconv"
)

/**
- Transposition of intervals to binary visual representation.
- Returns a map of positional binary "dots" for each musical note.
**/
func transposeIntervalsToBinary(majorIntervals []int) map[int]int {
	intervalMap := make(map[int]int)

	for idx := 0; idx < 12; idx++ {
		for _, note := range majorIntervals {
			if idx == note {
				intervalMap[idx] = 1
				break
			} else {
				intervalMap[idx] = 0
			}
		}
	}

	return intervalMap
}

/** TODO
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
	notes := make([]string, 12)

	notes[0] = "C"
	notes[1] = "C#"
	notes[2] = "D"
	notes[3] = "D#"
	notes[4] = "E"
	notes[5] = "F"
	notes[6] = "F#"
	notes[7] = "G"
	notes[8] = "G#"
	notes[9] = "A"
	notes[10] = "A#"
	notes[11] = "B"

	return notes
}

func majorScale(tonic string, notes []string) []string {
	fmt.Printf("For the key of : %s", tonic)

	var transposedScale []string

	for _, note := range notes {
		transposedScale = append(transposedScale, note)
	}

	return transposedScale
}

func main() {
	allNotes := chromatic()

	newScale := majorScale("C", allNotes)
	fmt.Printf("\n\n ... You have the transposed scale %v\n", newScale)

	majorIntervals := []int{0, 2, 4, 5, 7, 9, 11}
	fmt.Println(transposeIntervalsToBinary(majorIntervals))

	minorIntervals := []int{0, 2, 3, 5, 7, 8, 10}
	fmt.Println(transposeIntervalsToBinary(minorIntervals))
}
