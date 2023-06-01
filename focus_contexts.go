package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/itchyny/gojq"
)

// getContexts extracts the TaskPaper file tags from tpf.
func getContexts(tpf string) ([]string, error) {
	fd, err := os.Open(tpf)
	if err != nil {
		return []string{}, err
	}
	defer fd.Close()

	query, err := gojq.Parse(".data.[0].modeConfigurations.[].mode.name")
	if err != nil {
		return []string{}, fmt.Errorf("can't parse the gojq command: %v", err)
	}

	dec := json.NewDecoder(fd)
	var jsonpayload map[string]any

	if err := dec.Decode(&jsonpayload); err != nil {
		return []string{}, fmt.Errorf("can't decode %q: %v", tpf, err)
	}

	iter := query.Run(jsonpayload)
	focusmodes := make([]string, 0)
	for {
		v, ok := iter.Next()
		if !ok {
			break
		}
		if err, ok := v.(error); ok {
			return []string{}, fmt.Errorf("gojq iteration failed: %v", err)
		}
		if s, ok := v.(string); ok {
			focusmodes = append(focusmodes, s)
		}
	}

	return focusmodes, nil
}
