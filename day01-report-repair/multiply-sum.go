package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func scanNumbers(scanner *bufio.Scanner) []int64 {
    var numbers []int64

    for scanner.Scan() {
        numberText := scanner.Text()

        if numberText != "" {
            number, _ := strconv.ParseInt(numberText, 10, 0)
            numbers = append(numbers, number)
        }
    }

    return numbers
}

func buildRange(numIndices int) []int {
    var indices []int

    for i := 0; i < numIndices; i++ {
        indices = append(indices, i)
    }

    return indices
}

func buildGroups(indices []int, arity int) [][]int {
    var groups [][]int

    for i, index := range indices {
        if arity == 1 {
            groups = append(groups, []int{index})
        } else if arity > 1 {
            for _, subGroup := range buildGroups(indices[i:], arity - 1) {
                var group []int = []int{index}

                for _, element := range subGroup {
                    group = append(group, element)
                }

                groups = append(groups, group)
            }
        }
    }

    return groups
}

func findSums(numbers []int64, groups [][]int, expectedSum int64) [][]int {
    var sumGroups [][]int

    for _, group := range groups {
        sum := int64(0)

        for _, element := range group {
            sum += numbers[element]
        }

        if sum == expectedSum {
            sumGroups = append(sumGroups, group)
        }
    }

    return sumGroups
}

func product(numbers []int64, set []int) int64 {
    product := int64(1)

    for i := range set {
        product *= numbers[set[i]]
    }

    return product
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    numbers := scanNumbers(scanner)

    indices := buildRange(len(numbers))

    groupsOfTwo := buildGroups(indices, 2)
    sumsOfTwo := findSums(numbers, groupsOfTwo, 2020)

    if len(sumsOfTwo) == 0 {
        fmt.Println("Error: found no set of two")
    } else if len(sumsOfTwo) > 1 {
        fmt.Println("Warning: found multiple sets of two")
    }

    for _, set := range sumsOfTwo {
        fmt.Printf(
                "%d + %d == %d && %[1]d * %d == %[4]d\n",
                numbers[set[0]],
                numbers[set[1]],
                2020,
                product(numbers, set))
    }

    groupsOfThree := buildGroups(indices, 3)
    sumsOfThree := findSums(numbers, groupsOfThree, 2020)

    if len(sumsOfThree) == 0 {
        fmt.Println("Error: found no set of three")
    } else if len(sumsOfThree) > 1 {
        fmt.Println("Warning: found multiple sets of three")
    }

    for _, set := range sumsOfThree {
        fmt.Printf(
            "%d + %d + %d == %d && %[1]d * %d * %d == %[5]d\n",
            numbers[set[0]],
            numbers[set[1]],
            numbers[set[2]],
            2020,
            product(numbers, set))
    }
}
