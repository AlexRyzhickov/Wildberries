package main

import (
	"fmt"
	"sort"
	"strings"
)

type ByRune []rune

func (r ByRune) Len() int           { return len(r) }
func (r ByRune) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }
func (r ByRune) Less(i, j int) bool { return r[i] < r[j] }

func AnagramsSet(strs []string) map[string][]string {
	m := make(map[string][]string, 0)
	for _, str := range strs {
		runes := []rune(str)
		sort.Sort(ByRune(runes))
		key := strings.ToLower(string(runes))
		m[key] = append(m[key], strings.ToLower(str))
	}
	for k, v := range m {
		if len(v) <= 1 {
			delete(m, k)
		} else {
			newK := v[0]
			v = v[1:]
			sort.Strings(v)
			delete(m, k)
			m[newK] = v
		}
	}
	return m
}

func main() {
	//Out map[кот:[ток] листок:[слиток столик] пятак:[пятка тяпка]]
	anagrams := []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "кот", "ток", "профсоюз"}
	fmt.Println(AnagramsSet(anagrams))
}
