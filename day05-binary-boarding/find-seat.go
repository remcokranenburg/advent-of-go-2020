package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
)

func findSeat(s []rune, start int) int {
    if len(s) == 0 {
        return start
    } else {
        newStart := start

        if s[0] == 'B' || s[0] == 'R' {
            newStart = start + int(math.Pow(2, float64(len(s) - 1)))
        }

        return findSeat(s[1:], newStart)
    }
}

func main() {
    findMissing := len(os.Args) == 2 && os.Args[1] == "--find-missing"
    scanner := bufio.NewScanner(os.Stdin)
    highestId := 0
    lowestId := 1024
    missingId := -1
    filledSeats := make([]bool, 1024)

    for scanner.Scan() {
        line := scanner.Text()
        letters := []rune(line)
        row := findSeat(letters[:7], 0)
        column := findSeat(letters[7:], 0)
        id := row * 8 + column

        if id < lowestId {
            lowestId = id
        }

        if id > highestId {
            highestId = id
        }

        filledSeats[id] = true
    }

    for id, filled := range filledSeats {
        if id > lowestId && id < highestId && !filled {
            missingId = id
            break
        }
    }

    if findMissing {
        fmt.Println(missingId)
    } else {
        fmt.Println(highestId)
    }
}
