package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func start() (string, error) {

	args := os.Args[1:]

	if len(args) < 1 {
		return "", errors.New("please enter text, example: 'go run . \"Hello\"")
	}

	return args[0], nil

}

func readAscii() []string {

	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println(err)
	}
	str0 := ""
	arrSplit := strings.Split(string(file), "\n\n")
	for i, v := range arrSplit[0] {
		if i > 0 {
			str0 += string(v)
		}

	}
	arrSplit[0] = str0

	return arrSplit

}

func printAsciiArt(txt, arrSplit []string) {
	if txt[0] == "" {
		fmt.Println()
		return
	}

	for _, word := range txt {
		if word == "" {
			fmt.Println()
			continue
		}
		for k := 0; k < 8; k++ {
			for _, s := range word {

				strCut := strings.Split(arrSplit[s-32], "\n")
				for _, ascii := range strCut[k] {

					fmt.Print(string(ascii))

				}

			}

			fmt.Println()
		}

	}

}

func checkTxt(str string) error {

	noChar := ""
	for _, s := range str {
		if s < 0 || s > 127 {
			noChar += string(s)
		}

	}
	if noChar != "" {
		return errors.New(fmt.Sprintf("a character \"%v\" is not available.", noChar))
	}
	return nil
}

func main() {

	txt, err := start()
	if err != nil {
		log.Fatal(err)
	}
	if txt == "" {
		return
	}

	err = checkTxt(txt)
	if err != nil {
		fmt.Println(err)
		return
	}

	arrTxt := strings.Split(txt, "\\n")

	arrSplit := readAscii()
	printAsciiArt(arrTxt, arrSplit)

}
