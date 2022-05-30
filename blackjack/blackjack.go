package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	var value int

	switch {
	case card == "ace":
		value = 11
	case card == "two":
		value = 2
	case card == "three":
		value = 3
	case card == "four":
		value = 4
	case card == "five":
		value = 5
	case card == "six":
		value = 6
	case card == "seven":
		value = 7
	case card == "eight":
		value = 8
	case card == "nine":
		value = 9
	case card == "ten" || card == "jack" || card == "queen" || card == "king":
		value = 10
	default:
		value = 0
	}
	return value
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
	var decision string = "NONE"
	var totalValue = ParseCard(card1) + ParseCard(card2)
	var dealerValue = ParseCard(dealerCard)
	if card2 == "ace" && card1 == "ace" {
		decision = "P"
	} else if totalValue == 21 {
		if dealerValue < 10 {
			decision = "W"
		} else {
			decision = "S"
		}
	} else if 17 <= totalValue && totalValue <= 20 {
		decision = "S"
	} else if 12 <= totalValue && totalValue <= 16 {
		if dealerValue >= 7 {
			decision = "H"
		} else {
			decision = "S"
		}
	} else if totalValue <= 11 {
		decision = "H"
	}
	return decision

}
