package main

import (
	"encoding/csv"
	"flag"
	"fmt"
	"os"

	"github.com/Wide97/logcli/internal/analyzer"
	"github.com/Wide97/logcli/internal/formatter"
	"github.com/Wide97/logcli/internal/model"
)

// Contenitore per path del file, statistiche e possibili errori
type fileResult struct {
	path  string
	stats model.Stats
	err   error
}

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

	//SOSTITUISCO TUTTO IL BLOCCO [2]:

	// for _, path := range filePaths {
	// 	fmt.Println(" -- Lettura file: -- ", path, " --- ") //log di chiarimento
	// 	stats, err := analyzer.ReadFile(path, *summaryOnly, *onlyErrors)
	// 	if err != nil {
	// 		fmt.Println("Errore: ", err)
	// 		continue
	// 	}

	// 	//chiamo la nuova funzione per creare il report csv
	// 	if *csvOut {
	// 		exportCSV(stats)
	// 		continue
	// 	}
	// 	//Importo json encoding e quanto SEGUE:
	// 	// fmt.Println("--- Report per: ", path, "---")
	// 	// fmt.Println("Linee totali: ", stats.Lines)
	// 	// fmt.Println("Counts: ")
	// 	// for k, v := range stats.Counts {
	// 	// 	fmt.Printf(" %s: %d\n", k, v)
	// 	// }
	// 	// fmt.Println()
	// 	//DIVENTA:

	// 	if *jsonOut {
	// 		data, err := json.MarshalIndent(stats, "", "  ")
	// 		if err != nil {
	// 			fmt.Println("Errore json: ", err)
	// 			continue
	// 		}
	// 		fmt.Println(string(data))
	// 		continue
	// 	}

	// 	fmt.Println("--- Report per: ", path, "---")
	// 	fmt.Println("Linee totali: ", stats.Lines)
	// 	fmt.Println("Counts:")
	// 	for k, v := range stats.Counts {
	// 		fmt.Printf(" %s: %d\n", k, v)
	// 	}

	// 	fmt.Println()

	//CON [2]

	//channel per ricevere dati della go routine
	results := make(chan fileResult)
	// lancio una routine per ogni file
	for _, path := range filePaths {
		p := path //copia locale per evitare problemi di cattura variabile
		fmt.Println("-- Lettura file:", p, "---")

		go func() {
			stats, err := analyzer.ReadFile(p, *summaryOnly, *onlyErrors)
			results <- fileResult{
				path:  p,
				stats: stats,
				err:   err,
			}
		}()
	}
	i := 0
	//raccolgo i risultati, tante volte quanti sono i file
	for i = 0; i < len(filePaths); i++ {
		res := <-results //blocca finchè non arriva un risultato

		if res.err != nil {
			fmt.Println("errore su: ", res.path, ": ", res.err)
			continue
		}

		//Commento perchè, una volta creato il file json.go, me lo importo qui : [3]
		// if *jsonOut {
		// 	data, err := json.MarshalIndent(res.stats, "", "  ")
		// 	if err != nil {
		// 		fmt.Println("Errore JSON su", res.path, ":", err)
		// 		continue
		// 	}
		// 	fmt.Println("== JSON report per:", res.path, "==")
		// 	fmt.Println(string(data))
		// 	continue
		// }

		//e scrivo: [3]

		if *jsonOut {
			//Dentro ToJSON (funzione dell' altro file) mi viene restituita una stringa leggibile a blocchi
			jsonStr, err := formatter.ToJSON(res.stats)
			if err != nil {
				fmt.Println("Errore JSON su", res.path, ":", err)
				continue
			}
			//stampo il risultato formattato
			fmt.Println("== JSON report per:", res.path, "==")
			fmt.Println(jsonStr)
			continue
		}
		//Commento perchè, una volta creato il file csv.go, me lo importo qui : [4]
		// if *csvOut {
		// 	fmt.Println("== CSV report per:", res.path, "==")
		// 	exportCSV(res.stats)
		// 	continue
		// }

		//e scrivo: [4]

		if *csvOut {
			//to CSV crea string.Builder, aggiunge intestazione ecc.
			fmt.Println("== CSV report per:", res.path, "==")
			csvStr := formatter.ToCSV(res.stats)
			//scrive la stringa ricevuta
			fmt.Println(csvStr)
			continue
		}

		fmt.Println("--- Report per:", res.path, "---")
		fmt.Println("Linee totali:", res.stats.Lines)
		fmt.Println("Counts:")
		for k, v := range res.stats.Counts {
			fmt.Printf("  %s: %d\n", k, v)
		}
		fmt.Println()

	}

}
