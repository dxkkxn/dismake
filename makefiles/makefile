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
	../make.sh premier/Makefile

premier_benchmark:
	$(CC) premier/premier.c -o premier.exe -lm
	chmod +x premier.exe
	mv premier.exe $(SERVER_PATH)/premier
	../make.sh premier/Makefile
	../make.sh premier/Makefile
	../make.sh premier/Makefile

matrix:
	cp matrix/* $(SERVER_PATH)
	../make.sh matrix/Makefile

matrix_benchmark:
	cp matrix/* $(SERVER_PATH)
	../make.sh matrix/Makefile
	../make.sh matrix/Makefile
	../make.sh matrix/Makefile

blender:
	cp blender_2.59/*.blend $(SERVER_PATH)

.PHONY: premier matrix blender
