package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	nums := make([]int, 0)

	var currentInt int

	for {
		if _, err := fmt.Fscanf(os.Stdin, "%d\n", &currentInt); err != nil {
			if !errors.Is(err, io.EOF) {
				log.Fatalf("Error while reading the input: %v", err)
			}

			break
		}

		nums = append(nums, currentInt)
	}

	N := len(nums)

	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			ni := nums[i]
			nj := nums[j]

			if ni+nj == 2020 {
				log.Printf("Part one answer: %d", ni*nj)
			}
		}
	}

	for i := 0; i < N-2; i++ {
		for j := i + 1; j < N-1; j++ {
			for k := j + 1; k < N; k++ {
				ni := nums[i]
				nj := nums[j]
				nk := nums[k]

				if ni+nj+nk == 2020 {
					log.Printf("Part two answer: %d", ni*nj*nk)
				}
			}
		}
	}
}
