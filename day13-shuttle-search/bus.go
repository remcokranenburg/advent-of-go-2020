package main

import (
    "bufio"
    "fmt"
    "math"
    "strconv"
    "strings"
    "os"
)

// Insight: The interval is at least the product of all primes
// Note: the offset might have effect

func findNextDeparture(desiredDeparture, busInterval int64) int64 {
    if desiredDeparture % busInterval == 0 {
        return desiredDeparture
    } else {
        numDepartures := desiredDeparture / busInterval
        return busInterval * numDepartures + busInterval
    }
}

func validateDepartures(departures []int64, desiredDiffs []int64) bool {
    for i, departureTime := range departures {
        if departureTime != departures[0] + desiredDiffs[i] {
            return false
        }
    }

    return true
}

func max(l []int64) int64 {
    result := int64(math.MinInt64)
    for _, x := range l {
        if x > result {
            result = x
        }
    }
    return result
}

func sum(l []int64) int64 {
    result := int64(0)
    for _, x := range l {
        result += x
    }
    return result
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    departureText := scanner.Text()
    desiredDeparture, _ := strconv.ParseInt(departureText, 10, 0)
    scanner.Scan()
    busIntervalsText := scanner.Text()
    busIntervals := strings.Split(busIntervalsText, ",")

    busIds := []int64{}
    desiredDiffs := []int64{}

    for i, busIdText := range busIntervals {
        if busIdText == "x" {
            continue
        }

        busId, _ := strconv.ParseInt(busIdText, 10, 0)
        busIds = append(busIds, busId)
        desiredDiffs = append(desiredDiffs, int64(i))
    }

    earliestBus := int64(0)
    earliestBusDeparture := int64(math.MaxInt64)

    for _, busId := range busIds {
        nextDeparture := findNextDeparture(desiredDeparture, busId)

        if nextDeparture < earliestBusDeparture {
            earliestBus = busId
            earliestBusDeparture = nextDeparture
        }
    }

    fmt.Println(earliestBus * (earliestBusDeparture - desiredDeparture))

    t := int64(0)

    for {
        departures := []int64 {}

        firstDeparture := findNextDeparture(t, busIds[0])

        for i, busId := range busIds {
            departures = append(departures,
                    findNextDeparture(firstDeparture + desiredDiffs[i], busId))
        }

        if validateDepartures(departures, desiredDiffs) {
            fmt.Println(firstDeparture)
            break
        } else {
            fmt.Println("t:", t)
            fmt.Println("firstDeparture:", firstDeparture)
            fmt.Println("departures:", departures)
            fmt.Println("desiredDiffs:", desiredDiffs)
            fmt.Println("validateDepartures:", validateDepartures(departures, desiredDiffs))
            newT := max(departures) - desiredDiffs[len(desiredDiffs) - 1]
            if newT > t {
                t = newT
            } else {
                t += 1
            }
        }

        // if t > 3417 {
        //     break
        // }
    }
}
