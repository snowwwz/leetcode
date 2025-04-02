/**
 * 1342. Number of Steps to Reduce a Number to Zero
 * 
 * Given an integer num, return the number of steps to reduce it to zero.
 * In one step, if the current number is even, you have to divide it by 2, otherwise, you have to subtract 1 from it.
 */
class Solution {
    public int numberOfSteps(int num) {
        int count = 0;
        while (num > 0) {
            if (num % 2 == 0) {  
                num /= 2;
            } else {
                num--;
            }
            count++;
        }
        return count;
    }
}

// 偶数の最下位ビットは0、奇数は1、x & 1(0001)で奇数なら1、偶数なら0になる
// num % 2 == 0 よりnum & 1 の方が計算コストが低く、速い。
// %（剰余演算）より &（ビット演算）のほうがCPU命令的に単純
class Solution {
    public int numberOfSteps(int num) {
        int count = 0;
        while (num > 0) {
            if ((num & 2) == 0) {  
                num = num >> 1; // 右シフトを使うと2 の累乗での割り算、>> n = 2^n で割った結果
            } else {
                num--;
            }
            count++;
        }
        return count;
    }
}
