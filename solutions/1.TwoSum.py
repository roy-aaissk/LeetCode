class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        result = [0,0]
        for i in range(len(nums)):
            for j in range(len(nums)-i -1):
                if nums[i] + nums[j+i+1] == target:
                    result[0], result[1] = i, j+i+1
                    break
            if result[0] + result[1] == target:
                    break
            result.sort()
        return result
