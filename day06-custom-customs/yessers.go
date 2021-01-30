package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strings"
)

func personHasLetter(letter rune, person []rune) bool {
    for _, answer := range person {
        if answer == letter {
            return true
        }
    }

    return false
}

func allHaveLetter(letter rune, people []string) bool {
    for _, person := range people {
        if !personHasLetter(letter, []rune(person)) {
            return false
        }
    }

    return true
}

func main() {
    all := len(os.Args) == 2 && os.Args[1] == "--all"
    content, _ := ioutil.ReadAll(os.Stdin)
    groups := strings.Split(strings.TrimSpace(string(content)), "\n\n")

    totalYessed := 0

    for _, group := range groups {

        groupYessed := map[rune]bool{}
        people := strings.Split(group, "\n")

        if all {
            for i := 0; i < 26; i++ {
                letter := rune('a' + i)

                if allHaveLetter(letter, people) {
                    groupYessed[letter] = true
                }
            }
        } else {
            for _, person := range people {
                for _, answer := range []rune(person) {
                    groupYessed[answer] = true
                }
            }
        }

        totalYessed += len(groupYessed)
    }

    fmt.Println(totalYessed)
}
