package processor

import (
	"guap-statement-parser/config"
	"guap-statement-parser/internal/processor/utilities"
)
 
func AddTransactionType(config *config.GuapCSVMappingConfig, record []string) []string {
    return utilities.MapProcessingType(config, "Type", config.TransactionTypeMapping, record)
}
