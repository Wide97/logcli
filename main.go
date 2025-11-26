package main

import (
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/Wide97/logcli/internal/analyzer"
	"github.com/Wide97/logcli/internal/model"
)

func exportCSV(stats model.Stats) {
	w := csv.NewWriter(os.Stdout)

	//intestazione
	w.Write([]string{"category", "count"})

	//per ogni chiave, ne scriviamo il valore
	for k, v := range stats.Counts {
		w.Write([]string{k, fmt.Sprintf("%d", v)})
	}
	w.Flush()
}

func main() {
	summaryOnly := flag.Bool("summary-only", false, "mostra solo il report finale")
	onlyErrors := flag.Bool("only-errors", false, "mostra solo le righe di categoria error")
	jsonOut := flag.Bool("json", false, "esporta il report in formato json")
	csvOut := flag.Bool("csv", false, "esporta il report in formato csv")

	flag.Parse()

	//os.Args legge gli argomenti all' interno dell' arg che gli vengono passati, un
	//args[0] sarà il nome del programma
	// args := os.Args

	//Utilizziamo ora flag al posto di os:
	filePaths := flag.Args()

	//Almeno due elementi in args

	if len(filePaths) == 0 {
		fmt.Println("Inserisci più di due file.")
		flag.PrintDefaults() //stampa elenco flag disponibili
		os.Exit(1)           //Errore utente
	}

	//Inserisco il nome del programma e gli argomenti che verranno passati successivamente
	// filePaths := args[1:]

	for _, path := range filePaths {
		fmt.Println(" -- Lettura file: -- ", path, " --- ") //log di chiarimento
		stats, err := analyzer.ReadFile(path, *summaryOnly, *onlyErrors)
		if err != nil {
			fmt.Println("Errore: ", err)
			continue
		}

		//chiamo la nuova funzione per creare il report csv
		if *csvOut {
			exportCSV(stats)
			continue
		}
		//Importo json encoding e quanto SEGUE:
		// fmt.Println("--- Report per: ", path, "---")
		// fmt.Println("Linee totali: ", stats.Lines)
		// fmt.Println("Counts: ")
		// for k, v := range stats.Counts {
		// 	fmt.Printf(" %s: %d\n", k, v)
		// }
		// fmt.Println()
		//DIVENTA:

		if *jsonOut {
			data, err := json.MarshalIndent(stats, "", "  ")
			if err != nil {
				fmt.Println("Errore json: ", err)
				continue
			}
			fmt.Println(string(data))
			continue
		}

		fmt.Println("--- Report per: ", path, "---")
		fmt.Println("Linee totali: ", stats.Lines)
		fmt.Println("Counts:")
		for k, v := range stats.Counts {
			fmt.Printf(" %s: %d\n", k, v)
		}

		fmt.Println()

	}
}
