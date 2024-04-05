class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        result = [0,0]
        k = 0
        i = 1
        while k+2 < len(nums):
            if nums[k] + nums[i] == target:
                break
            if i < len(nums)-1:
                i = i+1
            else:
                k= k+1
                i= k+1
        result = [k, i]
        return result
