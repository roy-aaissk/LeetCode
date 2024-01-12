class Solution:
    def numberOfMatches(self, n: int) -> int:
        result = 0
        while n > 1:
            if n % 2 == 0:
                n = n/2
                result += n
            else:
                n = (n-1)/2
                result += n+1
        import math
        return math.floor(result)
