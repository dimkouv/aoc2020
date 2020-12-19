package onenine

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
)

func parseInput(r io.Reader, part2 bool) (map[int]*regexp.Regexp, chan string) {
	s := bufio.NewScanner(r)
	rules := parseRules(s, part2)
	stringStream := parseStrings(s)
	return rules, stringStream
}

type ruleT struct {
	ruleSets [][]int
	char     rune
}

type rulesT map[int]ruleT

func computeRegExps(rules rulesT, part2 bool) map[int]*regexp.Regexp {
	strRegExp := make(map[int]string)
	for id, rule := range rules {
		if rule.char != 0 {
			strRegExp[id] = "(" + string(rule.char) + ")"
		}
	}

	for len(strRegExp) < len(rules) {
		for id, rule := range rules {
			if _, exists := strRegExp[id]; exists {
				continue
			}
			computable := true
			ruleSetStrRegExps := make([]string, 0)
			for _, ruleSet := range rule.ruleSets {
				regExpParts := make([]string, 0)
				for _, ruleID := range ruleSet {
					rgx, exists := strRegExp[ruleID]
					if !exists {
						computable = false
						break
					}
					regExpParts = append(regExpParts, rgx)
				}
				if !computable {
					break
				}

				if part2 && id == 8 {
					a := regExpParts[0]
					regExpParts[0] = fmt.Sprintf(`((%s)+)`, a)
				} else if part2 && id == 11 {
					a, b := regExpParts[0], regExpParts[1]
					regExpParts = regExpParts[:1]
					regExpParts[0] = ""
					for n := 0; n < 5; n++ {
						regExpParts[0] += fmt.Sprintf(`(((%s){%d})((%s){%d}))|`, a, n+1, b, n+1)
					}
					regExpParts[0] = "(" + strings.TrimRight(regExpParts[0], "|") + ")"
				}
				ruleSetStrRegExps = append(ruleSetStrRegExps, "("+strings.Join(regExpParts, "")+")")
			}
			if computable {
				strRegExp[id] = "(" + strings.Join(ruleSetStrRegExps, "|") + ")"
			}
		}
	}

	compiledRegExps := make(map[int]*regexp.Regexp)
	for i, rgx := range strRegExp {
		compiledRegExps[i] = regexp.MustCompile("^" + rgx + "$")
	}
	return compiledRegExps
}

func parseRules(s *bufio.Scanner, part2 bool) map[int]*regexp.Regexp {
	rules := make(rulesT)
	for s.Scan() {
		if s.Text() == "" {
			return computeRegExps(rules, part2)
		}

		parts := strings.Split(s.Text(), ": ")

		ruleID, _ := strconv.Atoi(parts[0])
		rule := ruleT{ruleSets: make([][]int, 0), char: 0}

		if parts[1][0] == '"' {
			rule.char = rune(parts[1][1])
			rules[ruleID] = rule
		} else {
			rawRuleSets := strings.Split(parts[1], " | ")
			for i := range rawRuleSets {
				rawRuleSet := strings.Split(rawRuleSets[i], " ")
				ruleSet := make([]int, len(rawRuleSet))
				for j := range rawRuleSet {
					ruleSet[j], _ = strconv.Atoi(rawRuleSet[j])
				}
				rule.ruleSets = append(rule.ruleSets, ruleSet)
			}

			rules[ruleID] = rule
		}
	}
	return nil
}

func parseStrings(s *bufio.Scanner) chan string {
	ch := make(chan string)
	go func() {
		for s.Scan() {
			ch <- s.Text()
		}
		close(ch)
	}()
	return ch
}
