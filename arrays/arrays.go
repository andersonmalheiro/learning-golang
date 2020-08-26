package arrays

// Arrays cannot be resized and are prealocated with the length passed in declaration
var a [5]int // Creates an array with length of five, with zeroes as the values

// Slices are a reference to the sliced array
// They work exactly like in Python
var slice = a[1:3] // An slice of a, from index 1 to 3, but the value of index 3 is not included

// As a slice is a reference to the original array, changing one value inside the slice will update the original array
// slice[0] = 5 will change the corresponding index on original array

var array = [5]int{1, 2, 3, 4, 5} // Declaration of an array with initial values
// slice starting index is inclusive but the ending index is exclusive
// so in an array of [1,2,3,4,5], a slice of [0: 4] will return [1,2,3,4]

var length = len(array)   // Length of an array
var capacity = cap(array) // Capacity is the max number of values the slice can have

// 'make' method generates an slice based on the params passed
var dynamicSlice = make([]int, 2, 5) // This creates a slice with length of 5 and 2 positions filled

var x []int                       // An empty slice
var appended = append(x, 2, 3, 4) // Appends 2, 3, 4 to the slice
