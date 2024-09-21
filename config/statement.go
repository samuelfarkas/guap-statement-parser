package config

import (
	"encoding/json"
	"log"
	"os"
)

type GuapCSVMappingConfig struct {
    Delimiter     string            `json:"delimiter"`
	InputHeaders  []string          `json:"inputHeaders"`
	OutputHeaders []string          `json:"outputHeaders"`
	Mapping       map[string]string `json:"mapping"`
    AccountInfoMapping map[string]string `json:"accountInfoMapping"`
    TransactionTypeMapping map[string]string `json:"transactionTypeMapping"`
    ProcessingTypesMapping map[string]string `json:"processingTypesMapping"`
    PaymentMethodMapping map[string]string `json:"paymentMethodMapping"`
    SkipRows      int               `json:"skipRows"`
    DateFormat   string            `json:"dateFormat"`
}

type AccountInfo struct {
	AccountName string
}

type ProcessingTypes struct {
    Type string
}

func LoadGuapCSVMappingConfig(path string) (*GuapCSVMappingConfig, error) {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Unable to read configuration file: %v", err)
		return nil, err
	}

	config := &GuapCSVMappingConfig{}
	err = json.Unmarshal(file, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}
