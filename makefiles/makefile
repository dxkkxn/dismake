CC=gcc
MAKE=make
BIN=bin

all: mkdir premier_test

mkdir:
	if [ ! -d $(BIN) ]; then \
        mkdir $(BIN); \
        echo "$(BIN) created."; \
    fi

premier_test:premier/premier.c
	$(CC) $< -o $(BIN)/premier -lm
	chmod +x $(BIN)/premier
	PATH=${PATH}:./$(BIN)
	$(MAKE) -f premier/Makefile-small

clean:
	rm -f list*.txt 
	rm -rf $(BIN)

.PHONY: clean