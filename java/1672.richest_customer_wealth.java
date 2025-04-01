/**
 * 1672. Richest Customer Wealth
 * 
 * You are given an m x n integer grid accounts where accounts[i][j] is the amount of money the i​​​​​​​​​​​th​​​​ customer has in the j​​​​​​​​​​​th​​​​ bank. Return the wealth that the richest customer has.
 * A customer's wealth is the amount of money they have in all their bank accounts. The richest customer is the customer that has the maximum wealth.
 * 
 * time complexity: O(n x m)
 * space complexity: O(1) 
 */

class Solution {
    public int maximumWealth(int[][] accounts) {
        int max = 0;
        for (int i = 0; i < accounts.length; i++){
           int sum = 0;
           for (int j = 0; j < accounts[i].length; j++) {
            sum += accounts[i][j];
           }
           if (sum > max) max = sum;
        }
        return max;
    }
}

/**
 * alternative solution
 */
class Solution {
    public int maximumWealth(int[][] accounts) {
        int max = 0;
        for (int[] customer: accounts){
           int sum = 0;
           for (int bank: customer) {
            sum += bank;
           }
           max = Math.max(max, sum);
        }
        return max;
    }
}
