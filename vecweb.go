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

	"gopkg.in/danieldk/go2vec.v1"
)

type loadedEmbeddings map[string]*go2vec.Embeddings

func createAnalogy(embeddings loadedEmbeddings) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		embeddingsName := strings.TrimSpace(r.FormValue("embeddings"))
		curEmbeddings, ok := embeddings[embeddingsName]
		if !ok {
			http.Error(w, "Unknown embeddings", 500)
			return
		}

		word1 := strings.TrimSpace(r.FormValue("word1"))
		word2 := strings.TrimSpace(r.FormValue("word2"))
		word3 := strings.TrimSpace(r.FormValue("word3"))

		result, err := curEmbeddings.Analogy(word1, word2, word3, 10)
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

func createEmbeddings(embeddings loadedEmbeddings) func(http.ResponseWriter, *http.Request) {
	names := make([]string, 0, len(embeddings))
	for name := range embeddings {
		names = append(names, name)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		namesJSON, err := json.Marshal(names)
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}

		w.Write(namesJSON)
	}
}

func createSimilarity(embeddings loadedEmbeddings) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		embeddingsName := strings.TrimSpace(r.FormValue("embeddings"))
		curEmbeddings, ok := embeddings[embeddingsName]
		if !ok {
			http.Error(w, "Unknown embeddings", 500)
			return
		}

		word := strings.TrimSpace(r.FormValue("word"))

		result, err := curEmbeddings.Similarity(word, 10)
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

func loadEmbedding(e wordEmbedding) *go2vec.Embeddings {
	f, err := os.Open(e.Path)
	fatalIfErr("Cannot open word embeddings", err)
	defer f.Close()

	vecs, err := go2vec.ReadWord2VecBinary(bufio.NewReader(f), true)
	fatalIfErr("Could not read word embeddings", err)
	log.Printf("Loaded vectors from %s", e.Path)

	return vecs
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

	embeddings := make(loadedEmbeddings)

	for _, embeddingConf := range config.WordEmbedding {
		embeddings[embeddingConf.Name] = loadEmbedding(embeddingConf)
	}

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
	http.HandleFunc("/analogy", createAnalogy(embeddings))
	http.HandleFunc("/distance", createSimilarity(embeddings))
	http.HandleFunc("/embeddings", createEmbeddings(embeddings))

	log.Printf("Starting to serve from %s", config.HTTP)
	http.ListenAndServe(config.HTTP, nil)
}
