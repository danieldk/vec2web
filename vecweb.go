package main

//go:generate go-bindata -pkg $GOPACKAGE -o bindata.go static/

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/danieldk/go2vec"
)

func createAnalogy(vecs *go2vec.Vectors) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		word1 := strings.TrimSpace(r.FormValue("word1"))
		word2 := strings.TrimSpace(r.FormValue("word2"))
		word3 := strings.TrimSpace(r.FormValue("word3"))

		result, err := vecs.Analogy(word1, word2, word3, 10)
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

func createSimilarity(vecs *go2vec.Vectors) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		word := strings.TrimSpace(r.FormValue("word"))

		result, err := vecs.Similarity(word, 10)
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
		fmt.Fprintln(os.Stderr, "Usage: vecweb config")
		flag.PrintDefaults()
		os.Exit(1)
	}

	config := readConfigOrExit(flag.Arg(0))
	if len(config.WordEmbedding) == 0 {
		log.Fatal("No embeddings specified in the configuration file")
	}

	f, err := os.Open(config.WordEmbedding[0].Path)
	fatalIfErr("Cannot open word embeddings", err)
	defer f.Close()

	vecs, err := go2vec.ReadVectors(bufio.NewReader(f), true)
	fatalIfErr("Could not read word embeddings", err)
	log.Printf("Loaded vectors from %s", config.WordEmbedding[0].Path)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}

		assetOrError("static/index.html", w)
	})
	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/about" {
			http.NotFound(w, r)
			return
		}

		assetOrError("static/about.html", w)
	})
	http.HandleFunc("/analogy", createAnalogy(vecs))
	http.HandleFunc("/distance", createSimilarity(vecs))

	log.Printf("Starting to serve from %s", config.HTTP)
	http.ListenAndServe(config.HTTP, nil)
}
