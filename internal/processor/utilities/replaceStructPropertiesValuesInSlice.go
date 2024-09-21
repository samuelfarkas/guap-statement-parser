package utilities

import "reflect"

func ReplaceStructPropertiesValuesInSlice(properties reflect.Value, indexMap map[string]int, record []string) []string {
    for property, idx := range indexMap {
        value := properties.FieldByName(property).String()

        if properties.FieldByName(property).IsValid() && idx > -1 && idx < len(record) {
            record[idx] = value
        }
    }

    return record 
}
