package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Args legge gli argomenti all' interno dell' arg che gli vengono passati, un
	//args[0] sarà il nome del programma
	args := os.Args
	fmt.Println("logcli inizializzato!")
	fmt.Println("Numero di argomenti: ", len(args))
	fmt.Println("Contenuto args: ", args)
}
