package main

import (
    "fmt"
    "bufio"
    "os"
    "regexp"
    "strconv"
    "strings"
)

func containsColor(rules map[string]map[string]int64, containerColor,
        bagColor string) bool {
    _, containsDirectly := rules[containerColor][bagColor]

    if containsDirectly {
        return true
    }

    for subContainerColor, _ := range rules[containerColor] {
        if containsColor(rules, subContainerColor, bagColor) {
            return true
        }
    }

    return false
}

func numBagsContained(rules map[string]map[string]int64, bagColor string) int64 {
    numBags := int64(0)

    for subContainer, frequency := range rules[bagColor] {
        numBags += frequency
        numBags += frequency * numBagsContained(rules, subContainer)
    }

    return numBags
}

func main() {
    if len(os.Args) != 3 && len(os.Args) != 4 {
        fmt.Println("Usage: go run bag-containers.go [--find-contained] " +
                "COLOR TEXT < rules.txt")
        return
    }

    findContained := false
    bagColor := ""

    if len(os.Args) == 3 {
        bagColor = os.Args[1] + " " + os.Args[2]
    } else {
        if os.Args[1] == "--find-contained" {
            findContained = true
        }
        bagColor = os.Args[2] + " " + os.Args[3]
    }

    scanner := bufio.NewScanner(os.Stdin)

    rules := map[string]map[string]int64{}

    containerPattern := regexp.MustCompile(`^\w+ \w+`)
    elementsPattern := regexp.MustCompile(`(\d+ \w+ \w+ bags?(, |\.))+$`)
    colorFrequencyPattern := regexp.MustCompile(`^(\d+) (\w+ \w+)`)

    for scanner.Scan() {
        ruleText := scanner.Text()
        container := containerPattern.FindString(ruleText)

        rules[container] = map[string]int64{}

        elementsString := elementsPattern.FindString(ruleText)
        elements := strings.Split(elementsString, ", ")

        for _, element := range elements {
            parts := colorFrequencyPattern.FindStringSubmatch(element)

            if len(parts) == 3 {
                frequency, _ := strconv.ParseInt(parts[1], 10, 0)
                color := parts[2]
                rules[container][color] = frequency
            }
        }
    }

    if findContained {
        fmt.Println(numBagsContained(rules, bagColor))
    } else {
        numContainers := 0

        for containerColor, _ := range rules {
            if containsColor(rules, containerColor, bagColor) {
                numContainers++
            }
        }

        fmt.Println(numContainers)
    }
}
