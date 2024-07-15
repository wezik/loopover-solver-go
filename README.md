# Loopover solver

An app that solves [Loopover](https://loopover.xyz/) puzzle :)

## How it works

It takes in current state of the board, runs a bunch of logic on it, and returns a list of moves to complete the puzzle

## How to use

1. **Build** with `go build ./...`
2. **Run** with `go run .`
3. Send a **POST** request to `http://localhost:8080/v1/api/solve` with **JSON** body of the board structured like:
```json
{
  "board": [
    [20,6,9,24,15],
    [5,17,21,7,11],
    [13,22,4,25,2],
    [8,12,16,1,23],
    [18,14,10,3,19]
  ]
}
```
4. Read or interpret the list of moves returned as an array in **JSON** format

## How to understand the list of moves
Example:
```json
{
  "moves": [
    "2R3",
    "3U4",
    "1L0",
    "4U4",
    "1L0"
  ]
}
```
Structure of move string:
`{Repetitions} {Direction} {Index}`

How to read `2R3` then?
- **2** means it's repeated 2 times
- **R** means it's a move to the **R**ight
  - other options: **L**eft / **U**p / **D**own
- **3** means index of 3, or 4th row from the top (since we index from 0)
  - for Up and Down it is an index of column

Moving then 2 times, row of index 3, to the right, should look like this:  

**Rows ↓** **Columns →**
| Index |    |    |    |    |    | |    |    |    |    |    |
|-------|----|----|----|----|----|-|----|----|----|----|----|
| **0** |  1 |  2 |  3 |  4 |  5 | |  1 |  2 |  3 |  4 |  5 |
| **1** |  6 |  7 |  8 |  9 | 10 | |  6 |  7 |  8 |  9 | 10 |
| **2** | 11 | 12 | 13 | 14 | 15 | | 11 | 12 | 13 | 14 | 15 |
| **3** | **16** | **17** | **18** | **19** | **20** |→| **19** | **20** | **16** | **17** | **18** |
| **4** | 21 | 22 | 23 | 24 | 25 | | 21 | 22 | 23 | 24 | 25 |

## Configurating
This part is still work in progress as I haven't made it past local enviornment with this project yet
