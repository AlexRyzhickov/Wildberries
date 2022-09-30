package main

import (
	"flag"
	"fmt"
	"log"
	"strings"
)

func cut(str, sep string, f int, s bool) {
	if f >= 1 {
		f--
	}
	lines := strings.Split(str, "\n")
	strs := make([][]string, 0)
	for _, line := range lines {
		columns := strings.Split(line, sep)
		if f > len(columns) {
			log.Println("Invalid f flag")
			return
		}
		strs = append(strs, columns)
	}
	for _, columns := range strs {
		if !s {
			fmt.Println(columns[f])
		} else {
			if len(columns) > 1 {
				fmt.Println(columns[f])
			}
		}
	}

}

//echo "Winter: white: snow: frost
//Spring: green: grass: warm
//Summer: colorful: blossom: hot
//Autumn: yellow: leaves: cool
//Another" | cut -s -d ":" -f 1

//go run main.go -s -d ":" -f 1 "Winter: white: snow: frost
//Spring: green: grass: warm
//Summer: colorful: blossom: hot
//Autumn: yellow: leaves: cool
//Another"

func main() {
	f := flag.Int("f", 0, "\"fields\" - выбрать поля (колонки)")
	d := flag.String("d", "\t", "\"delimiter\" - использовать другой разделитель\n")
	s := flag.Bool("s", false, "\"separated\" - только строки с разделителем\n")
	flag.Parse()
	tail := flag.Args()
	cut(tail[0], *d, *f, *s)
}
