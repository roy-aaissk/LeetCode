class Solution:
    def freqAlphabets(self, s: str) -> str:
        alphabet_dict = {}
        result = ""

        # アルファベットのaからzまでをループで処理
        for i in range(26):
            # chrを使用してASCIIコードから文字に変換し、1を足して数字を得る
            letter = chr(97 + i)  # 97は小文字のaのASCIIコード

            # 辞書にキーと値を追加
            if i < 9:
                alphabet_dict[str(i + 1)] = letter
            else:
                alphabet_dict[str(i + 1) + "#"] = letter

        i = 0
        while i < len(s):
            if i + 2 < len(s) and s[i + 2] == '#':
                # 現在の数字の後に '#' が続いている場合の処理
                result += alphabet_dict[s[i:i + 3]]
                i += 3
            else:
                # 現在の数字の後に '#' が続いていない場合の処理
                result += alphabet_dict[s[i]]
                i += 1
        return result
