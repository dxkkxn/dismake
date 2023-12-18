#!/usr/bin/sh
set -xe

nodes=5

# Loop until the variable reaches 20
while [ $nodes -le 5 ]; do
    # Run the Python script with the current value of 'count' as an argument
    oarsub -l nodes=$nodes "cd makefiles; make premier_benchmark; make matrix_benchmark"

    # Increment the variable
    ((count++))
done
