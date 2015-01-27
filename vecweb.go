package main

import (
	"bufio"
	"flag"
	"encoding/json"
	"github.com/danieldk/go2vec"
	"net/http"
	"io"
	"log"
	"os"
	"strings"
)
func createAnalogy(vecs map[string]go2vec.Vector) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		word1 := strings.TrimSpace(r.FormValue("word1"))
		word2 := strings.TrimSpace(r.FormValue("word2"))
		word3 := strings.TrimSpace(r.FormValue("word3"))

		result, err := go2vec.Analogy(vecs, word1, word2, word3, 10)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		resultJSON, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(resultJSON)
	}
}

func createDistance(vecs map[string]go2vec.Vector) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		word := strings.TrimSpace(r.FormValue("word"))

		result, err := go2vec.Distance(vecs, word, 10)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		resultJSON, err := json.Marshal(result)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(resultJSON)
	}
}

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		os.Exit(1)
	}

	f, err := os.Open(flag.Arg(0))
	defer f.Close()
	if err != nil {
		log.Fatal(err)
	}

	vecs, err := go2vec.ReadVectors(bufio.NewReader(io.Reader(f)))
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
		http.ServeFile(w, r, "./index.html")
	})
	http.HandleFunc("/analogy", createAnalogy(vecs))
	http.HandleFunc("/distance", createDistance(vecs))
	http.ListenAndServe(":8080", nil)
}
