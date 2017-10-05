package com.kodango.leetcode;

import java.util.Arrays;

/*
 * Problem 665: Non-decreasing Array
 * https://leetcode.com/problems/non-decreasing-array/description/
 */
public class Problem665 {
    public boolean checkPossibility(int[] nums) {
        if (nums == null || nums.length < 3) {
            return true;
        }

        boolean found = false;

        for (int i = 1; i < nums.length; i++) {
            if (nums[i] >= nums[i-1]) {
                continue;
            }

            if (found) {
                return false;
            }

            if (i >= 2 && nums[i] < nums[i-2]) {
                nums[i] = nums[i-1];
            }

            found = true;
        }


        return true;
    }

    public static void main(String[] args) {
        Problem665 p = new Problem665();

        System.out.println(p.checkPossibility(new int[]{4, 2})); // true
        System.out.println(p.checkPossibility(new int[]{4, 2, 3})); // true
        System.out.println(p.checkPossibility(new int[]{4, 2, 1})); // false
    }
}