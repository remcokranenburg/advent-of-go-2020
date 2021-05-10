package main

import (
    "bufio"
    "fmt"
    "math"
    "os"
    "regexp"
    "strconv"
)

func powUint64(x, power uint64) uint64 {
    if power == 0 {
        return 1
    }

    result := x

    for i := uint64(2); i <= power; i++ {
        result *= x
    }

    return result
}

func parseMask(mask string) (uint64, uint64) {
    orMask := uint64(0)
    andMask := uint64(math.MaxUint64)

    for i, letter := range mask {
        if letter == '1' {
            orMask += powUint64(2, uint64(35 - i))
        } else if letter == '0' {
            andMask -= powUint64(2, uint64(35 - i))
        }
    }
    return orMask, andMask
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    maskPattern := regexp.MustCompile(`^mask = ([X10]+)$`)
    memPattern := regexp.MustCompile(`^mem\[([0-9]+)\] = ([0-9]+)$`)
    currentMask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
    currentOrMask := uint64(0)
    currentAndMask := uint64(math.MaxUint64)
    memory := map[uint64]uint64{}

    lineNumber := 1

    for scanner.Scan() {
        line := scanner.Text()

        memMatches := memPattern.FindStringSubmatch(line)

        if len(memMatches) > 2 {
            address, _ := strconv.ParseUint(memMatches[1], 10, 0)
            value, _ := strconv.ParseUint(memMatches[2], 10, 0)
            maskedValue := value & currentAndMask | currentOrMask

            fmt.Printf("value:  %036b (decimal %d)\n", value, value)
            fmt.Printf("mask:   %s\n", currentMask)
            fmt.Printf("|mask:  %036b\n", currentOrMask)
            fmt.Printf("&mask:  %036b\n", currentAndMask)
            fmt.Printf("result: %036b (decimal %d)\n", maskedValue, maskedValue)
            fmt.Println()

            memory[address] = maskedValue

        } else {
            maskMatches := maskPattern.FindStringSubmatch(line)

            if len(maskMatches) > 1 {
                currentMask = maskMatches[1]
                currentOrMask, currentAndMask = parseMask(currentMask)
                fmt.Printf("mask:   %s\n\n", currentMask)
            } else {
                fmt.Println("ERROR: line", lineNumber,
                    "did not match command:", line)
                os.Exit(1)
            }
        }

        lineNumber += 1
    }

    sum := uint64(0)

    for address, value := range memory {
        fmt.Printf("[%036b]: %036b (%d)\n", address, value, value)
        sum += value
    }

    fmt.Println(sum)
}
