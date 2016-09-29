[Problem: Two Sum](https://leetcode.com/problems/two-sum/) 
======
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution
##h2

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```
Tags: Hastables, Array
Difficulty: Easy
Runtime ~ 9 ms ; > 77.78% of (golang) submissions as of 9/28/2016
Explanation:
=====
The first problem in leet code asks users to, given a list, find the indicies of two numbers so that:

`nums[index1] + nums[index2] = target `.

Fairly straight forward problem, this can easily be done in O(n^2) time using a simple brute force problem. Obviously while this solution works, and is totally valid, we should strive to get a more optimal solution.

The solution I chose utilizes hashtables and allows us to solve this problem in O(n) time. The idea is to, given a number at some index i, calculate the complement of the target and that number. We then take that number and store it in a hashtable along with the current index, and as we go down the list we check if the number at index i is in the hash table.

If the number is in the hashtable, then that means we have already found the complement, and we can immediately return and give the two indicies. 

If you don't remember what complements are, in the context of this problem, they're simply the difference between the number at index i, and the target. If the target is 9, and the current number is 2, the complement is 7 ; 9-2 = 7

