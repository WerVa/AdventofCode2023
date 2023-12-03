import re
from functools import reduce
import operator

file_path = "day 2/input.in"

with open(file_path, "r") as file:
    lines = file.read().strip().split("\n")

sum = 0
for line in lines:
    number_cubes = {'red':0, 'blue':0, 'green':0}
    new_line = re.split('; |, |: | |\n',line)
    for i in range(3, len(new_line), 2):
        number_cubes[new_line[i]] = max(number_cubes[new_line[i]], int(new_line[i-1]))
    sum += reduce(operator.mul, number_cubes.values(), 1)

print(sum)