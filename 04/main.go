package zerofour

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
)

type passport struct {
	BirthYear      string `json:"byr" validate:"required"`
	IssueYear      string `json:"iyr" validate:"required"`
	ExpirationYear string `json:"eyr" validate:"required"`
	Height         string `json:"hgt" validate:"required"`
	HairColor      string `json:"hcl" validate:"required"`
	EyeColor       string `json:"ecl" validate:"required"`
	PassportID     string `json:"pid" validate:"required"`
	CountryID      string `json:"cid"`
}

func (p *passport) validateV1() error {
	return validator.New().Struct(*p)
}

func (p *passport) validateV2() error {
	if err := validator.New().Struct(*p); err != nil {
		return err
	}

	if num, err := strconv.Atoi(p.BirthYear); err != nil || num < 1920 || num > 2002 {
		return fmt.Errorf("invalid byr: %s", p.BirthYear)
	}

	if num, err := strconv.Atoi(p.IssueYear); err != nil || num < 2010 || num > 2020 {
		return fmt.Errorf("invalid iyr: %s", p.IssueYear)
	}

	if num, err := strconv.Atoi(p.ExpirationYear); err != nil || num < 2020 || num > 2030 {
		return fmt.Errorf("invalid eyr: %s", p.ExpirationYear)
	}

	switch {
	case strings.HasSuffix(p.Height, "cm"):
		if num, err := strconv.Atoi(strings.TrimRight(p.Height, "cm")); err != nil || num < 150 || num > 193 {
			return fmt.Errorf("invalid hgt: %s", p.Height)
		}
	case strings.HasSuffix(p.Height, "in"):
		if num, err := strconv.Atoi(strings.TrimRight(p.Height, "in")); err != nil || num < 59 || num > 76 {
			return fmt.Errorf("invalid hgt: %s", p.Height)
		}
	default:
		return fmt.Errorf("invalid hgt: %s", p.Height)
	}

	if match, err := regexp.MatchString(`^#[0-9a-f]{6}$`, p.HairColor); err != nil || !match {
		return fmt.Errorf("invalid hcl: %s", p.HairColor)
	}

	if match, err := regexp.MatchString(`^(amb|blu|brn|gry|grn|hzl|oth)$`, p.EyeColor); err != nil || !match {
		return fmt.Errorf("invalid ecl: %s", p.EyeColor)
	}

	if match, err := regexp.MatchString(`^[0-9]{9}$`, p.PassportID); err != nil || !match {
		return fmt.Errorf("invalid pid: %s", p.PassportID)
	}

	return nil
}

func passports(reader io.Reader) chan passport {
	ch := make(chan passport)

	go func() {
		parseAndSend := func(fields map[string]string) {
			b, _ := json.Marshal(fields)
			p := passport{}
			if err := json.Unmarshal(b, &p); err != nil {
				panic(err)
			}
			ch <- p
		}

		scanner := bufio.NewScanner(reader)
		fieldsMap := make(map[string]string)

		for i := 0; scanner.Scan(); i++ {
			if scanner.Text() == "" {
				parseAndSend(fieldsMap)
				fieldsMap = make(map[string]string)
			} else {
				keyValList := strings.Split(scanner.Text(), " ")
				for _, keyVal := range keyValList {
					parts := strings.Split(keyVal, ":")
					fieldsMap[parts[0]] = parts[1]
				}
			}
		}

		if len(fieldsMap) > 0 {
			parseAndSend(fieldsMap)
		}

		close(ch)
	}()

	return ch
}

func runP1(reader io.Reader) int {
	numValid := 0
	for passport := range passports(reader) {
		if err := passport.validateV1(); err == nil {
			numValid++
		}
	}
	return numValid
}

func runP2(reader io.Reader) int {
	numValid := 0
	i := 0
	for passport := range passports(reader) {
		if passport.validateV2() == nil {
			numValid++
		}

		i++
	}
	return numValid
}
