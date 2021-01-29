package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func validatePassword(min, max int64, letter rune, password string) bool {
    numLetters := int64(0)

    for _, x := range password {
        if x == letter {
            numLetters += 1
        }
    }

    return min <= numLetters && max >= numLetters
}

func main() {
    numValidPasswords := 0
    scanner := bufio.NewScanner(os.Stdin)
    pattern := regexp.MustCompile(`([0-9]+)-([0-9]+) ([a-zA-Z0-9]): ([a-zA-Z0-9]+)`)

    for scanner.Scan() {
        line := scanner.Text()
        matches := pattern.FindStringSubmatch(line)
        var minText = matches[1]
        var maxText = matches[2]
        var letter = []rune(matches[3])[0]
        var password = matches[4]
        min, _ := strconv.ParseInt(minText, 10, 0)
        max, _ := strconv.ParseInt(maxText, 10, 0)

        if validatePassword(min, max, letter, password) {
            numValidPasswords += 1
        }
    }

    fmt.Println(numValidPasswords)
}
