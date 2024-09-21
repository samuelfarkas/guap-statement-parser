package main

import (
	"flag"
	"guap-statement-parser/config"
	"guap-statement-parser/internal/csv"
	"guap-statement-parser/internal/processor"
	"log"
	"os"
    "path/filepath"
)


func main() {
    csvFilePath := flag.String("statement-csv", "", "Path to the CSV file to parse")
    statementFromBank := flag.String("statement-from", "", "Name of the bank statement is from")
    accountName := flag.String("account-name", "Default", "Name of the account")
    skipHeader := flag.Bool("skip-header", false, "Skip the header row")
    flag.Parse()

    accountInfo := config.AccountInfo{
        AccountName: *accountName,
    }

    if  *csvFilePath == "" || *statementFromBank == "" {
        log.Fatalf("Both --csv and --statement-from flags are required.")
    }

    var statementConfigPath string 

    switch *statementFromBank {
        case "fio":
            statementConfigPath = "data/fio.json" 
        case "revolut":
            statementConfigPath = "data/revolut.json"
    }

    mappingConfig, err := config.LoadGuapCSVMappingConfig(statementConfigPath)
    if err != nil {
        log.Fatalf("Unable to load mapping configuration: %v", err)
        return
    }


    parsedData, err := csv.ParseStatement(*csvFilePath, mappingConfig)

    if err != nil {
        log.Fatalf("Unable to parse statement: %v", err)
        return
    }

    if parsedData != nil {
        for _, record := range parsedData[1:] {
            processor.AddAccountInfo(mappingConfig, accountInfo, record)
            processor.AddTransactionType(mappingConfig, record)
            processor.AddPaymentMethod(mappingConfig, record)
            processor.ConvertAmountToMoney(mappingConfig, record)
            processor.MapDate(mappingConfig, record)
        }
        pwd, err := os.Getwd()
        if err != nil {
            log.Fatalf("Unable to get current working directory: %v", err)
            return
        }
        path := filepath.Join(pwd, accountInfo.AccountName + "-" + "output.csv")
        if *skipHeader {
            csv.WriteCSV(path, parsedData[1:])
        } else {
            csv.WriteCSV(path, parsedData)
        }
        log.Printf("Output written to: %v", path)
    }

    log.Println("Done!")
}
