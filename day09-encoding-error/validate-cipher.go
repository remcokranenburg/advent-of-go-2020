package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "strconv"
)

func isSumOfTwo(number int64, window []int64) bool {
    for i, x := range window {
        for j, y := range window {
            if i < j && x + y == number {
                return true
            }
        }
    }

    return false
}

func sum(arr []int64) int64 {
    result := int64(0)

    for _, x := range arr {
        result += x
    }

    return result
}

func max(arr []int64) int64 {
    result := int64(math.MinInt64)

    for _, x := range arr {
        if x > result {
            result = x
        }
    }

    return result
}

func min(arr []int64) int64 {
    result := int64(math.MaxInt64)

    for _, x := range arr {
        if x < result {
            result = x
        }
    }

    return result
}

func crack(number int64, window []int64) int64 {
    subset := []int64{}

    for _, x := range window {
        subset = append(subset, x)

        for sum(subset) > number {
            subset = subset[1:]
        }

        if sum(subset) == number {
            return min(subset) + max(subset)
        }
    }

    fmt.Println("Error: no weakness found")
    os.Exit(1)
    return 0
}

func main() {
    if len(os.Args) < 2 || len(os.Args) > 3 {
        fmt.Println(
            "Usage: go run validate-cipher.go <length> [--crack] < input.txt")
        os.Exit(1)
    }

    preambleLength, err := strconv.ParseInt(os.Args[1], 10, 0)

    if err != nil {
        fmt.Println("Error parsing command line argument 'length'")
    }

    shouldCrack := len(os.Args) == 3 && os.Args[2] == "--crack"

    scanner := bufio.NewScanner(os.Stdin)
    message := []int64{}

    for scanner.Scan() {
        line := scanner.Text()
        number, err := strconv.ParseInt(line, 10, 0)

        if err != nil {
            fmt.Printf("Error parsing input line as number: %s\n", line)
            os.Exit(1)
        }

        message = append(message, number)
    }

    window := []int64{}
    noTwoNumberSum := int64(0)

    for _, number := range message {
        if len(window) < int(preambleLength) {
            window = append(window, number)
        } else if isSumOfTwo(number, window) {
            window = append(window[1:], number)
        } else {
            noTwoNumberSum = number
            break
        }
    }

    if shouldCrack {
        crackedValue := crack(noTwoNumberSum, message)
        fmt.Println(crackedValue)
    } else {
        fmt.Println(noTwoNumberSum)
    }
}

