#!/usr/bin/env python3

import matplotlib.pyplot as plt
import sys
from collections import defaultdict

if __name__ == "__main__":
    workers = defaultdict(list)
    filename = sys.argv[1]
    with open(filename) as f:
        for line in f.readlines():
            worker, time = line.strip().split(', ')
            worker = int(worker)
            if 'm' in time:
                minutes = int(time[0])
                time = float(time[2:-1]) + (minutes * 60)
            else:
                time = float(time[:-1])
            workers[worker].append(time)
    keys = sorted(list(workers.keys()))
    times = [sum(workers[key])/len(workers[key]) for key in keys]
    plt.plot(keys, times, marker='o')
    plt.xlabel('Number of workers')
    plt.ylabel('Time (s)')
    plt.title('Time execution per number of worker')
    filename = filename[:-4]
    plt.savefig(filename)
    plt.close()
