package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func validatePasswordNewPolicy(indices map[int64]struct{}, letter rune, password string) bool {
    numLetters := 0

    for i, x := range password {
        _, containsIndex := indices[int64(i)]

        if containsIndex && x == letter {
            numLetters += 1
        }
    }

    return numLetters == 1
}

func main() {
    numValidPasswords := 0
    scanner := bufio.NewScanner(os.Stdin)
    pattern := regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

    for scanner.Scan() {
        line := scanner.Text()
        matches := pattern.FindStringSubmatch(line)
        first, _ := strconv.ParseInt(matches[1], 10, 0)
        second, _ := strconv.ParseInt(matches[2], 10, 0)
        indices := map[int64]struct{}{first - 1: {}, second - 1: {}}
        letter := []rune(matches[3])[0]
        password := matches[4]

        if validatePasswordNewPolicy(indices, letter, password) {
            numValidPasswords += 1
        }
    }

    fmt.Println(numValidPasswords)
}
