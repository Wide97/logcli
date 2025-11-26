package analyzer

import (
	"os"
	"testing"
)

func TestReadFile(t *testing.T) {
	//immettiamo un simil- file per verificare il funzionamento del codice
	logContent := `INFO Starting
WARN Disk low
ERROR Explosion
no keyword here
error again` // minuscolo per testare case-insensitive
	//Creiamo il file temporaneo
	tmpFile, err := os.CreateTemp("", "logcli_test_*.log")
	if err != nil {
		t.Fatalf("errore creazione file temporaneo: %v", err)
	}
	//una volta finito il test, rimuoviamo il file temporaneo
	defer os.Remove(tmpFile.Name())
	//scrivo contenuto nel file
	if _, err := tmpFile.Write([]byte(logContent)); err != nil {
		t.Fatalf("errore scrittura del file temporaneo: %v", err)
	}
	//chiudo il file temporaneo
	tmpFile.Close()

	stats, err := ReadFile(tmpFile.Name(), false, false)
	if err != nil {
		t.Fatalf("errore lettura del file %v", err)
	}
	// verifico le linee totali
	expectedLines := 5
	if stats.Lines != expectedLines {
		t.Fatalf("Lines= %d, expected %d", stats.Lines, expectedLines)
	}
	expected := map[string]int{
		"info":  1,
		"warn":  1,
		"error": 2,
		"other": 1,
	}

	for key, exp := range expected {
		got := stats.Counts[key]
		if got != exp {
			t.Fatalf("Counts[%q] = %d, expected %d", key, got, exp)
		}
	}

}
