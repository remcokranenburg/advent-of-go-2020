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

func findInDirection(seats [][]rune, row, seat, down, right int) rune {

    for i, j := row + down, seat + right; i >= 0 && i < len(seats) && j >= 0 && j < len(seats[i]); i, j = i + down, j + right {
        s := seats[i][j]

        if s != '.' {
            return s
        }
    }

    return 'X'
}

func numNeighbors(seats [][]rune, row, seat int, lineOfSight bool) int {
    result := 0

    if lineOfSight {
        for i := -1; i <= 1; i++ {
            for j := -1; j <= 1; j++ {
                if !(i == j && j == 0 ) && findInDirection(seats, row, seat, i, j) == '#' {
                    result++
                }
            }
        }
    } else {
        for i := row - 1; i <= row + 1; i++ {
            for j := seat - 1; j <= seat + 1; j++ {
                if !(i == row && j == seat) && clamp(seats, i, j) == '#' {
                    result++
                }
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

func predictSeating(seats [][]rune, maxNeighbors int,
        lineOfSight bool) [][]rune {
    newSeats := [][]rune{}

    for i, row := range seats {
        newRow := []rune{}

        for j, seat := range row {
            newSeat := 'X'

            if seat == 'L' && numNeighbors(seats, i, j, lineOfSight) == 0 {
                newSeat = '#'
            } else if seat == '#' && numNeighbors(seats, i, j, lineOfSight) > maxNeighbors {
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

func printNeighbors(seats [][]rune, lineOfSight bool) {
    for i, row := range seats {
        for j, _ := range row {
            fmt.Printf("%d", numNeighbors(seats, i, j, lineOfSight))
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
    verbose := len(os.Args) == 2 && os.Args[1] == "--verbose"
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

    seatsLineOfSight := seats

    if verbose {
        fmt.Println("Predict seating:")
        fmt.Println()
        printSeats(seats)
        printNeighbors(seats, false)
        fmt.Println()
    }

    previousSeats := [][]rune{}
    previousSeatsLineOfSight := [][]rune{}

    for i := 0; !seatingIsEqual(previousSeats, seats); i++ {
        previousSeats = seats
        seats = predictSeating(seats, 3, false)

        if verbose {
            printSeats(seats)
            printNeighbors(seats, false)
            fmt.Println()
        }
    }

    if verbose {
        fmt.Println("Predict seating (line of sight):")
        fmt.Println()
        printSeats(seatsLineOfSight)
        printNeighbors(seatsLineOfSight, false)
        fmt.Println()
    }

    for i := 0; !seatingIsEqual(previousSeatsLineOfSight, seatsLineOfSight); i++ {
        previousSeatsLineOfSight = seatsLineOfSight
        seatsLineOfSight = predictSeating(seatsLineOfSight, 4, true)

        if verbose {
            printSeats(seatsLineOfSight)
            printNeighbors(seatsLineOfSight, true)
            fmt.Println()
        }
    }

    fmt.Println("Number of occupied seats:", numSeated(seats))
    fmt.Println("Number of occupied seats (line of sight):",
            numSeated(seatsLineOfSight))
}
