package onezero

import (
	"fmt"
	"io"
	"sort"

	zeroone "github.com/dimkouv/AOC2K20/01"
)

func runP1(reader io.Reader) (int, error) {
	nums, err := zeroone.ReadNums(reader)
	if err != nil {
		return 0, err
	}

	sort.Ints(nums)
	nums = append([]int{0}, nums...)         // outlet
	nums = append(nums, nums[len(nums)-1]+3) // device's adapter

	diff1, diff3 := 0, 0
	for i := range nums[:len(nums)-1] {
		diff := nums[i+1] - nums[i]
		switch diff {
		case 3:
			diff3++
		case 0, 2:
			break
		case 1:
			diff1++
		default:
			return 0, fmt.Errorf("found an unusable adapter")
		}
	}

	return diff1 * diff3, nil
}

func runP2(reader io.Reader) (int, error) {
	nums, err := zeroone.ReadNums(reader)
	if err != nil {
		return 0, err
	}

	sort.Ints(nums)
	nums = append([]int{0}, nums...)         // outlet
	nums = append(nums, nums[len(nums)-1]+3) // device's adapter

	counter := &counter{counted: make(map[string]int)}
	return counter.count(nums) + 1, nil
}

type counter struct {
	counted map[string]int // cache for computed counts
}

func (c *counter) count(nums []int) int {
	key := fmt.Sprintf("%v", nums)
	if cnt, exists := c.counted[key]; exists {
		return cnt
	}

	cnt := 0
	for i := 0; i < len(nums)-1; i++ {
		for n := 1; i+n+1 < len(nums); n++ {
			diff := nums[i+n+1] - nums[i]
			if diff <= 3 {
				cnt = cnt + 1 + c.count(nums[i+n+1:])
			}
		}
	}

	c.counted[key] = cnt
	return cnt
}
