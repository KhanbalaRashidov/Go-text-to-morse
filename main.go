package main

import (
	"fmt"
	"regexp"
	"strings"
)

// Morse code dictionary
var morseCodeToText = map[string]string{
	".-": "A", "-...": "B", "-.-.": "C", "-..": "D", ".": "E", "..-.": "F",
	"--.": "G", "....": "H", "..": "I", ".---": "J", "-.-": "K", ".-..": "L",
	"--": "M", "-.": "N", "---": "O", ".--.": "P", "--.-": "Q", ".-.": "R",
	"...": "S", "-": "T", "..-": "U", "...-": "V", ".--": "W", "-..-": "X",
	"-.--": "Y", "--..": "Z",
	"-----": "0", ".----": "1", "..---": "2", "...--": "3", "....-": "4",
	".....": "5", "-....": "6", "--...": "7", "---..": "8", "----.": "9",
	".-.-.-": ".", "--..--": ",", "..--..": "?", ".----.": "'", "-.-.--": "!",
	"-..-.": "/", "-.--.": "(", "-.--.-": ")", ".-...": "&", "---...": ":",
	"-.-.-.": ";", "-...-": "=", ".-.-.": "+", "-....-": "-", "..--.-": "_",
	".-..-.": "\"", "...-..-": "$", ".--.-.": "@",
	"...---...": "SOS",
}

// Decode Morse code to text
func morseToText(morse string) string {
	words := strings.Split(morse, "   ") // 7 spaces between words
	var decodedWords []string

	for _, word := range words {
		var decodedWord []string
		for _, char := range strings.Split(word, " ") {
			if decodedChar, ok := morseCodeToText[char]; ok {
				decodedWord = append(decodedWord, decodedChar)
			}
		}
		decodedWords = append(decodedWords, strings.Join(decodedWord, ""))
	}

	return strings.Join(decodedWords, " ")
}

// Convert bit string to Morse code
func bitsToMorse(bits string) string {
	bits = strings.Trim(bits, "0")
	unitLength := findUnitLength(bits)

	bits = strings.ReplaceAll(bits, strings.Repeat("111", unitLength), "-")
	bits = strings.ReplaceAll(bits, strings.Repeat("1", unitLength), ".")
	bits = strings.ReplaceAll(bits, strings.Repeat("0000000", unitLength), "   ")
	bits = strings.ReplaceAll(bits, strings.Repeat("000", unitLength), " ")
	bits = strings.ReplaceAll(bits, strings.Repeat("0", unitLength), "")

	return bits
}

// Convert Morse code to bit string
func morseToBits(morse string) string {
	morse = strings.TrimSpace(morse)
	morse = strings.ReplaceAll(morse, "   ", "0000000") // 7 spaces between words
	morse = strings.ReplaceAll(morse, " ", "000")       // 3 spaces between characters
	morse = strings.ReplaceAll(morse, "-", "111")       // Dash
	morse = strings.ReplaceAll(morse, ".", "1")         // Dot

	return morse
}

// Find the length of the unit (1 or 0) in the bit string
func findUnitLength(bits string) int {
	ones := regexp.MustCompile(`1+`).FindAllString(bits, -1)
	zeros := regexp.MustCompile(`0+`).FindAllString(bits, -1)

	var lengths []int
	for _, s := range ones {
		lengths = append(lengths, len(s))
	}
	for _, s := range zeros {
		lengths = append(lengths, len(s))
	}

	minLength := lengths[0]
	for _, length := range lengths {
		if length < minLength {
			minLength = length
		}
	}

	return minLength
}

// Convert text to Morse code
func textToMorse(text string) string {
	var morseCode []string

	for _, char := range strings.ToUpper(text) {
		if char == ' ' {
			morseCode = append(morseCode, "   ")
			continue
		}
		for code, letter := range morseCodeToText {
			if letter == string(char) {
				morseCode = append(morseCode, code)
				break
			}
		}
	}

	return strings.Join(morseCode, " ")
}

// Test functions
func main() {
	// Test bit string
	bits := "1100110011001100000011000000111111001100111111001111110000000000000011001111110011111100111111000000110011001111110000001111110011001100000011"
	fmt.Println("Bits to Morse:", bitsToMorse(bits))

	morse := bitsToMorse(bits)
	fmt.Println("Morse to Text:", morseToText(morse))

	text := "HEY JUDE"
	morseFromText := textToMorse(text)
	fmt.Println("Text to Morse:", morseFromText)

	bitsFromMorse := morseToBits(morseFromText)
	fmt.Println("Morse to Bits:", bitsFromMorse)

	decodedText := morseToText(morseFromText)
	fmt.Println("Morse to Text:", decodedText)
}
