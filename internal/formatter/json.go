package formatter

import (
	"encoding/json"
	"fmt"

	"github.com/Wide97/logcli/internal/model"
)

func ToJSON(stats model.Stats) (string, error) {
	//json.MarshalIndent serializza la struttura in json e aggiunge indentazione carina per leggerlo meglio
	b, err := json.MarshalIndent(stats, "", "  ")
	if err != nil {
		return "", fmt.Errorf("errore serializzazione JSON: %w", err)
	}
	//Converte i byte in json e nil come errorre
	return string(b), nil
}
