#!/bin/bash

cd ../generator
gcc -w -o generator g2.c

MAX_PRICE=1000
RATIO=4
REPEAT_INSTANCE=10

mkdir ../input/ratio/
mkdir ../input/ratio/${RATIO}

for i in `seq 10 10 100`; do
    CLAUSES=$((i*RATIO))
    for j in `seq 1 1 ${REPEAT_INSTANCE}`; do
        ./generator $i $CLAUSES $MAX_PRICE > "../input/ratio/${RATIO}/input-${i}-${CLAUSES}-${j}"
    done
done