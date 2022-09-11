package main

import "fmt"

type Set struct {
	m map[string]struct{}
}

func InitializeSet() *Set {
	m := make(map[string]struct{})
	m["cat"] = struct{}{}
	m["dog"] = struct{}{}
	m["tree"] = struct{}{}
	return &Set{m: m}
}

func (s *Set) IsInSet(key string) bool {
	_, ok := s.m[key]
	return ok
}

func main() {
	set := InitializeSet()
	fmt.Println(set.IsInSet("cat"))
	fmt.Println(set.IsInSet("dog"))
	fmt.Println(set.IsInSet("tree"))
	fmt.Println(set.IsInSet("world"))
}
