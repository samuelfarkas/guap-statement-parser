package csv 

import (
    "encoding/csv"
    "os"
    "log"
)

func WriteCSV(filePath string, data [][]string) error {
    file, err := os.Create(filePath)
    if err != nil {
        log.Fatalf("Unable to create file: %v", err)
        return err
    }

    defer file.Close()

    writer := csv.NewWriter(file)
    defer writer.Flush()

    for _, record := range data {
        err := writer.Write(record)
        if err != nil {
            log.Printf("Error writing record: %v", err)
        }
    }

    return nil
}
