/**
 * 1432. Max Difference You Can Get From Changing an Integer
 * 
 * You are given an integer num. You will apply the following steps to num two separate times:
 * Pick a digit x (0 <= x <= 9).
 * Pick another digit y (0 <= y <= 9). Note y can be equal to x.
 * Replace all the occurrences of x in the decimal representation of num by y.
 * Let a and b be the two results from applying the operation to num independently.
 * 
 * Return the max difference between a and b.
 * Note that neither a nor b may have any leading zeros, and must not be 0.
 */

class Solution {
    public int maxDiff(int num) {
        String numStr = String.valueOf(num);
        int max = num;
        int min = num;

        boolean firstLetter = false;
        for (char x: numStr.toCharArray()) {
            for (char y = '0'; y <= '9'; y++) {
                String replaced = numStr.replace(x, y);
                int newNum = Integer.parseInt(replaced);

                // 桁が変わっていたらng
                if (numStr.length() != String.valueOf(newNum).length()) {
                    continue;
                }

                max = Math.max(max, newNum);
                min = (Math.min(min, newNum) == 0) ? min : Math.min(min, newNum); 
            }
        }
        return max - min;
    }
}

