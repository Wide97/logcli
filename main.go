package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

// func readFile(path string) (Stats, error) Aggiungo anche i flag:
func readFile(path string, summaryOnly bool, onlyErrors bool) (Stats, error) {
	stats := Stats{
		Counts: make(map[string]int),
		Lines:  0,
	}
	//Apriamo il file in input
	f, err := os.Open(path)
	if err != nil { //Se durante l' apertura non ci sono errori, prosegue skippando qui
		return stats, fmt.Errorf("Errore nell' apertura dle file %s: %w:", path, err)
	}
	defer f.Close() // Chidiamo f, altrimenti mi sembra di aver capito che occupiamo inutilmente memoria

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()         //leggiamo tutte le righe del file
		category := classifyLine(line) // chiamo la funzione per categorizzare ogni linea
		stats.Lines++                  //incrementa le line man mano che vengono lette
		stats.Counts[category]++       // conta per categoria
		// fmt.Printf("[%s] %s\n", category, line) // stampiamo solo la riga
		//modifichiamo al fly la parte sopra, introducendo:
		if !summaryOnly {
			if onlyErrors && category != "error" {
				continue
			}
			fmt.Printf("[%s] %s\n", category, line)
		}
	}

	if err := scanner.Err(); err != nil { //Nel caso ci fosse qualche errore durante lo scan del file
		return stats, fmt.Errorf("errore nella lettura del file %s: %w", path, err)
	}

	return stats, nil // torniamo null
}

func classifyLine(line string) string { //gli passo ogni line del file

	upper := strings.ToUpper(line) // confornto più robusto con caps

	//catcho ogni caso, che sia error, warn, info o altro
	switch {
	case strings.Contains(upper, "ERROR"):
		return "error"
	case strings.Contains(upper, "WARN"):
		return "warn"
	case strings.Contains(upper, "INFO"):
		return "info"
	default:
		return "other"
	}

}

func main() {
	summaryOnly := flag.Bool("summary-only", false, "mostra solo il report finale")
	onlyErrors := flag.Bool("only-errors", false, "mostra solo le righe di categoria error")

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
		stats, err := readFile(path, *summaryOnly, *onlyErrors)
		if err != nil {
			fmt.Println("Errore: ", err)
			continue
		}

		fmt.Println("--- Report per: ", path, "---")
		fmt.Println("Linee totali: ", stats.Lines)
		fmt.Println("Counts: ")
		for k, v := range stats.Counts {
			fmt.Printf(" %s: %d\n", k, v)
		}
		fmt.Println()

	}
}

type Stats struct {
	Counts map[string]int
	Lines  int
}
