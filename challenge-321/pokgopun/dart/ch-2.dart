//# https://theweeklychallenge.org/blog/perl-weekly-challenge-321/
/*#

Task 2: Backspace Compare

Submitted by: [45]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given two strings containing zero or more #.

   Write a script to return true if the two given strings are same by
   treating # as backspace.

Example 1

Input: $str1 = "ab#c"
       $str2 = "ad#c"
Output: true

For first string,  we remove "b" as it is followed by "#".
For second string, we remove "d" as it is followed by "#".
In the end both strings became the same.

Example 2

Input: $str1 = "ab##"
       $str2 = "a#b#"
Output: true

Example 3

Input: $str1 = "a#b"
       $str2 = "c"
Output: false
     __________________________________________________________________

   Last date to submit the solution 23:59 (UK Time) Sunday 18th May 2025.
     __________________________________________________________________

SO WHAT DO YOU THINK ?
#*/
//# solution by pokgopun@gmail.com
