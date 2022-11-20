# https://leetcode.com/problems/binary-search/
from typing import List

# using recursion
class Solution:
    def search(self, nums: List[int], target: int) -> int:
        return self.binarySearch(nums, target, 0, len(nums)-1)
    
    def binarySearch(self, nums: List[int], target: int, l: int, r: int) -> int:
        if l > r:
            return -1
        m = (l+r)//2
        if nums[m] == target:
            return m
        elif nums[m] > target:
            r = m - 1
        else:
            l = m + 1
        return self.binarySearch(nums, target, l, r)

# class Solution:
#     def search(self, nums: List[int], target: int) -> int:
#       l = 0
#       r = len(nums)-1

#       while l <= r:
#         m = (l+r)//2
#         if nums[m] == target:
#             return m
#         elif nums[m] > target:
#             r = m - 1
#         else:
#             l = m + 1
#       return -1

import unittest

class TestSearch(unittest.TestCase):
  def test_search(self):
    case = [
      {'arg1': [5], 'arg2': 5, 'expected': 0},
      {'arg1': [2,5], 'arg2': 2, 'expected': 0},
      {'arg1': [-1,0,3,5,9,12], 'arg2': 5, 'expected': 3},
    ]

    for c in case:
      self.assertEqual(c['expected'], Solution().search(c['arg1'], c['arg2']))
      
if __name__ == "__main__":
    unittest.main()
