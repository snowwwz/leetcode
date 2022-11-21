# https://leetcode.com/problems/squares-of-a-sorted-array/
from typing import List

class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        for i in range(len(nums)):
            nums[i] *= nums[i]
        nums.sort()
        return nums