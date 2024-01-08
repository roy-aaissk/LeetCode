class Solution:
    def minPartitions(self, n: str) -> int:
        # nを構成するために必要な最小のdeci-binary数
        min_decis = 0
        
        # nの各桁に対して処理を行う
        for digit in n:
            # 各桁の数字をdeci-binary数として加算
            min_decis = max(min_decis, int(digit))
        
        return min_decis

# 以下だと最小でできる(文字列をアンパックをしたlistを作り最大値を数値に変換して返している)
# return int(max[*n])
