class Solution:
    def differenceOfSums(self, n: int, m: int) -> int:
      notDivisible = 0
      divisible = 0
      for i in list(range(1, n+1)):
          if i % m == 0:
              divisible = divisible + i
              continue
          notDivisible = notDivisible + i
      return notDivisible- divisible
