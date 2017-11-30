package tournament

type PlayerInfo struct {
    playerId string `json:"playerid"`
    balance  int `json:"balance"` 
}
type JoinedPlayer struct {
    playerId string
    backerId []string
}
type Tournament struct {
    tournamentId int
    deposit int
    prize int
    winner string
    joinedPlayers   []JoinedPlayer
}

type Winner struct {
    PlayerId string
    Prize int
}