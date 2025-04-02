/**
 * 412. Fizz Buzz
 * 
 * Given an integer n, return a string array answer (1-indexed) where:

answer[i] == "FizzBuzz" if i is divisible by 3 and 5.
answer[i] == "Fizz" if i is divisible by 3.
answer[i] == "Buzz" if i is divisible by 5.
answer[i] == i (as a string) if none of the above conditions are true.
 */
class Solution {
    public List<String> fizzBuzz(int n) {
        List<String> answer = new ArrayList<>();
        for (int i = 1; i <= n; i++) {
            if (i % 3 == 0 && i % 5 == 0) {
                answer.add("FizzBuzz");
                continue;
            } 
            if (i % 3 == 0) {
                answer.add("Fizz");
                continue;
            } 
           if (i % 5 == 0) {
                answer.add("Buzz");
                continue;
            } 
            answer.add(String.valueOf(i));
        }
        return answer;
    }
}

class Solution {
    public List<String> fizzBuzz(int n) {
        List<String> answer = new ArrayList<>();
        for (int i = 1; i <= n; i++) {
            boolean devidedBy3 = i % 3 == 0;
            boolean devidedBy5 = i % 5 == 0;

            String currStr = "";
            if (devidedBy3) {
                currStr += "Fizz";
            } 
            if (devidedBy5) {
                currStr += "Buzz";
            } 

           if (currStr.isEmpty()) {
                currStr += String.valueOf(i);
            } 
        }
        return answer;
    }
}
