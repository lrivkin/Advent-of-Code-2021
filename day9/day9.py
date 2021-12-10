import math

brace_pairs = {
    '(': ')',
    '{': '}',
    '[': ']',
    '<': '>'
}
bad_points = {
    ')': 3,
    ']': 57,
    '}': 1197,
    '>': 25137
}
completion_points = {
    ')': 1,
    ']': 2,
    '}': 3,
    '>': 4
}


def find_corrupt(line):
    stack = []
    for c in line:
        if c in brace_pairs.keys():
            stack.append(c)
        else:
            if len(stack) == 0 or brace_pairs[stack[-1]] != c:
                return None, c
            stack.pop()

    completion_string = []
    while len(stack) > 0:
        x = stack.pop()
        completion_string.append(brace_pairs[x])
    return completion_string, None


def calculate_score(completion_string):
    score = 0
    for c in completion_string:
        score *= 5
        score += completion_points[c]
    # print(score)
    return score


if __name__ == "__main__":
    with open("../inputs/day9.txt") as file:
        lines = [x.strip() for x in file.readlines()]
    error_score = 0
    autocomplete_scores = []
    for line in lines:
        completion, bad = find_corrupt(line)
        if bad is not None:
            error_score += bad_points[bad]
        else:
            autocomplete_scores.append(calculate_score(completion))
    print(error_score)
    autocomplete_scores = sorted(autocomplete_scores)
    print(autocomplete_scores[math.floor(len(autocomplete_scores)/2)])
