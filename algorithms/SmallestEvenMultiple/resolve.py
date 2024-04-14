class Solution:
    def __init__(self) -> None:
      pass
    def smallestEvenMultiple(self, n: int) -> int:
      if 1 <= n <= 150:
        ValueError("invalid number %i", n)
      if n % 2 == 0:
        return n
      else:
        return n * 2
