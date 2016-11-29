package main

import (
	"bufio"
	"errors"
	"strings"

	"gopkg.in/yaml.v2"
)

// Config parsing errors
var (
	ErrNotFound = errors.New("not found")
)

// A ReifiedTypes is mapping of a generic type to a concrete type
type ReifiedTypes map[string]string

// NameExtension is the name extension to a function or type
func (t ReifiedTypes) NameExtension() string {
	var vs []string
	for _, v := range t {
		vs = append(vs, v)
	}
	return strings.Join(vs, "_")
}

// A ReifyConfig specifies how to reify types
type ReifyConfig struct {
	TypeSpecs map[string][]string
}

// Permutations returns all the possible permutations of the types
func (cfg *ReifyConfig) Permutations() []ReifiedTypes {
	type pair struct{ fst, snd string }
	var all []ReifiedTypes
	var permute func([]pair, []string)
	permute = func(set []pair, left []string) {
		switch len(left) {
		case 0:
			m := ReifiedTypes{}
			for _, p := range set {
				m[p.fst] = p.snd
			}
			all = append(all, m)
		default:
			for _, v := range cfg.TypeSpecs[left[0]] {
				permute(append(set, pair{left[0], v}), left[1:])
			}
		}
	}
	var allkeys []string
	for k := range cfg.TypeSpecs {
		allkeys = append(allkeys, k)
	}
	permute(nil, allkeys)
	return all
}

// ParseConfigFromComment parses the reify config from a comment
func ParseConfigFromComment(comment string) (*ReifyConfig, error) {
	s := bufio.NewScanner(strings.NewReader(comment))
	const (
		normal byte = iota
		startConfig
		inConfig
	)

	mode := normal
	indent := 0
	configtext := ""
	for s.Scan() {
		ln := s.Text()
		switch mode {
		case normal:
			if strings.TrimSpace(ln) == "reify:" {
				mode = startConfig
			}
		case startConfig:
			indent = getIndent(ln)
			if indent == 0 {
				mode = normal
			} else {
				mode = inConfig
				configtext += ln + "\n"
			}
		case inConfig:
			if getIndent(ln) < indent {
				mode = normal
			} else {
				configtext += ln + "\n"
			}
		}
	}

	if configtext == "" {
		return nil, ErrNotFound
	}

	var tmp struct {
		Types map[string]string `yaml:"types"`
	}
	err := yaml.Unmarshal([]byte(configtext), &tmp)
	if err != nil {
		return nil, err
	}

	cfg := &ReifyConfig{
		TypeSpecs: map[string][]string{},
	}
	for name, types := range tmp.Types {
		cfg.TypeSpecs[name] = expandTypes(strings.Split(types, ",")...)
	}
	return cfg, nil
}

func getIndent(txt string) int {
	return len(txt) - len(strings.TrimLeft(txt, " \t"))
}

func expandTypes(types ...string) []string {
	var all []string
	for _, typ := range types {
		switch typ {
		case "numeric":
			fallthrough
		case "numerics":
			all = append(all,
				expandTypes("integers", "floats", "complex values")...)
		case "integers":
			all = append(all,
				"int", "int8", "int16", "int32", "int64",
				"uint", "uint8", "uint16", "uint32", "uint64")
		case "floats":
			all = append(all, "float32", "float64")
		case "complex values":
			all = append(all, "complex64", "complex128")
		default:
			all = append(all, typ)
		}
	}
	return all
}
