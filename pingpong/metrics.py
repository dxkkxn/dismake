import sys
import re
import matplotlib.pyplot as plt
from datetime import datetime


def parse_log(line):
    pattern = r'time elapsed (.*?)s'
    match = re.search(pattern, line)
    if match:
        valor = match.group(1)
        valor = float(match.group(1)[:-1]) * (
            valor.endswith("m") * 1 +
            valor.endswith("µ") * 10**-3)
        return valor
    return None

def parse_debit(line):
    pattern = r'debit: (.*?)$'
    match = re.search(pattern, line)
    if match:
        valor = match.group(1)
        valor = float(match.group(1))
        return valor
    return None

def parse_size(line):
    pattern = r'message_size: (.*?)$'
    match = re.search(pattern, line)
    if match:
        valor = match.group(1)
        valor = float(match.group(1))
        return valor
    return None


if __name__ == "__main__":
    times_elapsed = []
    debits = []
    sizes = []
    for line in sys.stdin.readlines():
        if "time elapsed" in line:
            times_elapsed += [parse_log(line)]
        if "debit" in line:
            debits += [parse_debit(line)]
        if "message_size" in line:
            sizes += [parse_size(line)]

    plt.hist(times_elapsed[1:-2], bins=10, edgecolor='black')  # Adjust the number of bins as needed

    # Add labels and title
    plt.xlabel('Time (ms)')
    plt.ylabel('Frequency')
    plt.title('Histogram of Sample Data')

    file_name = "hist.png"

    plt.savefig(file_name)
    plt.close()

    plt.plot(sizes, debits, marker='o')
    plt.xlabel('Message size')
    plt.ylabel('Debit')
    plt.title('Debit per message size 1KB -> 3MB')

    file_name = "debit.png"

    # Save the plot as an image file
    plt.savefig(file_name)
    plt.close()


