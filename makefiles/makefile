CC=gcc
MAKE=make

MATRIX_SIZE ?= 4


all: premier_test


###### test makefiles/premier ######
premier_test:premier/premier.c
	$(CC) $< -o premier/premier -lm
	chmod +x premier/premier
	$(MAKE) --directory=premier -f Makefile-small

premier_test_long:premier/premier.c
	$(CC) $< -o premier/premier -lm
	chmod +x premier/premier
	$(MAKE) --directory=premier -f Makefile


####### test makefiles/matrix ######

matrix_test:matrix/
	./matrix/random_matrix.pl $(MATRIX_SIZE) $(MATRIX_SIZE) > ./matrix/a
	./matrix/random_matrix.pl $(MATRIX_SIZE) $(MATRIX_SIZE) > ./matrix/b
	./matrix/generate_makefile.pl $(MATRIX_SIZE) > ./matrix/Makefile
	$(MAKE) --directory=/Users/huth/Documents/ENSIMAG/ISI/distrib-make/makefiles/matrix

############## clean ###############
clean:
	rm -f premier/list*.txt 
	rm -rf premier/premier
	$(MAKE) clean --directory=/Users/huth/Documents/ENSIMAG/ISI/distrib-make/makefiles/matrix
.PHONY: clean