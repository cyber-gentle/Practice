OPERATION: GOPHER PROTOCOL 2026-03-28
A CodeCrafters Weekend Mission
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

CLASSIFICATION: WEEKEND ASSIGNMENT
STATUS: EARTH IS IN DANGER. STRINGS ARE THE WEAPON.



TRANSMISSION FROM HQ
  ─────────────────────────────────────────────────────

  CodeCrafters.

  SENTINEL intercepted a burst of corrupted
  intelligence reports this morning. Critical
  data — agent names, coordinates, threat
  descriptions, emergency directives — all
  mangled. Wrong casing. Broken formatting.
  Completely unreadable by our analysts.

  Our text processing unit is offline.

  You are going to rebuild it.

  One module. This weekend.
  Make strings bend to your will.

 Director X

BEFORE YOU WRITE A SINGLE LINE OF CODE
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
Non-negotiable. 

1. Completion of your cli-calculator and base-converter
is compulsory before commencing this mission.
 
2. Watch this:
  → Downfall of the 7-Hour Coding Tutorial 
  Lane Wagner on using AI the right way
  as a developer. If you cannot explain
  every line you submit, you did not write it.
  Director X will ask.

3. Review these W3Schools Go topics.
  Everything you need is in here:
  → Go Tutorial 

  Focus on:
  strings · functions · loops · slices
  variables · data types · if/else

REPOSITORY SETUP
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  Push to your the-codecrafters repo under the string-transformer directory:

       the-codecrafters/
       └──string-transformer/
             ├── main.go
             └── README.md

  Top of every .go file must have:

       // CodeCrafters — Operation Gopher Protocol
       // Module: String Transformer
       // Author: [Your Name]
       // Squad:  [Your Squad Name]

MODULE — THE STRING TRANSFORMER
  folder: module-string-transformer/

