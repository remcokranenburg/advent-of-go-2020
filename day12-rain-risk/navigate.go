package main

import (
    "bufio"
    "fmt"
    "regexp"
    "strconv"
    "os"
)

var (
    EAST = int64(0)
    SOUTH = int64(90)
    WEST = int64(180)
    NORTH = int64(270)
)

func mod(x, modulo int64) int64 {
    result := x

    for result < 0 || result >= modulo {
        if result < 0 {
            result += modulo
        } else if result >= modulo {
            result -= modulo
        }
    }

    return result
}

func abs(x int64) int64 {
    if x < 0 {
        return -x
    }

    return x
}

func main() {
    verbose := len(os.Args) == 2 && os.Args[1] == "--verbose"
    scanner := bufio.NewScanner(os.Stdin)
    pattern := regexp.MustCompile(`^([A-Z])(\d+)$`)
    direction := EAST
    positionNorth := int64(0)
    positionEast := int64(0)

    for scanner.Scan() {
        line := scanner.Text()
        matches := pattern.FindStringSubmatch(line)
        action := matches[1]
        value, _ := strconv.ParseInt(matches[2], 10, 0)

        if action == "N" || (action == "F" && direction == NORTH) {
            positionNorth += value
        } else if action == "S" || (action == "F" && direction == SOUTH) {
            positionNorth -= value
        } else if action == "E" || (action == "F" && direction == EAST) {
            positionEast += value
        } else if action == "W" || (action == "F" && direction == WEST) {
            positionEast -= value
        } else if action == "L" {
            direction = mod(direction - value, int64(360))
        } else if action == "R" {
            direction = mod(direction + value, int64(360))
        }

        if verbose {
            fmt.Println("direction:", direction)
            fmt.Println("position north:", positionNorth)
            fmt.Println("position east:", positionEast)
        }
    }

    fmt.Println(abs(positionNorth) + abs(positionEast))
}
