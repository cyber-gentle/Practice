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

