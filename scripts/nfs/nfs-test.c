#include <stdio.h>
#include <stdlib.h>
#include <fcntl.h>
#include <unistd.h>
#include <sys/time.h>

int main(int argc, char *argv[]) {
    if (argc != 5) {
        fprintf(stderr, "Usage: %s <block_size_KB> <total_memory_KB> <filename> <fsync_interval>\n", argv[0]);
        return 1;
    }

    // Get block size, total memory, filename, and fsync interval from command line
    int block_size_KB = atoi(argv[1]);
    int total_memory_KB = atoi(argv[2]);
    char *filename = argv[3];
    int fsync_interval = atoi(argv[4]);

    if (block_size_KB <= 0 || total_memory_KB <= 0 || fsync_interval <= 0) {
        fprintf(stderr, "Block size, total memory, and fsync interval must be positive integers\n");
        return 1;
    }

    // Convert block size, total memory, and fsync interval to bytes
    int block_size = block_size_KB * 1024;
    int total_memory = total_memory_KB * 1024;

    // Calculate the number of blocks to write
    int num_blocks = total_memory / block_size;

    // Allocate memory on the heap
    char *data = (char *)malloc(block_size);
    if (data == NULL) {
        perror("Memory allocation failed");
        return 1;
    }

    // Initialize the memory (you can modify this part as needed)
    for (int i = 0; i < block_size; ++i) {
        data[i] = i % 256; // Fill with some sample data
    }

    // Open the file using creat with binary mode
    int file_descriptor = creat(filename, S_IRUSR | S_IWUSR | S_IRGRP | S_IROTH | O_CREAT | O_WRONLY | O_TRUNC);
    if (file_descriptor == -1) {
        perror("Error opening file");
        free(data);
        return 1;
    }

    // Record start time
    struct timeval start, end;
    gettimeofday(&start, NULL);

    // Write the data to the file
    for (int i = 0; i < num_blocks; ++i) {
        ssize_t bytes_written = write(file_descriptor, data, block_size);
        if (bytes_written == -1) {
            perror("Error writing to file");
            close(file_descriptor);
            free(data);
            return 1;
        }

        // Check if all bytes were written
        if (bytes_written != block_size) {
            fprintf(stderr, "Error: Not all bytes were written to the file\n");
            close(file_descriptor);
            free(data);
            return 1;
        }

        // Perform fsync at every k blocks
        if ((i + 1) % fsync_interval == 0) {
            if (fsync(file_descriptor) == -1) {
                perror("Error syncing file to disk");
                close(file_descriptor);
                free(data);
                return 1;
            }
        }
    }

    // Record end time
    gettimeofday(&end, NULL);

    // Calculate elapsed time in microseconds
    long elapsed_time = (end.tv_sec - start.tv_sec) * 1000000 + (end.tv_usec - start.tv_usec);

    // Print the time taken
    // printf("Time taken for writing: %ld microseconds\n", elapsed_time);
    printf("%ld;%d;%d;%d\n", elapsed_time, block_size_KB, total_memory_KB, fsync_interval);
    // Close the file and free allocated memory
    close(file_descriptor);
    free(data);

    return 0;
}

