package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strconv"
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

	file, err := os.ReadFile("banners/thinkertoy.txt")
	if err != nil {
		fmt.Println(err)
	}
	str0 := ""
	file = []byte(strings.ReplaceAll(string(file), "\r", ""))
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

		if word == "\n" {
			fmt.Println()

			continue
		}
		firstLine := true
		for k := 0; k < 8; k++ {
			str := ""

			for _, s := range word {

				strCut := strings.Split(arrSplit[s-32], "\n")
				for _, ascii := range strCut[k] {

					str += string(ascii)

				}

			}
			if firstLine {
				firstLine = false
				fitConsole(str)

			}

			fmt.Print(str)

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
func check(errMsg string, err error) {
	if err != nil {
		fmt.Println(errMsg, err)
		os.Exit(1)
	}
}

func fitConsole(s string) {

	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	check("Error measuring console size:", err)

	outStr := string(out)
	outStr = strings.TrimSpace(outStr)
	heightWidth := strings.Split(outStr, " ")
	width, err := strconv.Atoi(heightWidth[1])
	check("Error measuring console size:", err)

	if len(s) > width {
		fmt.Println("The input string doesn't fit into terminal.")
		os.Exit(1)
	}
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
	fmt.Println(txt)

	arrTxt := strings.Split(txt, "\\n")

	arrSplit := readAscii()
	printAsciiArt(arrTxt, arrSplit)

}
