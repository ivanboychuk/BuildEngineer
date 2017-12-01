package tournament

import (
	"log"
	"errors"
	"strconv"
	"os"
	"time"
	"bufio"
	"fmt"
)

var players []PlayerInfo
var tournaments []Tournament
var tournamentsId int
var transactionfile string

// Just for testing.
func DataInitialization() {
	transactionfile = "today.log"
	if _, err := os.Stat(transactionfile); !os.IsNotExist(err) {
		logTransaction("0", "A new run: ", time.Now().String())
	} else {
		os.Create(transactionfile)
	}
	DatacheckLatestTransaction()
	logMsg(transactionfile)
	tournamentsId = 0
	for i := 1; i < 6; i++ {
		DataAddPlayer(PlayerInfo{"P"+strconv.Itoa(i), 0})
	}
	DataInitTournament(1000, 2000)
	DataIncreasePoint("P1", 300)
	DataIncreasePoint("P2", 300)
	DataIncreasePoint("P3", 300)
	DataIncreasePoint("P4", 500)
	DataIncreasePoint("P5", 1000)
}

// Checks latest transaction and re-run if it is unsuccessful
func DatacheckLatestTransaction(){
	file, err := os.Open(transactionfile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

// Add a new player in scope
func DataAddPlayer(player PlayerInfo){
	players = append(players, player)
}

// Find and return point to the player by id
func DataFindPlayer(id string) (*PlayerInfo, error) {
	for index, player := range players{
		if player.playerId == id{
			return &players[index], nil
		}
	}
	return &PlayerInfo{}, errors.New("No such player " + id)
}

// Increase player's point by value
func DataIncreasePoint(id string, points int) {
	logTransaction("1", id, "+"+strconv.Itoa(points))
	player, err := DataFindPlayer(id)
	if err != nil {
		log.Fatal(err)
	}
	player.balance += points
	logTransaction("0", id, "+"+strconv.Itoa(points))
}

// Decrease player's point by value
func DataDecreasePoint(id string, points int) {
	logTransaction("1", id, "-"+strconv.Itoa(points))
	player, err := DataFindPlayer(id)
	if err != nil {
		log.Fatal(err)
	}
	player.balance -= points
	logTransaction("0", id, "-"+strconv.Itoa(points))
}

// Add new tournament in the scope
// Tournaments never end
func DataInitTournament(deposit int, prize int) []Tournament{
	tournamentsId += 1
	// Yep, winner is always P1. Lucky guy
	tournaments = append(tournaments, Tournament{tournamentId: tournamentsId,deposit:deposit, prize:prize, winner:"P1"})

	return tournaments
}

// Return info about tournament
func DataGetTournametInfo(id int) (Tournament, error){
	for _, tournament := range tournaments{
		if tournament.tournamentId == id{
			return tournament, nil
		}
	}
	return Tournament{}, errors.New("No tournament found")
}

// Join a player into tournament
func DataAddPlayerToTournament(playerId string, tournamentId int){
	_, err := DataFindPlayer(playerId)
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
func DataAddPlayerToTournamentWithBacker(playerId string, tournamentId int, backers []string){
	_, err := DataFindPlayer(playerId)
	if err != nil {
		log.Fatal(err)
	}
	for index, tournament := range tournaments{
		if tournament.tournamentId == tournamentId{
			player := JoinedPlayer{playerId:playerId}
			for _, backer := range backers{
				_, err := DataFindPlayer(backer)
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
func DataGetWinners(id int) (JoinedPlayer, int){
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