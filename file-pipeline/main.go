// CodeCrafters — Operation Gopher Protocol
// Module: File Pipeline
// Author: Abraham David
// Squad: The Gophers

// ═══════════════════════════════════════════
// SQUAD PIPELINE CONTRACT
// Squad: Gophers
// ───────────────────────────────────────────
// Input line types:
// Number of lines: 20
// Normal report lines
// Lines in ALL CAPS
// Lines in all lowercase
// Lines starting with TODO:
// Lines with extra leading/trailing spaces
//
// Transformation rules (in order):
// 1. Trim all leading and trailing whitespace
// 2. Replace TODO: with ✦ ACTION:
// 3. Convert ALL CAPS lines to Title Case
// 4. Convert all lowercase lines to uppercase
// 5. Reverse the words in any line that contains the word REVERSE
//
// Output format:
// Header: Yes, Exact Text: "Gopher's Sentinel Field Report - Processed"
// Line numbering format : "1."
// Summary block: yes
//  	Fields :
//		✦ Lines read    :
//		✦ Lines written :
//		✦ Lines removed :
//		✦ Rules applied : [our 5 rules]
//
// Terminal summary fields:
//		✦ Lines read    :
//		✦ Lines written :
//		✦ Lines removed :
//		✦ Rules applied : [our 5 rules]
// ═══════════════════════════════════════════

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

// ─────────────────────────────────────────────────────
// RULE 1 — Trim all leading and trailing whitespace
// ─────────────────────────────────────────────────────
// This function removes any extra spaces at the
// beginning or end of a line.
// Example: "   hello world   " → "hello world"
func trimWhitespace(line string) string {
	return strings.TrimSpace(line)
}

// ─────────────────────────────────────────────────────
// RULE 2 — Replace TODO: with ✦ ACTION:
// ─────────────────────────────────────────────────────
// If a line starts with "TODO:" (in any casing — "todo:", "Todo:", etc.)
// we replace it with "✦ ACTION:" to mark it as an action item.
// Example: "TODO: send backup" → "✦ ACTION: send backup"
// Example: "todo: check logs" → "✦ ACTION: check logs"
func replaceTodo(line string) string {
	// We check the uppercased version so we catch "todo:", "Todo:", "TODO:" etc.
	if strings.HasPrefix(strings.ToUpper(line), "TODO:") {
		// "TODO:" is always 5 characters long.
		// We grab everything after those 5 characters and attach our prefix.
		rest := line[5:] // everything after "TODO:" (keeps original casing of the rest)
		line = "✦ ACTION:" + rest
	}
	return line
}

// ─────────────────────────────────────────────────────
// RULE 3 — Convert ALL CAPS lines to Title Case
// ─────────────────────────────────────────────────────
// If every letter in the line is uppercase,
// we convert it to Title Case (first letter of
// each word is capital, the rest are lowercase).
// Example: "ALERT LEVEL FIVE" → "Alert Level Five"
func convertAllCapsToTitleCase(line string) string {
	// First check if the line is actually ALL CAPS
	if isAllCaps(line) {
		return toTitleCase(line)
	}
	return line
}

// isAllCaps checks if all the letters in a line are uppercase.
// It ignores spaces, numbers, and symbols — only checks letters.
func isAllCaps(line string) bool {
	// We need at least one letter to consider it "all caps"
	hasLetter := false

	for _, ch := range line {
		if unicode.IsLetter(ch) {
			hasLetter = true
			// If we find even one lowercase letter, it's NOT all caps
			if unicode.IsLower(ch) {
				return false
			}
		}
	}

	// Return true only if there was at least one letter
	return hasLetter
}

// toTitleCase converts a string so the first letter of
// each word is uppercase and the rest are lowercase.
func toTitleCase(line string) string {
	// Split the line into individual words
	words := strings.Fields(line)

	// Go through each word and capitalise its first letter
	for i, word := range words {
		if len(word) == 0 {
			continue
		}
		// Make the whole word lowercase first, then
		// capitalise just the first character
		words[i] = strings.ToUpper(string(word[0])) + strings.ToLower(word[1:])
	}

	// Join the words back together with spaces
	return strings.Join(words, " ")
}

// ─────────────────────────────────────────────────────
// RULE 4 — Convert all lowercase lines to UPPERCASE
// ─────────────────────────────────────────────────────
// If every letter in the line is lowercase,
// we convert the whole line to uppercase.
// Example: "the perimeter was breached" → "THE PERIMETER WAS BREACHED"
func convertAllLowerToUpper(line string) string {
	if isAllLower(line) {
		return strings.ToUpper(line)
	}
	return line
}

// isAllLower checks if all letters in the line are lowercase.
// Like isAllCaps, it ignores spaces, numbers, and symbols.
func isAllLower(line string) bool {
	hasLetter := false

	for _, ch := range line {
		if unicode.IsLetter(ch) {
			hasLetter = true
			// If we find even one uppercase letter, it's NOT all lower
			if unicode.IsUpper(ch) {
				return false
			}
		}
	}

	return hasLetter
}

// ─────────────────────────────────────────────────────
// RULE 5 — Reverse the words in lines containing REVERSE
// ─────────────────────────────────────────────────────
// If the line contains the word "REVERSE" anywhere,
// we flip the order of all the words in that line.
// Example: "send backup REVERSE now" → "now REVERSE backup send"
func reverseWords(line string) string {
	// strings.Contains checks if "REVERSE" is anywhere in the line
	if strings.Contains(line, "REVERSE") {
		// Split the line into words
		words := strings.Fields(line)

		// Swap words from both ends moving toward the middle
		// This is a standard technique to reverse a slice
		left := 0
		right := len(words) - 1
		for left < right {
			// Swap the two words
			words[left], words[right] = words[right], words[left]
			left++
			right--
		}

		return strings.Join(words, " ")
	}
	return line
}

