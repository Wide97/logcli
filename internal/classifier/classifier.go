package classifier

import "strings"

type LineClassifier interface {
	Classify(line string) string
}

type SimpleClassifier struct{}

func NewSimpleClassifier() *SimpleClassifier {
	return &SimpleClassifier{}

}

// Classify implementa l'interfaccia LineClassifier per SimpleClassifier.
func (c *SimpleClassifier) Classify(line string) string {
	lower := strings.ToLower(line)

	switch {
	case strings.Contains(lower, "error"):
		return "error"
	case strings.Contains(lower, "warn"):
		return "warn"
	case strings.Contains(lower, "info"):
		return "info"
	default:
		return "other"
	}
}
