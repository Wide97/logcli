package cli

import (
	"flag"
	"fmt"
	"os"
)

type Options struct {
	SummaryOnly bool
	OnlyErrors  bool
	JsonOutput  bool
	CsvOutput   bool
	Files       []string
}

func ParseArgs() Options {
	summaryOnly := flag.Bool("summary-only", false, "mostra solo il report finale")
	onlyErrors := flag.Bool("only-errors", false, "mostra solo le righe di categoria error")
	jsonFlag := flag.Bool("json", false, "esporta il report in formato json")
	csvFlag := flag.Bool("csv", false, "esporta il report in formato csv")

	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		fmt.Println("Uso: nessun file selezionato")
		os.Exit(1)
	}

	return Options{
		SummaryOnly: *summaryOnly,
		OnlyErrors:  *onlyErrors,
		JsonOutput:  *jsonFlag,
		CsvOutput:   *csvFlag,
		Files:       files,
	}
}
