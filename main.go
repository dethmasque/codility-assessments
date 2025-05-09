package main

import (
    "fmt"
    "sort"
    "strings"
    "time"
)

// best set name and the chosen matching cards
// from the provided hand
type Results struct {
    setName       string
    selectedCards []string
}

// corresponding integer value for rank strings
var rankMap = map[string]int{
    "2":  2,
    "3":  3,
    "4":  4,
    "5":  5,
    "6":  6,
    "7":  7,
    "8":  8,
    "9":  9,
    "10": 10,
    "J":  11,
    "Q":  12,
    "K":  13,
    "A":  14,
}

// suit rank matters when suit matters
var suitRankMap = map[string]int{
    "S": 4,
    "H": 3,
    "D": 2,
    "C": 1,
}

// parse card takes in a card string, ensures it is an appropriate playing
// card length, and that its rank and suit is within the allowed mapped
// ranges
func parseCard(card string) (rank string, suit string, playingCard bool) {
    if len(card) < 2 || len(card) > 3 {
        // "10" is the longest rank string, suit string is always 1,
        // and there's always a rank and a suit. ex: "2S"
        return "", "", false
    }
    suit = card[len(card)-1:] // suit is always the last character
    rank = card[:len(card)-1] // support rank 10! lol
    _, rankOk := rankMap[rank]
    _, suitOk := suitRankMap[suit]
    if !rankOk || !suitOk {
        // rank or suit is not in range of valid mapped
        // ranks and suits
        return "", "", false
    }
    // parsed valid playing card
    return rank, suit, true
}

