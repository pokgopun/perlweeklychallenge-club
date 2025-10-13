### https://theweeklychallenge.org/blog/perl-weekly-challenge-330/
=begin

Task 1: Clear Digits

Submitted by: [44]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given a string containing only lower case English letters and
   digits.

   Write a script to remove all digits by removing the first digit and the
   closest non-digit character to its left.

Example 1

Input: $str = "cab12"
Output: "c"

Round 1: remove "1" then "b" => "ca2"
Round 2: remove "2" then "a" => "c"

Example 2

Input: $str = "xy99"
Output: ""

Round 1: remove "9" then "y" => "x9"
Round 2: remove "9" then "x" => ""

Example 3

Input: $str = "pa1erl"
Output: "perl"

Task 2: Title Capital
=end
### solution by pokgopun@gmail.com
