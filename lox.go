package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Fingel/golox/lox"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Lox struct {
	hadError bool
}

func (l *Lox) Main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: glox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 {
		l.runFile(os.Args[1])
	} else {
		l.runPrompt()
	}
}

func (l *Lox) runFile(path string) {
	fmt.Println(path)
	dat, err := ioutil.ReadFile(path)
	check(err)
	l.run(string(dat))
	if l.hadError {
		os.Exit(65)
	}
}

func (l *Lox) runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		l.run(text)
		l.hadError = false
	}
}

func (l *Lox) run(source string) {
	scanner := lox.NewScanner(source)
	tokens, errors := scanner.ScanTokens()
	for _, error := range errors {
		l.report(error.Line, error.Where, error.Message)
	}
	for _, token := range tokens {
		fmt.Println(token.String())
	}
}

func (l Lox) Error(line int, message string) {
	l.report(line, "", message)
}

func (l *Lox) report(line int, where string, message string) {
	fmt.Printf("[line %d] Error %s: %s\n", line, where, message)
	l.hadError = true
}

func main() {
	lox := new(Lox)
	lox.Main()
}
