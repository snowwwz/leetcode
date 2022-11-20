# https://leetcode.com/problems/first-bad-version

# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

class Solution:
    def isBadVersion(self, n: int, m: int):
      if n == 5 and m == 4:
        return True
      if n == 1 and m == 1:
        return True
      return False

    def firstBadVersion(self, n: int) -> int:
        l = 1
        r = n
        re = 1
        while l <= r:
            m = (l+r)//2
            if self.isBadVersion(n, m):
                r = m - 1
                re = m
            else:
                l = m + 1
        return re

import unittest

class TestFirstBadVersion(unittest.TestCase):
  def test_firstBadVersion(self):
    case = [
      {'arg': 5, 'expected': 4},
      {'arg': 1, 'expected': 1},
    ]

    for c in case:
      self.assertEqual(c['expected'], Solution().firstBadVersion(c['arg']))
      
if __name__ == "__main__":
    unittest.main()