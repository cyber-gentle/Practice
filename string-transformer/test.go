// CodeCrafters — Operation Gopher Protocol
// Module: String Transformer
// Author: Abraham David
// Squad:  The Gopher's

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ─────────────────────────────────────────────────────
// toUpperCase — converts every letter to UPPERCASE
// You wrote this one. It works perfectly.
// ─────────────────────────────────────────────────────
func toUpperCase(s string) string {
	return strings.ToUpper(s)
}

// ─────────────────────────────────────────────────────
// toLowerCase — converts every letter to lowercase
// You wrote this one too. Also works perfectly.
// ─────────────────────────────────────────────────────
func toLowerCase(s string) string {
	return strings.ToLower(s)
}

// ─────────────────────────────────────────────────────
// capitalizeCase — capitalises the first letter of
// every word. All other letters go lowercase.
// Example: "THREAT LEVEL elevated" → "Threat Level Elevated"
// ─────────────────────────────────────────────────────
func capitalizeCase(s string) string {
	// Split the input into individual words
	words := strings.Fields(s)

	for i, word := range words {
		if len(word) == 0 {
			continue
		}
		// Make the whole word lowercase first,
		// then capitalise just the very first character
		words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	}

	// Join the words back into a single string
	return strings.Join(words, " ")
}

// ─────────────────────────────────────────────────────
// titleCase — like capitalizeCase but small connector
// words stay lowercase unless they are the first word.
// Small words: a, an, the, and, but, or, for, nor,
//
//	on, at, to, by, in, of, up, as, is, it
//
// Example: "the fall of the western power grid"
//
//	→ "The Fall of the Western Power Grid"
//
// ─────────────────────────────────────────────────────
func titleCase(s string) string {
	// This map lets us quickly check if a word is a small word
	// map[string]bool means: key is a string, value is true/false
	smallWords := map[string]bool{
		"a": true, "an": true, "the": true, "and": true,
		"but": true, "or": true, "for": true, "nor": true,
		"on": true, "at": true, "to": true, "by": true,
		"in": true, "of": true, "up": true, "as": true,
		"is": true, "it": true,
	}

	words := strings.Fields(s)

	for i, word := range words {
		if len(word) == 0 {
			continue
		}
		lower := strings.ToLower(word)

		// The first word is ALWAYS capitalised — even if it's a small word
		// For all other words: capitalise unless it's a small word
		if i == 0 || !smallWords[lower] {
			words[i] = strings.ToUpper(string(lower[0])) + lower[1:]
		} else {
			// It's a small word and not the first — keep it lowercase
			words[i] = lower
		}
	}

	return strings.Join(words, " ")
}

// ─────────────────────────────────────────────────────
// toSnakeCase — converts text to snake_case.
// All lowercase, spaces replaced with underscores.
// Any character that is not a letter, digit, or
// underscore is removed.
// Example: "Alert! Level 5 detected." → "alert_level_5_detected"
// ─────────────────────────────────────────────────────
func toSnakeCase(s string) string {
	// First replace spaces with underscores
	word := strings.ReplaceAll(s, " ", "_")

	// Now go through every character and keep only
	// letters, digits, and underscores
	var result strings.Builder // Builder lets us build a string piece by piece

	for _, char := range word {
		// Check if the character is a letter, a digit, or an underscore
		isLetter := (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z')
		isDigit := char >= '0' && char <= '9'
		isUnderscore := char == '_'

		if isLetter || isDigit || isUnderscore {
			result.WriteRune(char) // Keep this character
		}
		// Otherwise: skip it (this removes !, ., , etc.)
	}

	// Finally convert the whole thing to lowercase
	return strings.ToLower(result.String())
}

// ─────────────────────────────────────────────────────
// reverseWords — reverses each word individually.
// Word order stays the same. Only the letters inside
// each word are reversed.
// Example: "Lagos Nigeria" → "sogaL airegiN"
// ─────────────────────────────────────────────────────
func reverseWords(s string) string {
	words := strings.Fields(s)

	for i, word := range words {
		// Convert the word to a slice of runes so we can index into it
		// (rune handles special characters safely)
		runes := []rune(word)

		// Swap from both ends moving toward the middle
		left := 0
		right := len(runes) - 1
		for left < right {
			runes[left], runes[right] = runes[right], runes[left]
			left++
			right--
		}

		words[i] = string(runes)
	}

	return strings.Join(words, " ")
}

// ─────────────────────────────────────────────────────
// palindrome — checks if the text reads the same
// forwards and backwards (ignoring spaces and casing).
// Example: "never odd or even" → is a palindrome
// Example: "sentinel"          → is NOT a palindrome
// ─────────────────────────────────────────────────────
func palindrome(s string) string {
	// Remove spaces and convert to lowercase for a fair comparison
	cleaned := strings.ToLower(strings.ReplaceAll(s, " ", ""))

	// Reverse the cleaned string
	runes := []rune(cleaned)
	left := 0
	right := len(runes) - 1
	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}
	reversed := string(runes)

	// If the cleaned string equals its reverse — it's a palindrome
	if cleaned == reversed {
		return fmt.Sprintf("✦ %q is a palindrome!", s)
	}
	return fmt.Sprintf("✗ %q is not a palindrome.", s)
}

// ─────────────────────────────────────────────────────
// MAIN — where the program starts
// ─────────────────────────────────────────────────────
func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("\n SENTINEL STRING TRANSFORMER — ONLINE")
	fmt.Println(" ──────────────────────────────────────")

	// This loop keeps the program running until the user types "exit"
	for {
		fmt.Print("\n > ")

		// Read the full line the user typed
		inputWord, _ := reader.ReadString('\n')
		inputWord = strings.TrimSpace(inputWord)

		// Edge case 8: user pressed Enter with nothing typed — just loop again
		if inputWord == "" {
			continue
		}

		// Split the input into a slice of words
		split := strings.Fields(inputWord)

		// The first word is the command (always lowercased so UPPER = upper)
		cmd := strings.ToLower(split[0])

		// Everything after the command is the text to transform
		wordSlice := split[1:]

		// Edge case 2 & 6: command given but no text after it
		if len(wordSlice) == 0 && cmd != "exit" {
			fmt.Printf(" ✗ No text provided. Usage: %v <text>\n", cmd)
			continue
		}

		// Join the word slice back into a single string
		words := strings.Join(wordSlice, " ")

		switch cmd {
		case "upper":
			fmt.Printf("   → %v\n", toUpperCase(words))

		case "lower":
			fmt.Printf("   → %v\n", toLowerCase(words))

		case "cap":
			fmt.Printf("   → %v\n", capitalizeCase(words))

		case "title":
			fmt.Printf("   → %v\n", titleCase(words))

		case "snake":
			fmt.Printf("   → %v\n", toSnakeCase(words))

		case "reverse":
			fmt.Printf("   → %v\n", reverseWords(words))

		case "palindrome":
			fmt.Printf("   → %v\n", palindrome(words))

		case "exit":
			fmt.Println(" Shutting down String Transformer. Goodbye.")
			return

		default:
			// Edge case 1: unknown command
			fmt.Printf(" ✗ Unknown command: %q\n", cmd)
			fmt.Println("   Valid commands: upper, lower, cap, title, snake, reverse, palindrome, exit")
		}
	}
}
