//# https://theweeklychallenge.org/blog/perl-weekly-challenge-323/
/*#

Task 1: Increment Decrement

Submitted by: [47]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given a list of operations.

   Write a script to return the final value after performing the given
   operations in order. The initial value is always 0.
Possible Operations:
++x or x++: increment by 1
--x or x--: decrement by 1

Example 1

Input: @operations = ("--x", "x++", "x++")
Output: 1

Operation "--x" =>  0 - 1 => -1
Operation "x++" => -1 + 1 =>  0
Operation "x++" =>  0 + 1 =>  1

Example 2

Input: @operations = ("x++", "++x", "x++")
Output: 3

Example 3

Input: @operations = ("x++", "++x", "--x", "x--")
Output: 0

Operation "x++" => 0 + 1 => 1
Operation "++x" => 1 + 1 => 2
Operation "--x" => 2 - 1 => 1
Operation "x--" => 1 - 1 => 0

Task 2: Tax Amount
#*/
//# solution by pokgopun@gmail.com
