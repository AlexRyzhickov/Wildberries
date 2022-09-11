package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func highBoundary(number float64) float64 {
	n := int(number)
	if number >= 0 {
		n = (int(number/10) + 1) * 10
	} else {
		n = (int(number / 10)) * 10
	}
	return float64(n)
}

func lowBoundary(number float64) float64 {
	n := int(number)
	if number >= 0 {
		n = int(number/10) * 10
	} else {
		n = (int(number/10) - 1) * 10
	}
	return float64(n)
}

func main() {
	nums := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	sort.Float64s(nums)
	fmt.Println("Sorted nums:", nums)

	l := lowBoundary(nums[0])
	h := highBoundary(nums[len(nums)-1])

	j := 0
	for i := l; i < h; i = i + 10 {
		interval := make([]string, 0)
		for ; j < len(nums); j++ {
			if nums[j] >= i && nums[j] < (i+10) {
				interval = append(interval, strconv.FormatFloat(nums[j], 'f', -1, 64))
			} else {
				break
			}
		}
		if len(interval) > 0 {
			fmt.Print(i, ":")
			fmt.Print("{", strings.Join(interval, ","), "}\n")
		}
	}
}
