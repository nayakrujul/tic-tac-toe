def check_win(b: list[list[int]]) -> int:
    wins = {1}, {2}
    for i in range(len(b)):
        if {*b[i]} in wins:
            return b[i][0]
        if {*[*zip(*b)][i]} in wins:
            return b[0][i]
    if {b[j][j] for j in range(len(b))} in wins:
        return b[0][0]
    if {b[k][-k - 1] for k in range(len(b))} in wins:
        return b[0][-1]
    return (0 in sum(b, [])) - 1


def take_input(b: list[list[int]]) -> int:
    s = input(" > ")
    if not s.isdigit():
        return -1
    n = int(s) - 1
    if not (0 <= n <= 8):
        return -1
    if b[n // 3][n % 3] != 0:
        return -1
    return n


def print_board(b: list[list[int]]) -> None:
    for x, row in enumerate(b):
        for y, i in enumerate(row):
            print(" XO"[i], end=" | " * (y != len(b) - 1))
        print(("\n" + "--|-" * (len(b) - 1) + "-") * (x != len(b) - 1))


board = \
    [[0, 0, 0],
     [0, 0, 0],
     [0, 0, 0]]


player = 1
while (winner := check_win(board)) == 0:
    print("\033[H\033[2JPlayer", player, "\n")
    print_board(board)
    print()
    while (inp := take_input(board)) == -1:
        pass
    board[inp // 3][inp % 3] = player
    player = player % 2 + 1

print(end="\033[H\033[2J")

if winner == -1:
    print("Draw!\n")
else:
    print("Player", winner, "wins!\n")

print_board(board)
