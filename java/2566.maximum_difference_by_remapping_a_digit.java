/**
 * 2566. Maximum Difference by Remapping a Digit
 * 
 * You are given an integer num. You know that Bob will sneakily remap one of the 10 possible digits (0 to 9) to another digit.
 * Return the difference between the maximum and minimum values Bob can make by remapping exactly one digit in num.
 * 
 * Notes:
 * When Bob remaps a digit d1 to another digit d2, Bob replaces all occurrences of d1 in num with d2.
 * Bob can remap a digit to itself, in which case num does not change.
 * Bob can remap different digits for obtaining minimum and maximum values respectively.
 * The resulting number after remapping can contain leading zeroes.
 */

class Solution {
    public int minMaxDifference(int num) {
        String numStr = String.valueOf(num);
        int max = num;
        int min = num;

        for (char d1 : numStr.toCharArray()) {
            for (char d2 = '0'; d2 <= '9'; d2++) {
                if (d1 == d2) continue;

                String remapped = numStr.replace(d1, d2);
                int newNum = Integer.parseInt(remapped);
                max = Math.max(max, newNum);
                min = Math.min(min, newNum);
            }
        }
        return max - min;
    }
}
