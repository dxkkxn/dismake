%{
package main

var filesSlice []string
var allRules []rule

type rule struct {
	target string
	requisites []string
	cmd string
}

%}

%union{
String string
StringSlice []string
Rule rule
RuleSlice []rule
}


%error-verbose
%token<String> FILE CMD
%type<String> cmd
%type<StringSlice> files
%type<Rule> target
%type<RuleSlice> expr start

%%
start: expr { $$ = $1 }

expr :
        target cmd expr {
          $1.cmd = $2
          allRules = append(allRules, $1)
          $$ = allRules
        }
     | empty_lines expr { $$ = allRules}
     | /* empty */ { $$ = allRules };

target:
        FILE ':' files '\n' {
            $$ = rule{
              target: $1,
              requisites: $3,
            }
        }
        | FILE ':' '\t' files '\n' {
            $$ = rule{
              target: $1,
              requisites: $4,
            }
        }
        ;
cmd: '\t' CMD '\n' {$$ = $2}

files :
      /* empty */ {
          // empty is called first each time
          filesSlice = []string{}
          $$ = filesSlice
      }
      | FILE files {
            filesSlice = append(filesSlice, $1)
            $$ = filesSlice
      }
      ;

empty_lines: /* empty */ | '\n' empty_lines;

%%
