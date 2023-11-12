%{
package main

%}

%union{
String string
}


%error-verbose
%token FILE CMD

%%
start: expr;

expr :
       target cmd expr
     | empty_lines expr
     | /* empty */;

target: FILE ':' files '\n';
cmd: '\t' CMD '\n'

files : /* empty */ | FILE files;

empty_lines: /* empty */ | '\n' empty_lines;

%%
