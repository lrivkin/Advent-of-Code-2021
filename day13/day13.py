def print_paper(paper, max_x, max_y, fold_dir="", fold_line=-1):
    paper_grid = []
    for i in range(max_y+1):
        paper_grid.append(["." for _ in range(max_x+1)])

    for x, y in paper:
        paper_grid[y][x] = "#"

    if fold_dir == "x":
        for g in paper_grid:
            g[fold_line] = "|"
    if fold_dir == "y":
        paper_grid[fold_line] = ["-" for _ in range(max_x+1)]

    for g in paper_grid:
        print("".join(g))
    print()


def fold_over(paper, fold_dir, fold_line):


if __name__ == "__main__":
    with open("../tests/day13.txt") as file:
        lines = [x.strip() for x in file.readlines()]
    i = 0
    max_x = 0
    max_y = 0
    positions = set()
    for input_spot in lines:
        if input_spot == '':
            break
        input_spot = input_spot.split(",")
        x = int(input_spot[0])
        y = int(input_spot[1])
        positions.add((x, y))
        max_x = max(x, max_x)
        max_y = max(y, max_y)
        i += 1

    instructions = []
    for c in range(i+1, len(lines)):
        instruction = lines[c].split(' ')[-1]
        instructions.append(instruction)

    print(positions)
    print(instructions)

    for fold in instructions:
        fold = fold.split("=")
        direction = fold[0]
        line_num = int(fold[1])
