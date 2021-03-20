package main

import (
	"fmt"
	"os"
)

type line struct {
	Start  int
	Height int
}

type lines []line

func (result lines) Append(line line) lines {
	size := len(result)
	if size > 0 && result[size-1].Height == line.Height {
		return result
	}

	if size > 0 && result[size-1].Start == line.Start {
		if line.Height > result[size-1].Height {
			result[size-1].Height = line.Height
		}
		return result
	}

	return append(result, line)
}

func main() {
	var (
		lines  []line
		start  int
		height int
		end    int
	)

	for {
		size, _ := fmt.Fscanf(os.Stdin, "%d %d %d\n", &start, &height, &end)
		if size <= 0 {
			break
		}
		lines = append(lines, line{start, height}, line{end, 0})
	}

	result := mergeSort(lines)
	for i, line := range result {
		fmt.Print(line.Start, " ", line.Height)
		if i == len(result)-1 {
			fmt.Println()
		} else {
			fmt.Print(" ")
		}
	}
}

func mergeSort(items []line) lines {
	var size = len(items)

	if size == 1 {
		return items
	}

	middle := int(size / 2)
	var (
		left  = make([]line, middle)
		right = make([]line, size-middle)
	)

	for i := 0; i < size; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			right[i-middle] = items[i]
		}
	}

	return merge(mergeSort(left), mergeSort(right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func merge(left, right []line) (result lines) {
	result = make(lines, 0)
	heightLeft := 0
	heightRight := 0

	for len(left) > 0 && len(right) > 0 {
		if left[0].Start < right[0].Start {
			heightLeft = left[0].Height
			result = result.Append(line{left[0].Start, max(heightLeft, heightRight)})
			left = left[1:]
		} else {
			heightRight = right[0].Height
			result = result.Append(line{right[0].Start, max(heightLeft, heightRight)})
			right = right[1:]
		}
	}

	for j := 0; j < len(left); j++ {
		result = result.Append(left[j])
	}

	for j := 0; j < len(right); j++ {
		result = result.Append(right[j])
	}

	return
}
