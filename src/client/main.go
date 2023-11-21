package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"regexp"
	"time"

	pb "dismake/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func check(e error) {
	if e != nil {
		fmt.Println(e)
		panic("panicking")
	}
}

func main() {
	if len(os.Args) != 2 {
		panic("only one argument pls :(")
	}
	var file string = os.Args[1]
	body, err := os.ReadFile(file)
	check(err)

	interpreter := interpreter{}

	interpreter.input = string(body)
	interpreter.evaluationFailed = false

	yyParse(&interpreter)
	rulesMap := make(map[string]rule)
	for _, rule := range allRules {
		rulesMap[rule.target] = rule
	}
	mainTarget := allRules[len(allRules)-1].target
	// execMakeSeq(mainTarget, rulesMap)

	var server string
	flag.StringVar(&server, "server", "localhost", "Specify the server")
	flag.Parse()
	addr := flag.String("addr", server+":50051", "the address to connect to")
	conn, err := grpc.Dial(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	execMakeDistrib(mainTarget, rulesMap, conn)
	// Set up a connection to the server.
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
}

func execMakeSeq(target string, graph map[string]rule) {
	for _, req := range graph[target].requisites {
		execMakeSeq(req, graph)
	}
	cmd := exec.Command("bash", "-c", graph[target].cmd)
	stdout, err := cmd.Output()
	check(err)
	fmt.Println(graph[target].cmd)
	if len(stdout) != 0 {
		fmt.Println(stdout)
	}
}

func execMakeDistrib(target string, graph map[string]rule, conn *grpc.ClientConn) {
	for _, req := range graph[target].requisites {
		execMakeDistrib(req, graph, conn)
	}
	c := pb.NewCommandRemoteExecClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.CmdRemoteExec(ctx, &pb.CmdRequest{Cmd: graph[target].cmd})
	if err != nil {
		log.Fatalf("could not execute function: %v", err)
	}
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
	var targetToken = tokens[0]
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
	last_returned_value = rune(l.input[0])
	// fmt.Printf("ret: %v %v", ret, '\n')

	l.input = l.input[1:]
	return ret
}
