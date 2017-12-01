package tournament

import (
	"os"
)

func logMsg(message string){
	print("Messages!\n", message)
}

// Logs the transactions like: 1:P1+100
// 1 - started
// 0 - finished
func logTransaction(status string, playerId string, operation string){
	message := status+":"+playerId+operation+"\n"

	f, err := os.OpenFile(transactionfile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(message); err != nil {
		panic(err)
	}
}