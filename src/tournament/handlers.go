package tournament

import (
	"log"
	"net/http"
	"fmt"
	"strconv"
	"encoding/json"
	"errors"
)

var percent float64

func fatal(err error){
	if err != nil{
		log.Fatal(err)
	}
}

// Checks is player can join tournament with his own points or
// Do backers have enough poiints to fund player
func checkBalancePlayers(deposit int, balance int, backers []string) bool{
	if balance >= deposit{
		return true
	}
	percent = float64((deposit/(len(backers)+1)))/float64(deposit)
	log.Print(percent)
	for _,back := range backers{
		backer, err := DataFindPlayer(back)
		fatal(err)
		if float64(backer.balance) < float64(deposit)*percent{
			log.Printf("You have not enough money")
			return false
		}
	}
	return true
}

// Return balance of a player
func balance(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	playersId := r.FormValue("playerId")
	player, err := DataFindPlayer(playersId)
	fatal(err)
	json.NewEncoder(w).Encode(players)
	fmt.Fprintf(w, "Player %v has balance: %v", player.playerId, player.balance)
}

// Increase player's points by value
func fundPlayer(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	playerId := r.FormValue("playerId")
	points, err := strconv.Atoi(r.FormValue("points"))
	fatal(err)
	if points < 0{
		fatal(errors.New("Points must be greater than 0"))
	}
	fmt.Fprintf(w, "Fund player %v with %v points\n", playerId, points)
	DataIncreasePoint(playerId, points)
	// Show balance to be sure points are changed
	balance(w, r)
}

// Decrease player's points by value.
func takePlayer(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	playerId := r.FormValue("playerId")
	points, err := strconv.Atoi(r.FormValue("points"))
	fatal(err)
	if points < 0{
		fatal(errors.New("Points must be greater than 0"))
	}
	fmt.Fprintf(w, "Take player %v with %v points\n", playerId, points)
	DataDecreasePoint(playerId, points)
	// Show balance to be sure points are changed
	balance(w, r)
}

func joinTournament(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	playerId := r.FormValue("playerId")
	backers := r.Form["backerId"]
	tournamentId, err := strconv.Atoi(r.FormValue("tournamentId"))
	fatal(err)
	t_balance, err := DataGetTournametInfo(tournamentId)
	fatal(err)
	t_player, err := DataFindPlayer(playerId)
	fatal(err)
	// Check if player and backers have enough points
	if checkBalancePlayers(t_balance.deposit, t_player.balance, backers) {
		log.Printf("Possible")
		fmt.Fprintf(w, "Player %q joins ", playerId)
		if r.FormValue("backerId") != "" {
			decrease_size := float64(t_balance.deposit)*percent
			fmt.Fprintf(w, "with backers:")
			for _, backerId := range r.Form["backerId"] {
				fmt.Fprintf(w, "%q, ", backerId)
				log.Printf("%v",decrease_size)
				DataDecreasePoint(backerId, int(decrease_size))
			}
			DataDecreasePoint(playerId, int(decrease_size))
			DataAddPlayerToTournamentWithBacker(playerId, tournamentId, backers)
		} else {
			fmt.Fprintf(w, "on his own")
			DataDecreasePoint(playerId, t_balance.deposit)
			DataAddPlayerToTournament(playerId, tournamentId)
		}
	}
}

// Initialize new tournament
func announceTournament(w http.ResponseWriter, r *http.Request){
	r.ParseForm()
	tournamentId, err := strconv.Atoi(r.FormValue("tournamentId"))
	fatal(err)
	t_deposit, err := strconv.Atoi(r.FormValue("deposit"))
	fatal(err)
	if t_deposit < 0 || tournamentId < 0{
		fatal(errors.New("Deposit and tournament id must be greater than 0"))
	}
	DataInitTournament(t_deposit, tournamentId)
}

// Get result of a tournament
func resultTournament(w http.ResponseWriter, r *http.Request){
	// Winner is always P1
	for _, tournament := range tournaments {
		winners, prize := DataGetWinners(tournament.tournamentId)
		if prize == 0 {
			fatal(errors.New("no prize"))
		}
		if len(winners.backerId) > 0 {
			prize = int(float64(prize) * percent)
			for _, backer := range winners.backerId {
				DataIncreasePoint(backer, prize)
			}
		}
		DataIncreasePoint(winners.playerId, prize)
		// Json preparation
		profile := Winner{winners.playerId, prize}
		js, err := json.Marshal(profile)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}
}

func Index(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "Hello!")
}

