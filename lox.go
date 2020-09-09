package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Lox struct {
	hadError bool
}

func (lox *Lox) Main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		lox.runFile(os.Args[1])
	} else {
		lox.runPrompt()
	}
}

func (lox *Lox) runFile(path string) {
	fmt.Println(path)
	dat, err := ioutil.ReadFile(path)
	check(err)
	lox.run(string(dat))
	if lox.hadError {
		os.Exit(65)
	}
}

func (lox *Lox) runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		lox.run(text)
		lox.hadError = false
	}
}

func (lox *Lox) run(source string) {
	fmt.Print(source)
}

func (lox Lox) Error(line int, message string) {
	lox.report(line, "", message)
}

func (lox *Lox) report(line int, where string, message string) {
	fmt.Println("[line " + string(line) + "] Error" + where + ": " + message)
	lox.hadError = true
}

func main() {
	lox := new(Lox)
	lox.Main()
}
