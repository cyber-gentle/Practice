# String Transformer — Operation Gopher Protocol

**Author:** Abraham David  
**Squad:** The Gophers

---

## What This Program Does

SENTINEL's intercepted field reports were corrupted — wrong casing, broken formatting, unreadable text. This program fixes them. It takes a command and a string, applies the correct transformation, and prints the result. It keeps running until you type `exit`.

---

## How to Run It

```bash
go run . main.go
```

Then type commands at the prompt.

---

## The Commands

### `upper <text>`
Converts every letter to uppercase.
```
> upper sentinel is online
  → SENTINEL IS ONLINE
```

### `lower <text>`
Converts every letter to lowercase.
```
> lower ALERT LEVEL FIVE DETECTED
  → alert level five detected
```

### `cap <text>`
Capitalises the first letter of every word. All other letters go lowercase.
```
> cap director adaeze okonkwo
  → Director Adaeze Okonkwo

> cap THREAT LEVEL elevated
  → Threat Level Elevated
```

### `title <text>`
Like `cap`, but small connector words stay lowercase unless they are the first word.

Small words: `a, an, the, and, but, or, for, nor, on, at, to, by, in, of, up, as, is, it`
```
> title the fall of the western power grid
  → The Fall of the Western Power Grid

> title a storm is coming
  → A Storm is Coming
```

### `snake <text>`
Converts to snake_case. All lowercase, spaces replaced with underscores. Any character that is not a letter, digit, or underscore is removed.
```
> snake Operation Gopher Protocol!
  → operation_gopher_protocol

> snake Alert! Level 5 detected.
  → alert_level_5_detected
```

### `reverse <text>`
Reverses each word individually. Word order stays the same.
```
> reverse Lagos Nigeria
  → sogaL airegiN

> reverse Go is fun
  → oG si nuf
```

### `palindrome <text>`
Checks if the text reads the same forwards and backwards. Ignores spaces and casing.
```
> palindrome never odd or even
  → ✦ "never odd or even" is a palindrome!

> palindrome sentinel
  → ✗ "sentinel" is not a palindrome.
```

### `exit`
Shuts the program down cleanly.
```
> exit
  Shutting down String Transformer. Goodbye.
```

---

## Edge Cases Handled

| Situation | What happens |
|---|---|
| Unknown command | Prints error and lists valid commands |
| Command with no text | Prints usage hint for that command |
| Extra spaces in input | `strings.Fields` handles them automatically |
| Numbers and symbols in text | Pass through safely — only snake removes non-alphanumeric |
| Single word input | Works correctly for all commands |
| Only whitespace entered | Treated as empty — shows usage error |
| Mixed case commands (`UPPER`, `Upper`) | Lowercased before checking — all work |
| Empty line (just Enter) | Prompt appears again, no crash |
| `title` — first word always capitalised | Even small words like "a" or "the" get capitalised if first |
| `reverse` — single letter words | Single letter reversed is itself — correct |

---

## Example Full Session

```
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

 > palindrome never odd or even
   → ✦ "never odd or even" is a palindrome!

 > exit
   Shutting down String Transformer. Goodbye.
```