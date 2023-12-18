#!/usr/bin/env python3

import matplotlib.pyplot as plt
import sys
from collections import defaultdict

if __name__ == "__main__":
    workers = defaultdict(list)
    with open(sys.argv[1]) as f:
        for line in f.readlines():
            worker, time = line.strip().split(', ')
            workers[worker].append(float(time[:-1]))
    keys = sorted(list(workers.keys()))
    times = [sum(workers[key])/len(workers[key]) for key in keys]
    plt.plot(keys, times, marker='o')
    plt.xlabel('Number of workers')
    plt.ylabel('Time (s)')
    plt.title('Time execution per number of worker')
    plt.savefig("workers")
    plt.close()
