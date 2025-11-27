package analyzer

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Wide97/logcli/internal/classifier"
	"github.com/Wide97/logcli/internal/model"
)

// func readFile(path string) (Stats, error) Aggiungo anche i flag:
func ReadFile(path string, summaryOnly bool, onlyErrors bool, c classifier.LineClassifier) (model.Stats, error) {
	stats := model.Stats{
		Counts: make(map[string]int),
		Lines:  0,
		Errors: []model.ErrorDetail{},
	}
	//Apriamo il file in input
	f, err := os.Open(path)
	if err != nil { //Se durante l' apertura non ci sono errori, prosegue skippando qui
		return stats, fmt.Errorf("Errore nell' apertura dle file %s: %w:", path, err)
	}
	defer f.Close() // Chidiamo f, altrimenti mi sembra di aver capito che occupiamo inutilmente memoria

	scanner := bufio.NewScanner(f)
	//aggiungo un contatore di linea
	lineNo := 0

	for scanner.Scan() {
		//incremento il contatore
		lineNo++
		line := scanner.Text()       //leggiamo tutte le righe del file
		category := c.Classify(line) // chiamo la funzione per categorizzare ogni linea
		stats.Lines++                //incrementa le line man mano che vengono lette

		stats.Counts[category]++ // conta per categoria
		// fmt.Printf("[%s] %s\n", category, line) // stampiamo solo la riga
		//modifichiamo al fly la parte sopra, introducendo:
		// if !summaryOnly {
		// 	if onlyErrors && category != "error" {
		// 		continue
		// 	}
		// 	fmt.Printf("[%s] %s\n", category, line)
		// }
		//se vogliamo solo il riepilogo, non stampiamo le singole righe

		if category == "error" {
			stats.Errors = append(stats.Errors, model.ErrorDetail{
				Line: lineNo,
				Text: line,
			})
		}
		if summaryOnly {
			continue
		}
		//se vogliamo solo gli errori, skippiamo il resto
		if onlyErrors && category != "error" {
			continue
		}

		if category == "error" {
			//per errori mostriamo il numero di riga
			fmt.Printf("[error] (linea %d) %s\n", lineNo, line)
		} else {
			//per le altre categorie lascio come era
			fmt.Printf("[%s] %s\n", category, line)
		}
	}

	if err := scanner.Err(); err != nil { //Nel caso ci fosse qualche errore durante lo scan del file
		return stats, fmt.Errorf("errore nella lettura del file %s: %w", path, err)
	}

	return stats, nil // torniamo null
}

func ClassifyLine(line string) string { //gli passo ogni line del file

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