// return the best set of cards (name of set and the cards themselves)
// from a provided hand of cards.
func Solution(cards []string) Results {
    // should we check for more conditions? parseCard() handles add'l cases
    if len(cards) == 0 {
        // no cards provided
        return Results{"Best set from an empty hand is an empty set.", []string{}}
    }

    // for returning a card after working with ranks (triple, etc.)
    rankToCard := map[string][]string{}
    rankCount := map[string][]string{} // for triple and pair
    suitCount := map[string][]string{} // for five of a kind ("suit")
    uniqueRanks := map[int]string{}    // for five in a row

    // parse received cards, store mapped values for comparisons and
    // storing results
    for _, card := range cards {
        rank, suit, valid := parseCard(strings.ToUpper(card))
        if !valid {
            return Results{"Received invalid playing card. No best set.", []string{}}
        }
        fmt.Printf("Rank parsed as: %s Suit parsed as: %s \n", rank, suit)
        // add card to map of received cards by rank
        rankCount[rank] = append(rankCount[rank], card)
        // add card to map of received cards by suit
        suitCount[suit] = append(suitCount[suit], card)
        // get the mapped rank value for parsed rank
        rankVal := rankMap[rank]
        // add card to other map of received cards by rank
        rankToCard[rank] = append(rankToCard[rank], card)
        // store unique parsed ranks at their mapped int value
        uniqueRanks[rankVal] = rank
    }
    fmt.Printf("Received %d cards.\n", len(cards))

    // in descending order of precedence because we want
    // to return the best card set given the hand
    // set 6: triple (same rank) and a pair (same rank)
    var tripleRankVal, pairRankVal int
    var tripleRankStr, pairRankStr string
    // should we check []card length up front?
    for rank, cards := range rankCount {
        if len(cards) >= 3 {
            rankVal := rankMap[rank]
            if rankVal > tripleRankVal {
                tripleRankVal = rankVal
                tripleRankStr = rank
            }
        }
    }
    for rank, cards := range rankCount {
        if len(cards) >= 2 {
            rankVal := rankMap[rank]
            if rankVal > pairRankVal && rankVal != tripleRankVal {
                pairRankVal = rankVal
                pairRankStr = rank
            }
        }
    }

    if tripleRankVal > 0 && pairRankVal > 0 {
        result := append(rankToCard[tripleRankStr][:3], rankToCard[pairRankStr][:2]...)
        return Results{"a triple and a pair", result}
    }

    // set 5: five of a kind (same suit)
    bestSuit := ""
    for suit, cards := range suitCount {
        if len(cards) >= 5 {
            // if it's our first encounter with five of a kind or if
            // it's a higher ranking suit than the previously
            // encountered set
            if bestSuit == "" || suitRankMap[suit] > suitRankMap[bestSuit] {
                bestSuit = suit
            }
        }
    }
    if bestSuit != "" {
        return Results{"suit", suitCount[bestSuit][:5]}
    }

    // set 4: five in a row
    sortedRanks := make([]int, 0, len(uniqueRanks))
    for rank := range uniqueRanks {
        sortedRanks = append(sortedRanks, rank)
    }
    sort.Sort(sort.Reverse(sort.IntSlice(sortedRanks)))
    for i := 0; i <= len(sortedRanks)-5; i++ {
        selectedRanks := sortedRanks[i : i+5]
        // assume in a row
        isInARow := true
        for j := 1; j < 5; j++ {
            // previous rank in slice should be -1, if not
            // it's not in a row
            if selectedRanks[j-1]-1 != selectedRanks[j] {
                isInARow = false
                break
            }
        }
        if isInARow {
            result := make([]string, 0, 5)
            for _, rankVal := range selectedRanks {
                rank := uniqueRanks[rankVal]
                result = append(result, rankToCard[rank][0])
            }
            return Results{"five in a row", result}
        }
    }

    // set 3: triple
    bestTriple := 0
    tripleRank := ""
    for rank, cards := range rankCount {
        if len(cards) >= 3 {
            rankVal := rankMap[rank]
            if rankVal > bestTriple {
                bestTriple = rankVal
                tripleRank = rank
            }
        }
    }
    if bestTriple > 0 {
        return Results{"triple", rankToCard[tripleRank][:3]}
    }

    // set 2: pair
    bestPair := 0
    pairRank := ""
    for rank, cards := range rankCount {
        if len(cards) >= 2 {
            rankVal := rankMap[rank]
            if rankVal > bestPair {
                bestPair = rankVal
                pairRank = rank
            }
        }
    }
    if bestPair > 0 {
        return Results{"pair", rankToCard[pairRank][:2]}
    }

    // set 1: single card
    bestVal := 0
    bestRank := ""
    for rank, rankVal := range rankMap {
        if len(rankToCard[rank]) > 0 && rankVal > bestVal {
            bestVal = rankVal
            bestRank = rank
        }
    }
    return Results{"single card", []string{rankToCard[bestRank][0]}}
}

func main() {
    // define test hands to cover each supported set and a couple
    // of error conditions
    type hand []string
    hand1 := hand{"10H", "10C", "10S", "2H", "2S", "6S", "KH"}    // triple and pair
    hand2 := hand{"KS"}                                           // single
    hand3 := hand{"3H", "4S", "5S", "6C", "7C", "8H", "AC", "AS"} // five in a row and a pair
    hand4 := hand{"2D", "4D", "6D", "8D", "10D", "QD"}            // five of a suit
    hand5 := hand{"JH", "JD", "JC", "7H", "7C", "8H", "8C"}       // triple and pair (two pairs)
    hand6 := hand{"5H", "5D", "5C", "5S", "9D"}                   // triple
    hand7 := hand{}                                               // no hand
    hand8 := hand{"4", "D"}                                       // bad length (short)
    hand9 := hand{"324S", "502C"}                                 // bad length (long)
    hand10 := hand{"3S", "3H", "JH", "JD"}                        // two pairs

    testHands := []hand{hand1, hand2, hand3, hand4, hand5, hand6, hand7, hand8, hand9, hand10}

    // loop over test hands, print results
    for _, hand := range testHands {
        bestSet := Solution(hand)
        if len(bestSet.selectedCards) > 0 {
            fmt.Printf("Your best set from the given hand is a %s. \n", bestSet.setName)
            fmt.Printf("The selected cards are: %s \n", bestSet.selectedCards)
        } else {
            fmt.Println(bestSet.setName) // error message stored as setName
        }
    }

    time.Unix()
}
