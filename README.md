# Number of Lakes

## Description
The user will provide an m x n grid map consisting of cells marked with # (representing land) and . (representing water). 
The function will calculate the number of lakes in the grid. A lake is defined as a connected group of water cells (.) 
that are adjacent either horizontally, vertically, or diagonally. The function employs the Depth-First Search (DFS) 
algorithm to solve this problem.  

## Input Format
- The first line contains a single integer, `m`, the map's width.
- The second line contains a single integer, `n`, the map’s height.
- The following lines represent the map grid.

## Constraints
- 1 <= `m` <= 100
- 1 <= `n` <= 100

## Sample Input
```
5
4
#.###
..###
##.##
####.
```

## Sample Output
```
TOTAL: 2
```

## Explanation
### Lake = 1
first lake represented by a set of water that connected vertically, horizontally, and diagonally. 
by the example below the first lake represented by `+`
```
#+###
++###
##+##
####.
```

### Lake = 2
second lake represented by a single of water that not connected with the other waters.
```
#.###
..###
##.##
####+
```