# File Pipeline — Operation Gopher Protocol

**Author:** Abraham David  
**Squad:** The Gophers  
**Module:** File Pipeline

---

## What This Program Does

This program reads a raw messy field report (input.txt), cleans it up using 5 transformation rules, and writes the processed version to output.txt. It also prints a summary to the terminal when it's done.

---

## How to Run It

```bash
go run . input.txt output.txt
```

- First argument = the file to read  
- Second argument = the file to write to  

---

## The 5 Transformation Rules (in order)

These rules are applied to every line, one after the other, in this exact order:

**Rule 1 — Trim whitespace**  
Removes any extra spaces at the beginning or end of a line.  
`"   hello world   "` → `"hello world"`

**Rule 2 — Replace TODO:**  
If a line starts with `TODO:`, it gets replaced with `✦ ACTION:`.  
`"TODO: send backup"` → `"✦ ACTION: send backup"`

**Rule 3 — ALL CAPS → Title Case**  
If every letter in the line is uppercase, the line is converted to Title Case.  
`"ALERT LEVEL FIVE"` → `"Alert Level Five"`

**Rule 4 — all lowercase → UPPERCASE**  
If every letter in the line is lowercase, the whole line becomes uppercase.  
`"the perimeter was breached"` → `"THE PERIMETER WAS BREACHED"`

**Rule 5 — Reverse words**  
If the line contains the word `REVERSE`, the order of all words in the line is flipped.  
`"send backup REVERSE now"` → `"now REVERSE backup send"`

---

## What Happens When Rules Overlap

Rules are applied in order (1 → 2 → 3 → 4 → 5). This means:

- A line that is **ALL CAPS and starts with TODO:** — Rule 2 fires first (replaces TODO:), then Rule 3 checks if the result is still all caps. Since `✦ ACTION:` is not all caps, Rule 3 won't fire. The order matters.

- A line that is **all lowercase and contains REVERSE** — Rule 4 fires first, making the line all uppercase. By the time Rule 5 runs, the line is uppercase but still contains the word `REVERSE`, so Rule 5 also fires and reverses the words.

- A line with **only spaces** — Rule 1 trims it to an empty string. The pipeline then removes it and counts it as a removed line. Rules 2–5 never see it.

The key principle: **each rule sees the output of the rule before it, not the original line.**

---

## Output Format

The output file looks like this:

```
Gopher's Sentinel Field Report - Processed
───────────────────────────────────────────
1. Alert Level Five Detected In Northern Sector
2. ✦ ACTION: confirm coordinates with Agent Bulus
3. THE PERIMETER WAS BREACHED AT 03:47 LOCAL TIME
...
───────────────────────────────────────────
SUMMARY
✦ Lines read    : 20
✦ Lines written : 18
✦ Lines removed : 2
✦ Rules applied : Trim whitespace | Replace TODO | ALL CAPS → Title Case | all lower → UPPER | Reverse words
```

---

## Terminal Summary

After writing the output file, the program prints this to the screen:

```
✦ Lines read    : 20
✦ Lines written : 18
✦ Lines removed : 2
✦ Rules applied : Trim whitespace | Replace TODO | ALL CAPS → Title Case | all lower → UPPER | Reverse words
```

---

## Edge Cases Handled

| Situation | What happens |
|---|---|
| Input file doesn't exist | Prints error and exits cleanly |
| Input file is empty | Prints warning, writes empty output, shows zero counts |
| Input and output are the same file | Detects it and rejects before doing anything |
| Output path is a directory | Detects it and prints a clear error |
| Lines with only whitespace | Trimmed to empty, then removed and counted |
| Windows line endings (\r\n) | bufio.Scanner handles these automatically |
| No arguments given | Prints usage instructions and exits |
| Only one argument given | Prints usage instructions and exits |

---

## Example Run 1

**Input (input.txt):**
```
ALERT LEVEL FIVE DETECTED IN NORTHERN SECTOR
todo: confirm coordinates with Agent Bulus
the perimeter was breached at 03:47 local time
```

**Output (output.txt):**
```
Gopher's Sentinel Field Report - Processed
───────────────────────────────────────────
1. Alert Level Five Detected In Northern Sector
2. ✦ ACTION: confirm coordinates with Agent Bulus
3. THE PERIMETER WAS BREACHED AT 03:47 LOCAL TIME
───────────────────────────────────────────
SUMMARY
✦ Lines read    : 3
✦ Lines written : 3
✦ Lines removed : 0
✦ Rules applied : Trim whitespace | Replace TODO | ALL CAPS → Title Case | all lower → UPPER | Reverse words
```

**Terminal:**
```
✦ Lines read    : 3
✦ Lines written : 3
✦ Lines removed : 0
✦ Rules applied : Trim whitespace | Replace TODO | ALL CAPS → Title Case | all lower → UPPER | Reverse words
```

---

## Example Run 2

**Input (input.txt):**
```
   extra spaces on both ends   
this line contains the word REVERSE in it
```

**Output (output.txt):**
```
Gopher's Sentinel Field Report - Processed
───────────────────────────────────────────
1. EXTRA SPACES ON BOTH ENDS
2. it in REVERSE word the contains line this
───────────────────────────────────────────
SUMMARY
✦ Lines read    : 2
✦ Lines written : 2
✦ Lines removed : 0
✦ Rules applied : Trim whitespace | Replace TODO | ALL CAPS → Title Case | all lower → UPPER | Reverse words
```

**Terminal:**
```
✦ Lines read    : 2
✦ Lines written : 2
✦ Lines removed : 0
✦ Rules applied : Trim whitespace | Replace TODO | ALL CAPS → Title Case | all lower → UPPER | Reverse words
```