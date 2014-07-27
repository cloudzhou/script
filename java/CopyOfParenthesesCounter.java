package test_jcseg;

public class CopyOfParenthesesCounter {

    private static final char LEFT  = '(';
    private static final char RIGHT = ')';

    public static void main(String[] args) {
        CopyOfParenthesesCounter counter = new CopyOfParenthesesCounter();

        System.out.println("ParenthesesCounter number is " + counter.cal("((ab)c(dfcc)()))"));
        System.out.println("ParenthesesCounter number is " + counter.cal("((ab)c(dfcc)()))(("));
        System.out.println("ParenthesesCounter number is " + counter.cal("((ab)c(dfcc)())"));
        System.out.println("ParenthesesCounter number is " + counter.cal("abc"));
        System.out.println("ParenthesesCounter number is " + counter.cal("(abc"));
        System.out.println("ParenthesesCounter number is " + counter.cal(")abc"));
        System.out.println("ParenthesesCounter number is " + counter.cal("()abc"));
        System.out.println("ParenthesesCounter number is " + counter.cal(""));
        String param = null;
        System.out.println("ParenthesesCounter number is " + counter.cal(param));
    }

    private int cal(String param) {
        if (param == null || param.isEmpty()) {
            return 0;
        }
        int stack = 0;
        int totalNum = 0;
        int compare = 0;
        for (char c : param.toCharArray()) {
            if (c == LEFT) {
                compare++;
                stack++;
            }
            if (c == RIGHT) {
                compare--;
                if (stack > 0) {
                    stack--;
                    totalNum++;
                }
            }
        }

        if (stack == 0) {
            System.out.println("the number of () is " + totalNum + ", the whole string are valid syntax");
        } else {
            System.out.print("the number of () is " + totalNum + ", the syntax is invalid. ");
            if (compare > 0) {
                System.out.println("there are " + compare + " consecutive left parenthesis");
            } else {
                System.out.println("there are " + (-compare) + " consecutive right parenthesis");
            }

        }
        return totalNum;
    }
}

