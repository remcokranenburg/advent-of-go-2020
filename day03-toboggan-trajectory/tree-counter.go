package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "strconv"
)

func main() {
    filename := os.Args[1]
    numRight, _ := strconv.ParseInt(os.Args[2], 10, 0)
    numDown, _ := strconv.ParseInt(os.Args[3], 10, 0)

    content, _ := ioutil.ReadFile(filename)

    var field [][]bool
    field = append(field, []bool{})

    for _, value := range(content) {
        if value == '\n' {
            field = append(field, []bool{})
        } else if value == '#' {
            row := len(field) - 1
            field[row] = append(field[row], true)
        } else if value == '.' {
            row := len(field) - 1
            field[row] = append(field[row], false)
        } else {
            fmt.Printf("Warning: unexpected character: %s\n", value)
        }
    }

    numTrees := 0

    for y, x := 0, 0; y < len(field) - 1; y, x = y + int(numDown), x + int(numRight) {
        if field[y][x % len(field[y])] {
            numTrees += 1
        }
    }

    fmt.Println(numTrees)
}
