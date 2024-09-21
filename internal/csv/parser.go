package csv

import (
    "encoding/csv"
    "os"
    "log"
    "guap-statement-parser/config"
)


func ParseStatement(path string, config *config.GuapCSVMappingConfig) ([][]string, error) {
    file, err := os.Open(path)

    if err != nil {
        log.Fatalf("Unable to open file: %v", err)
        return nil, err
    }
    defer file.Close()

    reader := csv.NewReader(file)
    if config.Delimiter == "" || len(config.Delimiter) > 1 {
        log.Fatalf("Delimiter is required and must be a single character.")
    }
    reader.Comma = rune(config.Delimiter[0])

    records := [][]string{}

    if config.SkipRows > 0 {
        for i := 0; i < config.SkipRows; i++ {
            _, err := reader.Read()
            if err != nil {
                log.Printf("Error skipping row %d: %v", i+1, err)
                continue
            }
        }
    }

    reader.FieldsPerRecord = -1; // reader sets this to the number of fields in the first record
    header, err := reader.Read()
    if err != nil {
        log.Fatalf("Unable to read header: %v", err)
        return nil, err
    }
    // set the number of fields per record to the number of fields in the header
    reader.FieldsPerRecord = len(header)

    for {
        record, err := reader.Read()
        if err != nil {
            if err.Error() == "EOF" {
                break // End of file, stop reading
            }
            log.Printf("Error reading line: %v. Skipping...\n", err)
            continue 
        }
        records = append(records, record)
    }

    data := records[1:]
    outputCSV := [][]string{config.OutputHeaders}

    headerIndexMap := make(map[string]int)
    for idx, header := range header {
        headerIndexMap[header] = idx
    }


    for _, row := range data {
        outputRow := make([]string, len(config.OutputHeaders))
        for i, outputHeader := range config.OutputHeaders {
            inputHeader := config.Mapping[outputHeader]
            if idx, ok := headerIndexMap[inputHeader]; ok {
                outputRow[i] = row[idx]
            } else {
                outputRow[i] = ""
            }
        }
        outputCSV = append(outputCSV, outputRow)
    }

    return outputCSV, nil
}
