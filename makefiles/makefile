##
# dismake tests
#
# @file
# @version 0.1
CC=gcc
MAKE=make

MATRIX_SIZE ?= 4


SERVER_PATH=../dismake/server

premier:
	$(CC) premier/premier.c -o premier.exe -lm
	chmod +x premier.exe
	mv premier.exe $(SERVER_PATH)/premier

matrix:
	cp matrix/* $(SERVER_PATH)

blender:
	cp blender_2.59/*.blend $(SERVER_PATH)

.PHONY: premier matrix blender
