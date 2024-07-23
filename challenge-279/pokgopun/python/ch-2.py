### https://theweeklychallenge.org/blog/perl-weekly-challenge-279/
"""

Task 2: Split String

Submitted by: [46]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given a string, $str.

   Write a script to split the given string into two containing exactly
   same number of vowels and return true if you can otherwise false.

Example 1

Input: $str = "perl"
Ouput: false

Example 2

Input: $str = "book"
Ouput: true

Two possible strings "bo" and "ok" containing exactly one vowel each.

Example 3

Input: $str = "good morning"
Ouput: true

Two possible strings "good " and "morning" containing two vowels each or "good m
" and "orning" containing two vowels each.
     __________________________________________________________________

   Last date to submit the solution 23:59 (UK Time) Sunday 28th July 2024.
     __________________________________________________________________

SO WHAT DO YOU THINK ?
"""
### solution by pokgopun@gmail.com

def splitString0(string: str):
    return sum( 1 for i in range(len(string)) if string[i] in "aeiou" ) % 2 == 0

def splitString1(string: str):
    idx = tuple( i for i in range(len(string)) if string[i] in "aeiou" )
    l = len(idx)
    h = int(l/2)
    b = l % 2 == 0
    if b:
        h -= 1
    i = idx[h+1] if l > 1 else idx[h]+1
    return (string[:i], string[i:], b)

def splitString2(string: str):
    lol = len(string)
    mol = int(lol/2)
    eol = lol % 2 == 0
    if eol:
        mol -= 1
    idx = tuple( i for i in range(lol) if string[i] in "aeiouAEIOU" )
    lov = len(idx)
    mov = int(lov/2)
    eov = lov % 2 == 0
    if lov < 2:
        return (string[:mol+1],string[mol+1:],eov)
    if eov:
        mov -= 1
    i = idx[mov]
    if mol > idx[mov]:
        if mol < idx[mov+1]:
            i = mol
        else:
            i = idx[mov+1]-1
    return (string[:i+1], string[i+1:], eov)

import unittest

class TestSplitString(unittest.TestCase):
    def test0(self):
        for inpt,otpt in {
                "perl": False,
                "book": True,
                "good morning": True,
                }.items():
            self.assertEqual(splitString0(inpt),otpt)
    def test1(self):
        for inpt,otpt in {
                "perl": ("pe","rl",False),
                "book": ("bo","ok",True),
                "good morning": ("good m","orning",True),
                }.items():
            self.assertEqual(splitString1(inpt),otpt)
    def test2(self):
        for inpt,otpt in {
                "perl": ("pe","rl",False),
                "book": ("bo","ok",True),
                "good morning": ("good m","orning",True),
                "good morning bird": ("good morn","ing bird",False),
                "go": ("g","o",False),
                "golang": ("gol","ang",True),
                "lua": ("lu","a",True),
                "bash": ("ba","sh",False),
                "powershell": ("power","shell",False),
                "javascript": ("javas","cript",False),
                "python": ("pyt","hon",False),
                "swift": ("swi","ft",False),
                "bqn": ("bq","n",True),
                "haskell": ("hask","ell",True),
                "c": ("c","",True),
                "Assembly": ("Ass","embly",True),
                "WebAssembly": ("WebAss","embly",False),
                "Pascal": ("Pas","cal",True),
                "Fortan": ("For","tan",True),
                "ShellScript": ("ShellS","cript",True),
                "Scala": ("Sca","la",True),
                "flutter": ("flut","ter",True),
                "Progressive Rock": ("Progressi","ve Rock",False), ### should better be ("Progress","ive Rock",False)
                }.items():
            print(splitString2(inpt))
            self.assertEqual(splitString2(inpt),otpt)

unittest.main()
