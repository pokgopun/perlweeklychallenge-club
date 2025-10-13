### https://theweeklychallenge.org/blog/perl-weekly-challenge-335/
=begin

Task 2: Find Winner

Submitted by: [42]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given an array of all moves by the two players.

   Write a script to find the winner of the TicTacToe game if found based
   on the moves provided in the given array.

   UPDATE: Order move is in the order - A, B, A, B, A, ….

Example 1

Input: @moves = ([0,0],[2,0],[1,1],[2,1],[2,2])
Output: A

Game Board:
[ A _ _ ]
[ _ A _ ]
[ B B A ]

Example 2

Input: @moves = ([0,0],[1,1],[0,1],[0,2],[1,0],[2,0])
Output: B

Game Board:
[ A A B ]
[ A B _ ]
[ B _ _ ]

Example 3

Input: @moves = ([0,0],[1,1],[2,0],[1,0],[1,2],[2,1],[0,1],[0,2],[2,2])
Output: Draw

Game Board:
[ A A B ]
[ B B A ]
[ A B A ]

Example 4

Input: @moves = ([0,0],[1,1])
Output: Pending

Game Board:
[ A _ _ ]
[ _ B _ ]
[ _ _ _ ]

Example 5

Input: @moves = ([1,1],[0,0],[2,2],[0,1],[1,0],[0,2])
Output: B

Game Board:
[ B B B ]
[ A A _ ]
[ _ _ A ]
     __________________________________________________________________

   Last date to submit the solution 23:59 (UK Time) Sunday 25th August
   2025.
     __________________________________________________________________

SO WHAT DO YOU THINK ?
=end
### solution by pokgopun@gmail.com
