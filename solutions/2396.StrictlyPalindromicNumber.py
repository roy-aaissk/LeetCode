class Solution:
    def is_palindrome(s):
        return s == s[::-1]
    def isStrictlyPalindromic(self, n: int) -> bool:
        for base in range(2, n - 1):
        # nをbase進数に変換
        # base進数の変換は元の値が0より小さくなる前までn % baseの余りを求めると対象の進数の値に変換できる
            representation = ""
            temp_n = n
            while temp_n > 0:
                representation = str(temp_n % base) + representation
                temp_n //= base
        
        # 回文かどうか判定
        if representation != representation[::-1]:
            return False

        return True
