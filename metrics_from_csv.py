#!/usr/bin/env python3

import matplotlib.pyplot as plt
import sys

if __name__ == "__main__":
    with open(sys.argv[1]) as f:
        workers = []
        times = []
        for line in f.readlines():
            worker, time = line.strip().split(', ')
            workers.append(int(worker))
            times.append(float(time[:-1]))
    plt.plot(workers, times, marker='o')
    plt.xlabel('Number of workers')
    plt.ylabel('Time (s)')
    plt.title('Time execution per number of worker')
    plt.savefig("workers")
    plt.close()
