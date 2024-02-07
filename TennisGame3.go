package tenniskata

import (
	"fmt"
	"math"
)

type tennisGame3 struct {
	playerOne player
	playerTwo player
}

type player struct {
	name  string
	score int
}

const (
	duecePhaseCondition   = 6
	initialPhaseCondition = 4
	dueceTemplate         = "Deuce"
	advantageTemplate     = "Advantage %s"
	winTemplate           = "Win for %s"
	scoreTemplate         = "%s-%s"
	tieTemplate           = "%s-All"
)

func TennisGame3(playerOneName string, playerTwoName string) TennisGame {
	return &tennisGame3{
		playerOne: player{name: playerOneName},
		playerTwo: player{name: playerTwoName},
	}
}

func (game *tennisGame3) GetScore() string {
	if game.playerOne.score+game.playerTwo.score >= duecePhaseCondition {
		return game.renderScoreDuecePhase()
	}

	if game.playerOne.score < initialPhaseCondition && game.playerTwo.score < initialPhaseCondition {
		return game.renderScoreInitialPhase()
	}

	return fmt.Sprintf(winTemplate, game.getWinner())

}

func (game *tennisGame3) renderScoreInitialPhase() string {
	var scoreMap = map[int]string{
		0: "Love",
		1: "Fifteen",
		2: "Thirty",
		3: "Forty",
	}

	if game.playerOne.score == game.playerTwo.score {
		return fmt.Sprintf(tieTemplate, scoreMap[game.playerOne.score])
	}
	return fmt.Sprintf(scoreTemplate, scoreMap[game.playerOne.score], scoreMap[game.playerTwo.score])

}

func (game *tennisGame3) renderScoreDuecePhase() string {

	if game.playerOne.score == game.playerTwo.score {
		return dueceTemplate
	}

	if math.Abs(float64(game.playerOne.score-game.playerTwo.score)) == 1 {
		return fmt.Sprintf(advantageTemplate, game.getWinner())
	}
	return fmt.Sprintf(winTemplate, game.getWinner())
}

func (game *tennisGame3) getWinner() string {
	if game.playerOne.score > game.playerTwo.score {
		return game.playerOne.name
	}
	return game.playerTwo.name
}
func (game *tennisGame3) WonPoint(playerName string) {
	if playerName == game.playerOne.name {
		game.playerOne.score += 1
	} else {
		game.playerTwo.score += 1
	}
}
