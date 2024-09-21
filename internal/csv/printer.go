package csv 

import (
    "fmt"
)

func PrintCSV(data [][]string) {
    for idx, record := range data {
        fmt.Printf("Row %d: %v\n", idx + 1, record)
    }
}
