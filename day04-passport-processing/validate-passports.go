package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func existsInArray(arr []string, element string) bool {
	for i := range arr {
		if arr[i] == element {
			return true
		}
	}
	return false
}

func validatePassport(passport map[string]string, validateFields bool) bool {
	required := []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
	optional := []string{"cid"}

	// check that the required fields are in the passport
	for _, k := range required {
		_, ok := passport[k]

		if !ok {
			return false
		}
	}

	// check that only required or optional fields are in the passport
	for k := range passport {
		if !existsInArray(required, k) && !existsInArray(optional, k) {
			return false
		}
	}

	if validateFields {
		if !validateBirthYear(passport["byr"]) {
			return false
		}
		if !validateIssueYear(passport["iyr"]) {
			return false
		}
		if !validateExpirationYear(passport["eyr"]) {
			return false
		}
		if !validateHeight(passport["hgt"]) {
			return false
		}
		if !validateHairColor(passport["hcl"]) {
			return false
		}
		if !validateEyeColor(passport["ecl"]) {
			return false
		}
		if !validatePassportId(passport["pid"]) {
			return false
		}
	}

	return true
}

func validatePassportId(s string) bool {
	pattern := regexp.MustCompile(`^[0-9]{9}$`)
	return pattern.Match([]byte(s))
}

func validateEyeColor(s string) bool {
	for _, c := range []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"} {
		if s == c {
			return true
		}
	}

	return false
}

func validateHairColor(s string) bool {
	pattern := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	return pattern.Match([]byte(s))
}

func validateHeight(s string) bool {
	pattern := regexp.MustCompile(`^([0-9]+)(cm|in)$`)

	if !pattern.Match([]byte(s)) {
		return false
	}

	matches := pattern.FindStringSubmatch(s)
	height, _ := strconv.ParseInt(matches[1], 10, 0)

	if matches[2] == "cm" {
		return height >= 150 && height <= 193
	} else if matches[2] == "in" {
		return height >= 59 && height <= 76
	} else {
		panic("huh, it's not cm or in after all? BUG")
	}
}

func validateExpirationYear(s string) bool {
	return validateYear(s, 2020, 2030)
}

func validateIssueYear(s string) bool {
	return validateYear(s, 2010, 2020)
}

func validateBirthYear(s string) bool {
	return validateYear(s, 1920, 2002)
}

func validateYear(s string, min, max int64) bool {
	pattern := regexp.MustCompile(`^[0-9]{4}$`)

	if !pattern.Match([]byte(s)) {
		return false
	}

	year, _ := strconv.ParseInt(s, 10, 0)

	return year >= min && year <= max
}

func main() {
	validateFields := len(os.Args) == 2 && os.Args[1] == "--validate-fields"
	content, _ := ioutil.ReadAll(os.Stdin)
	passportStrings := strings.Split(string(content), "\n\n")
	numValidPassports := 0

	for _, passportString := range passportStrings {
		passport := map[string]string{}
		for _, recordString := range strings.Fields(passportString) {
			record := strings.Split(recordString, ":")
			passport[record[0]] = record[1]
		}

		if validatePassport(passport, validateFields) {
			numValidPassports++
		}
	}

	fmt.Println(numValidPassports)
}