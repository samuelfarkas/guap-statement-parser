package processor

import (
	"guap-statement-parser/config"
	"guap-statement-parser/internal/processor/utilities"
)
 
func AddPaymentMethod(config *config.GuapCSVMappingConfig, record []string) []string {
    return utilities.MapProcessingType(config, "Payment Method", config.PaymentMethodMapping, record)
}

