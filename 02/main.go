package zerotwo

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

type passwordPolicy struct {
	n1     int
	n2     int
	letter string
}

func (policy passwordPolicy) validateV1(password string) error {
	cnt := 0

	for _, c := range password {
		if string(c) == policy.letter {
			cnt++
		}
	}

	if cnt >= policy.n1 && cnt <= policy.n2 {
		return nil
	}
	return fmt.Errorf("cannot validate %s with policy %v", password, policy)
}

func (policy passwordPolicy) validateV2(password string) error {
	if policy.n1 < 1 || policy.n2 < 1 {
		return fmt.Errorf("invalid positions: %v", policy)
	}

	if policy.n1 > len(password) || policy.n1 > len(password) {
		return fmt.Errorf("invalid positions: %v for %s", policy, password)
	}

	r1 := string(password[policy.n1-1]) == policy.letter
	r2 := string(password[policy.n2-1]) == policy.letter
	if r1 && !r2 || r2 && !r1 {
		return nil
	}
	return fmt.Errorf("cannot validate %s with policy %v", password, policy)
}

func parsePasswordPolicy(policy string) (*passwordPolicy, error) {
	parts := strings.SplitN(policy, " ", 2)
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid policy: %v", policy)
	}

	nums, letter := parts[0], parts[1]

	rngParts := strings.SplitN(nums, "-", 2)
	n1, err := strconv.Atoi(rngParts[0])
	if err != nil {
		return nil, err
	}

	n2, err := strconv.Atoi(rngParts[1])
	if err != nil {
		return nil, err
	}

	return &passwordPolicy{
		n1:     n1,
		n2:     n2,
		letter: letter,
	}, nil
}

func run(reader io.Reader, policyVersion int) (int, error) {
	numValid := 0
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		parts := strings.SplitN(scanner.Text(), ": ", 2)
		if len(parts) != 2 {
			return 0, fmt.Errorf("cannot parse: %v", scanner.Text())
		}

		policyStr, password := parts[0], parts[1]
		policy, err := parsePasswordPolicy(policyStr)
		if err != nil {
			return 0, err
		}

		switch {
		case policyVersion == 1 && policy.validateV1(password) == nil:
			numValid++
		case policyVersion == 2 && policy.validateV2(password) == nil:
			numValid++
		}
	}

	return numValid, nil
}
