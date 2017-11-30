package tournament

import (
	"log"
	"errors"
	"strconv"
)

var players []PlayerInfo
var tournaments []Tournament
var TournamentId int

// Just for testing.
func datainit() {
	TournamentId = 0
	for i := 1; i < 6; i++ {
		dataAddPlayer(PlayerInfo{"P"+strconv.Itoa(i), 0})
	}
	dataInitTournament(1000, 2000)
	dataIncreasePoint("P1", 300)
	dataIncreasePoint("P2", 300)
	dataIncreasePoint("P3", 300)
	dataIncreasePoint("P4", 500)
	dataIncreasePoint("P5", 1000)
}

// Add a new player in scope
func dataAddPlayer(player PlayerInfo){
	players = append(players, player)
}

// Find and return point to the player by id
func dataFindPlayer(id string) (*PlayerInfo, error) {
	for index, player := range players{
		if player.playerId == id{
			return &players[index], nil
		}
	}
	return &PlayerInfo{}, errors.New("No such player " + id)
}

// Increase player's point by value
func dataIncreasePoint (id string, points int) {
	player, err := dataFindPlayer(id)
	if err != nil {
		log.Fatal(err)
	}
	player.balance += points
}

// Decrease player's point by value
func dataDecreasePoint (id string, points int) {
	player, err := dataFindPlayer(id)
	if err != nil {
		log.Fatal(err)
	}
	player.balance -= points
}

// Add new tournamnet in the scope
// Tournaments never end
func dataInitTournament (deposit int, prize int) []Tournament{
	TournamentId += 1
	// Yep, winner is always P1. Lucky guy
	tournaments = append(tournaments, Tournament{tournamentId:TournamentId,deposit:deposit, prize:prize, winner:"P1"})

	return tournaments
}

// Return info about tournament
func dataGetTournametInfo(id int) (Tournament, error){
	for _, tournament := range tournaments{
		if tournament.tournamentId == id{
			return tournament, nil
		}
	}
	return Tournament{}, errors.New("No tournament found")
}

// Join a player into tournament
func dataAddPlayerToTournament (playerId string, tournamentId int){
	_, err := dataFindPlayer(playerId)
	if err != nil {
		log.Fatal(err)
	}
	for index, tournament := range tournaments{
		if tournament.tournamentId == tournamentId{
			tournaments[index].joinedPlayers = append(tournaments[index].joinedPlayers, JoinedPlayer{playerId:playerId})
		}
	}
}

// Join a player with its backers
func dataAddPlayerToTournamentWithBacker (playerId string, tournamentId int, backers []string){
	_, err := dataFindPlayer(playerId)
	if err != nil {
		log.Fatal(err)
	}
	for index, tournament := range tournaments{
		if tournament.tournamentId == tournamentId{
			player := JoinedPlayer{playerId:playerId}
			for _, backer := range backers{
				_, err := dataFindPlayer(backer)
				if err != nil {
					log.Fatal(err)
				}
				player.backerId = append(player.backerId, backer)
			}
			tournaments[index].joinedPlayers = append(tournaments[index].joinedPlayers, player)
		}
	}
}

// Return winner for tournament by id of this one
func dataGetWinners(id int) (JoinedPlayer, int){
	for _, tournament := range tournaments {
		if tournament.tournamentId == id {
			for _, player := range tournament.joinedPlayers {
				if player.playerId == tournament.winner{
					return player, tournament.prize
				}
			}
		}
	}
	return JoinedPlayer{}, 0
}