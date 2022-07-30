package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

type MSQ struct {
	Notes []interface{} `json:"notes"`
	Loop  bool          `json:"loop"`
	End   int           `json:"end"`
	Tempo interface{}   `json:"tempo"`
	Beats int           `json:"beats"`
}

func (m *MSQ) Log() {
	log.Printf("notes: %v", m.Notes)
	log.Printf("loop: %v", m.Loop)
	log.Printf("end: %v", m.End)
	log.Printf("tempo: %v", m.Tempo)
	log.Printf("beats: %v", m.Beats)
}

func main() {
	var notes = map[string][]int{
		"NOTE_B3":  {12},
		"NOTE_C4":  {11},
		"NOTE_CS4": {74, 139},
		"NOTE_D4":  {10},
		"NOTE_DS4": {73, 138},
		"NOTE_E4":  {9},
		"NOTE_F4":  {8},
		"NOTE_FS4": {71, 136},
		"NOTE_G4":  {7},
		"NOTE_GS4": {70, 135},
		"NOTE_A4":  {6},
		"NOTE_AS4": {69, 134},
		"NOTE_B4":  {5},
		"NOTE_C5":  {4},
		"NOTE_CS5": {67, 132},
		"NOTE_D5":  {3},
		"NOTE_DS5": {66, 131},
		"NOTE_E5":  {2},
		"NOTE_F5":  {1},
		"NOTE_FS5": {64, 129},
		"NOTE_G5":  {0},
		"NOTE_GS5": {128},
	}

	filename := flag.String("file", "", "File to parse")
	flag.Parse()

	data, err := ioutil.ReadFile(*filename)
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	var msq MSQ
	err = json.Unmarshal(data, &msq)
	if err != nil {
		log.Fatal("Error parsing JSON data: ", err)
	}

	msq.Log()
	log.Print(notes)
}
