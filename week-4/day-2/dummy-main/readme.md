# Question: Sum of Odd Numbers

Write a Go program that receives a slice containing integers and returns the sum of odd numbers present in that slice.

**Skeleton code :**

```go

package main

func SumOddNumbers(numbers []int) int {
    // your code here
    return 0
}

```

**Test-Case :**

```go
SumOddNumbers([]int{1, 2, 3, 4, 5})    // Should return 9, as the sum of odd numbers is 1+3+5=9.
SumOddNumbers([]int{0, 2, 4, 6, 8})    // Should return 0, as there are no odd numbers in the slice.
SumOddNumbers([]int{1, 3, 5, 7, 9})    // Should return 25, as the sum of odd numbers is 1+3+5+7+9=25.
SumOddNumbers([]int{10, 11, 12, 13, 14, 15})    // Should return 39, as the sum of odd numbers is 11+13+15=39.
```
