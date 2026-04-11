package main

import (
	"fmt"
	"os"
	"reloaded/gen2"
	"strconv"
	"strings"
)

func readFile(rawFile string) ([]byte, error) {
	data, err := os.ReadFile(rawFile)
	if err != nil {
		return nil, fmt.Errorf("read file %q: %w", rawFile, err)
	}
	return data, nil
}

func writeFile(outputFileName string, data []byte, permission os.FileMode) error {
	err := os.WriteFile(outputFileName, []byte(data), permission)
	if err != nil {
		fmt.Printf(" %v\n", err)
		return fmt.Errorf("file could not be written: %w", err)
	}
	return err
}

func convertBase(data string) (string, error, error) {
	word := strings.Fields(data)

	for i := 0; i < len(word); i++ {
		if word[i] == "(hex)" || word[i] == "(bin)" {
			switch word[i] {
			case "(hex)":
				decimalNumber, err := strconv.ParseInt(word[i-1], 16, 64)
				if err != nil {
					return "", fmt.Errorf("Conversion failed: %q is not a valid HexaDecimal Number\n", word[i-1]), nil
				}
				word[i-1] = strconv.Itoa(int(decimalNumber))
				word = append(word[:i], word[i+1:]...)

			case "(bin)":
				binaryNumber, err := strconv.ParseInt(word[i-1], 2, 64)
				if err != nil {
					return "", nil, fmt.Errorf("Conversion failed: %q is not a valid Binary Number\n", word[i-1])
				}
				word[i-1] = strconv.Itoa(int(binaryNumber))
				word = append(word[:i], word[i+1:]...)

			}
		}
	}
	data = strings.Join(word, " ")
	return data, nil, nil
}

func transformLastWord(data string) string {

	word := strings.Fields(data)
	for i := 0; i < len(word); i++ {
		switch word[i] {
		case "(up)":
			lastWord := gen2.UpperCaseLastWord(word[i-1 : i+1])
			word[i-1] = lastWord
			word = append(word[:i], word[i+1:]...)

		case "(low)":
			lastWord := gen2.LowerCaseLastWord(word[i-1 : i+1])
			word[i-1] = lastWord
			word = append(word[:i], word[i+1:]...)

		case "(cap)":
			lastWord := gen2.CapitaliseLastWord(word[i-1 : i+1])
			word[i-1] = lastWord
			word = append(word[:i], word[i+1:]...)
		}
	}
	data = strings.Join(word, " ")
	return data
}

func transformLastWordN(data string) string {

	word := strings.Fields(data)
	for i := 0; i < len(word); i++ {
		switch word[i] {
		case "(up,":
			N, _ := strconv.Atoi(string(word[i+1][0]))
			lastWordN := gen2.UpperCaseLastWordN(word[i-N : i])
			word1 := append(word[:i-N], lastWordN)
			fullWord := append(word1, word[i+2:]...)
			word = fullWord

		case "(low,":
			N, _ := strconv.Atoi(string(word[i+1][0]))
			lastWordN := gen2.LowerCaseLastWordN(word[i-N : i])
			word1 := append(word[:i-N], lastWordN)
			fullWord := append(word1, word[i+2:]...)
			word = fullWord

		case "(cap,":
			N, _ := strconv.Atoi(string(word[i+1][0]))
			lastWordN := gen2.CapitaliseLastWordN(word[i-N : i])
			word1 := append(word[:i-N], lastWordN)
			fullWord := append(word1, word[i+2:]...)
			word = fullWord

		}
	}
	data = strings.Join(word, " ")
	return data
}

// func fixPunctuation(data string) string {
// 	var count int = 0
// 	var word string

// 	words := strings.Fields(data)
// 	for _, word = range words {
// 		if strings.ContainsAny(word, "., ,, !, ?, :, ;") {
// 			strings.Replace(word, " ,", ",", count+1)
// 		}
// 	}
// 	word +=  word
// 	return word
// }

func applyTransformation(s []byte) string {
	data := string(s)

	data, err1, err2 := convertBase(data)
	if err1 != nil && err2 != nil {
		fmt.Println(err1, err2)
	} else if err1 != nil && err2 == nil {
		fmt.Println(err1)
	} else if err2 != nil && err1 == nil {
		fmt.Println(err2)
	}

	data = transformLastWord(data)
	data = transformLastWordN(data)
	//	data = fixPunctuation(data)

	return data
}

func main() {
	if len(os.Args) != 3 {
		fmt.Print(" Usage: go run . <name of file to read> <name to save the processed file.>\n go run . sample.doc result.docx \n \n")
		os.Exit(1)
	}

	rawFile := os.Args[1]
	outputFileName := os.Args[2]

	data, err := readFile(rawFile) // called my readFile to collect the pre-processed file.
	if err != nil {
		fmt.Println(err) //If the function readFile returns an error, print it to the screen.
	}

	processedData := []byte(applyTransformation(data)) //

	err = writeFile(outputFileName, []byte(processedData), 0664) // called my WriteFile to collect the name the written file should be saved with; collect the data to write and permission for the file
	if err != nil {
		fmt.Println(err)
	}

	// ---------------------------------------------------------------------------//
	// data, err := os.ReadFile(rawFile)
	// 	if err != nil {
	// 		fmt.Errorf("file could not be read: %w", err)
	// 	}

	// err = os.WriteFile(outputFileName, []byte(data), 0777)
	// ---------------------------------------------------------------------------//

}
