import re
from functools import reduce

file_path = "day 2/input.in"

with open(file_path, "r") as file:
    lines = file.read().strip().split("\n")

sum = 0
possible_cubes = {"red": 12, "blue": 14, "green": 13}
for line in lines:
    game = re.split("; |, |: | |\n", line)
    for i in range(3, len(game), 2):
        if int(game[i - 1]) > possible_cubes[game[i]]:
            sum += int(game[1])
            break
print(int(len(lines) * (len(lines) + 1) / 2 - sum))
