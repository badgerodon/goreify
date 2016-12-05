package main

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

// Config parsing errors
var (
	ErrNotFound = errors.New("not found")
)

// A ReifiedTypes is mapping of a generic type to a concrete type
type ReifiedTypes map[string]string

// NameExtension is the name extension to a function or type
func (t ReifiedTypes) NameExtension() string {
	ks := make([]string, 0, len(t))
	for k := range t {
		ks = append(ks, k)
	}
	sort.Strings(ks)

	if options.namingConvention == NamingConventionUnderscore {
		var vs []string
		for _, k := range ks {
			vs = append(vs, t[k])
		}
		return "_" + strings.Join(vs, "_")
	}

	var vs []string
	for _, k := range ks {
		vs = append(vs, strings.Title(t[k]))
	}
	return strings.Join(vs, "")
}

// A ReifyConfig specifies how to reify types
type ReifyConfig struct {
	TypeSpecs map[string][]string
}

// Set sets the cfg from a string
func (cfg *ReifyConfig) Set(str string) error {
	for _, str := range strings.Fields(str) {
		if strings.Contains(str, ":") {
			xs := strings.SplitN(str, ":", 2)
			cfg.TypeSpecs[xs[0]] = expandTypes(strings.Split(xs[1], ",")...)
		} else {
			i := len(cfg.TypeSpecs) + 1
			cfg.TypeSpecs[fmt.Sprint("T", i)] = expandTypes(strings.Split(str, ",")...)
		}
	}
	return nil
}

// String returns the config as a string
func (cfg *ReifyConfig) String() string {
	var keys []string
	for k := range cfg.TypeSpecs {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var kvs []string
	for _, k := range keys {
		v := strings.Join(cfg.TypeSpecs[k], ",")
		kv := fmt.Sprintf("%v:%v", k, v)
		kvs = append(kvs, kv)
	}
	return strings.Join(kvs, " ")
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
