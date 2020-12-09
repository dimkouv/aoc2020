package zeronine

import (
	"errors"
	"io"

	zeroone "github.com/dimkouv/AOC2K20/01"
)

func runP1(reader io.Reader, preambleSize int) (int, error) {
	nums, err := zeroone.ReadNums(reader)
	if err != nil {
		return 0, err
	}
	return findInvalidNum(nums, preambleSize)
}

func findInvalidNum(nums []int, preambleSize int) (int, error) {
	for i := preambleSize + 1; i < len(nums); i++ {
		if !hasTwoIntsThatSumUpTo(nums[i-preambleSize:i], nums[i]) {
			return nums[i], nil
		}
	}
	return 0, errors.New("not found")
}

func hasTwoIntsThatSumUpTo(nums []int, target int) bool {
	lookup := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if lookup[nums[i]] {
			return true
		}
		lookup[target-nums[i]] = true
	}
	return false
}

func runP2(reader io.Reader, preambleSize int) (int, error) {
	nums, err := zeroone.ReadNums(reader)
	if err != nil {
		return 0, err
	}

	invalidNum, err := findInvalidNum(nums, preambleSize)
	if err != nil {
		return 0, err
	}

	return findSequence(nums, invalidNum), nil
}

func findSequence(nums []int, target int) int {
	res := make(chan int)

	for i := range nums[:len(nums)-1] {
		if nums[i] >= target {
			continue
		}

		go func(pos0 int) {
			min, max, sum := nums[pos0], nums[pos0], nums[pos0]

			for pos1 := pos0 + 1; ; pos1++ {
				if nums[pos1] < min {
					min = nums[pos1]
				}
				if nums[pos1] > max {
					max = nums[pos1]
				}

				sum += nums[pos1]

				if sum == target {
					res <- min + max
				} else if sum > target {
					return
				}
			}
		}(i)
	}

	num := <-res
	close(res)
	return num
}
