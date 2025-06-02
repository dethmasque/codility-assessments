# Air Traffic Control Problem — “Runway Scheduler”

## Task Description

An airport has a single runway and multiple planes scheduled to land. Each plane is assigned a **landing window**, defined by an earliest time it can land and the latest time it must land.

A plane can only land if the runway is **free**, and only **one plane can land at any given time**. The goal is to **schedule the maximum number of planes** to land safely without overlapping landing times, and still within their allowed landing windows.

You are given two arrays `A` and `B` of the same length `N`:

- `A[i]` is the **earliest minute** plane `i` can land.  
- `B[i]` is the **latest minute** plane `i` must land.

Your task is to return the **maximum number of planes** that can be scheduled to land on the runway.

---

## Example

```go
A = [1, 3, 0, 5, 8, 5]
B = [2, 4, 6, 7, 9, 9]
```

In this case, the optimal number of landings is `4`.

One possible valid schedule:
- Plane 0 lands at minute 2
- Plane 1 lands at minute 4
- Plane 3 lands at minute 7
- Plane 4 lands at minute 9

---

## Function Signature

```go
func MaxLandings(A []int, B []int) int
```

Given:
- Two arrays `A` and `B` of equal length `N`
- Returns the **maximum number of planes** that can land

---

## Constraints

- `N` is an integer within the range `[1..100,000]`
- `0 ≤ A[i] ≤ B[i] ≤ 1,000,000,000`

---

## Hints

- Consider using a **greedy algorithm**.
- Try **sorting planes by their latest landing time** (`B[i]`) and then selecting non-overlapping windows.

# Solution

## Output

```
dethmasque@MacBookPro max-landings % go run main.go                            
Checking plane 0 with start 4 and end 4 
Found free landing time at 4 
Checking plane 1 with start 4 and end 5 
Found free landing time at 5 
Checking plane 2 with start 0 and end 1 
Found free landing time at 0 
Checking plane 3 with start 1 and end 3 
Found free landing time at 1 
Checking plane 4 with start 10 and end 12 
Found free landing time at 10 
Checking plane 5 with start 2 and end 5 
Found free landing time at 2 
Checking plane 6 with start 2 and end 7 
Found free landing time at 3 
Checking plane 7 with start 8 and end 15 
Found free landing time at 8 
Checking plane 8 with start 0 and end 15 
Found free landing time at 6 
Number of planes that can land safely: 9 
```
