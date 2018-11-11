package main

import (
	"encoding/json"
	"html/template"
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

	switch r.URL.Path {
	case "/api":
		switch r.FormValue("output") {
		case "prettyjson":
			var b []byte
			b, err := json.MarshalIndent(resp, "", " ")
			if err == nil {
				_, err = w.Write(b)
			}
		default:
			json.NewEncoder(w).Encode(resp)
		}
	default:
		t, err := template.ParseFiles("template.html")
		if err != nil {
			log.Panic(err)
		}
		t.Execute(w, resp)
	}
}

func main() {
	http.HandleFunc("/", getAnswer)
	log.Fatal(http.ListenAndServe(":8080", nil))
	http.HandleFunc("/api", getAnswer)
	log.Fatal(http.ListenAndServe(":8080/api", nil))
}
