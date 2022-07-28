package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type MSQ struct {
	Notes [][]int `json:"notes"`
	Loop  bool    `json:"loop"`
	End   int     `json:"end"`
	Tempo int     `json:"tempo"`
	Beats int     `json:"beats"`
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

	log.Printf("notes: %v", msq.Notes)
	log.Printf("loop: %v", msq.Loop)
	log.Printf("end: %v", msq.End)
	log.Printf("tempo: %v", msq.Tempo)
	log.Printf("beats: %v", msq.Beats)
}
