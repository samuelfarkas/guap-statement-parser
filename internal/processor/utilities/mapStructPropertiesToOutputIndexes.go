package utilities

import (
	"guap-statement-parser/pkg/utilities"
	"reflect"
)

func MapOutputIndexesToStructProperties(properties reflect.Value, mapping map[string]string, outputHeaders []string) map[string]int {
    indexMap := make(map[string]int)

    for idx :=0; idx < properties.NumField(); idx++ {
        propertyOutputName := mapping[properties.Type().Field(idx).Name]
        propertyIndex := utilities.IndexOf(outputHeaders, propertyOutputName)
        if propertyIndex != -1 {
            indexMap[properties.Type().Field(idx).Name] = propertyIndex
        }
    }

    return indexMap 
}
