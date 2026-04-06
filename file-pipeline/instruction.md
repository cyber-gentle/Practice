OPERATION: GOPHER PROTOCOL
     Mission - The File Pipeline

  CLASSIFICATION: ACTIVE MISSION
  MODE: SQUAD PLANNING → SOLO BUILD → SQUAD ASSEMBLY
  STATUS: THE ARCHIVE IS BURNING.



 TRANSMISSION FROM DIRECTOR X (HQ)
  ─────────────────────────────────────────────────────

  CodeCrafters.

  SENTINEL's intelligence archive holds
  thousands of raw field reports. Text files.
  Messy ones. Logged by agents under pressure
  in the dark with no formatting guidelines
  and no time to be careful.

  We need them cleaned. Transformed. Numbered.
  Categorised. Written back to disk in a format
  our analysts can actually use.

  But here is the problem.

  No single agent can process all of them alone.
  The archive is too large. The rules are too
  specific. And every file that comes out of
  your pipeline feeds directly into SENTINEL's
  next subsystem.

  So you will plan together.
  You will build alone.
  And at the end — your pieces will be
  assembled into one working pipeline.

  If one piece is broken, the whole thing fails.
  Act accordingly.

  — Director X

  

PREREQUISITES

  You must have completed and pushed:
  → Mission 01 — The String Transformer
  → Mission 02 — The Word Surgeon

  If either is missing — finish them first.
  Director X does not run incomplete
  pipelines.


REPOSITORY SETUP

  Push to your the-codecrafters repo under:

       the-codecrafters/
       └── file-pipeline/
             ├── main.go
             ├── input.txt         ← you create this
             ├── output.txt        ← your program writes this
             └── README.md

  Top of every .go file:

       // CodeCrafters — Operation Gopher Protocol
       // Module: File Pipeline
       // Author: [Your Name]
       // Squad:  [Your Squad Name]
















