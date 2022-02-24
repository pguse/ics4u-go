# Algorithms - Searching - Using Functions & Slices

## Linear Search

A linear search is an algorithm that we can use to find an value in an slice.  The algorithm begin by looking at the first element in the slice, and then moves to the next element and the next element, in sequence, until it finds the value or reaches the end of the slice.

The main advantage of the linear search include:  1.  It is a simple algorithm; 2. It does not require the slice to be sorted in any order.  The main disadvantage is that if you have many elements in an slice, it is a slow algorithm.  In the worst case, if the element is at the end of the slice, every element must be looked at.  On average, half the element must be looked at to find the value being searched for, if it exists in the slice.

## Linear Search

<a href="http://www.youtube.com/watch?feature=player_embedded&v=TwsgCHYmbbA
" target="_blank"><img src="http://img.youtube.com/vi/TwsgCHYmbbA/0.jpg" 
alt="IMAGE ALT TEXT HERE" width="240" height="180" border="10" /></a>


## Binary Search

A binary search is also an algorithm that we can use to find an value in an slice.  However, it goes about the process in a very different way compared with a linear search.  First of all, the slice must be sorted - this will be used to our advantage.  When looking for a value, we can first look at the middle element.  If the middle element is lower than the value being searched for, we know we only need to look at the second half of the array in the next step.  If the middle element is higher than the value being searched for, we only need to look in the first half of the slice in the next step.  So, with each step, we reduce to number of elements to be searched by half.  In each step, we look at the middle element of the portion of the slice we are searching, and divide the slice in half depending on whether the middle element is higher or lower than the value being searched for.  This search algorithm is the best way of playing the guessing game that you may have created in 11 Computer Science.  The main advantage of this algorithm is its speed.  For example, to find a value in an slice of 100 elements only requires 7 steps, since 2^7= 128, and our array is smaller than this power of 2.  Its main disadvantages is 1. The array must first be sorted; 2. The algorithm is more complex than the linear search.

Here is a visual description of how the algorithm works:

For example, let's say we are looking for the value 25.  Here is visually how the algorithm would work.

-4 	-2 	0 	3 	4 	5 	7 	18 	25 
-4 	-2 	0 	3 	4 	5 	7 	18 	25 
-4 	-2 	0 	3 	4 	5 	7 	18 	25 
-4 	-2 	0 	3 	4 	5 	7 	18 	25 

There are 4 steps in this example, since 2^4= 16 and our array has a length of 9.  The number of steps lies between 2^3= 8 and 2^4= 16, so the larger number of steps is required.


## Binary Search


## Tasks

	1. Write a function **linSearch(n int, values []int)**, that searches for the integer n in the integer slice values and returns the index where the item is found or -1 if **n** is not in the slice.
	2. Write a function **binSearch(n int, values []int)**, that searches for the integer item in the integer slice values and returns the index where the item is found or -1 if **n** is not in the slice.
	3. Demonstrate the use of both functions on an integer slice of your own. 



