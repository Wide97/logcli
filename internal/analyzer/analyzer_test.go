package analyzer

import (
	"testing"

	"github.com/Wide97/logcli/internal/classifier"
)

func TestClassifyLine(t *testing.T) {
	//Imposto la struttura
	test := []struct {
		name     string
		input    string
		expected string
	}{
		//Fornisco dei casi di test
		{"info uppercase", "INFO Starting application", "info"},
		{"info lowercase", "info: something", "info"},
		{"warn uppercase", "WARN Disk almost full", "warn"},
		{"warn mixed", "WaRn something", "warn"},
		{"error uppercase", "ERROR Fatal exception", "error"},
		{"error lowercase", "error occurred", "error"},
		{"no keyword", "this line has no level", "other"},
	}
	//Verifico il funzionamento chiamando il metodo classifyline
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			got := classifier.NewSimpleClassifier().Classify(tt.input)
			if got != tt.expected {
				t.Fatalf("ClassifyLine(%q) = %q, expected %q", tt.input, got, tt.expected)
			}
		})
	}
}
