file_path = "day 1/input.in"

sum = 0

with open(file_path, "r") as file:
    for line in file:
        words = line.split()
        for word in words:
            numbers = [letter for letter in word if letter.isdigit()]
            if numbers:
                extnum = int(numbers[0] + numbers[-1])
                sum += extnum

print(sum)
