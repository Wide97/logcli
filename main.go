package main

import (
	"fmt"
	"os"
)

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
	programName := args[0]
	filePaths := args[1:]

	//Stampo a console il nome del programma e i file passati in input
	fmt.Println("Programma: ", programName)
	fmt.Println("Files: ", filePaths)
}
