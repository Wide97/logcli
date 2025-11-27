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
	//se non ci sono errori finisco qui
	if len(stats.Errors) == 0 {
		return sb.String()
	}
	//separo con riga vuota
	sb.WriteString("\n")

	//seconda tabella con errori:
	sb.WriteString("error_line, error_text\n")
	for _, e := range stats.Errors {
		sb.WriteString(fmt.Sprintf("%d,\"%s\"\n", e.Line, e.Text))
	}

	return sb.String()

}
