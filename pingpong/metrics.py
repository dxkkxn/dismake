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


if __name__ == "__main__":
    times_elapsed = []
    for line in sys.stdin.readlines():
        if "time elapsed" in line:
            times_elapsed += [parse_log(line)]

    plt.hist(times_elapsed[1:-2], bins=10, edgecolor='black')  # Adjust the number of bins as needed

    # Add labels and title
    plt.xlabel('Time (ms)')
    plt.ylabel('Frequency')
    plt.title('Histogram of Sample Data')

    file_name = datetime.now().strftime("%Y%m%d_%H%M%S") + "_hist.png"

    plt.savefig(file_name)
