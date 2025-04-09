/**
 * 1295. Find Numbers with Even Number of Digits
 * 
 * Given an array nums of integers, return how many of them contain an even number of digits.
 */
class Solution {
    public int findNumbers(int[] nums) {
        int evenCnt = 0;
        for (int i = 0; i < nums.length; i++) {
            if (String.valueOf(nums[i]).length() % 2 == 0) {
                evenCnt++;
            }
        }
        return evenCnt;
    }
}
