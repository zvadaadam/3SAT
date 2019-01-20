import matplotlib
import matplotlib.pyplot as plt
import matplotlib.style as style
import numpy as np
import pandas as pd
import seaborn as sns

style.use('seaborn-poster') #sets the size of the charts
style.use('ggplot')
style.use('fivethirtyeight')
matplotlib.rcParams['lines.linewidth'] = 2
#matplotlib.rcParams['font.family'] = "roboto"


def satisfablelity_histogram(base_name, from_num, to_num, step, x_label, title, instance_info):

    print("Satisfablelity Plot")

    unsolved = []
    generations = []
    for i in range(from_num, to_num + step, step):
        filename = f'{base_name}_{i}.csv'
        print(filename)
        df = pd.read_csv(filename)

        unsolved_count = len(np.where(df['fitness'] < 0)[0])
        unsolved.append(unsolved_count)
        generations.append(i)

    plt.clf()

    sns_plot = sns.barplot(x=generations, y=unsolved, label=f'Instances - {instance_info}')
    sns_plot.set_title(f'{title} reflecting unsolved instances')
    sns_plot.set_xlabel(x_label)
    sns_plot.set_ylabel('Percentage number of unsolved instances.')
    sns_plot = sns_plot.get_figure()
    sns_plot.savefig('unsolved.png')


def duration_plot(base_name, from_num, to_num, step, x_label, title, instance_info):

    print("Duration Plot")

    generations = []
    durations = []
    for i in range(from_num, to_num + step, step):
        filename = f'{base_name}_{i}.csv'
        print(filename)

        df = pd.read_csv(filename)
        generations.append(i)
        durations.append(df['duration'].mean())

    plt.clf()

    sns_plot = sns.lineplot(x=generations, y=durations, label=f'Instances - {instance_info}')
    sns_plot.set_title(f'{title} reflecting duration for instance')
    sns_plot.set_xlabel(x_label)
    sns_plot.set_ylabel('Duration in seconds')
    sns_plot = sns_plot.get_figure()
    sns_plot.savefig('duration.png')



def fitness_plot(base_name, from_num, to_num, step, num_instance, x_label, title, instance_info):

    #sns.set(style="fivethirtyeight", rc={"lines.linewidth": 1})

    plt.clf()
    print("Fitness Plot")
    for j in range(num_instance):
        fitnesses_instance = []
        params_instance = []
        for i in range(from_num, to_num + step, step):
            filename = f'{base_name}_{i}.csv'
            print(filename)
            df = pd.read_csv(filename)

            fitness_val = df.at[j, 'fitness']
            fitnesses_instance.append(fitness_val)
            params_instance.append(i)

        sns_plot = sns.lineplot(x=params_instance, y=fitnesses_instance, label=f'Instance #{j} - {instance_info}')

    sns_plot.set_xlabel(x_label)
    sns_plot.set_ylabel('Fitness')
    sns_plot.set_title(f'{title} reflecting fitness value')
    sns_plot = sns_plot.get_figure()
    sns_plot.savefig('fitness.png')





# print(style.available)
#
# data_path = 'instance_data.csv'
# df = pd.read_csv(data_path)
# df.reset_index()
#
# #sns.set(style="darkgrid")
# sns_plot = sns.lineplot(x=df.index, y='fitness', label='Genetic SAT', data=df)
# sns_plot.set_xlabel('instance')
# sns_plot.set_title('Fitness for 100 variables and 100 clauses')
# sns_plot = sns_plot.get_figure()
# sns_plot.savefig('graph.png')


if __name__ == '__main__':

    base_name = 'instance_data'
    from_num = 10
    to_num = 250
    step = 10

    #satisfablelity_histogram(base_name, from_num, to_num, step)

    base_name = '../output/generations/instance_data'
    from_num = 100
    to_num = 1500
    step = 100
    num_instance = 10
    x_label = 'Generations'
    y_label = 'Fitness'
    title = 'Fitness for instance with increasing generations'

    #fitness_plot(base_name, from_num, to_num, step, num_instance, x_label, y_label, title, instance_info)

    base_name = '../output/population/instance_data'
    from_num = 10
    to_num = 250
    step = 10
    num_instance = 2
    x_label = 'Population Size'
    y_label = 'Fitness'
    title = 'Fitness for instance with increasing population size'
    instance_info = 'v100, c150, g500'

    #fitness_plot(base_name, from_num, to_num, step, num_instance, x_label, y_label, title, instance_info)


    base_name = '../output/random/instance_data'
    from_num = 0
    to_num = 8
    step = 1
    num_instance = 3
    x_label = 'Percentage of Random Individuals in Population'
    y_label = 'Fitness'
    title = 'Fitness for instance with increasing random individuals'
    instance_info = 'v100, c150, g500, p100'

    #fitness_plot(base_name, from_num, to_num, step, num_instance, x_label, y_label, title, instance_info)

    base_name = '../output/elitism/e_instance_data'
    from_num = 0
    to_num = 16
    step = 1
    num_instance = 2
    x_label = 'Percentage of Elitism Individuals in Population'
    y_label = 'Fitness'
    title = 'Fitness for instance with increasing Elitism individuals'
    instance_info = 'v100, c150, g500, p100'

    #fitness_plot(base_name, from_num, to_num, step, num_instance, x_label, y_label, title, instance_info)

    base_name = '../output/tournament/instance_data'
    from_num = 10
    to_num = 100
    step = 10
    num_instance = 2
    title = 'Tournament size from given population'
    x_label = 'Tournament size'
    instance_info = 'v100, c150, g500, p100, e5, r10'

    fitness_plot(base_name, from_num, to_num, step, num_instance, x_label, title, instance_info)
    satisfablelity_histogram(base_name, from_num, to_num, step, x_label, title, instance_info)
    duration_plot(base_name, from_num, to_num, step, x_label, title, instance_info)
