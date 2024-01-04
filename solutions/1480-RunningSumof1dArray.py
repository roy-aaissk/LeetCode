class Solution:
    def runningSum(self, nums: List[int]) -> List[int]:
        result = []
        for i in range(len(nums)):
            numbers = 0
            if i == 0:
                result.append(nums[i])
                continue
            for j in range(len(nums)):
                if j <= i:
                    numbers += nums[j]
            result.append(numbers)
        return result

# こんな書き方をするともっとスッキリして速くなる
# class Solution:
#     def runningSum(self, nums: List[int]) -> List[int]:
#         for i in range(1, len(nums)):
#             nums[i] = nums[i] + nums[i - 1]
#         return nums
