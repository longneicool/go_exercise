package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

var movies = []Movie{
	{Title: "Casablanca", Year: 1942, Color: false, Actors: []string{"Humphrey Bogart", "Ingrid Bergman"}},
	{Title: "Cool Hand Luke", Year: 1967, Color: true, Actors: []string{"Paul Newman"}},
	{Title: "Bullitt", Year: 1968, Color: true, Actors: []string{"Steven", "Bisset"}}}

func main() {
	data, err := json.MarshalIndent(movies, "", "    ")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to Marshal! err %s", err)
		os.Exit(1)
	}

	fmt.Printf("data: %s\n", data)

	//通过Unmarshal来获取Ｊｓｏｎ中想要的成分
	var titles []struct{ Title string }
	if err := json.Unmarshal(data, &titles); err != nil {
		log.Fatalf("JSON unmarshaling failed: %s ", err)
	}

	fmt.Println(titles)
}
