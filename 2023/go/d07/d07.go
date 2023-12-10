package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type play struct {
	hand [5]int;
	bid int;
	cardsCount map[int]int;
	handType int
}

var cardTier = make(map[byte]int)

func main() {
	// Part 2
	cardList := "J23456789TQKA"
	for i, card := range cardList {
		cardTier[byte(card)] = i
	}

	plays := getData("assets/input1")

	fmt.Printf("Plays: %v\n", plays)

	score := getScore(plays)

	fmt.Printf("Score: %v\n", score)
}

func getData(filename string) []play {
	fileData, _ := os.ReadFile(filename)

	playsStr := strings.Split(string(fileData), "\n")
	playsStr = playsStr[:len(playsStr)-1]

	playsSize := len(playsStr)
	plays := make([]play, playsSize)

	for i := 0; i < playsSize; i++ {
		splitPlay := strings.Split(playsStr[i], " ")

		handStr := splitPlay[0]
		bid, _ := strconv.ParseInt(splitPlay[1], 10, 64)
		hand := [5]int{}
		cardsCount := make(map[int]int)

		highestRepeated := []int{0, 0}
		jCount := 0
		for c := 0; c < 5; c++ {
			hand[c] = cardTier[handStr[c]]
			cardsCount[hand[c]]++

			// Part 2
			isHighest := cardsCount[hand[c]] > highestRepeated[1] || (cardsCount[hand[c]] > highestRepeated[1]  && hand[c] > highestRepeated[0])
			if hand[c] == 0 {
				jCount++
			} else if isHighest {
				highestRepeated[0] = hand[c]
				highestRepeated[1] = cardsCount[hand[c]]
			}
		}

		// Part 2
		if highestRepeated[0] != 0 {
			cardsCount[highestRepeated[0]] += jCount
			delete(cardsCount, 0)
		}

		plays[i] = play {
			bid: int(bid),
			hand: hand,
			cardsCount: cardsCount,
		}
	}

	return plays
}

func getScore(plays []play) int {
	score := 0

	playsSize := len(plays)

	for i := 0; i < playsSize; i++ {
		plays[i].handType = getHandType(plays[i].hand, plays[i].cardsCount)
		fmt.Printf("Hand type %v(%v): %v\n", i, plays[i].hand, plays[i].handType)
	}

	ranksMap := make(map[int]int)
	maxRank := playsSize

	for i := 0; i < playsSize; i++ {
		rank := calcRank(plays, i, maxRank)

		fmt.Printf("[%v]Rank %v | %v\n", i, rank, plays[i].hand)

		ranksMap[rank]++

		score += plays[i].bid * rank
	}

	fmt.Printf("(%v/%v)Ranks map: %v", len(ranksMap), playsSize, ranksMap)

	return score
}

func calcRank(plays []play, playIndex int, maxRank int) int {
	playsSize := len(plays)
	rank := maxRank // Max rank

	currentPlay := plays[playIndex]

	for i := 0; i < playsSize; i++ {
		if i == playIndex {
			continue
		}

		if currentPlay.handType == plays[i].handType {
			for c, currentHCard := range currentPlay.hand {
				if currentHCard == plays[i].hand[c] {
					continue
				}

				if currentHCard < plays[i].hand[c] {
					rank--
				}
				break
			}
		} else if currentPlay.handType < plays[i].handType {
			rank--
		}
	}

	return rank
}

func getHandType(hand [5]int, cards map[int]int) int {
	combo := [5]int{0, 0, 0, 0, 0}

	for card, repeated := range cards {
		fmt.Printf("Card %v repeated %v times\n", card, repeated)
		combo[repeated-1]++
	}

	fmt.Printf("Combo: %v\n", combo)

	for i := 4; i >= 3; i-- {
		if combo[i] == 1 {
			return i+1
		}
	}

	if combo[2] == 1 && combo[1] == 1 {
		return 3
	}

	if combo[2] == 1 {
		return 2
	}

	if combo[1] == 2 {
		return 1
	}

	if combo[1] == 1 {
		return 0
	}

	return -1
}
