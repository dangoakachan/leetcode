package com.kodango.leetcode;

public class Problem5 {
    private static String changeString(String s) {
        StringBuffer sb = new StringBuffer("#");

        for (int i = 0, len = s.length(); i < len; i++) {
            sb.append(s.charAt(i));
            sb.append("#");
        }

        return sb.toString();
    }

    private static String revertString(String s) {
        StringBuffer sb = new StringBuffer();

        for (int i = 1, len = s.length(); i < len; i += 2) {
            sb.append(s.charAt(i));
        }

        return sb.toString();
    }

    // Version: v1
    public String longestPalindromeV1(String s) {
        int len = s.length();

        if (len == 0) {
            return "";
        } else if (len == 1) {
            return s;
        }

        int maxStart = 0, maxEnd = 0;

        s = changeString(s);
        len = s.length();

        for (int i = 1; i < len; i++) {
            int start = i, end = i;

            while (start - 1 >= 0 && end + 1 < len && s.charAt(start - 1) == s.charAt(end + 1)) {
                start--;
                end++;
            }

            if (end - start > maxEnd - maxStart) {
                maxStart = start;
                maxEnd = end;
            }

            if (end == len - 1)
                break;
        }

        return revertString(s.substring(maxStart, maxEnd + 1));
    }

    // Version V2
    public String longestPalindromeV2(String s) {
        return s;
    }

    public void runCase(String s) {
        System.out.println(longestPalindromeV1(s));
        //System.out.println(longestPalindromeV2(s));
    }

    public static void main(String[] args) {
        Problem5 p = new Problem5();

        p.runCase("");
        p.runCase("b");
        p.runCase("bb");
        p.runCase("bbb");
        p.runCase("cbbd");
        p.runCase("babad");
        p.runCase("abcdefghijk");
        p.runCase("2017102");
        p.runCase("20171021");
        p.runCase("002017102");
    }
}
