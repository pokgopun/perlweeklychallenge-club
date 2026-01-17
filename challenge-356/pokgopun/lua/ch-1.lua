--# https://theweeklychallenge.org/blog/perl-weekly-challenge-356/
--[[

Task 1: Kolakoski Sequence

Submitted by: [51]Mohammad Sajid Anwar
     __________________________________________________________________

   You are given an integer, $int > 3.

   Write a script to generate the Kolakoski Sequence of given length $int
   and return the count of 1 in the generated sequence. Please follow the
   [52]wikipedia page for more informations.

Example 1

Input: $int = 4
Output: 2

(1)(22)(11)(2) => 1221

Example 2

Input: $int = 5
Output: 3

(1)(22)(11)(2)(1) => 12211

Example 3

Input: $int = 6
Output: 3

(1)(22)(11)(2)(1)(22) => 122112

Example 4

Input: $int = 7
Output: 4

(1)(22)(11)(2)(1)(22)(1) => 1221121

Example 5

Input: $int = 8
Output: 4

(1)(22)(11)(2)(1)(22)(1)(22) => 12211212

Task 2: Who Wins
--]]
--# solution by pokgopun@gmail.com

--@param n int
local function ksCo(n) --@yield int
	local s = {1,2,2}
	local i = 0
	local v
	while i < 3 do
		v = s[1]
		coroutine.yield(v)
		i = i + 1
		if i == n then
			return
		end
		table.remove(s,1)
	end
	local t, e
	t = i
	while i < n do
		e = 2 - (t % 2)
		for j=1, v do
			coroutine.yield(e)
			i = i + 1
			if i == n then
				return
			end
			table.insert(s,e)
		end
		t = t + 1
		v = s[1]
		table.remove(s,1)
	end
end

--@param n int
local function ks(n) --@return int
	local c = 0
	for v in coroutine.wrap(function() ksCo(n) end) do
		if v == 1 then
			c = c + 1
		end
	end
	return c
end

local lu = require("luaunit")

function TestKs()
	local data = {4, 2, 5, 3, 6, 3, 7, 4, 8, 4}
	for i=2, #data, 2 do
		local inpt, otpt = data[i-1], data[i]
		lu.assertEquals(ks(inpt), otpt)
	end
end

lu.run()
