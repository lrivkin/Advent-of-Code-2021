import collections

if __name__ == "__main__":
    print("hello")
    with open("../inputs/day3.txt") as file:
        lines = [x.strip() for x in file.readlines()]

    total = len(lines)
    size = len(lines[0])
    gamma = ""
    epsilon = ""
    for i in range(size):
        totalOnes = sum([int(b[i]) for b in lines])
        if totalOnes >= total/2:
            # print(f"for position {i} using 1")
            gamma += "1"
            epsilon += "0"
        else:
            # print(f"for position {i} using 0")
            gamma += "0"
            epsilon += "1"
    print(f"gamma = {int(gamma, 2)}")
    print(f"epsilon = {int(epsilon, 2)}")
    print(f"part1 = {int(gamma, 2)*int(epsilon, 2)}")

    # part 2
    oxygen = lines
    for i in range(size):
        totalOnes = sum([int(b[i]) for b in oxygen])
        target = 0
        if totalOnes >= len(oxygen)/2:
            target = 1
        oxygen = [o for o in oxygen if int(o[i]) == target]
        # print(oxygen)
        if len(oxygen) == 1:
            break
    co2 = lines
    for i in range(size):
        totalOnes = sum([int(b[i]) for b in co2])
        target = 1
        if totalOnes >= len(co2)/2:
            target = 0
        co2 = [o for o in co2 if int(o[i]) == target]
        # print(co2)
        if len(co2) == 1:
            break
    print(f"part2 = {int(oxygen[0], 2)*int(co2[0], 2)}")
