package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"sort"

	"github.com/PuerkitoBio/goquery"
)

var kmlFile = flag.String("kmlfile", "", "path to kml file")

type KmlSnippet struct {
	Name        string
	Coordinates string
}

type KmlSnippets []KmlSnippet

func (k KmlSnippets) Len() int {
	return len(k)
}

func (k KmlSnippets) Less(i, j int) bool {
	return k[i].Name < k[j].Name
}

func (k KmlSnippets) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func dumpToCsv(data []KmlSnippet) error {
	return nil
}

func extractSnippet(file string) (KmlSnippets, error) {
	var k []KmlSnippet

	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		return nil, err
	}
	doc.Find("placemark").Each(func(i int, s *goquery.Selection) {
		name := s.Find("name").Text()
		coordinates := s.Find("coordinates").Text()
		if name != "" {
			k = append(k, KmlSnippet{name, coordinates})
		}
	})

	return k, nil
}

func main() {
	flag.Parse()
	if *kmlFile == "" {
		log.Fatal("kmlfile flag missing")
	}
	snippets, err := extractSnippet(*kmlFile)
	if err != nil {
		log.Fatal(err)
	}
	sort.Sort(snippets)
	fmt.Println("\nSorted")
	for k, v := range snippets {
		fmt.Println(k, v.Name)
	}
}
