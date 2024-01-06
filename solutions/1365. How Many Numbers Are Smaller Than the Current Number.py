class Solution:
    def smallerNumbersThanCurrent(self, nums: List[int]) -> List[int]:
        result = []
        for i in range(len(nums)):
            num = 0
            for j in range(len(nums)):
                if i == j:
                    continue
                if nums[j] < nums[i]:
                    num += 1
            result.append(num)
        return result

# こんな書き方もある
# class Solution:
    # def smallerNumbersThanCurrent(self, nums: List[int]) -> List[int]:
    #     result = nums[::]
    #     result.sort()
    #     position = {}
    #     for i, n in enumerate(result):
    #         if n not in position:
    #             position[n] = i
    #     for i, n in enumerate(nums):
    #         result[i] = position[n]
    #     return result
