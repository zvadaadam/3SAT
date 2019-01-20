#!/bin/bash

# LOAD OPTIONS
while getopts v:c:g:p:r:e:t:m:R:f:- opt
do
	case "$opt" in
	v)
	    variables="$OPTARG" # -v
	;;

	c)
	    clauses="$OPTARG" # -c
	;;

	g)
		generations="$OPTARG" # -g
	;;

	p)
		population="$OPTARG" # -p
	;;

	r)
		random="$OPTARG" # -r
	;;

	e)
		elitism="$OPTARG" # -e
	;;

	t)
		tournament="$OPTARG" # -t
	;;

	m)
	    mutation="$OPTARG" # -m
	;;
	R)
	    repeat="$OPTARG" # -m
	;;
	f)
	    input_folder="$OPTARG" # -m
	;;
	esac
done
shift "$(( OPTIND - 1 ))"


echo "__________________________________"
echo "variables: $variables"
echo "clauses: $clauses"
echo "generation: $generations"
echo "population: $population"
echo "random: $random"
echo "elitism: $elitism"
echo "tournament: $tournament"
echo "mutation: $mutation"
echo "repeat: $repeat"
echo "input_folder: $input_folder"
echo "__________________________________"


cd ..
echo "Build Go-SAT solver project..."
go build main.go
cd scripts

rm -rf instance_data.csv
printf -- "variables,clauses,fitness,duration\n" >> instance_data.csv

for i in `seq 1 1 ${repeat}`; do

    if [ $input_folder == "/Users/adamzvada/go/src/SAT/scripts" ]
    then
        rm -rf input-${i}
        ../helper/generator $variables $clauses 100 >> "input-${i}"
    fi

    echo "Running SAT Solver #${i} with ${variables} variables and ${clauses} clauses."

    printf -- "${variables},${clauses}," >> instance_data.csv
    ../main -input "${input_folder}/input-${i}" \
        -generation $generations \
        -population $population \
        -random $random -elitism \
        $elitism -tournament \
        $tournament -mutation \
        $mutation  \
        >> instance_data.csv

    rm -rf input
done