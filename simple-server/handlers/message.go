package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/InfiniteCodeMonkeys/simple-server/types"
)

func MessageHandler(w http.ResponseWriter, r *http.Request) {
	msg := types.Message{Text: "Welcome to the Go server!",
		Time: time.Now().Format(time.RFC3339)}
	json.NewEncoder(w).Encode(msg)
}
