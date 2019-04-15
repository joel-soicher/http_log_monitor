package main

import (
	"sort"
	"strconv"
	"strings"
)

// Implements the section checker
type Section struct {
	cfg      *Config
	sections map[string]int
}

func NewSection(cfg *Config) *Section {
	return &Section{
		cfg:      cfg,
		sections: make(map[string]int),
	}
}

func (s *Section) AddRequest(req *Request) {
	if req == nil {
		return
	}
	split := strings.Split(req.Resource, "/")
	if len(split) < 2 {
		return // Should return an error
	}
	s.sections[split[1]]++
}

func (s *Section) Compute() {
}

func (s *Section) DisplayString() string {
	if len(s.sections) == 0 {
		return ""
	}

	// Temporary map used to switch key and value for sorting
	tmp := map[int]string{}
	// This algorithm is not optimized, but as there are just a few elements in the map,
	// it will be more than enough for our needs
	tmpKeys := []int{}
	for k, v := range s.sections {
		tmp[v] = k
		tmpKeys = append(tmpKeys, v)
	}
	sort.Ints(tmpKeys)

	result := make([]string, len(tmpKeys))
	nb := len(tmpKeys)
	if nb > s.cfg.MaxSections {
		nb = s.cfg.MaxSections
	}
	for i := 0; i < nb; i++ {
		k := tmpKeys[i]
		result[len(tmpKeys)-i-1] = tmp[k] + "(" + strconv.FormatInt(int64(k), 10) + ")"
	}

	return "MostHits sections: " + strings.Join(result, ",")
}

func (s *Section) Flush() {
	s.sections = make(map[string]int)
}
