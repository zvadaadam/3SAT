import os
from os import listdir
from os.path import isfile, join
import pandas as pd

ABS_PATH = os.path.abspath(os.path.dirname(__file__))


def mesure_instance(num_variables, num_clauses, generation, population, random, elitism, tournament, mutation):

    df = pd.DataFrame()

    for f in listdir(input_folder):
        filename = join(input_folder, f)

        if isfile(filename):
            print(filename)



if __name__ == "__main__":

    generation = 500
    population = 50
    random = 10
    elitism = 5
    tournament = 10
    mutation = 0.8

    input_folder = '/Users/adamzvada/go/src/SAT/input/ratio/1'

    mesure_instance(input_folder, generation, population, random, elitism, tournament, mutation)



