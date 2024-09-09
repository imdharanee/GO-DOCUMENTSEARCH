package main

import (
	"flag"
	"log"
	"time"
	utils "GO-TEXTSEARCH",
)

func main() {
	var dumpath, query string
	flag.StringVar(&dumpath, "p", "file-path-of-dataset-dump")
	flag.StringVar(&query, "Small wild cat", "searchquery")
	flag.Parse()
	log.Println("Full tex search is in progress")
	start := time.Now()
	docs, err := utils.LoadDocuments(dumpath)

	if err != nil {
		log.Fatal(err)

	}
	log.Printf("Loaded %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	idx := make(utils.Index)
	idx.Add(docs)
	log.Printf("Indexed %d documents in %v", len(docs), time.Since(start))
	start = time.Now()
	matchedId := idx.Search(query)
	log.Printf("Search found %d documents in %v", len(matchedId), time.Since(start))
	for _, id := range matchedId {

		doc := docs[id]
		log.Printf("%d\t%s\n", id, doc.Text)

	}

}
