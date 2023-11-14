package gelo

import "math"

type Elo struct {
	// K-Factor
	K int
	C int
}

// default values for K and C
const (
	defaultK int = 32
	defaultC int = 400
)

func NewElo() *Elo {
	return &Elo{defaultK, defaultC}
}

// Specify values for k and c
func NewEloCustom(k int, c int) *Elo {
	return &Elo{k, c}
}

// returns the expected score of a player with ratingA
func (elo *Elo) ExpectedScore(ratingA int, ratingB int) float64 {
	return 1 / (1 + math.Pow(10, float64(ratingB-ratingA)/float64(elo.C)))
}

// returns the new rating of a player with ratingA
func (elo *Elo) NewRating(ratingA int, ratingB int, actualScoreA float64) int {
	expectedScoreA := elo.ExpectedScore(ratingA, ratingB)
	return ratingA + int(float64(elo.K)*(actualScoreA-expectedScoreA))
}

// returns the new rating of the winner and the new rating of the loser
func (elo *Elo) CalculateNewRatings(winnerRating int, loserRating int) (int, int) {
	newRatingWinner := elo.NewRating(winnerRating, loserRating, 1)
	newRatingLoser := elo.NewRating(loserRating, winnerRating, 0)

	return newRatingWinner, newRatingLoser
}
