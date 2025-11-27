📄 README.md — logcli

Analizzatore di file di log da riga di comando scritto in Go

# logcli

**logcli** è uno strumento da riga di comando scritto in Go per analizzare file di log.  
Consente di:

- leggere uno o più file di log
- classificare automaticamente le righe (`INFO`, `WARN`, `ERROR`, `OTHER`)
- mostrare solo gli errori o il file completo
- generare statistiche
- esportare report in **JSON** o **CSV**
- vedere **linea + testo** di ogni errore individuato

Il tool supporta l’analisi concorrente di più file in parallelo.

---

## 🚀 Installazione

### Compilazione locale

Se hai Go installato:

```bash
go build -o logcli .


Su Windows verrà creato logcli.exe.

Installazione via go install (da GitHub)

Funziona solo dopo aver pubblicato il progetto su GitHub.

go install github.com/Wide97/logcli@latest


Assicurati che ~/go/bin sia nel tuo PATH.

🧠 Utilizzo

La forma generale è:

logcli [opzioni] file1.log [file2.log ...]

Esempi
Analisi semplice
logcli app.log

Mostrare solo gli errori
logcli --only-errors app.log

Non mostrare il dettaglio riga-per-riga, solo il riepilogo
logcli --summary-only app.log

Export JSON
logcli --json app.log


Output tipo:

{
  "counts": {
    "error": 2,
    "warn": 1,
    "info": 3,
    "other": 4
  },
  "lines": 10,
  "errors": [
    { "line": 3, "text": "ERROR DB Failure" },
    { "line": 8, "text": "ERROR Timeout" }
  ]
}

Export CSV
logcli --csv app.log


Output tipo:

category,count
error,2
warn,1
info,3
other,4
lines,10

error_line,error_text
3,ERROR DB Failure
8,ERROR Timeout

🏷️ Opzioni disponibili
Flag	Descrizione
--summary-only	Mostra solo il riepilogo finale, senza stampare ogni riga
--only-errors	Mostra (o esporta) solo le righe in categoria error
--json	Esporta il report in formato JSON
--csv	Esporta il report in formato CSV
🧱 Struttura del progetto
logcli/
├── cmd/           (eventuale, non usato ora)
├── internal/
│   ├── analyzer/   → logica di lettura e classificazione file
│   ├── cli/        → parsing dei flag e validazione argomenti
│   ├── formatter/  → generazione JSON e CSV
│   └── model/      → definizione strutture dati (Stats, ErrorDetail)
├── go.mod
├── go.sum
└── main.go

⚙️ Funzionalità principali

Classificazione automatica
Riconosce parole chiave in modo case-insensitive:
ERROR, WARN, INFO.

Dettagli errori
Per ogni errore salva:

numero di linea

testo originale

(in futuro potremo aggiungere timestamp o regex custom)

Analisi concorrente
Ogni file viene processato in una goroutine separata.

Formati di export

JSON ben formattato (MarshalIndent)

CSV leggibile con blocco riepilogo + blocco errori

🧪 Test

I test sono eseguibili con:

go test ./...


Attualmente i test coprono la logica di analisi (internal/analyzer).

📦 Cross-compilazione (opzionale)

Per creare binari per altri sistemi operativi:

# Linux x64
GOOS=linux GOARCH=amd64 go build -o bin/logcli-linux .

# macOS ARM (M1/M2)
GOOS=darwin GOARCH=arm64 go build -o bin/logcli-mac .

# Windows x64
GOOS=windows GOARCH=amd64 go build -o bin/logcli-windows.exe .

📌 Da fare / future features

parsing regex personalizzabile per categorie

supporto a log multi-linea (stacktrace)

esportazione HTML

worker pool configurabile

configurazione tramite file .yaml

GUI minimale per Windows

👨‍💻 Autore

Progetto creato da Marco Widesott
Scrivo Go per imparare divertendomi e costruire tool utili nella vita reale.
