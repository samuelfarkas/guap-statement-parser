package processor

import (
	"guap-statement-parser/config"
	"guap-statement-parser/internal/processor/utilities"
	"reflect"
)

func AddAccountInfo(config *config.GuapCSVMappingConfig, accountInfo config.AccountInfo, record []string) []string {
    accountProperties := reflect.ValueOf(accountInfo)
    accountInfoIndexMap := utilities.MapOutputIndexesToStructProperties(accountProperties, config.AccountInfoMapping, config.OutputHeaders) 

    return utilities.ReplaceStructPropertiesValuesInSlice(accountProperties, accountInfoIndexMap, record)
}
