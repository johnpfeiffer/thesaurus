package main

import (
	"github.com/johnpfeiffer/thesaurus/my-go-app/words"
	"html/template"
	"log"
	"net/http"
	"os"
)

type PageData struct {
	Word       string
	Synonym    string
	HasSynonym bool
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/md", mdHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data := PageData{}
	if r.Method == http.MethodPost {
		word := r.FormValue("word")
		data.Word = word
		if synonym, ok := words.IsSynonymOfTwoLetterWord(word); ok {
			data.Synonym = synonym
			data.HasSynonym = true
		}
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
