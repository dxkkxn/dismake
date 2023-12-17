package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"sync"
	"regexp"
	"time"
	"strings"

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

type remoteConn struct {
	conn *grpc.ClientConn
	used bool
	serverName string
}

func main() {
	log.SetPrefix("[client] ")
	log.SetFlags(0)

	var serversStr string
	flag.StringVar(&serversStr, "server", "localhost:50051", "Specify the servers as \"server1 server2 server3\" ")
	flag.Parse()

	servers := strings.Split(serversStr, " ")

	log.Printf("provided servers: %v\n", servers)

	if len(flag.Args()) != 1 {
		panic("one argument pls :(")
	}

	var file string = flag.Args()[0]
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

	connections := make([]remoteConn, len(servers));
	for i, server := range servers {
		log.Println(server)
		// Set up a connection to the server.
		conn, err := grpc.Dial(server, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		connections[i] = remoteConn{ conn, false, server }
	}
	mech := syncMech{sync.WaitGroup{}, make(chan message, len(servers))}
	available := len(servers)
	execMakeDistrib(mainTarget, rulesMap, &connections, &mech, &available)
	wg.Wait()
}

var wg sync.WaitGroup;
type syncMech struct {
	wg sync.WaitGroup
	done chan message

}


func execCmd(serverNum int, conn *grpc.ClientConn, cmd string, done chan<-message) {
	log.Printf("sending command: %v", cmd)
	c := pb.NewCommandRemoteExecClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := c.CmdRemoteExec(ctx, &pb.CmdRequest{Cmd: cmd})
	if err != nil {
		log.Printf("could not execute function: %v", err)
	}
	log.Printf("execution finished for command: %v", cmd)
	done <- message{serverNum, 0};
}


func execMakeDistrib(target string, graph map[string]rule, connections *[]remoteConn, mech *syncMech, available *int) {
	for _, req := range graph[target].requisites {
		execMakeDistrib(req, graph, connections, mech, available)
	}
	if *available <= 0 {
		m := <-mech.done
		// available++
		(*connections)[m.serverNum].used = false
	}

	for i, _:= range *connections {
		remote := &((*connections)[i])
		if (*remote).used == false {
			(*remote).used = true
			*available--
			log.Printf("executing function at server: %v", (*connections)[i].serverName)
			wg.Add(1)
			go func() {
				defer wg.Done()
				i := i
				execCmd(i, remote.conn, graph[target].cmd, mech.done)
			}()
			break;
		}
	}
}

const EOF = 0

type interpreter struct {
	input            string
	evaluationFailed bool
}

type message struct {
	serverNum int
	status int
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
		regex: regexp.MustCompile(`[a-zA-Z0-9\.\-]*`),
		token: FILE,
	},
	{
		regex: regexp.MustCompile(`[a-zA-z0-9;\-\|\/\*\.\<\>\ "]*`),
		token: CMD,
	},
}

var cleaner = regexp.MustCompile(`(#.*\n)*|^\n$`) // checks for comments and empty lines

var last_returned_value [2]rune;

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
	if last_returned_value[0] == '\n' && last_returned_value[1] == '\t' {
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
	last_returned_value[0] = last_returned_value[1]
	last_returned_value[1] = rune(l.input[0])

	l.input = l.input[1:]
	return ret
}
