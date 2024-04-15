import unittest
import sys
import random

sys.path.append("..")
from algorithms.NumberofGoodPairs_1512.resolve import Solution

class TestSolution(unittest.TestCase):
    def test_no_pair(self):
      n = [1,2,3]
      excepted = 0
      result = Solution.numIdenticalPairs(self, n)
      self.assertEqual(result, excepted)
    def test_1_pair(self):
      n = [1,2,3,4,5,6,7,8,9,10,1]
      excepted = 1
      result = Solution.numIdenticalPairs(self, n)
      self.assertEqual(result, excepted)
    def test_3_pair(self):
      n = [1,1,1]
      excepted = 3
      result = Solution.numIdenticalPairs(self, n)
      self.assertEqual(result, excepted)

if __name__ == '__main__':
    unittest.main()
