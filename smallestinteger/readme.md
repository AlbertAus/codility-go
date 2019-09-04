Smallest positive integer or Missing Integer

Write a function:

func Solution(A []int) int

that, given an array A of N integers, returns the smallest positive integer (greater than 0) that does not occur in A.

For example, given A = [1, 3, 6, 4, 1, 2], the function should return 5.

Given A = [1, 2, 3], the function should return 4.

Given A = [−1, −3], the function should return 1.

Write an efficient algorithm for the following assumptions:

N is an integer within the range [1..100,000];
each element of array A is an integer within the range [−1,000,000..1,000,000].


#####################################################################################################
Test result:

The solution obtained perfect score.

Analysis
Detected time complexity:
O(N) or O(N * log(N))
expand allExample tests
▶ example1 
first example test ✔OK
▶ example2 
second example test ✔OK
▶ example3 
third example test ✔OK
expand allCorrectness tests
▶ extreme_single 
a single element ✔OK
▶ simple 
simple test ✔OK
▶ extreme_min_max_value 
minimal and maximal values ✔OK
▶ positive_only 
shuffled sequence of 0...100 and then 102...200 ✔OK
▶ negative_only 
shuffled sequence -100 ... -1 ✔OK
expand allPerformance tests
▶ medium 
chaotic sequences length=10005 (with minus) ✔OK
▶ large_1 
chaotic + sequence 1, 2, ..., 40000 (without minus) ✔OK
▶ large_2 
shuffled sequence 1, 2, ..., 100000 (without minus) ✔OK
▶ large_3 
chaotic + many -1, 1, 2, 3 (with minus) ✔OK
