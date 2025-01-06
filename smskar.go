package main

import (
	"fmt"
	"os"
	"strings"
)

// Map of words to Unicode symbols
var wordToUnicodeMap = map[string]string{
	"sync":     "🔄",
	"download": "⬇️",
	"upload":   "⬆️",
	"refresh":  "🔄",
	"save":     "💾",
	"delete":   "❌",
	"play":     "▶️",
	"pause":    "⏸️",
	"stop":     "⏹️",
	"home":     "🏠",
	"check":    "✔️",
	"error":    "❗",
	"warning":  "⚠️",
	"heart":    "❤️",
}

func convertToUnicode(word string) string {

	// For each word, check if there's a Unicode symbol mapped to it
	var result string

	// If the word has a mapped Unicode symbol, append it
	if symbol, exists := wordToUnicodeMap[word]; exists {
		result += symbol
	} else {
		// If no mapped symbol exists, keep the word as is
		result += word
	}

	// Return the result without trailing space
	return strings.TrimSpace(result)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide input words as command-line arguments.")
		return
	}

	for _, arg := range os.Args[1:] {
		unicodeWord := convertToUnicode(arg)
		fmt.Printf("%16s\t%16s\n", arg, unicodeWord)
	}
}
