package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/InfiniteCodeMonkeys/simple-server/types"
)

func EchoHandler(w http.ResponseWriter, r *http.Request) {
	var msg types.Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	msg.Time = time.Now().Format(time.RFC3339)
	json.NewEncoder(w).Encode(msg)
}
