package tournament

import (
	"testing"
)

// Adds player, then tries to find
func TestDataAddFindPlayer(t *testing.T){
	playersID := []string{"P1", "P2", "P3", "P4", "P5"}
	for _, playerId := range playersID {
		// add player
		DataAddPlayer(PlayerInfo{playerId, 0})
		// find just added player
		_, err := DataFindPlayer(playerId)
		if err != nil {
			t.Error("Add player failed. No player found")
		}
	}
}

// Increase, decrease points
func TestDataIncDecPoints(t *testing.T){
	playersID := []string{"P1", "P2", "P3", "P4", "P5"}
	for _, playerId := range playersID {
		player, err := DataFindPlayer(playerId)
		t_balance := player.balance
		if err != nil {
			t.Error("Could not find player by id")
		}
		DataIncreasePoint(player.playerId, 100)
		if t_balance == player.balance{
			t.Error("Balance were not increased")
		}
		DataDecreasePoint(player.playerId, 100)
		if t_balance != player.balance{
			t.Error("Balance were not decreased")
		}
	}
}

// Init and Get tournament info
func TestDataInitGetTournament(t *testing.T){
	tournamentsID := []int{1, 2, 3, 4, 5}
	for _, tournamentId := range tournamentsID {
		DataInitTournament(tournamentId, 1000)
		_, err := DataGetTournametInfo(tournamentId)
		if err != nil{
			t.Error("Init tournament fialed. Tournament not found")
		}
	}
}

// Add player into tournament
func TestDataAddPlayer(t *testing.T){
	tournamentsID := []int{1, 2, 3, 4, 5}
	playersID := []string{"P1", "P2", "P3", "P4", "P5"}
	var okCount int
	for _, tournamentId := range tournamentsID{
		okCount = 0
		for _, playerId := range playersID{
			DataAddPlayerToTournament(playerId, tournamentId)
		}
		tournament, err := DataGetTournametInfo(tournamentId)
		if err != nil{
			t.Error("Some error while test add player. Tournament not found")
		}
		for index, playerId := range playersID{
			if playerId == tournament.joinedPlayers[index].playerId{
				okCount += 1
			}
		}
		if okCount != 5{
			t.Error("Some player not added to tournament")
		}
	}
}

// Add player with backers into tournament
func TestDataAddPlayerBackers(t *testing.T){
	tournamentsID := []int{1, 2, 3, 4, 5}
	playersID := []string{"P1", "P2", "P3", "P4", "P5"}
	var okCount int
	for _, tournamentId := range tournamentsID{
		okCount = 0
		for _, playerId := range playersID{
			DataAddPlayerToTournamentWithBacker(playerId, tournamentId, playersID)
		}
		tournament, err := DataGetTournametInfo(tournamentId)
		if err != nil{
			t.Error("Some error while test add player with backer. Tournament not found")
		}
		for index, playerId := range tournament.joinedPlayers{
			for _, backerId := range tournament.joinedPlayers[index].backerId {
				if playerId.playerId == backerId {
						okCount += 1
				}
			}
		}
		if okCount != 5{
			t.Error("Some player not added to tournament", okCount)
		}
	}
}

// Get winners
func TestGetWinner(t *testing.T){
	tournamentsID := []int{1, 2, 3, 4, 5}
	for _, tournament := range tournamentsID{
		winner, prize := DataGetWinners(tournament)
		if winner.playerId != "P1" || prize == 0{
			t.Error("Winner are not P1")
		}
	}
}