package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
)

type MSQ struct {
	Notes [][]interface{} `json:"notes"`
	Loop  bool            `json:"loop"`
	End   int             `json:"end"`
	Tempo interface{}     `json:"tempo"`
	Beats int             `json:"beats"`
}

func (m *MSQ) Log() {
	log.Printf("notes: %v", m.Notes)
	log.Printf("loop: %v", m.Loop)
	log.Printf("end: %v", m.End)
	log.Printf("tempo: %v", m.Tempo)
	log.Printf("beats: %v", m.Beats)
}

func main() {
	var notes = map[int]string{
		0:   "NOTE_G5",
		1:   "NOTE_F5",
		2:   "NOTE_E5",
		3:   "NOTE_D5",
		4:   "NOTE_C5",
		5:   "NOTE_B4",
		6:   "NOTE_A4",
		7:   "NOTE_G4",
		8:   "NOTE_F4",
		9:   "NOTE_E4",
		10:  "NOTE_D4",
		11:  "NOTE_C4",
		12:  "NOTE_B3",
		64:  "NOTE_FS5",
		66:  "NOTE_DS5",
		67:  "NOTE_CS5",
		69:  "NOTE_AS4",
		70:  "NOTE_GS4",
		71:  "NOTE_FS4",
		73:  "NOTE_DS4",
		74:  "NOTE_CS4",
		128: "NOTE_GS5",
		129: "NOTE_FS5",
		131: "NOTE_DS5",
		132: "NOTE_CS5",
		134: "NOTE_AS4",
		135: "NOTE_GS4",
		136: "NOTE_FS4",
		138: "NOTE_DS4",
		139: "NOTE_CS4",
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
