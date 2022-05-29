class Solution(object):
    def canReach(self, s, minJump, maxJump):
        """
        :type s: str
        :type minJump: int
        :type maxJump: int
        :rtype: bool
        """
        n = len(s)
        dp = [0] * n
        dp[0]= 1
        diff = 0
        for i in range(1, n):
            left = i-maxJump-1
            right = i-minJump
            # print('i',i, left, right, dp[left:right])
            if right >= 0 : diff += dp[right]
            if left >= 0 : diff -= dp[left]
            if s[i] == '0':
                dp[i] = 1 if diff > 0 else 0
            # print(dp, diff)
        return dp[-1] > 0