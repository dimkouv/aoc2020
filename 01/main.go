package zeroone

import (
	"errors"
	"io"
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadNums(reader io.Reader) ([]int, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	s := string(b)
	var nums []int
	for _, line := range strings.Split(s, "\n") {
		if line == "" {
			continue
		}
		num, err := strconv.Atoi(line)
		if err != nil {
			return nil, err
		}
		nums = append(nums, num)
	}
	return nums, nil
}

func runP1(reader io.Reader) (int, error) {
	nums, err := ReadNums(reader)
	if err != nil {
		return 0, err
	}

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == 2020 {
				return nums[i] * nums[j], nil
			}
		}
	}

	return 0, errors.New("not found")
}

func runP2(reader io.Reader) (int, error) {
	nums, err := ReadNums(reader)
	if err != nil {
		return 0, err
	}

	for i := range nums {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] > 2020 {
				continue
			}

			for z := j + 1; z < len(nums); z++ {
				if nums[i]+nums[j]+nums[z] == 2020 {
					return nums[i] * nums[j] * nums[z], nil
				}
			}
		}
	}

	return 0, errors.New("not found")
}
