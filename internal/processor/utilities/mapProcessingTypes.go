package utilities 

import (
	"guap-statement-parser/config"
	"guap-statement-parser/pkg/utilities"
	"log"
)

func MapProcessingType(config *config.GuapCSVMappingConfig, processingType string, processingMap map[string]string, record []string) []string {
    processingTypeValue := config.ProcessingTypesMapping[processingType]

    if processingType == "" {
        log.Fatalf("Unable to find processing type: %v", processingTypeValue)
    }

	typeIndex := utilities.IndexOf(config.OutputHeaders, processingTypeValue)

	if typeIndex == -1 {
		log.Printf("Unable to find %s index in output headers", processingTypeValue)
		return record 
	}

    mappedValue := processingMap[record[typeIndex]]
    if mappedValue == "" {
        log.Printf("Unable to find mapping for %s: %v", processingTypeValue, record[typeIndex])
    }
    record[typeIndex] = mappedValue

	return record 
}
