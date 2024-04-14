class Solution:
    def findWordsContaining(self, words: List[str], x: str) -> List[int]:
      result = []
      for i , word in enumerate(words):
        for key in list(word):
          if key == x:
            result += [i]
            break
      return result