HOW THIS MISSION WORKS
  This mission has three phases.
  Do not skip phases. Do not reverse them.
  The order is the mission.

  ─────────────────────────────────────────────────────
  PHASE 1 — SQUAD PLANNING SESSION  (30 minutes)
  ─────────────────────────────────────────────────────

  Before anyone writes a single line of Go,
  your entire squad sits together and makes
  four decisions as a group.

  These decisions become your squad's
  Pipeline Contract — a shared agreement that
  every member's program must honour.

  ┌─────────────────────────────────────────────────┐
  │  DECISION 1 — THE INPUT FORMAT                  				│
  │                                                 							│
  │  What does your squad's raw input.txt look      					│
  │  like? Agree on:                                						│
  │                                                 							│
  │  → Minimum number of lines (at least 15)        					│
  │  → What kinds of lines exist — you must         					│
  │    include at least four of these:              						│
  │                                                 							│
  │      Normal report lines                        						│
  │      Lines in ALL CAPS                          						│
  │      Lines in all lowercase                     						│
  │      Lines starting with TODO:                  						│
  │      Lines starting with CLASSIFIED:            					│
  │      Lines that are only dashes or blanks       					│
  │      Lines with extra leading/trailing spaces   					│
  │      Lines containing numbers and symbols       					│
  │                                                 							│
  │  Every member writes their OWN input.txt        					│
  │  but it must contain the same line TYPES        					│
  │  your squad agreed on. The content can differ.  					│
  └─────────────────────────────────────────────────┘

 
  ┌─────────────────────────────────────────────────┐
  │  DECISION 2 — THE TRANSFORMATION RULES          			│
  │                                                 							│
  │  Agree on exactly 5 transformation rules        					│
  │  your pipeline must apply to every line.        					│
  │  Choose 5 from this list:                      				 		│
  │                                                 							│
  │  1. Trim all leading and trailing whitespace     					│
  │  2. Replace TODO: with ✦ ACTION:                 					│
  │  3. Replace CLASSIFIED: with [REDACTED]:         				│
  │  4. Convert ALL CAPS lines to Title Case         					│
  │  5. Convert all lowercase lines to uppercase    					│
  │  6. Remove lines that are only dashes or blanks 				│
  │  7. Add a line number prefix to every           					│
  │     remaining line: "001.", "002." etc.         						│
  │  8. Append a timestamp to the end of every      					│
  │     CLASSIFIED: line before redacting it       					│
  │  9. Reverse the words in any line that          					│
  │     contains the word REVERSE                  					│
  │  10. Flag any line longer than 80 characters     					│
  │     with [TRUNCATED] at the end                					│
  │                                                 							│
  │  Every member implements the SAME 5 rules.     				│
  │  This is non-negotiable. If your output        					│
  │  does not match your squadmate's output        					│
  │  on the same input — someone got it wrong.     					│
  └─────────────────────────────────────────────────┘

  ┌─────────────────────────────────────────────────┐
  │  DECISION 3 — THE OUTPUT FORMAT                 				│
  │                                                 							│
  │  Agree on exactly what your output.txt          					│
  │  must look like. Specifically:                  						│
  │                                                 							│
  │  → Does the file start with a header?           					│
  │    e.g. "SENTINEL FIELD REPORT — PROCESSED"    			│
  │    If yes — what does it say exactly?           					│
  │                                                 							│
  │  → Does the file end with a summary block?      					│
  │    If yes — what does it include?               					│
  │    (Lines read, lines written, lines removed,   					│
  │    rules applied, timestamp of processing)      					│
  │                                                 							│
  │  → How are line numbers formatted?              					│
  │    "1." or "01." or "001." — agree on one.     					│
  │                                                 							│
  │  Every member's output.txt must follow          					│
  │  the same format. When assembled, they          					│
  │  should be indistinguishable in structure.     					│
  └─────────────────────────────────────────────────┘

  ┌─────────────────────────────────────────────────┐
  │  DECISION 4 — THE SUMMARY BLOCK                 				│
  │                                                 							│
  │  At the end of every output.txt your program   					│
  │  must print a summary to the terminal           					│
  │  (not to the file) when it finishes running.   					│
  │                                                 							│
  │  Agree as a squad on exactly what it shows:    					│
  │                                                 							│
  │  → At minimum, these four lines:               					│
  │                                                 							│
  │    ✦ Lines read    : [number]                  						│
  │    ✦ Lines written : [number]                  						│
  │    ✦ Lines removed : [number]                  					│
  │    ✦ Rules applied : [list your 5 rules]       						│
  │                                                 							│
  │  → Any additional stats your squad decides     					│
  │    to include must be included by everyone.    					│
  └─────────────────────────────────────────────────┘

  ─────────────────────────────────────────────────────
  BEFORE PHASE 1 ENDS:
  ─────────────────────────────────────────────────────

  Your squad lead writes the Pipeline Contract
  as a single comment block at the top of
  every member's main.go — before anyone
  starts coding.

 
 It looks like this:

  // ═══════════════════════════════════════════
  // SQUAD PIPELINE CONTRACT
  // Squad: [Squad Name]
  // ───────────────────────────────────────────
  // Input line types:
  //   [list what your squad agreed on]
  //
  // Transformation rules (in order):
  //   1. [Rule 1]
  //   2. [Rule 2]
  //   3. [Rule 3]
  //   4. [Rule 4]
  //   5. [Rule 5]
  //
  // Output format:
  //   [Header: yes/no — exact text if yes]
  //   [Line numbering format]
  //   [Summary block: yes/no — fields if yes]
  //
  // Terminal summary fields:
  //   [List what your squad agreed on]
  // ═══════════════════════════════════════════

  If any member's code does not match this
  contract — they broke the pipeline.

  ─────────────────────────────────────────────────────
  PHASE 2 — SOLO BUILD  (main task)
  ─────────────────────────────────────────────────────

  Everyone builds alone. Same contract.
  Different files. Different code.

  ─────────────────────────────────────────────────────
  PHASE 3 — SQUAD ASSEMBLY  (final step)
  ─────────────────────────────────────────────────────

  Squads test each other's programs.
  Details below.



