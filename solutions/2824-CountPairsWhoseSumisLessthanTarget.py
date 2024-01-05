class Solution:
    def countPairs(self, nums: List[int], target: int) -> int:
        result = 0
        for i in range(len(nums)):
            for j in range(len(nums)):
                if i < j and nums[i] + nums[j] < target:
                    result += 1
        return result
