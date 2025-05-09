# Codility 90-minute Technical Assessment - Golang

## Assessment overview

Given a hand of 1 to 10 cards, return the name of the 
best matching set along with the selected cards. 

Supported sets:

* single
* pair
* triple
* five in a row
* five of a kind (suit)
* triple and a pair

## How to run the Go example

The test cards are hard-coded and can be updated in 
`main()`. 

From the repo root, `go run ./main.go`

```
> go run .\main.go
Executing test case: Triple and a pair
Your best set from the given hand is a a triple and a pair.
The selected cards are: [10H 10C 10S 2H 2S]
Executing test case: Single
Your best set from the given hand is a single card.
The selected cards are: [KS]
Executing test case: Five in a row and a pair
Your best set from the given hand is a five in a row.
The selected cards are: [8H 7C 6C 5S 4S]
Executing test case: Five of a suit
Your best set from the given hand is a suit.
The selected cards are: [2D 4D 6D 8D 10D]
Executing test case: Triple and a pair
Your best set from the given hand is a a triple and a pair.
The selected cards are: [JH JD JC 8H 8C]
Executing test case: Triple
Your best set from the given hand is a triple.
The selected cards are: [5H 5D 5C]
Executing test case: No hand
Best set from an empty hand is an empty set.
Executing test case: Bad card length (short)
Received invalid playing card. No best set.
Executing test case: Bad card lenth (long)
Received invalid playing card. No best set.
Executing test case: Two pairs
Your best set from the given hand is a pair.
The selected cards are: [JH JD]
```

Non-ASCII characters are not supported. Hand input is 
checked for length (no cards received) as well as 
individual cards (minimum length of 2 and maximum length 
of 3).
