class Solution:
    def mostWordsFound(self, sentences: List[str]) -> int:
        max = 0
        for i in range(len(sentences)):
            target = sentences[i].count(' ')+ 1
            if target > max:
                max = target
        return max