SITUATION:
  ─────────────────────────────────────────────────────

  SENTINEL's intercepted reports are a mess.
  Some are screaming in ALL CAPS. Some are
  whispering in all lowercase. Agent names
  are uncapitalized. Coordinates are written
  in snake_case when they should be readable.
  Words are reversed. Entire phrases are
  mangled beyond recognition.

  Your String Transformer will take any
  corrupted input and produce clean,
  correctly formatted intelligence.

  Six transformation commands. One program.
  Zero crashes.

  ─────────────────────────────────────────────────────
  THE SIX TRANSFORMATIONS:
  ─────────────────────────────────────────────────────

  Your program accepts a command followed by
  a string and applies the correct transformation.

  ┌─────────────────────────────────────────────────┐
  │  upper <text>                                   						│
  │  → Convert every letter to UPPERCASE            					│
  │                                                 							│
  │  Input : "sentinel is online"                   						│
  │  Output: "SENTINEL IS ONLINE"                   					│
  └─────────────────────────────────────────────────┘

  ┌─────────────────────────────────────────────────┐
  │  lower <text>                                   						│
  │  → Convert every letter to lowercase            					│
  │                                                 							│
  │  Input : "ALERT LEVEL FIVE DETECTED"            				│
  │  Output: "alert level five detected"            						│
  └─────────────────────────────────────────────────┘

  ┌─────────────────────────────────────────────────┐
  │  cap <text>                                     						│
  │  → Capitalise the first letter of every word.  					│
  │    All other letters go lowercase.              						│
  │                                                 							│
  │  Input : "director adaeze okonkwo"              					│
  │  Output: "Director Adaeze Okonkwo"              					│
  │                                                 							│
  │  Input : "THREAT LEVEL elevated"               					│
  │  Output: "Threat Level Elevated"                					│
  └─────────────────────────────────────────────────┘

  ┌─────────────────────────────────────────────────┐
  │  title <text>                                   							│
  │  → Title Case — like cap, but small connector  					│
  │    words stay lowercase unless they are the    					│
  │    first word.                                  							│
  │                                                 							│
  │  Small words: a, an, the, and, but, or,        					│
  │               for, nor, on, at, to, by, in,    						│
  │               of, up, as, is, it               							│
  │                                                 							│
  │  Input : "the fall of the western power grid"   					│
  │  Output: "The Fall of the Western Power Grid"   					│
  │                                                 							│
  │  Input : "a threat in the north"                						│
  │  Output: "A Threat in the North"                					│
  └─────────────────────────────────────────────────┘

  ┌─────────────────────────────────────────────────┐
  │  snake <text>                                   						│
  │  → Convert to snake_case.                       					│
  │    All lowercase, spaces replaced with _.       					│
  │    Remove any character that is not a letter,  					│
  │    digit, or underscore.                        						│
  │                                                 							│
  │  Input : "Operation Gopher Protocol"            					│
  │  Output: "operation_gopher_protocol"            					│
  │                                                 							│
  │  Input : "Alert! Level 5 detected."             						│
  │  Output: "alert_level_5_detected"               					│
  └─────────────────────────────────────────────────┘

 
  ┌─────────────────────────────────────────────────┐
  │  reverse <text>                                 						│
  │  → Reverse each word individually.             					│
  │    Word order stays the same.                   					│
  │    Spaces between words are preserved.          					│
  │                                                 							│
  │  Input : "Lagos Nigeria"                        						│
  │  Output: "sogaL airegiN"                        						│
  │                                                 							│
  │  Input : "Go is fun"                            						│
  │  Output: "oG si nuf"                            						│
  └─────────────────────────────────────────────────┘

  ─────────────────────────────────────────────────────
  HOW THE PROGRAM RUNS:
  ─────────────────────────────────────────────────────

  The program starts and shows a prompt.
  The user types a command and a string.
  The program prints the result.
  Then it waits for the next command.
  It keeps running until the user types: exit

  Example session:

  SENTINEL STRING TRANSFORMER — ONLINE
  ──────────────────────────────────────
  > upper sentinel is watching
    → SENTINEL IS WATCHING

  > cap director adaeze okonkwo
    → Director Adaeze Okonkwo

  > title the rise of the eastern threat
    → The Rise of the Eastern Threat

  > snake Operation Gopher Protocol!
    → operation_gopher_protocol

  > reverse Go saves the world
    → oG sevas eht dlrow

  > exit
    Shutting down String Transformer. Goodbye.

  ─────────────────────────────────────────────────────
  EDGE CASES YOU MUST HANDLE:
  ─────────────────────────────────────────────────────

  Every one of these must be handled cleanly.
  No panics. No silent failures. No blank output.

  1. Unknown command:
     The user types a command that is not one
     of the six. Tell them it is unrecognised
     and list the valid commands.

     > fly sentinel is down
       ✗ Unknown command: "fly"
         Valid commands: upper, lower, cap,
         title, snake, reverse, exit

  2. Missing text:
     The user types a command with nothing after it.

     > upper
       ✗ No text provided. Usage: upper <text>

  3. Extra spaces:
     "  upper   sentinel online  "
     Leading, trailing, and multiple internal
     spaces must not break the transformation
     or produce weird output.

  4. Numbers and symbols in the text:
     Your transformations must handle strings
     that contain digits and punctuation — they
     should pass through without crashing.

     > upper alert! level 5 breach at 03:47
       → ALERT! LEVEL 5 BREACH AT 03:47

  5. Single word input:
     > reverse Okonkwo
       → ownokO

 

 6. All spaces — only whitespace entered:
     > cap           
       ✗ No text provided. Usage: cap <text>

  7. Mixed case commands:
     > UPPER sentinel is watching
     > Upper sentinel is watching
     Both must work exactly like:
     > upper sentinel is watching

  8. Empty lines:
     User presses Enter with nothing typed.
     Do not crash. Do not print an error.
     Just show the prompt again.

  9. title — first word is always capitalised:
     Even if it is a small word like "a" or "the".

     > title a storm is coming
       → A Storm Is Coming

  10. reverse — a string with only one letter
     per word:
     > reverse a b c
       → a b c
     (Each single letter reversed is itself.)

  ─────────────────────────────────────────────────────
  THINK ABOUT:
  ─────────────────────────────────────────────────────

  → Each transformation must be its own function.
    If your entire program lives inside main(),
    it is not done.

  → For cap and title — how do you isolate the
    first letter of each word, change only that,
    and leave the rest of the letters alone?
    Think before you type.

  → For title — how do you store and check the
    list of small words efficiently?
    And how do you handle the first word rule
    cleanly — not as an afterthought?

  → For snake — what exactly counts as a character
    to remove? Define your rule clearly in your
    README before you code it.

  → For reverse — you are reversing each word,
    not the whole string. What does "word" mean
    when there are multiple spaces between words?

  → How do you split the user's input into a
    command and the rest of the text — especially
    when the text itself might contain spaces?

  ─────────────────────────────────────────────────────
  BONUS — IF YOU FINISH EARLY:
  ─────────────────────────────────────────────────────

  ✦ Add a history command that shows the last
    5 transformations the user ran — command,
    input, and output.

  ✦ Add a count <text> command that prints:
       - Total characters
       - Total letters only
       - Total words
       - Total spaces

  ✦ Add a palindrome <text> command that checks
    whether the input reads the same forwards
    and backwards (ignoring spaces and casing).

       > palindrome never odd or even
         → ✦ "never odd or even" is a palindrome!

       > palindrome sentinel
         → ✗ "sentinel" is not a palindrome.








━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
  RULES OF ENGAGEMENT
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

  → Standard library only. No external packages.

  → Every transformation must be its own function.

  → The program must never crash — no matter
    what the user types.

  → README.md must include:
       - What the program does
       - How to run it
       - One example interaction per command

  → ALL THE CODE MUST BE WRITTEN BY YOU.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━
DEADLINE
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

  Monday 2026-03-30 · 8:59 PM (WAT)

  Commits beyond 8:59 PM are out of range.
  SENTINEL does not wait. Neither does Earth.

  The corrupted reports are piling up.
  Analysts are standing by.
  Director X is watching the clock.

  Go fix the strings. Save the world.


