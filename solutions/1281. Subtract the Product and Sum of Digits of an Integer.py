class Solution:
    def subtractProductAndSum(self, n: int) -> int:
        x =[int(a) for a in str(n)]
        productDigits = 1
        sumDigits = 0
        for i in range(len(x)):
            productDigits = productDigits * x[i]
            sumDigits += x[i]
        result = productDigits - sumDigits
        return result
