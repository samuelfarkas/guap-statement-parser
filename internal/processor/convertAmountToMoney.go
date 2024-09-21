package processor

import (
	"guap-statement-parser/config"
	"guap-statement-parser/pkg/utilities"
)

func ConvertAmountToMoney(config *config.GuapCSVMappingConfig, record []string) []string {
    amountFieldIndex := utilities.IndexOf(config.OutputHeaders, config.ProcessingTypesMapping["Amount"]);
    if amountFieldIndex == -1 {
        return record 
    }

    record[amountFieldIndex] = utilities.ReplaceCommaWithDot(record[amountFieldIndex])

    return record
}
