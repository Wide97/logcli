package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(path string) error {
	//Apriamo il file in input
	f, err := os.Open(path)
	if err != nil { //Se durante l' apertura non ci sono errori, prosegue skippando qui
		return fmt.Errorf("Errore nell' apertura dle file %s: %w:", path, err)
	}
	defer f.Close() // Chidiamo f, altrimenti mi sembra di aver capito che occupiamo inutilmente memoria

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		line := scanner.Text()                  //leggiamo tutte le righe del file
		category := classifyLine(line)          // chiamo la funzione per categorizzare ogni linea
		fmt.Printf("[%s] %s\n", category, line) // stampiamo solo la riga
	}

	if err := scanner.Err(); err != nil { //Nel caso ci fosse qualche errore durante lo scan del file
		return fmt.Errorf("errore nella lettura del file %s: %w", path, err)
	}

	return nil // torniamo null
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
	//os.Args legge gli argomenti all' interno dell' arg che gli vengono passati, un
	//args[0] sarà il nome del programma
	args := os.Args

	//Almeno due elementi in args

	if len(args) < 2 {
		fmt.Println("Inserisci più di due file.")
		os.Exit(1) //Errore utente
	}

	//Inserisco il nome del programma e gli argomenti che verranno passati successivamente
	filePaths := args[1:]

	for _, path := range filePaths {
		fmt.Println(" -- Lettura file: -- ", path, " --- ") //log di chiarimento
		err := readFile(path)                               // chiamo la funzione readfile
		if err != nil {
			fmt.Println("Errore : ", err) //chiamata solo in caso di errore
		}
	}
}
