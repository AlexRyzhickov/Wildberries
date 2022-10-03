package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

func sortFile(s [][]string, n bool, k int) [][]string {
	if !n {
		sort.SliceStable(s, func(i, j int) bool {
			a := s[i][k]
			b := s[j][k]
			if len(a) > 0 && len(b) > 0 {
				l1 := []rune(a)[0]
				l2 := []rune(b)[0]
				if unicode.IsLetter(l1) && unicode.IsLetter(l2) {
					if strings.ToLower(a) == strings.ToLower(b) {
						if unicode.IsLower(l1) && unicode.IsUpper(l2) {
							return true
						} else {
							return false
						}
					}
				}
			}
			return strings.ToLower(a) < strings.ToLower(b)
		})
	} else {
		sort.SliceStable(s, func(i, j int) bool {
			num1, _ := strconv.Atoi(s[i][k])
			num2, _ := strconv.Atoi(s[j][k])
			return num1 < num2
		})
	}
	return s
}

// Example: go run main.go -r -filePath test.txt

func main() {
	k := flag.Int("k", 0, "help message for flag k")
	n := flag.Bool("n", false, "help message for flag n")
	r := flag.Bool("r", false, "help message for flag r")
	fileName := flag.String("filePath", "", "help message for flag filePath")
	flag.Parse()

	f, err := os.Open(*fileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer f.Close()

	lines := make([][]string, 0)
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := strings.Split(sc.Text(), " ")
		lines = append(lines, line)
	}
	sortFile(lines, *n, *k)
	for _, line := range lines {
		fmt.Println(strings.Join(line, " "))
	}

	f, err = os.Create(*fileName)

	if !(*r) {
		for _, line := range lines {
			_, err = f.WriteString(strings.Join(line, " ") + "\n")
			if err != nil {
				log.Println(err)
				return
			}
		}
	} else {
		for i := len(lines) - 1; i > 0; i-- {
			_, err = f.WriteString(strings.Join(lines[i], " ") + "\n")
			if err != nil {
				log.Println(err)
				return
			}
		}
	}
}
