CC=gcc
MAKE=make
BIN=bin

MATRIX_SIZE ?= 4


all: mkdir_bin matrix_test


###### test makefiles/premier ######
mkdir_bin:
	if [ ! -d $(BIN) ]; then \
        mkdir $(BIN); \
        echo "$(BIN) created."; \
    fi

premier_test:premier/premier.c
	$(CC) $< -o $(BIN)/premier -lm
	chmod +x $(BIN)/premier
	PATH=${PATH}:./$(BIN)
	$(MAKE) -f premier/Makefile-small


####### test makefiles/matrix ######

matrix_test:matrix/
	./matrix/random_matrix.pl $(MATRIX_SIZE) $(MATRIX_SIZE) > ./matrix/a
	./matrix/random_matrix.pl $(MATRIX_SIZE) $(MATRIX_SIZE) > ./matrix/b
	./matrix/generate_makefile.pl $(MATRIX_SIZE) > ./matrix/Makefile
	$(MAKE) --directory=/Users/huth/Documents/ENSIMAG/ISI/distrib-make/makefiles/matrix

############## clean ###############
clean:
	rm -f list*.txt 
	rm -rf $(BIN)
	$(MAKE) clean --directory=/Users/huth/Documents/ENSIMAG/ISI/distrib-make/makefiles/matrix
.PHONY: clean