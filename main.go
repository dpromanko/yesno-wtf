package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type yesNoResponse struct {
	Answer string
	Forced bool
	Image  string
}

func getAnswer(w http.ResponseWriter, r *http.Request) {
	const url = "http://yesno.wtf/api"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err)
	}

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	defer resp.Body.Close()

	var answer yesNoResponse

	if err := json.NewDecoder(resp.Body).Decode(&answer); err != nil {
		log.Panic(err)
	}

	w.Write([]byte(answer.Answer))
}

func main() {
	http.HandleFunc("/", getAnswer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
