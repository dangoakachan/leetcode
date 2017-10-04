package com.kodango.leetcode;

/*
 * Problem 27: Remove Element
 * https://leetcode.com/problems/remove-element/description/
 */

import java.util.Arrays;

public class Problem27 {
    public int removeElement1(int[] nums, int val) {
        int i = 0;
        int j = nums.length;

        while (i < j) {
            if (nums[i] != val) { // point to target value
                i++;
                continue;
            }

            if (nums[j-1] == val) { // point to non-target value
                j--;
                continue;
            }

            // Swap the target value backward
            nums[i] ^= nums[j-1];
            nums[j-1] ^= nums[i];
            nums[i] ^= nums[j-1];

            // Skip the non-target value
            i++;
        }

        return i;
    }

    public int removeElement2(int[] nums, int val) {
        int i = 0;

        for (int j = 0; j < nums.length; j++) {
            if (nums[j] != val) {
                nums[i++] = nums[j];
            }
        }

        return i;
    }

    public int removeElement(int[] nums, int val) {
        return removeElement2(nums, val);
    }

    public void runCase(int[] nums, int val) {
        int ret = removeElement(nums, val);
        System.out.println("Result: " + Arrays.toString(nums) + ", new length: " + ret);
    }

    public static void main(String[] args) {
        Problem27 p = new Problem27();

        p.runCase(new int[]{}, 3);
        p.runCase(new int[]{3}, 3);
        p.runCase(new int[]{2}, 3);
        p.runCase(new int[]{3, 2, 2, 3}, 3);
        p.runCase(new int[]{3, 2, 2, 3}, 4);
        p.runCase(new int[]{1, 2, 2, 3}, 3);
        p.runCase(new int[]{3, 3, 1, 3}, 3);
        p.runCase(new int[]{3, 3, 1, 3, 4}, 3);
        p.runCase(new int[]{2, 3, 1, 3, 4}, 3);
        p.runCase(new int[]{2, 3, 1, 3, 4}, 5);
    }
}