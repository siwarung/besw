package utils

import (
	"strings"

	"github.com/texttheater/golang-levenshtein/levenshtein"
)

// NormalizeText menghapus spasi ekstra dan membuat teks menjadi lowercase
func NormalizeText(text string) string {
	return strings.ToLower(strings.TrimSpace(text))
}

// IsSimilarString membandingkan dua string dan mengecek apakah mereka mirip (berdasarkan Levenshtein Distance)
func IsSimilarString(a, b string, threshold int) bool {
	a = NormalizeText(a)
	b = NormalizeText(b)

	// Hitung jarak Levenshtein
	distance := levenshtein.DistanceForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
	return distance <= threshold
}
