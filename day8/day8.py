from collections import defaultdict


def match_digits(line_input):
    line_input = line_input.split(" ")
    by_size = defaultdict(list)
    for x in line_input:
        by_size[len(x)].append(x)

    # handle the 8
    digits_map = {by_size[7][0]: 8}

    # handle the 1 digit
    one_places = set(by_size[2][0])
    digits_map[by_size[2][0]] = 1

    # handle the 7 -- we know whatever is in 7 that is NOT in 1 is the top
    seven_places = set(by_size[3][0])
    digits_map[by_size[3][0]] = 7

    # handle the 4 -- we know that the 4 shares 2 characters w/ the 1 and 2 on its own
    four_places = set(by_size[4][0])
    digits_map[by_size[4][0]] = 4

    # These are the ones that will make up the bottom left and bottom segments
    #   since those are the only ones NOT used in 1, 4, or 7
    leftovers = set("abcdefg") - seven_places - four_places

    # length 6 -> digits 0, 6, and 9
    #   0 and 9 have both "one_places"
    #   0 and 6 will have both "leftovers"
    for digit in by_size[6]:
        d = set(digit)
        if not one_places.issubset(digit):
            # this is the 0
            digits_map[digit] = 6
            pass
        elif not leftovers.issubset(digit):
            # this is the 9
            digits_map[digit] = 9
        else:
            # this is the 6
            digits_map[digit] = 0

    # 2, 3, and 5 all have length 5
    for digit in by_size[5]:
        # 3 has top and both 1 digits
        if one_places.issubset(digit):
            digits_map[digit] = 3
        # 2 has both "leftover" segments (BL and B)
        elif leftovers.issubset(digit):
            digits_map[digit] = 2
        else:
            digits_map[digit] = 5

    print(digits_map)
    return digits_map


def calculate_output(output, mapping):
    split_out = output.split(" ")
    val = 0
    m = 1000
    for x in split_out:
        for k, v in mapping.items():
            if set(k) == set(x):
                val += m * v
                break
        m = m / 10
    print(val)
    return val


if __name__ == "__main__":
    with open("../inputs/day8.txt") as file:
        lines = [x.strip() for x in file.readlines()]

    total = 0
    for line in lines:
        splitLines = line.split(" | ")
        mapping = match_digits(splitLines[0])
        total += calculate_output(splitLines[1], mapping)
    print(total)