PHASE 2 — SOLO BUILD
     What every fellow must build individually

  SITUATION:
  ─────────────────────────────────────────────────────

  SENTINEL's archive processor takes a raw
  field report as input, runs it through your
  squad's agreed transformation pipeline, and
  writes a clean processed report to disk.

  It runs like this:

       go run . input.txt output.txt

  ─────────────────────────────────────────────────────
  WHAT TO BUILD:
  ─────────────────────────────────────────────────────

  1. CREATE YOUR OWN input.txt
    Write a raw field report of at least 15 lines.
    It must contain all the line types your squad
    agreed on in Phase 1. Be creative — write it
    like a real field agent's messy notes.

    Example of what a raw field report looks
    like inside SENTINEL's archive:

    ──────────────────────────────────────────────
    ALERT LEVEL FIVE DETECTED IN NORTHERN SECTOR
    todo: confirm coordinates with Agent Bulus
    classified: target has been relocated
    -----------------------------------------------
    the perimeter was breached at 03:47 local time
       extra spaces were not cleaned before logging
    CLASSIFIED: secondary site compromised
    normal report line with no issues
    todo: send backup to eastern checkpoint
    EVERYTHING WAS FINE UNTIL IT WASN'T
    another normal line from the field
    classified: the director has been briefed
    SHORT LINE
    this line contains the word REVERSE in it so it gets special treatment
    a very long line that goes on and on past eighty characters and should be flagged accordingly by your pipeline when it encounters it during processing
    ──────────────────────────────────────────────

  2. READ THE INPUT FILE
    Your program takes the input filename as
    the first argument:

       go run . input.txt output.txt

    → If input.txt does not exist, print a
      clear error and exit immediately.
      Do not create the file. Do not guess.

    → If no arguments are provided, print
      usage instructions and exit:
      "Usage: go run . <input.txt> <output.txt>"

    → If only one argument is provided,
      same — print usage and exit cleanly.

  3. APPLY YOUR SQUAD'S 5 TRANSFORMATION RULES
    Process every line through all 5 rules
    in the exact order your squad agreed on
    in the Pipeline Contract.

    Each rule must be its own function.
    No exceptions. Your pipeline is a chain
    of function calls — not a wall of if-else
    inside main().

  4. WRITE THE OUTPUT FILE
    Write the processed lines to output.txt.

    → If output.txt already exists — overwrite it.
      No prompt. No warning. Just overwrite.

    → If the output file cannot be written
      (permissions, bad path) — print a clear
      error and exit. Do not silently fail.

    → The output must match your squad's
      agreed format exactly:
      header (if agreed), numbered lines,
      summary block (if agreed).

  5. PRINT THE TERMINAL SUMMARY
    After writing the file, print your squad's
    agreed summary block to the terminal.
    Not to the file. To the terminal.

    At minimum:
    ✦ Lines read    : 15
    ✦ Lines written : 11
    ✦ Lines removed : 4
    ✦ Rules applied : [your 5 rule names]

  ─────────────────────────────────────────────────────
  EDGE CASES YOU MUST HANDLE:
  ─────────────────────────────────────────────────────

  These are individual. Every fellow handles
  all of them — regardless of squad contract.

  1. Input file does not exist:
    → Clear error. Exit. Do not crash.
    → "✗ File not found: input.txt"

  2.  Input file exists but is completely empty:
    → Do not crash. Write an empty output.txt.
    → Terminal summary shows: Lines read: 0
    → Print a warning:
      "⚠ Input file is empty. Nothing to process."

  3.  Input file has only blank lines or dashes:
    → All lines get removed by the pipeline.
    → Output file is written but contains
      only the header and summary (if agreed).
    → Terminal summary reflects 0 lines written.

  4.  A line that triggers multiple rules
    at the same time:
    e.g. a line that is ALL CAPS AND starts
    with TODO: AND is longer than 80 characters.
    → All applicable rules must fire.
    → The order matters — apply them in the
      exact order of your contract.
    → Document in your README what happens
      when rules overlap.

  5. Output file path is a directory, not a file:
    → "✗ Cannot write to output: path is a
         directory, not a file."
    → Exit cleanly.

  6. The program is run with the same file
    as both input and output:
    go run . report.txt report.txt
    → Detect this. Reject it with a clear
      message before any processing begins.
    → "✗ Input and output cannot be the
         same file."

  7. A line that is just whitespace:
    → After trimming — it becomes empty.
    → Treat it the same as a blank line.
    → Remove it. Count it as removed.

  8. The input file has Windows-style line
    endings (\r\n):
    → Your program must handle both \n and
      \r\n without producing garbled output
      or extra blank lines.

  ─────────────────────────────────────────────────────
  THINK ABOUT:
  ─────────────────────────────────────────────────────

  → How do you read a file line by line without
    loading the entire thing into memory at once?
    What is the difference between os.ReadFile
    and bufio.Scanner — and which is better here?

  → Your transformations run in a specific order.
    If rule 3 changes a line that rule 5 also
    targets — does the order affect the result?
    Test it. Know it. Document it.

  → How do you count lines removed vs lines written?
    Where in your program do you track this?
    Is it inside the transformation functions
    or outside them?

  → When you overwrite output.txt — what mode
    do you open the file in? What happens if
    you get it wrong?

  → You are printing a summary to the terminal
    AND writing processed content to a file.
    These are two different destinations.
    Make sure you never mix them up.



 PHASE 3 — SQUAD ASSEMBLY
     When every member's build is pushed

  SITUATION:
  ─────────────────────────────────────────────────────

  Every member has built their own pipeline.
  Now SENTINEL needs to know they all
  speak the same language.

  ─────────────────────────────────────────────────────
  THE ASSEMBLY TEST:
  ─────────────────────────────────────────────────────

  1. One squad member writes a single shared
    test file — squad_test_input.txt —
    with exactly 20 lines covering every
    line type in your contract.
    This file is the same for every member.
    Push it to the squad lead's repo.

  2. Every member runs their own program
    against this shared input:

       go run . squad_test_input.txt \
                squad_test_output.txt

  3. Compare the output files.
    Every member's squad_test_output.txt
    must be identical in structure and content.

    If two members produce different output
    on the same input — one of them broke
    the contract. Find it. Fix it.

  ─────────────────────────────────────────────────────
  THE ASSEMBLY CHECKLIST:
  ─────────────────────────────────────────────────────

  Before your squad calls Phase 3 complete,
  every member answers YES to all of these:

  □ My program runs without errors on
    squad_test_input.txt

  □ My output.txt matches at least 4 out of 5
    other members' outputs exactly

  □ My terminal summary shows the correct
    counts for the shared test input

  □ My program handles all 8 edge cases
    without crashing

  □ Every transformation is its own function
    — I can point to it in my code

  □ My README explains my 5 transformation
    rules and documents what happens when
    they overlap

  Plan carefully in Phase 1.
  Code carefully in Phase 2.
  Test honestly in Phase 3.




RULES OF ENGAGEMENT
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

  → Standard library only. No external packages.

  → Every transformation rule is its own function.
    main() reads, calls, writes. That is all.

  → The program must never crash — on any input.

  → The Pipeline Contract comment must be at
    the top of every member's main.go — identical
    across the squad. If contracts differ between
    members, Phase 1 was not done properly.

  → README.md must include:
       - Your squad's 5 chosen transformation rules
       - The output format your squad agreed on
       - How to run the program
       - What happens when rules overlap
       - Two real example runs with input and output

  → You can explain every line you wrote.
    If you cannot — it does not count.


 DEADLINE
  Phase 1 — Squad planning:  Before you code.
  Phase 2 - 3 — Solo build & Assembly test:      Tuesday 2026.03.31 · 9:59 PM (WAT)

  The Pipeline Contract must be in your code
  before your first commit.
  If your first commit has no contract —
  you skipped Phase 1.
  Director X will know.


  The archive is waiting.
  The analysts are waiting.
  SENTINEL is waiting.


