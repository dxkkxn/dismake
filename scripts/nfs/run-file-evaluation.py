import random
import os
import string
import subprocess

test_path = "./nfs-test"
filename_length = 10
power_2 = [2**i for i in range(17)] 
commands = []

block_size = [4, 8, 16, 1024, 4096, 8192]

output_file = "./results.csv"

for s in power_2:
    for b in power_2:
        csv_file_path = f'./results/{s}_{b}.csv'
        command = f'echo "elapsed_time_us;block_size_KB;total_memory_KB;fsync_interval" > {csv_file_path}'
        p = subprocess.Popen(command, shell=True)
        p.wait()
        if b <= s:
            # for f in power_2:
                # if f <= s//b:
            file_path = "./"+''.join(random.choice(string.ascii_lowercase) for i in range(filename_length))
            commands.append(f'{test_path} {b} {s} {file_path} {1} >> {csv_file_path}')
            commands.append(f'rm {file_path}')
                # else:
                #     continue
        else:
            continue

# for c in commands:
#     for i in range(32):
#         print(f'about to run: {c}')
#         p = subprocess.Popen(c, shell=True)
#         p.wait()

i = 0
j = 1
while i < len(commands):
    for k in range(32):
        c = commands[i]
        print(f'about to run: {c}')
        p = subprocess.Popen(c, shell=True)
        p.wait()
        rm = commands[j]
        print(f'about to run: {rm}')
        p = subprocess.Popen(c, shell=True)
        p.wait()
    i += 2
    j += 2
