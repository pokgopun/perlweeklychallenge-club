### https://theweeklychallenge.org/blog/perl-weekly-challenge-341/
=begin

Task 1: Broken Keyboard

Submitted by: [39]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given a string containing English letters only and also you are
   given broken keys.

   Write a script to return the total words in the given sentence can be
   typed completely.

Example 1

Input: $str = 'Hello World', @keys = ('d')
Output: 1

With broken key 'd', we can only type the word 'Hello'.

Example 2

Input: $str = 'apple banana cherry', @keys = ('a', 'e')
Output: 0

Example 3

Input: $str = 'Coding is fun', @keys = ()
Output: 3

No keys broken.

Example 4

Input: $str = 'The Weekly Challenge', @keys = ('a','b')
Output: 3

Example 5

Input: $str = 'Perl and Python', @keys = ('p')
Output: 1

Task 2: Reverse Prefix
=end
### solution by pokgopun@gmail.com
