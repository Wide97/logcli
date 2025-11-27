package formatter

import (
	"fmt"
	"strings"

	"github.com/Wide97/logcli/internal/model"
)

// se uso iniziale maiuscola, dichiaro la funzione essere pubblica, accetta in entrata un model.Stats
// restituisce una string
func ToCSV(stats model.Stats) string {
	//al posto di fare concatenazioni, posso buildare stringhe pezzo per pezzo.
	var sb strings.Builder
	//Scrive la prima e la seconda colonna
	sb.WriteString("category, count \n")
	// Ciclo sulla mappa [string]di int Counts--> k= nome categoria (error, info)
	//v= numero di occorrenze
	for k, v := range stats.Counts {
		//Produce una stringa come error,10 (a capo) info, 25
		sb.WriteString(fmt.Sprintf("%s, %d \n", k, v))
	}
	//aggiunge una riga finale con il numero di lines processate
	sb.WriteString(fmt.Sprintf("lines, %d\n", stats.Lines))
	//converte tutto in stringa
	return sb.String()
}
