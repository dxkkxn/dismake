CC=gcc
MAKE=make


compile_premier:premier_time_operations/premier.c
	$(CC) $< -o premier_time_operations/premier -lm
	chmod +x premier_time_operations/premier

premier_test_partial_times:premier_time_operations/premier
	$(MAKE) --directory=premier_time_operations -f Makefile

ma:premier_time_operations/premier
	time $(MAKE) --directory=premier_time_operations -f Makefile_total_time

clean:
	rm -f premier_time_operations/list*.txt 
