package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bmizerany/pat"
	"github.com/mvrilo/go-cpf"
)

func generate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(cpf.GeneratePretty()))
}

func validate(w http.ResponseWriter, r *http.Request) {
	var value = r.URL.Query().Get(":cpf")

	if value == "about" || value == "info" || value == "source" || value == "src" {
		http.Redirect(w, r, "https://github.com/mvrilo/go-cpf", 302)
		return
	}

	if ok, _ := cpf.Valid(value); ok {
		w.Write([]byte("yep"))
		return
	}
	w.Write([]byte("nope"))
}

func main() {
	var m = pat.New()
	m.Get("/", http.HandlerFunc(generate))
	m.Get("/:cpf", http.HandlerFunc(validate))

	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	println("-- cpf web is running at " + port + " --")
	log.Fatal(http.ListenAndServe(":"+port, m))
}
