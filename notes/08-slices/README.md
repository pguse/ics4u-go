# Slices

In Go, arrays are actually not used very often for a couple of reasons.  First, an array has a fixed size chosen when the array is first created. You cannot remove or add an element to an array, except by creating another array with this new size.

 Second, because its size is part of its type, i