package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strings"
)

func isMatched(str, regExpr string, v, i, F bool) bool {
	if !F {
		if i {
			regExpr = strings.ToLower(regExpr)
			str = strings.ToLower(str)
		}
		matched, _ := regexp.MatchString(regExpr, str)
		return !v && matched || v && !matched
	}
	return strings.Contains(str, regExpr)
}

func grep(regExpr string, strings []string, after, before, context int, v, c, ignore, f, n bool) {
	nums := make([]int, 0)
	m := make(map[int]struct{}, 0)
	for i, str := range strings {
		if isMatched(str, regExpr, v, ignore, f) {
			nums = append(nums, i)
		}
	}
	if c {
		fmt.Println(len(nums))
		return
	}
	if after < context {
		after = context
	}
	if before < context {
		before = context
	}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		min := num - before
		max := num + after
		if min < 0 {
			min = 0
		}
		if max >= len(strings) {
			max = len(strings) - 1
		}
		for j := min; j <= max; j++ {
			m[j] = struct{}{}
		}
	}
	allNums := make([]int, 0)
	for k, _ := range m {
		allNums = append(allNums, k)
	}
	sort.Ints(allNums)
	if n {
		for i := 0; i < len(allNums); i++ {
			fmt.Printf("%d: %s\n", allNums[i], strings[allNums[i]])
		}
	} else {
		for i := 0; i < len(allNums); i++ {
			fmt.Println(strings[allNums[i]])
		}
	}
}

// Example 1
// grep -B 20 ...z test.txt
// go run main.go -B 20 -F ...z test.txt

//Example 2
// grep  -F -n 7 test.txt
// go run main.go  -F -n 7 test.txt

func main() {
	A := flag.Int("A", 0, "\"after\" печатать +N строк после совпадения")
	B := flag.Int("B", 0, "\"before\" печатать +N строк до совпадения")
	C := flag.Int("C", 0, "\"context\" (A+B) печатать ±N строк вокруг совпадения")
	c := flag.Bool("c", false, "\"count\" (количество строк)")
	i := flag.Bool("i", false, "\"ignore-case\" (игнорировать регистр)")
	v := flag.Bool("v", false, "\"invert\" (вместо совпадения, исключать)")
	F := flag.Bool("F", false, "\"fixed\", точное совпадение со строкой, не паттерн")
	n := flag.Bool("n", false, "\"line num\", напечатать номер строки")
	flag.Parse()
	tail := flag.Args()
	if len(tail) < 2 {
		log.Println("Invalid input")
		return
	}
	regExpr := tail[0]
	fileName := tail[1]
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()
	lines := make([]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		lines = append(lines, sc.Text())
	}
	grep(regExpr, lines, *A, *B, *C, *v, *c, *i, *F, *n)
}
