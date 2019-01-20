import os
import subprocess
import pandas as pd


ABS_PATH = os.path.abspath(os.path.dirname(__file__))


def mesure_instance(num_variables, num_clauses, generation, population, random, elitism, tournament, mutation, input_folder, repeats):


    subprocess.call(['./generate_and_run.sh',
                         '-v', f'{num_variables}',
                         '-c', f'{num_clauses}',
                         '-g', f'{generation}',
                         '-p', f'{population}',
                         '-r', f'{random}',
                         '-e', f'{elitism}',
                         '-t', f'{tournament}',
                         '-m', f'{mutation}',
                         '-R', f'{repeats}',
                         '-f', f'{input_folder}'])

    df = pd.read_csv('instance_data.csv')

    return df

def generation_test(fromGeneration, toGeneration, step, num_variables, num_clauses, population, random, elitism, tournament, mutation, input_folder, repeats):

    num_steps = int((toGeneration - fromGeneration)/step)
    for i in range(num_steps + 1):

        num_generations = fromGeneration + (i*step)
        print('______________________________________________')
        print(f'Increasing Generations to {num_generations}.')

        df = mesure_instance(num_variables, num_clauses, num_generations, population, random, elitism, tournament, mutation, input_folder, repeats)

        df.to_csv(f'instance_data_{num_generations}.csv', sep=',')



def population_test(from_population, to_population, step, num_variables, num_clauses, generation, random_ratio, elitism_ratio, tournament, mutation, input_folder, repeats):


    for i in range(from_population, to_population + step, step):

        random = int(random_ratio * i)
        elitism = int(elitism_ratio * i)

        print('______________________________________________')
        print(f'Increasing Population to {i} with random individuals {random} and number of elitism {elitism}.')

        df = mesure_instance(num_variables, num_clauses, generation, i, random, elitism, tournament, mutation, input_folder, repeats)

        df.to_csv(f'instance_data_{i}.csv', sep=',')

def population_random_test(from_random, to_random, step, num_variables, num_clauses, generation, population, elitism_ratio, tournament, mutation, input_folder, repeats):

    num_steps = int((to_random - from_random)/step)
    print(num_steps)
    for i in range(num_steps + 1):

        random = (from_random + (i*step))*population
        #random = int(random_ratio * i)
        elitism = int(elitism_ratio*population)

        print('______________________________________________')
        print(f'Increasing random ratio to {i} with {random} individuals.')

        df = mesure_instance(num_variables, num_clauses, generation, population, int(random), elitism, tournament, mutation, input_folder, repeats)

        df.to_csv(f'instance_data_{i}.csv', sep=',')

def population_elitism_test(from_elitism, to_elitism, step, num_variables, num_clauses, generation, population, random_ratio, tournament, mutation, input_folder, repeats):


    num_steps = int((to_elitism - from_elitism)/step)
    print(num_steps)
    for i in range(num_steps + 1):

        elitism = int((from_elitism + (i*step))*population)
        #random = int(random_ratio * i)
        random = int(random_ratio*population)

        print('______________________________________________')
        print(f'Increasing elitism ratio to {i} with {elitism} individuals.')

        df = mesure_instance(num_variables, num_clauses, generation, population, random, elitism, tournament, mutation, input_folder, repeats)

        df.to_csv(f'e_instance_data_{i}.csv', sep=',')



def tournament_mesure(from_tournament, to_tournament, step, num_variables, num_clauses, generation, population, random, elistism, mutation, input_folder, repeats):


    for i in range(from_tournament, to_tournament + step, step):

        print('______________________________________________')
        print(f'Increasing tournament size to {i}.')

        df = mesure_instance(num_variables, num_clauses, generation, population, random, elitism, i, mutation, input_folder, repeats)

        df.to_csv(f'instance_data_{i}.csv', sep=',')



if __name__ == "__main__":

    num_variables = 100
    num_clauses = 150

    generation = 500
    population = 50
    random = 10
    elitism = 5
    tournament = 10
    mutation = 0.8

    #input_folder = "/Users/adamzvada/go/src/SAT/scripts"
    input_folder = "/Users/adamzvada/go/src/SAT/input/100-150"
    repeats = 100

    # from_generation = 100
    # to_generation = 1500
    # generation_step = 100

    #df = mesure_instance(num_variables, num_clauses, generation, population, random, elitism, tournament, mutation, input_folder, repeats)
    #generation_test(from_generation, to_generation, generation_step, num_variables, num_clauses, population, random, elitism, tournament, mutation, input_folder, repeats)

    # from_population = 10
    # to_population = 250
    # population_step = 10
    #
    # random_ratio = 0.2
    # elitism_ratio = 0.1

    #population_test(from_population, to_population, population_step, num_variables, num_clauses, generation, random_ratio, elitism_ratio, tournament, mutation, input_folder, repeats)

    #population = 100
    #elitism_ratio = 0.05

    #population_random_test(0.1, 0.9, 0.1, num_variables, num_clauses, generation, population, elitism_ratio, tournament, mutation, input_folder, repeats)

    # population = 100
    # random_ratio = 0.1
    # population_elitism_test(0.05, 0.9, 0.05, num_variables, num_clauses, generation, population, random_ratio, tournament, mutation, input_folder, repeats)

    population = 100
    tournament_mesure(10, population, 10, num_variables, num_clauses, generation, population, random, elitism, mutation, input_folder, repeats)
