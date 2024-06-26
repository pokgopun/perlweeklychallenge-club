function array_sum(t)
   assert(type(t) == "table", "t must be a table!")
   local sum = 0
   for i=1, #t do sum = sum + t[i] end
   return sum
end

function equilibrium_index(t)
   assert(type(t) == "table", "t must be a table!")
   local left, right, ret = 0, array_sum(t), -1
   for i,j in pairs(t) do
      right = right - j
      if left == right then
	 ret = i
	 break
      end
      left = left + j
   end
   return ret
end

print(equilibrium_index({1,3,5,7,9}))
print(equilibrium_index({1,2,3,4,5}))
print(equilibrium_index({2,4,2}))
