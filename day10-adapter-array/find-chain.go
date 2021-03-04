package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
    "strconv"
)

func classifyAdapter(a, b int64, diff1, diff2, diff3 *int) {
    if b - a == 3 {
        *diff3 += 1
    } else if b - a == 2 {
        *diff2 += 1
    } else if b - a == 1 {
        *diff1 += 1
    } else {
        panic("Adapters don't fit!")
    }
}

func findArrangements(current, final int64, adapters []int64,
        cache map[int64]int) int {
    if current == final {
        return 1
    } else {
        arrangements, in_cache := cache[current]

        if in_cache {
            return arrangements
        } else {
            arrangements = 0

            for i, next := range adapters {
                if next - current > 0 && next - current <= 3 {
                    remainingAdapters := adapters[i + 1:]
                    arrangements += findArrangements(next, final,
                            remainingAdapters, cache)
                }
            }

            cache[current] = arrangements
            return arrangements
        }
    }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    adapters := []int64{}

    for scanner.Scan() {
        line := scanner.Text()
        joltage, err := strconv.ParseInt(line, 10, 0)

        if err != nil {
            fmt.Printf("Error parsing input line as number: %s\n", line)
            os.Exit(1)
        }

        adapters = append(adapters, joltage)
    }

    sort.Slice(adapters, func(i, j int) bool {
        return adapters[i] < adapters[j]
    })

    diff1, diff2, diff3 := 0, 0, 0

    for i, adapter := range adapters {
        if i == 0 {
            classifyAdapter(0, adapter, &diff1, &diff2, &diff3)
        } else {
            classifyAdapter(adapters[i - 1], adapter, &diff1, &diff2, &diff3)
        }
    }

    diff3 += 1

    fmt.Println("Number of 1-3 joltage differences:", diff1, diff2, diff3)
    fmt.Println("Diff1 * diff3:", diff1 * diff3)

    largestAdapter := adapters[len(adapters) - 1]
    cache := map[int64]int{}
    arrangements := findArrangements(0, largestAdapter, adapters, cache)
    fmt.Println("Number of adapter arrangements:", arrangements)
}

