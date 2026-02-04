# lem-in 🐜

A Go program that simulates ants moving from `##start` to `##end` using the fastest path.

This project reads a file describing rooms and tunnels, then shows how ants move step by step.

---

## Features

- Reads an ant farm from a file
- Finds the shortest path(s)
- Simulates ant movements
- Handles invalid input
- Written in Go using standard packages only

---

## Input Format

The input file contains:

- Number of ants
- Rooms with coordinates
- Links between rooms
- Special commands: `##start` and `##end`

Example:
3
##start
A 1 2
B 3 4
##end
Z 5 6
A-B
B-Z


---

## Output Format

The program prints:

1. The original input
2. Ant movements like this:

L1-B L2-A
L1-Z L2-B
L2-Z


Each line represents one turn.

---

## Usage

### Run the program
```bash
go run . map.txt
```

Build and run
go build -o lem-in
```bash
./lem-in map.txt
```

