#!/bin/bash

cd ../generator
gcc -w -o generator g2.c

MAX_PRICE=1000

for i in `seq 10 10 150`; do
    CLAUSES=$((i*3))
    ./generator $i $CLAUSES $MAX_PRICE > "../input/input-${i}-${CLAUSES}"
done