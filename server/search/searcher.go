package search

import (
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type Searcher struct {
	Path string
}

type Entry struct {
	File  string `json:"file"`
	Count int    `json:"count"`
	Slug  string `json:"slug"`
}

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type SortByCount []Entry

func (a SortByCount) Len() int           { return len(a) }
func (a SortByCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByCount) Less(i, j int) bool { return a[i].Count > a[j].Count }

type SortByFile []Entry

func (a SortByFile) Len() int           { return len(a) }
func (a SortByFile) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a SortByFile) Less(i, j int) bool { return strings.Compare(a[i].File, a[j].File) == 1 }

func glob(dir string, ext string) ([]string, error) {

	files := []string{}
	err := filepath.Walk(dir, func(path string, f os.FileInfo, err error) error {
		if filepath.Ext(path) == ext {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func ReadText(file string) (string, error) {
	data, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	text := string(data)
	// TODO: remove frontmatter +++ ---

	return text, nil
}

func extractSlug(file string) string {
	text, err := ReadText(file)
	if err != nil {
		log.Printf("[ERROR] extractSlug fail for %s: %v", file, err)
		return ""
	}
	runs := []rune(text)

	return string(runs[0:IntMin(240, len(runs))])
}

func (s *Searcher) Find(query string, limit int) []Entry {
	files, err := glob(s.Path, ".md")
	if err != nil {
		log.Printf("[ERROR] cannot read files at %s: %v", s.Path, err)
		return nil
	}

	// TODO: this should be smart algorythm to range files based on the query
	//  Algorythm counts number of tokens the text contains
	//  Range files based on this counter
	tokens := strings.Split(query, " ")
	counter := make(map[string]int)
	for _, file := range files {
		text, err := ReadText(file)
		if err != nil {
			log.Printf("[ERROR] Fail to read file %s: %v", file, err)
			continue
		}

		for _, token := range tokens {
			if strings.Contains(text, token) {
				counter[file] += 1
			}
		}
	}

	tuples := make([]Entry, 0)
	for file, count := range counter {
		if count > 0 {
			tuples = append(tuples, Entry{
				File:  file,
				Count: count,
			})
		}
	}

	sort.Sort(SortByFile(tuples))
	sort.Stable(SortByCount(tuples))

	results := make([]Entry, IntMin(limit, len(tuples)))
	for i := range results {
		log.Printf("%d %s\n", tuples[i].Count, tuples[i].File)

		results[i] = tuples[i]
		results[i].Slug = extractSlug(tuples[i].File)
	}

	return results
}
