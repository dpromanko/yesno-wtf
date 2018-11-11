package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type yesNoAPI interface {
	getYesNo()
}

type yesNoResponse struct {
	Answer string
	Forced bool
	Image  string
}

func (r *yesNoResponse) getYesNo() {
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

	if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
		log.Panic(err)
	}
}

func getAnswer(w http.ResponseWriter, r *http.Request) {
	resp := yesNoResponse{}
	var api yesNoAPI = &resp
	api.getYesNo()

	w.Write([]byte(resp.Answer))
}

func main() {
	http.HandleFunc("/", getAnswer)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
