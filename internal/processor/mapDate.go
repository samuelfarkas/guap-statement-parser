package processor

import (
    "guap-statement-parser/config"
    "guap-statement-parser/pkg/utilities"
    "time"
    "log"
)

func MapDate(config *config.GuapCSVMappingConfig, record []string) []string {
    dateFieldIndex := utilities.IndexOf(config.OutputHeaders, config.ProcessingTypesMapping["Date"]);
    if dateFieldIndex == -1 {
        return record 
    }

    parsedDate, err := time.Parse(config.DateFormat, record[dateFieldIndex])

    if err != nil {
        log.Println("Error parsing date: ", err)
        return record
    }

    record[dateFieldIndex] = parsedDate.Format("01/02/2006")

    return record
}
