package main

import (
	"encoding/json"
	"log"
	"os"
	"sort"
	"strings"
)

// tpf needs to be the path to the file containing the contexts (a
// complex JSON format. args is the argument to the command.
func genAlfredResult(tpf string, args []string) {
	//log.Println("args", args)
	var result Result

	encoder := json.NewEncoder(os.Stdout)
	encoder.SetIndent("", "	")

	// args should be a single string corresponding
	pr := strings.Join(args, " ")

	tags, err := getContexts(tpf)
	if err != nil {
		log.Printf("can't read contexts from %q: %v", tpf, tags)
		return
	}

	th := make(map[string]int)
	// Keep a running count of all of the match lengths
	for _, t := range tags {
log.Println("tag", t, pr)
		if (strings.HasPrefix(t, pr) || strings.HasPrefix(strings.ToUpper(t), strings.ToUpper(pr))) && pr != t {
log.Println("prefix matched")
			// TODO(rjk): I have something weird happening here?
			if s, ok := th[t]; !ok || s < len(pr) {
				th[t] = len(pr)
			}
		}
	}

	//log.Println(th)

	// Create a result based on the hash
	for t, v := range th {
		// Exclude the Uid field to make sure that the items aren't re-ordered.
		result.Items = append(result.Items, &Item{
			Title:        t,
			Arg:          t,
			Autocomplete: t,
			relevance:    v,
			Valid:        true,
		})
	}
	sort.Sort(result.Items)

	if err := encoder.Encode(result); err != nil {
		log.Fatalf("can't write json %v", err)
	}
}

type Item struct {
	Uid          string `json:"uid,omitempty"`
	Title        string `json:"title"`
	Subtitle     string `json:"subtitle,omitempty"`
	Arg          string `json:"arg"`
	Autocomplete string `json:"autocomplete"`
	relevance    int
	Valid        bool `json:"valid"`
}

type Result struct {
	Items ItemCollection `json:"items"`
}

type ItemCollection []*Item

func (c ItemCollection) Len() int {
	return len(c)
}

func (c ItemCollection) Less(i, j int) bool {
	return c[i].relevance > c[j].relevance
}

func (c ItemCollection) Swap(i, j int) {
	tmp := c[i]
	c[i] = c[j]
	c[j] = tmp
}

var _ = ItemCollection(nil)
