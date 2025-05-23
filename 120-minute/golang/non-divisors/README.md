## Count Non-Divisible — Problem Description

You are given an array `A` consisting of `N` integers.

For each number `A[i]` such that `0 ≤ i < N`, we want to count the number of elements of the array that are **not** the divisors of `A[i]`.  
We say that these elements are **non-divisors**.

---

### Example

Consider integer `N = 5` and array `A` such that:

```
A[0] = 3  
A[1] = 1  
A[2] = 2  
A[3] = 3  
A[4] = 6
```

For the following elements:

- `A[0] = 3`, the non-divisors are: `2, 6`
- `A[1] = 1`, the non-divisors are: `3, 2, 3, 6`
- `A[2] = 2`, the non-divisors are: `3, 3, 6`
- `A[3] = 3`, the non-divisors are: `2, 6`
- `A[4] = 6`, there aren't any non-divisors

---

## Task

Write a function:

```go
func Solution(A []int) []int
```

that, given an array `A` consisting of `N` integers, returns a sequence of integers representing the amount of **non-divisors** for each element.

The result array should be returned as a slice of integers.

---

## Example

Given:

```
A[0] = 3  
A[1] = 1  
A[2] = 2  
A[3] = 3  
A[4] = 6
```

The function should return:

```
[2, 4, 3, 2, 0]
```

as explained above.

---

## Constraints

- `N` is an integer within the range `[1..50,000]`;
- Each element of array `A` is an integer within the range `[1..2 * N]`

