# SLICE WINDOW

## Signs to Use Sliding Window
Consecutive Array or String:

1. You are working with an array or string and need to process consecutive elements.
    - Example: Summing the elements in a subarray, finding the length of the subarray that satisfies a certain condition, etc.
Fixed Size "Window":

2. When you need to work with a window of fixed size or sliding window, such as summing m consecutive elements in the array.
    - Example: Calculating the sum of m consecutive elements in an array.
Quickly Update Values:

3. You need to quickly update values for each element as the window slides. The sliding window technique allows you to efficiently add the new element and remove the old element without recalculating everything from the beginning.
    - Example: If you need to calculate the sum of elements in the window, each time the window slides right, you just add the next element and subtract the first element from the sum.
Search Within Consecutive Subarrays:

4. Problems that require searching or checking a subarray in the array that has a specific property, for example, finding the subarray whose sum is less than or equal to a value, or finding subarrays that satisfy a condition.
    - Example: Finding the number of ways to split an array into subgroups with a sum equal to a given value, such as in the birthday cake problem you asked about.


## Summary
The Sliding Window technique is suitable when you:

- Need to process consecutive elements in an array.
- Need to optimize re-computing values for each element in the window (adding/removing elements).
- Need to work with a fixed-size window or one that can change with clear conditions.<br>

If your problem has these characteristics, Sliding Window is a great choice and can significantly optimize the running time of your algorithm.