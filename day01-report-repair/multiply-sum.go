package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func scan_numbers(scanner: int) -> []int64 {
    var numbers []int64

    for scanner.Scan() {
        numberText := scanner.Text()
        number, _ := strconv.ParseInt(numberText, 10, 0)
        numbers = append(numbers, number)
    }

    return numbers
}

func find_two_sum(numbers: []int64, expected_sum: int) -> [][]int64 {
    var sets = [][]int64

    for i, number1 := range numbers {
        for j, number2 := range numbers {
            if i < j && number1 + number2 == expected_sum {
                append(sets, [i, j])
            }
        }
    }

    return sets
}

func find_three_sum() -> []int64 {
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    numbers := scan_numbers(scanner)
    sets := find_two_sum(numbers, 2020)

    if len(sets) == 0 {
        fmt.Println("Error: found no answer")
        // TODO: exit 1
    else if len(sets) > 1 {
        fmt.Println("Warning: found multiple answers")
    }

    for i, set := range sets {
        fmt.Println(numbers[set[0]] * numbers[set[1]])
    }
}
