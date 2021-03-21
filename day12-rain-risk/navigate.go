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

func rotateLeft(north, east int64, rotations int) (int64, int64) {
    for i := 0; i < rotations; i++ {
        tmpEast := -north
        north = east
        east = tmpEast
    }

    return north, east
}

func rotateRight(north, east int64, rotations int) (int64, int64) {
    for i := 0; i < rotations; i++ {
        tmpEast := north
        north = -east
        east = tmpEast
    }

    return north, east
}

func main() {
    verbose := len(os.Args) == 2 && os.Args[1] == "--verbose"
    scanner := bufio.NewScanner(os.Stdin)
    pattern := regexp.MustCompile(`^([A-Z])(\d+)$`)

    direction := EAST
    positionNorth := int64(0)
    positionEast := int64(0)

    waypointNorth := int64(1)
    waypointEast := int64(10)
    waypointShipNorth := int64(0)
    waypointShipEast := int64(0)

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
        } else {
            panic("Unexpected action (direction)")
        }

        if verbose {
            fmt.Println("direction:", direction, "position:", positionNorth,
                    "N", positionEast, "E")
        }

        if action == "N" {
            waypointNorth += value
        } else if action == "S" {
            waypointNorth -= value
        } else if action == "E" {
            waypointEast += value
        } else if action == "W" {
            waypointEast -= value
        } else if action == "L" {
            waypointNorth, waypointEast = rotateLeft(waypointNorth,
                    waypointEast, int(value) / 90)
        } else if action == "R" {
            waypointNorth, waypointEast = rotateRight(waypointNorth,
                    waypointEast, int(value) / 90)
        } else if action == "F" {
            waypointShipNorth += waypointNorth * value
            waypointShipEast += waypointEast * value
        } else {
            panic("Unexpected action (waypoint)")
        }

        if verbose {
            fmt.Println("waypoint:", waypointNorth, "N", waypointEast, "E",
                    "position:", waypointShipNorth, "N", waypointShipEast, "E")
        }
    }

    fmt.Println(abs(positionNorth) + abs(positionEast))
    fmt.Println(abs(waypointShipNorth) + abs(waypointShipEast))
}
