class Solution:
    def decode(self, encoded: List[int], first: int) -> List[int]:
        n = len(encoded) + 1
        arr = [first] + [0] * (n-1)
        for i in range(1, n):
            arr[i] = arr[i-1] ^ encoded[i-1]
        return arr
