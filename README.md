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

Non-ASCII characters are not supported. Hand input is 
checked for length (no cards received) as well as 
individual cards (minimum length of 2 and maximum length 
of 3).
