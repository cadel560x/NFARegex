package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"

	"./regex"
	// "./tests"
)

func main() {
	// tests.Tests()

	fmt.Println()
	fmt.Println("        Regex to NFA")
	fmt.Println("        ------------")
	fmt.Println()
	fmt.Println("Special symbols:")
	fmt.Println("(    start of subgroup")
	fmt.Println(")    end of subgroup")
	fmt.Println("*    zero or more occurences")
	fmt.Println("?    zero or one occurence")
	fmt.Println(".    concatenation")
	fmt.Println("|    logical OR operator")
	
	fmt.Print("Please enter a regular expression, e.g 'ab|c*': ")
	regexString, err := userInput()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Print("Please string to validate, e.g 'ab': ")
	language, err := userInput()
	if err != nil {
		fmt.Println("Error: ", err.Error())
		return
	}

	fmt.Println()

	status := " not"
	if regex.RegexNFA(regexString, language) {
		status = ""
	}

	msg := "String '" + language + "'" + status + " validated"
	fmt.Println(msg)

	fmt.Println()

} // end main

// User input function based on
// https://medium.com/@matryer/golang-advent-calendar-day-seventeen-io-reader-in-depth-6f744bb4320b
func userInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	s, err := reader.ReadString('\n')

	return strings.TrimSpace(s), err
}