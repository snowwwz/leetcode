# https://leetcode.com/problems/search-insert-position/
from typing import List
import unittest

class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        l = 0 
        r = len(nums) - 1
        while l <= r:
            m = (l+r)//2
            if nums[m] == target:
                return m
            elif nums[m] < target:
                l = m + 1
            else:
                r = m - 1
        return l

class TestSearchInsert(unittest.TestCase):
  def test_searchInsert(self):
    case = [
      {'arg1': [1,3,5,6], 'arg2': 5, 'expected': 2},
      {'arg1': [1,3,5,6], 'arg2': 2, 'expected': 1},
      {'arg1': [1,3,5,6], 'arg2': 7, 'expected': 4},
    ]

    for c in case:
      self.assertEqual(c['expected'], Solution().searchInsert(c['arg1'], c['arg2']))
      
if __name__ == "__main__":
    unittest.main()