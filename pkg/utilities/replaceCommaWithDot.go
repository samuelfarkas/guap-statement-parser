package utilities 

import "strings"

func ReplaceCommaWithDot(numberStr string) string {
    return strings.Replace(numberStr, ",", ".", 1) 
}

