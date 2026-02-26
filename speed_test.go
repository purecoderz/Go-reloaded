package main

import (
	"strings"
	"testing"
)

func BenchmarkRealSingleLoop(b *testing.B) {
	// The massive Harold Wilson test string!
	rawText := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair. Simply add 42 (hex) and 10 (bin) and you will see the result is 68. There is no greater agony than bearing a untold story inside you. Punctuation tests are ... kinda boring ,what do you think ? a"
	originalSlice := strings.Fields(rawText)

	// b.ResetTimer() tells the referee to IGNORE the time it took to setup the string above
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// 🚨 CRITICAL RULE: Because your function modifies the words,
		// we MUST give it a fresh copy of the sentence every single lap.
		wordsCopy := make([]string, len(originalSlice))
		copy(wordsCopy, originalSlice)

		// Call your REAL function!
		FormatWords(wordsCopy)
	}
}

func BenchmarkRealMultiLoop(b *testing.B) {
	// The massive Harold Wilson test string!
	rawText := "it (cap) was the best of times, it was the worst of times (up) , it was the age of wisdom, it was the age of foolishness (cap, 6) , it was the epoch of belief, it was the epoch of incredulity, it was the season of Light, it was the season of darkness, it was the spring of hope, IT WAS THE (low, 3) winter of despair. Simply add 42 (hex) and 10 (bin) and you will see the result is 68. There is no greater agony than bearing a untold story inside you. Punctuation tests are ... kinda boring ,what do you think ? a"
	originalSlice := strings.Fields(rawText)

	// b.ResetTimer() tells the referee to IGNORE the time it took to setup the string above
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// 🚨 CRITICAL RULE: Because your function modifies the words,
		// we MUST give it a fresh copy of the sentence every single lap.
		wordsCopy := make([]string, len(originalSlice))
		copy(wordsCopy, originalSlice)

		// Call your REAL function!
		FormatWords2(wordsCopy)
	}
}
