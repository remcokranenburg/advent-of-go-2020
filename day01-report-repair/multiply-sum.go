package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
)

func main() {
    var numbers []int64

    scanner := bufio.NewScanner(os.Stdin)

    for scanner.Scan() {
        numberText := scanner.Text()
    number, _ := strconv.ParseInt(numberText, 10, 0)
    numbers = append(numbers, number)
    }

    for i, number1 := range numbers {
        for j, number2 := range numbers {
        if i < j && number1 + number2 == 2020 {
                fmt.Println(number1 * number2)
        }
    }
    }
}
