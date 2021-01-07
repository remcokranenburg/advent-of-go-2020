package main

import (
    "bufio"
    "fmt"
    "os"
    "regexp"
    "strconv"
)

func validatePassword(min, max int64, letter rune, password string) bool {
    var numLetters int64 = 0

    for _, x := range(password) {
        if x == letter {
            numLetters += 1
        }
    }

    return min <= numLetters && max >= numLetters
}

func main() {
    var numValidPasswords int = 0
    var scanner = bufio.NewScanner(os.Stdin)
    var pattern = regexp.MustCompile(`([0-9]+)-([0-9]+) ([a-zA-Z0-9]): ([a-zA-Z0-9]+)`)

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
