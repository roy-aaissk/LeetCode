import unittest
import sys
import random

sys.path.append("..")
from algorithms.SmallestEvenMultiple.resolve import Solution

class TestSolution(unittest.TestCase):
    def randomNumber(min, max :int) -> int:
      return random.uniform(min, max)
    def test_even_number(self):
        n = 6
        expected = 6
        result = Solution.smallestEvenMultiple(self,n)
        self.assertEqual(result, expected)
    def test_max_number(self):
        n = 150
        expected = 150
        result = Solution.smallestEvenMultiple(self,n)
        self.assertEqual(result, expected)
    def test_min_number(self):
        n = 2
        expected = 2
        result = Solution.smallestEvenMultiple(self,n)
        self.assertEqual(result, expected)

if __name__ == '__main__':
    unittest.main()