// ─────────────────────────────────────────────────────
// applyAllRules runs a single line through all 5 rules
// in the exact order of the Pipeline Contract.
// ─────────────────────────────────────────────────────
func applyAllRules(line string) string {
	line = trimWhitespace(line)            // Rule 1
	line = replaceTodo(line)               // Rule 2
	line = convertAllCapsToTitleCase(line) // Rule 3
	line = convertAllLowerToUpper(line)    // Rule 4
	line = reverseWords(line)              // Rule 5
	return line
}

// ─────────────────────────────────────────────────────
// MAIN — This is where the program starts
// ─────────────────────────────────────────────────────
func main() {
	// ── Step 1: Check that the user gave us two arguments ──
	// os.Args holds everything the user typed when running the program.
	// os.Args[0] is always the program name itself.
	// os.Args[1] should be input file, os.Args[2] should be output file.
	if len(os.Args) < 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	// ── Step 2: Make sure input and output are not the same file ──
	if inputFile == outputFile {
		fmt.Println("✗ Input and output cannot be the same file.")
		os.Exit(1)
	}

	// ── Step 3: Check that the output path is not a directory ──
	info, err := os.Stat(outputFile)
	if err == nil && info.IsDir() {
		fmt.Println("✗ Cannot write to output: path is a directory, not a file.")
		os.Exit(1)
	}

	// ── Step 4: Open the input file ──
	file, err := os.Open(inputFile)
	if err != nil {
		// The file does not exist or cannot be opened
		fmt.Println("✗ File not found:", inputFile)
		os.Exit(1)
	}
	// defer means: close the file when main() finishes
	defer file.Close()

	// ── Step 5: Read the file line by line using bufio.Scanner ──
	// bufio.Scanner is better than os.ReadFile because it reads
	// one line at a time — it doesn't load the whole file into memory.
	scanner := bufio.NewScanner(file)

	// We will store each processed line in this slice
	var processedLines []string

	// Counters to track what happened
	linesRead := 0
	linesRemoved := 0

	for scanner.Scan() {
		// Get the current line (Scanner handles \r\n and \n automatically)
		line := scanner.Text()
		linesRead++

		// Run the line through all 5 transformation rules
		transformed := applyAllRules(line)

		// If the line is empty after trimming — remove it
		if transformed == "" {
			linesRemoved++
			continue // skip to the next line, don't add this one
		}

		// Otherwise, keep it
		processedLines = append(processedLines, transformed)
	}

	// Handle the special case where the input file was completely empty
	if linesRead == 0 {
		fmt.Println("⚠ Input file is empty. Nothing to process.")
		// Still create an empty output file
		emptyOut, err := os.Create(outputFile)
		if err != nil {
			fmt.Println("✗ Could not write output file:", err)
			os.Exit(1)
		}
		emptyOut.Close()

		// Print summary and exit
		fmt.Println("✦ Lines read    :", 0)
		fmt.Println("✦ Lines written :", 0)
		fmt.Println("✦ Lines removed :", 0)
		fmt.Println("✦ Rules applied : Trim whitespace, Replace TODO, ALL CAPS → Title Case, all lower → UPPER, Reverse words")
		return
	}

	// ── Step 6: Open (or create) the output file for writing ──
	// os.Create will create the file if it doesn't exist,
	// or overwrite it completely if it does.
	out, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("✗ Could not write output file:", err)
		os.Exit(1)
	}
	defer out.Close()

	// We use bufio.Writer so we can write to the file efficiently
	writer := bufio.NewWriter(out)

	// ── Step 7: Write the header ──
	fmt.Fprintln(writer, "Gopher's Sentinel Field Report - Processed")
	fmt.Fprintln(writer, "───────────────────────────────────────────")

	// ── Step 8: Write each processed line with a number ──
	// Example: "1. Alert Level Five Detected"
	for i, line := range processedLines {
		// i starts at 0, so we add 1 to make the first line "1."
		fmt.Fprintf(writer, "%d. %s\n", i+1, line)
	}

	// ── Step 9: Write the summary block to the output file ──
	fmt.Fprintln(writer, "───────────────────────────────────────────")
	fmt.Fprintln(writer, "SUMMARY")
	fmt.Fprintf(writer, "✦ Lines read    : %d\n", linesRead)
	fmt.Fprintf(writer, "✦ Lines written : %d\n", len(processedLines))
	fmt.Fprintf(writer, "✦ Lines removed : %d\n", linesRemoved)
	fmt.Fprintln(writer, "✦ Rules applied : Trim whitespace | Replace TODO | ALL CAPS → Title Case | all lower → UPPER | Reverse words")

	// writer.Flush() makes sure everything is actually written to the file
	// Without this, some lines might be left in memory and never saved
	writer.Flush()

	// ── Step 10: Print the terminal summary ──
	// This goes to the screen, NOT to the file
	fmt.Println()
	fmt.Println("✦ Lines read    :", linesRead)
	fmt.Println("✦ Lines written :", len(processedLines))
	fmt.Println("✦ Lines removed :", linesRemoved)
	fmt.Println("✦ Rules applied : Trim whitespace | Replace TODO | ALL CAPS → Title Case | all lower → UPPER | Reverse words")
}
