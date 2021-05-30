package main

import (
	"fmt"
	// "strconv"
)

/**
- Transposition of intervals to binary visual representation.
- Returns a map of positional binary "dots" for each musical note.
**/
func transposeIntervalsToBinary(intervals []int) map[int]int {
	intervalMap := make(map[int]int)

	for idx := 0; idx < 12; idx++ {
		for _, note := range intervals {
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

/**
- Use Go slices, modulo to emulate a circular ring buffer with wraparound ...
**/
func shiftNotes(tonic string, chromatic []string) []string {
	fmt.Printf("For the key of : %s", tonic)

	var shiftedNotes []string
	var offset int
	var moduloOffset int

	for idx, note := range chromatic {
		if note == tonic {
			offset = idx
		}

		firstSlice := chromatic[offset:]
		shiftedNotes = firstSlice

		moduloOffset = offset % 12
		wrappedSlice := chromatic[0:moduloOffset]
		shiftedNotes = append(shiftedNotes, wrappedSlice...)
	}

	return shiftedNotes
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

func main() {
	allNotes := chromatic()

	majorIntervals := []int{0, 2, 4, 5, 7, 9, 11}
	fmt.Println(transposeIntervalsToBinary(majorIntervals))

	minorIntervals := []int{0, 2, 3, 5, 7, 8, 10}
	fmt.Println(transposeIntervalsToBinary(minorIntervals))

	// WIP apply those mappings to the new shifted note system:
	shiftedNotes := shiftNotes("G", allNotes)
	fmt.Printf("\n\n ... You have the newly shifted 12 notes of: %v\n", shiftedNotes)

	fmt.Println("\n\n ... The next thing to do is to apply the mapped intervals to the notes\n")
}
