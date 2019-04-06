package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const url = "https://xkcd.com/571/info.0.json"

func main() {
	resp, _ := http.Get(url)
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Query failed! status:%s", resp.Status)
		os.Exit(1)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to readall!")
		os.Exit(1)
	}

	fmt.Printf("%s\n", body)

	var titles []struct{ month string }
	jsonerr := json.Unmarshal(body, titles)
	if jsonerr != nil {
		fmt.Fprintf(os.Stderr, "Failed to Unmarshal!")
		os.Exit(1)
	}

	fmt.Println(titles)
}
