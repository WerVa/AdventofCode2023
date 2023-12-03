import re

file_path = "day 1/input.in"

sum = 0

number_dict = {
    "one": "o1ne",
    "two": "t2wo",
    "three": "t3hree",
    "four": "f4our",
    "five": "f5ive",
    "six": "s6ix",
    "seven": "s7even",
    "eight": "e8ight",
    "nine": "n9ine",
}


with open(file_path, "r") as file:
    for line in file:
        words = line.split()
        for word in words:
            pattern = re.findall(
                r"(?=([0-9]|one|two|three|four|five|six|seven|eight|nine))", word
            )
            for match in pattern:
                if match in number_dict:
                    word = word.replace(match, number_dict[match])
            numbers = [letter for letter in word if letter.isdigit()]
            if numbers:
                extnum = int(numbers[0] + numbers[-1])
                sum += extnum
print(sum)
