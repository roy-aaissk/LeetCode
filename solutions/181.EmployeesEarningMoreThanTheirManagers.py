class Solution:
    def isPalindrome(self, x: int) -> bool:
        # sliceのstepを反転して実行することで取得できる
        reverse = str(x)[::-1]
        if reverse == str(x):
            return True
        return False
        
# best answer
# class Solution:
    # def isPalindrome(self, x: int) -> bool:
    # 負の値の場合は回文にならないので除外
    #     if x < 0:
    #         return False

    #     temp = x
    #     reversed_no = 0 
    #     while temp != 0:
    #         digit = temp % 10
    #         reversed_no = reversed_no * 10 + digit 
    #         temp //= 10
        
    #     return x == reversed_no

# digit = temp % 10: tempを10で割った余りをdigitに代入します。これにより、tempの最後の桁がdigitに格納されます。
# reversed_no = reversed_no * 10 + digit: reversed_noを10倍してからdigitを加えます。これにより、reversed_noが徐々に逆転していきます。例えば、最初の桁が1であれば、reversed_noは1です。次に最後の桁が2であれば、reversed_noは12になります。この操作を繰り返すことで、reversed_noは元の数を逆順にした値になります。

# temp //= 10: tempを10で割った商をtempに再代入します。これにより、tempから最後の桁が取り除かれ、次の桁に進みます。

# このプロセスを繰り返すことで、元の数xが逆順になった数reversed_noが構築されます。最終的に、xとreversed_noが等しいかどうかを比較し、等しい場合は回文であると判定されます。

# 例えば、x = 121の場合：
# 最初のループでは digit = 1、reversed_no = 1 になります。
# 2回目のループでは digit = 2、reversed_no = 12 になります。
# 3回目のループでは digit = 1、reversed_no = 121 になります。
