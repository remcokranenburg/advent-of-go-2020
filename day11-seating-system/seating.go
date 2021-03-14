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

func clamp(seats [][]rune, row, seat int) rune {
    if row < 0 || row >= len(seats) {
        return 'X'
    }

    if seat < 0 || seat >= len(seats[row]) {
        return 'X'
    }

    return seats[row][seat]
}

func numNeighbors(seats [][]rune, row, seat int) int {
    result := 0

    for i := row - 1; i <= row + 1; i++ {
        for j := seat - 1; j <= seat + 1; j++ {
            if i != row && j != seat && clamp(seats, i, j) == '#' {
                result++
            }
        }
    }

    return result
}

func numSeated(seats [][]rune) int {
    result := 0

    for _, row := range seats {
        for _, seat := range row {
            if seat == '#' {
                result++
            }
        }
    }

    return result
}

func predictSeating(seats [][]rune) [][]rune {
    newSeats := [][]rune{}

    for i, row := range seats {
        newRow := []rune{}

        for j, seat := range row {
            newSeat := 'X'

            if seat == 'L' && numNeighbors(seats, i, j) == 0 {
                newSeat = '#'
            } else if seat == '#' && numNeighbors(seats, i, j) >= 4 {
                newSeat = 'L'
            } else {
                newSeat = seat
            }

            newRow = append(newRow, newSeat)
        }

        newSeats = append(newSeats, newRow)
    }

    return newSeats
}

func printSeats(seats [][]rune) {
    for _, row := range seats {
        for _, seat := range row {
            fmt.Printf("%c", seat)
        }

        fmt.Println()
    }
}

func seatingIsEqual(a, b [][]rune) bool {
    if len(a) != len(b) {
        return false
    }

    for i, _ := range a {
        if len(a[i]) != len(b[i]) {
            return false
        }

        for j, _ := range a[i] {
            if a[i][j] != b[i][j] {
                return false
            }
        }
    }

    return true
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

    printSeats(seats)
    fmt.Println()

    previousSeats := [][]rune{}

    for i := 0; !seatingIsEqual(previousSeats, seats); i++ {
        previousSeats = seats
        seats = predictSeating(seats)
        printSeats(seats)
        fmt.Println()
    }

    fmt.Println(numSeated(seats))
}
