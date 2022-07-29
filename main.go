package main

import (
	"encoding/json"
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
	data, err := ioutil.ReadFile("./demo-msq.json")
	if err != nil {
		log.Fatal("Error opening file: ", err)
	}

	var msq MSQ
	err = json.Unmarshal(data, &msq)
	if err != nil {
		log.Fatal("Error parsing JSON data: ", err)
	}

	msq.Log()
}
