class Solution:
    def interpret(self, command: str) -> str:
        result = ""
        for i, c in enumerate(command):
            if c == "G":
                result += "G"
            elif c == "(" and command[i+1] == ")":
                result += "o"
            elif c == "(" and command[i+1] == "a" and command[i+2] == "l" and command[i+3] == ")":
                result += "al"
        return result
