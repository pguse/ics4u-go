# Algorithms - Sorting - Using Functions & Slices

In this course we need to investigate different types of algorithms that are commonly used to manipulate data.  One topic that is quite relevant is how to efficiently sort a large amount of data, to make it easier to work with.  Read following examples of sorting algorithms and watch the accompanying video:

## Selection Sort

The selection sort is a sorting algorithm where each item in a list is examined to find the item that should be first. For an slice of numbers this would be the smallest. For strings, it would be the alphabetically least. When the smallest item is found it is moved to the beginning of the list. Then the slice from the 2nd entry to the last is examined to select the smallest (least) item, which is then moved to the second position. This continues until the second last entry is reached.

<a href="http://www.youtube.com/watch?feature=player_embedded&v=3hH8kTHFw2A
" target="_blank"><img src="http://img.youtube.com/vi/3hH8kTHFw2A/0.jpg" 
alt="IMAGE ALT TEXT HERE" width="240" height="180" border="10" /></a>

## Bubble Sort

The bubble sort works by swapping adjacent pairs in a slice until all adjacent pairs are in order, at which time the entire slice is sorted. It does this by making repeated passes through the slice. The first pass compares element 1 and 2 and swaps them if they are out of order, then compares element 2 to element 3 and swaps them if they are out of order, and so on. Each pass moves from left to right across the slice, comparing the two elements in each pair and swapping the elements if they are out of order.

<a href="http://www.youtube.com/watch?feature=player_embedded&v=RT-hUXUWQ2I
" target="_blank"><img src="http://img.youtube.com/vi/RT-hUXUWQ2I/0.jpg" 
alt="IMAGE ALT TEXT HERE" width="240" height="180" border="10" /></a>

## Sorting Exercises

1. **Selection Sort**. Write a function called selSort(a []int) that will sort a slice **a** of integers , using the selection sort algorithm described above, and **return a sorted slice**.
	
2. **Bubble Sort**. Write a function called bubSort(a []int) that will sort a slice **a** of integers , using the bubble sort algorithm described above, and **return a sorted slice**.
	
3. **Shuffle** _(Demo in Class)_:  We now want to take a **sorted slice** and shuffle it so that the values are in a somewhat random order.  Design an algorithm that would allow you to shuffle a slice.  Write a function called shuffle(a []int) that will implement your algorithm, **returning a shuffled slice**. 
