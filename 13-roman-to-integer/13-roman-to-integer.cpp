class Solution {
public:
    int romanToInt(string s) {
        map<char, int> mappedValues = {
          {'I', 1}, 
          {'V', 5},
          {'X', 10},
          {'L', 50},
          {'C', 100},
          {'D', 500},
          {'M', 1000}
        };

        int decimal = 0;

        for(int i = 0; i < s.length(); i++){

          // sum the characters
          decimal += mappedValues[s[i]];

          if(i == 0) {
            continue;
          }

          bool isEarlierCharSmall = mappedValues[s[i-1]] < mappedValues[s[i]];

          if(isEarlierCharSmall) {
            // remove the earlier summation and substract that character value
            decimal -= 2 * mappedValues[s[i-1]];
          }
        }
        return decimal;
    }
};