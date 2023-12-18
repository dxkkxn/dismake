#!/bin/bash

# Initialize the variable
count=1

# Loop until the variable reaches 20
while [ $count -le 20 ]; do
    # Run the Python script with the current value of 'count' as an argument
    python3 benchmark.py $count

    # Increment the variable
    ((count++))
done