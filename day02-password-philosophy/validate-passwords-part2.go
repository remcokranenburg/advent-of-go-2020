package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func validatePassword(indices map[int64]struct{}, letter rune, password string) bool {
    var numLetters int = 0

    for i, x := range(password) {
        _, containsIndex := indices[int64(i)]

        if containsIndex && x == letter {
            numLetters += 1
        }
    }

    return numLetters == 1
}

func main() {
    var numValidPasswords int = 0
    var scanner = bufio.NewScanner(os.Stdin)
    var pattern = regexp.MustCompile(`(\d+)-(\d+) (\w): (\w+)`)

    for scanner.Scan() {
        line := scanner.Text()
        matches := pattern.FindStringSubmatch(line)
        first, _ := strconv.ParseInt(matches[1], 10, 0)
        second, _ := strconv.ParseInt(matches[2], 10, 0)
        var indices = map[int64]struct{} {
            (first - 1): struct{}{},
            (second - 1): struct{}{},
        }
        var letter = []rune(matches[3])[0]
        var password = matches[4]

        if validatePassword(indices, letter, password) {
            numValidPasswords += 1
        }
    }

    fmt.Println(numValidPasswords)
}
