package main

import (
	"fmt"
)

/**
- For any mode or scale, transpose music intervals to binary representation.
- Returns a map of positional binary "dots" for each musical note.
- Would maybe be cool? Using bits and bit masks instead of int ranges.
**/
func binaryMusicIntervals(intervals []int) map[int]int {
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

func mapIntervalsToNotes(shiftedNotes []string, intervalMap map[int]int) []string {
	var finalCut []string
	for idx, note := range shiftedNotes {
		if intervalMap[idx] == 1 {
			finalCut = append(finalCut, note)
		}
	}
	return finalCut
}

/**
- Shift 12 notes to select from. Emulate ring buffer w/ Go slices, modulo.
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

	ionianIntervals := []int{0, 2, 4, 5, 7, 9, 11}
	aeolianIntervals := []int{0, 2, 3, 5, 7, 8, 10}

	ionianBinary := binaryMusicIntervals(ionianIntervals)
	aeolianBinary := binaryMusicIntervals(aeolianIntervals)

	shiftedNotes := shiftNotes("G", allNotes)

	keySignatureMajor := mapIntervalsToNotes(shiftedNotes, ionianBinary)
	keySignatureMinor := mapIntervalsToNotes(shiftedNotes, aeolianBinary)
	fmt.Printf("\n The major ionian key signature is %v\n", keySignatureMajor)
	fmt.Printf("\n and the minor aeolian %v\n", keySignatureMinor)
}
