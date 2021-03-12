package main

import (
    "bufio"
    "fmt"
    "os"
)

func copySeating(seats [][]rune) [][]rune {
    newSeats := [][]rune{}

    for _, row := range seats {
        newRow := []rune{}

        for _, letter := range row {
            newRow = append(newRow, letter)
        }

        newSeats = append(newSeats, newRow)
    }

    return newSeats
}

func predictSeating(seats [][]rune) [][]rune {
    newSeats := [][]rune{}

    for _, row := range seats {
        newRow := []rune{}

        for _, seat := range row {
            newSeat := 'X'

            if seat == 'L' {
                newSeat = '#'
            } else if seat == '#' {
                newSeat = 'L'
            } else if seat == '.' {
                newSeat = '.'
            }

            newRow = append(newRow, newSeat)
        }

        newSeats = append(newSeats, newRow)
    }

    return newSeats
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    seats := [][]rune{}

    for scanner.Scan() {
        line := scanner.Text()
        row := []rune{}

        for _, letter := range line {
            row = append(row, letter)
        }

        seats = append(seats, row)
    }

    newSeats := predictSeating(seats)

    fmt.Println(seats)
    fmt.Println(newSeats)
}
