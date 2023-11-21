##
# distribmake
#
# @file
# @version 0.1

build: main.go y.go y.output
	go build

y.go y.output: make_grammar.y
	goyacc make_grammar.y

clean:
	rm -rf y.go y.output dismake

# end
