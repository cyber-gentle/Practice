2000 2026.03.30
CODECRAFTERTHON
     OPERATION: COMMAND & CONTROL

MISSION BRIEFING
  ─────────────────────────────────────────────────────

  Director X is back.
  And this time he is not asking nicely.

  SENTINEL's three subsystems — the calculator,
  the base converter, and the string transformer —
  were built as separate tools by separate teams.

  That was a mistake.

  In the field, analysts waste critical minutes
  switching between programs. They need one
  unified console. One prompt. Everything in
  one place. Fast, reliable, impossible to break.

  Your squad has one job today.

  Build SENTINEL's Command & Control Console —
  a single Go program that brings all three
  tools together under one roof, with a clean
  interface, bulletproof validation, and the
  kind of attention to detail that keeps
  people alive.

  Two hours and thirty minutes.
  One winner.
  No excuses.

  — Director X

  ─────────────────────────────────────────────────────


THE RULES OF ENGAGEMENT

  Read every word. Ignorance is not a defence.

  ─────────────────────────────────────────────────────
  THE NON-NEGOTIABLES:
  ─────────────────────────────────────────────────────

  1. ONE machine per squad. One coder types
    at a time. Everyone else navigates.
    Rotate the coder every 20 minutes —
    the instructor will call the switch.
    No rotation = disqualification.

  2. Standard library only. No external packages.
    Any squad caught importing a non-standard
    package loses 20 points. Immediately.

  3. The program must compile and run at
    submission time. A program that does not
    compile scores zero — no matter how much
    code is in it.

  4. You may use your own previous code as
    reference — but you may not copy-paste it
    directly. Every line written today must
    be written today. In this room.

  5. AI tools are not permitted during the
    hackathon. You watched Lane Wagner's video.
    You know why. Your brain is the tool today.

  6. When the buzzer sounds — hands off keyboards.
    Any squad still typing after the buzzer
    loses 15 points.

  7. Every squad must do a live demo at the end.
    If you cannot demo it, you did not build it.

 









 ─────────────────────────────────────────────────────
  THE SUBMISSION RULES:
  ─────────────────────────────────────────────────────

  → Push to your the-codecrafters repo under:

         the-codecrafters/
         └── command-and-control/
               └── main.go

NB: You can have as many files as you desire in the command-and-control directory.

  → One commit. Clean. Before the buzzer.
    Multiple last-minute commits look desperate.
    One commit looks deliberate.

  → Top of main.go must have:

         // CodeCrafters — Hackathon 002
         // Squad: [Squad Name]
         // Members: [All names]

  ─────────────────────────────────────────────────────
  THE SCORING SYSTEM:
  ─────────────────────────────────────────────────────

       Category                         Points
       ───────────────────────────────────────────
       Calculator — core features          15
       Base Converter — core features      15
       String Transformer — core features  15
       Integration & unified console       15
       Edge cases handled (per case)        2
       Code is clean and readable           5
       Live demo — no crashes               5
       Fastest squad to full completion    +10
       Second fastest                       +5
       Best error messages (voted)          +5
       ───────────────────────────────────────────
       Maximum possible                    90+

 




 ─────────────────────────────────────────────────────
  INSTANT DEDUCTIONS:
  ─────────────────────────────────────────────────────

       External package imported          -20
       Typing after buzzer                -15
       No rotation of coder               -15
       Program does not compile           -ALL
       Demo crashes on basic input        -10
WHAT YOU ARE BUILDING
━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━

  THE SENTINEL COMMAND & CONTROL CONSOLE

  One program. One prompt. Three tools inside it.

  When the program launches, the user sees this:

  ════════════════════════════════════════════════
    SENTINEL — COMMAND & CONTROL CONSOLE
       All systems nominal. Type 'help' to begin.
  ════════════════════════════════════════════════
  C&C>

  From this single prompt, the user can access
  all three tools using a module prefix:

       calc   <command>   → the calculator
       base   <command>   → the base converter
       str    <command>   → the string transformer
       help               → shows all commands
       history            → shows last 10 inputs
       exit               → shuts down the console

  ─────────────────────────────────────────────────────
  The console never drops back to a file or a
  separate program. Everything happens here.
  One prompt. One session. All three tools alive
  at the same time.
  ─────────────────────────────────────────────────────













 MODULE A — THE CALCULATOR
  Prefix: calc

  COMMANDS:
  ─────────────────────────────────────────────────────

  calc add <a> <b>      → addition
  calc sub <a> <b>      → subtraction
  calc mul <a> <b>      → multiplication
  calc div <a> <b>      → division
  calc mod <a> <b>      → remainder
  calc pow <a> <b>      → a to the power of b
  calc last             → show the last result
  calc history          → show last 5 calc results

  Example:
  C&C> calc add 42 58
       ✦ Result: 100

  C&C> calc pow 2 8
       ✦ Result: 256

  C&C> calc div 10 0
       ✗ Error: cannot divide by zero.

  ─────────────────────────────────────────────────────
  EDGE CASES — CALCULATOR:
  ─────────────────────────────────────────────────────

  → Division or mod by zero:
    clear error, no crash.

  → Non-numeric arguments:
    "calc add five 3" → clear error, no crash.

  → Missing arguments:
    "calc add 5" → tell them what is missing.

  → calc last with no previous result:
    "No previous result in this session."

  → Negative numbers must work:
    calc add -10 -5 → -15

