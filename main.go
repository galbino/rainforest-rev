package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	url := "https://letsrevolutionizetesting.com/challenge.json"
	var newResp map[string]string
	for url != "" {
		resp, err := http.Get(url)
		if err != nil {
			log.Fatal("error during http request:%w", err)
			continue
		}
		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal("error during http request:%w", err)
			continue
		}
		newResp = make(map[string]string, 0)
		json.Unmarshal(body, &newResp)
		url = strings.Replace(newResp["follow"], "challenge", "challenge.json", 1)
	}
	fmt.Println(newResp)
}
