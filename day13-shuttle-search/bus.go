package main

import (
    "bufio"
    "fmt"
    "math"
    "strconv"
    "strings"
    "os"
)

func findNextDeparture(desiredDeparture, busInterval int64) int64 {
    numDepartures := desiredDeparture / busInterval
    return busInterval * numDepartures + busInterval
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    scanner.Scan()
    departureText := scanner.Text()
    desiredDeparture, _ := strconv.ParseInt(departureText, 10, 0)
    scanner.Scan()
    busIntervalsText := scanner.Text()
    busIntervals := strings.Split(busIntervalsText, ",")

    earliestBus := int64(0)
    earliestBusDeparture := int64(math.MaxInt64)

    for _, busId := range busIntervals {
        if busId == "x" {
            continue
        }

        busInterval, _ := strconv.ParseInt(busId, 10, 0)
        nextDeparture := findNextDeparture(desiredDeparture, busInterval)

        if nextDeparture < earliestBusDeparture {
            earliestBus = busInterval
            earliestBusDeparture = nextDeparture
        }
    }

    fmt.Println(earliestBus * (earliestBusDeparture - desiredDeparture))
}
