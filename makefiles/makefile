##
# dismake tests
#
# @file
# @version 0.1
CC=gcc
MAKE=make

MATRIX_SIZE ?= 4


SERVER_PATH=../dismake/server

# all: mkdir_bin matrix_test


premier: premier/premier.c
	$(CC) $< -o premier.exe -lm
	chmod +x premier.exe
	mv premier.exe $(SERVER_PATH)/premier
	../run_app.sh premier/Makefile-small

matrix:
	cp matrix/multiply $(SERVER_PATH)/multiply
	../run_app.sh multiply/Makefile


# clean:
# 	rm -f list*.txt
# 	rm -rf $(BIN)
# .PHONY: clean
