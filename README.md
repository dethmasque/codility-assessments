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
Executing test case: Triple and a pairRank parsed as: 10 Suit parsed as: H
Rank parsed as: 10 Suit parsed as: C
Rank parsed as: 10 Suit parsed as: S
Rank parsed as: 2 Suit parsed as: H
Rank parsed as: 2 Suit parsed as: S
Rank parsed as: 6 Suit parsed as: S
Rank parsed as: K Suit parsed as: H
Received 7 cards.
Your best set from the given hand is a a triple and a pair.
The selected cards are: [10H 10C 10S 2H 2S]
Executing test case: SingleRank parsed as: K Suit parsed as: S
Received 1 cards.
Your best set from the given hand is a single card.
The selected cards are: [KS]
Executing test case: Five in a row and a pairRank parsed as: 3 Suit parsed as: H
Rank parsed as: 4 Suit parsed as: S
Rank parsed as: 5 Suit parsed as: S
Rank parsed as: 6 Suit parsed as: C
Rank parsed as: 7 Suit parsed as: C
Rank parsed as: 8 Suit parsed as: H
Rank parsed as: A Suit parsed as: C
Rank parsed as: A Suit parsed as: S
Received 8 cards.
Your best set from the given hand is a five in a row.
The selected cards are: [8H 7C 6C 5S 4S]
Executing test case: Five of a suitRank parsed as: 2 Suit parsed as: D
Rank parsed as: 4 Suit parsed as: D
Rank parsed as: 6 Suit parsed as: D
Rank parsed as: 8 Suit parsed as: D
Rank parsed as: 10 Suit parsed as: D
Rank parsed as: Q Suit parsed as: D
Received 6 cards.
Your best set from the given hand is a suit.
The selected cards are: [2D 4D 6D 8D 10D]
Executing test case: Triple and a pairRank parsed as: J Suit parsed as: H
Rank parsed as: J Suit parsed as: D
Rank parsed as: J Suit parsed as: C
Rank parsed as: 7 Suit parsed as: H
Rank parsed as: 7 Suit parsed as: C
Rank parsed as: 8 Suit parsed as: H
Rank parsed as: 8 Suit parsed as: C
Received 7 cards.
Your best set from the given hand is a a triple and a pair.
The selected cards are: [JH JD JC 8H 8C]
Executing test case: TripleRank parsed as: 5 Suit parsed as: H
Rank parsed as: 5 Suit parsed as: D
Rank parsed as: 5 Suit parsed as: C
Rank parsed as: 5 Suit parsed as: S
Rank parsed as: 9 Suit parsed as: D
Received 5 cards.
Your best set from the given hand is a triple.
The selected cards are: [5H 5D 5C]
Executing test case: No handBest set from an empty hand is an empty set.
Executing test case: Bad card length (short)Received invalid playing card. No best set.
Executing test case: Bad card lenth (long)Received invalid playing card. No best set.
Executing test case: Two pairsRank parsed as: 3 Suit parsed as: S
Rank parsed as: 3 Suit parsed as: H
Rank parsed as: J Suit parsed as: H
Rank parsed as: J Suit parsed as: D
Received 4 cards.
Your best set from the given hand is a pair.
The selected cards are: [JH JD]
```

Non-ASCII characters are not supported. Hand input is 
checked for length (no cards received) as well as 
individual cards (minimum length of 2 and maximum length 
of 3).
