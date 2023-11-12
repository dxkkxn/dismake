package main

import (
	"fmt"
	"os"
	"regexp"
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic("panicking");
	}
}

func main() {
	if len(os.Args) != 2 {
		panic("only one argument pls :(")
	}
	var file string = os.Args[1]
	body, err := ioutil.ReadFile(file)
	check(err)

	// reader := bufio.NewReader(os.Stdin)
	// yyErrorVerbose = true
	interpreter := interpreter{}


	// for {
	// 	fmt.Print("> ")

	// 	input, err := reader.ReadString('\n')
	// 	if err != nil {
	// 		fmt.Println("Bye.")
	// 		return
	// 	}

	interpreter.input = string(body)
	interpreter.evaluationFailed = false

	cont := yyParse(&interpreter)
	fmt.Println(cont)
	// }
}

const EOF = 0

type interpreter struct {
	input            string
	evaluationFailed bool
}

func (i *interpreter) Error(e string) {
	fmt.Println(e)
	i.evaluationFailed = true
}

type tokenDef struct {
	regex *regexp.Regexp
	token int
}

var tokens = []tokenDef{
	{
		regex: regexp.MustCompile(`[a-zA-Z0-9\.]*`),
		token: FILE,
	},
	{
		regex: regexp.MustCompile(`[a-zA-z0-9\.\<\>\ ]*`),
		token: CMD,
	},
}

var cleaner = regexp.MustCompile(`(#.*\n)*|^\n$`) // checks for comments and empty lines

var last_returned_value = rune(0)

func (l *interpreter) Lex(lval *yySymType) int {
	finished := false
	// skip spaces, empty lines and comments
	for len(l.input) > 0 && !finished {
		finished = true
		// spaces
		if l.input[0] == ' ' {
			l.input = l.input[1:]
			finished = false
		}
		str := cleaner.FindString(l.input)
		if str != "" {
			l.input = l.input[len(str):]
			finished = false
		}
	}

	// Check if the input has ended.
	if len(l.input) == 0 {
		return EOF
	}

	// try to match files except when last token is '\t'
	var targetToken = tokens[0];
	if last_returned_value == '\t' {
		targetToken = tokens[1]
	}
	str := targetToken.regex.FindString(l.input)
	if str != "" {
		// Pass string content to the parser.
		lval.String = str
		l.input = l.input[len(str):]
		return targetToken.token
	}

	// Otherwise return the next letter.

	ret := int(l.input[0])
	last_returned_value = rune(l.input[0]);
	// fmt.Printf("ret: %v %v", ret, '\n')

	l.input = l.input[1:]
	return ret
}
