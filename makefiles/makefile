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

####### test makefiles/matrix ######

# matrix_test:matrix/
# 	./matrix/random_matrix.pl $(MATRIX_SIZE) $(MATRIX_SIZE) > ./matrix/a
# 	./matrix/random_matrix.pl $(MATRIX_SIZE) $(MATRIX_SIZE) > ./matrix/b
# 	./matrix/generate_makefile.pl $(MATRIX_SIZE) > ./matrix/Makefile
# 	$(MAKE) --directory=/Users/huth/Documents/ENSIMAG/ISI/distrib-make/makefiles/matrix

############# clean ###############
clean:
	rm -f premier/list*.txt 
	rm -rf $(BIN)
# .PHONY: clean