MODULE B — THE BASE CONVERTER
  Prefix: base

  COMMANDS:
  ─────────────────────────────────────────────────────

  base dec <number>   → convert decimal to
                        binary AND hex

  base hex <number>   → convert hex to decimal

  base bin <number>   → convert binary to decimal

  Example:
  C&C> base dec 255
       ✦ Binary : 11111111
       ✦ Hex    : FF

  C&C> base hex 1F
       ✦ Decimal: 31

  C&C> base bin 1010
       ✦ Decimal: 10

  ─────────────────────────────────────────────────────
  EDGE CASES — BASE CONVERTER:
  ─────────────────────────────────────────────────────

  → Invalid hex characters:
    "base hex 1G" → "1G is not valid hex."

  → Invalid binary digits:
    "base bin 10201" → "10201 is not valid binary."

  → Negative decimal input:
    Handle it or reject it clearly.
    Whichever you choose — document it.

  → Missing number:
    "base hex" → tell them what is missing.

  → Letters where a number is expected:
    No crash. Clean error.

  → All hex output must be uppercase.


MODULE C — THE STRING TRANSFORMER
  Prefix: str

  COMMANDS:
  ─────────────────────────────────────────────────────

  str upper  <text>    → ALL UPPERCASE
  str lower  <text>    → all lowercase
  str cap    <text>    → Every Word Capitalised
  str title  <text>    → Title Case With Small Words
  str snake  <text>    → convert_to_snake_case
  str reverse <text>   → esreveR hcaE droW

  Example:
  C&C> str upper sentinel is watching
       ✦ SENTINEL IS WATCHING

  C&C> str title the fall of the eastern grid
       ✦ The Fall of the Eastern Grid

  C&C> str snake Operation Gopher Protocol!
       ✦ operation_gopher_protocol

  ─────────────────────────────────────────────────────
  EDGE CASES — STRING TRANSFORMER:
  ─────────────────────────────────────────────────────

  → Command with no text after it:
    "str upper" → "No text provided."

  → Extra spaces in the input must not
    produce extra spaces in the output.

  → Numbers and symbols pass through cleanly.

  → Commands are case-insensitive:
    "str UPPER text" works like "str upper text"

  → Empty input after trimming:
    Show the prompt again. No crash.

THE INTEGRATION REQUIREMENTS
  (This is where squads separate from each other)

  Building three tools that each work is good.
  Building three tools that work together
  is what earns the integration points.

  Your console must support these cross-module
  features:

  ─────────────────────────────────────────────────────
  1. PIPE — pass output from one module into another:
  ─────────────────────────────────────────────────────

  Use the | symbol to pipe the output of one
  command as the input to another.

  C&C> calc add 200 55 | base dec
       ✦ Result : 255
       ✦ Binary : 11111111
       ✦ Hex    : FF

  C&C> base hex FF | calc add last 1
       ✦ Decimal: 255
       ✦ Result : 256

  C&C> calc mul 3 4 | str upper result is last
       ✦ Result : 12
       ✦ RESULT IS 12

  The pipe takes the numeric result or text
  output of the left command and makes it
  available to the right command automatically.

  ─────────────────────────────────────────────────────
  2. LAST — shared across modules:
  ─────────────────────────────────────────────────────

  The keyword last always refers to the most
  recent numeric result — whether it came from
  the calculator or the base converter.

  C&C> base hex 1F
       ✦ Decimal: 31

  C&C> calc mul last 2
       ✦ Result: 62

  C&C> calc pow last 2
       ✦ Result: 3844

  ─────────────────────────────────────────────────────
  3. HISTORY — shared session log:
  ─────────────────────────────────────────────────────

  Typing history at the C&C> prompt shows
  the last 10 commands entered across all
  three modules — not just one.

  C&C> history
       1.  calc add 42 58          → 100
       2.  base dec 255            → 11111111 / FF
       3.  str upper sentinel      → SENTINEL
       4.  calc pow 2 8            → 256
       5.  base hex 1F             → 31
       ...

  ─────────────────────────────────────────────────────
  4. CLEAR — resets the session:
  ─────────────────────────────────────────────────────

  C&C> clear

  Clears the shared history and resets last
  to zero. Prompts for confirmation first:

  "Type CONFIRM to clear the session: "

  Any other input cancels the clear.




