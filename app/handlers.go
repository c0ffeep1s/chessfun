package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"chessfun/engine/board"
	"chessfun/engine/search"
)

type Data struct {
	Fen string `json:"fen"`
}

// get the fen and put it on the board broh
func makemove(rw http.ResponseWriter, r *http.Request) {
	// ok for now
	rw.Header().Set("Access-Control-Allow-Origin", "*")
	rw.Header().Add("Content-Type", "application/json")

	// extract FEN str from POST
	var req Data

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(rw, err.Error(), http.StatusBadRequest)
		return
	}

	// init Board instance
	var pos board.PositionStruct
	var searchInfo search.InfoStruct
	searchInfo.Depth = board.MaxDepth
	var moveTime = 1000
	forceMode := false
	moveMade := true

	board.Initialize()
	// Init hash tables size with 2 MB's
	pos.HashTable.Init(16)

	err2 := pos.LoadFEN(req.Fen)
	if err2 != nil {
		fmt.Println(err2.Error())
	}

	// search for best move
	var fenStr string = " "
	if moveMade && !forceMode {
		searchInfo.StartTime = time.Now().UnixNano() / int64(time.Millisecond)
		searchInfo.StopTime = searchInfo.StartTime + int64(moveTime)
		searchInfo.TimeSet = true
		err := searchInfo.SearchPosition(&pos)
		if err != nil {
			fmt.Println(err.Error())
		}
		// update internal Board state
		_, err = pos.MakeMove(pos.PvArray[0])
		if err != nil {
			fmt.Println(err.Error())
		}

		// increment full move counter
		if pos.Side == 0 {
			pos.FullMove += 1
		}
		// extract FEN from current Board state
		fenStr = pos.ExtractFEN()
	}

	// extract FEN
	var currfen = Data{Fen: fenStr}

	rw.Header().Add("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(currfen)
}
