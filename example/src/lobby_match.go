// package src

// import (
// 	"context"
// 	"database/sql"

// 	"github.com/heroiclabs/nakama-common/runtime"
// )

// type LobbyMatch struct{}

// type LobbyMatchState struct {
// 	presences  map[string]runtime.Presence
// 	emptyTicks int
// }

// func (m *LobbyMatch) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
// 	state := &LobbyMatchState{
// 		emptyTicks: 0,
// 		presences:  map[string]runtime.Presence{},
// 	}
// 	tickRate := 1 // 1 tick per second = 1 MatchLoop func invocations per second
// 	label := ""
// 	return state, tickRate, label
// }

// func (m *LobbyMatch) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
// 	lobbyState, ok := state.(*LobbyMatchState)
// 	if !ok {
// 		logger.Error("state not a valid lobby state object")
// 		return nil
// 	}

// 	for i := 0; i < len(presences); i++ {
// 		lobbyState.presences[presences[i].GetSessionId()] = presences[i]
// 	}

// 	return lobbyState
// }

// func (m *LobbyMatch) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
// 	lobbyState, ok := state.(*LobbyMatchState)
// 	if !ok {
// 		logger.Error("state not a valid lobby state object")
// 		return nil
// 	}

// 	for i := 0; i < len(presences); i++ {
// 		delete(lobbyState.presences, presences[i].GetSessionId())
// 	}

// 	return lobbyState
// }

// func (m *LobbyMatch) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
// 	lobbyState, ok := state.(*LobbyMatchState)
// 	if !ok {
// 		logger.Error("state not a valid lobby state object")
// 		return nil
// 	}

// 	// If we have no presences in the match according to the match state, increment the empty ticks count
// 	if len(lobbyState.presences) == 0 {
// 		lobbyState.emptyTicks++
// 	}

// 	// If the match has been empty for more than 100 ticks, end the match by returning nil
// 	if lobbyState.emptyTicks > 100 {
// 		return nil
// 	}

// 	return lobbyState
// }

// // 实现 MatchJoinAttempt 方法
// func (m *LobbyMatch) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
// 	lobbyState, ok := state.(*LobbyMatchState)
// 	if !ok {
// 		logger.Error("state not a valid lobby state object")
// 		return lobbyState, ok, ""
// 	}

// 	return lobbyState, ok, ""
// }

// func (m *LobbyMatch) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
// 	lobbyState, ok := state.(*LobbyMatchState)
// 	if !ok {
// 		logger.Error("state not a valid lobby state object")
// 		return lobbyState, ""
// 	}

// 	return lobbyState, ""
// }

// func (m *LobbyMatch) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
// 	lobbyState, ok := state.(*LobbyMatchState)
// 	if !ok {
// 		logger.Error("state not a valid lobby state object")
// 		return nil
// 	}

// 	return lobbyState
// }

// // func SendMsg() {
// // 	const MATCH_START_OPCODE = 7

// // 	matchStartData := &map[string]interface{}{
// // 		"started":    true,
// // 		"roundTimer": 100,
// // 	}

// // 	data, err := json.Marshal(matchStartData)
// // 	if err != nil {
// // 		logger.Error("error marshaling match start data", err)
// // 		return nil
// // 	}

// // 	reliable := true
// // 	dispatcher.BroadcastMessage(MATCH_START_OPCODE, data, nil, nil, reliable)
// // }

package src

import (
	"context"
	"database/sql"
	"strconv"

	"github.com/heroiclabs/nakama-common/runtime"
)

type MatchState struct {
	presences map[string]runtime.Presence
}

type Match struct{}

func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	logger.Info("Hello Multiplayer!")
	err := initializer.RegisterMatch("standard_match", newMatch)

	if err != nil {
		logger.Error("[RegisterMatch] error: ", err.Error())
		return err
	}

	return nil
}

func newMatch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (m runtime.Match, err error) {
	return &Match{}, nil
}

func (m *Match) MatchInit(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, params map[string]interface{}) (interface{}, int, string) {
	state := &MatchState{
		presences: make(map[string]runtime.Presence),
	}

	tickRate := 1
	label := ""

	return state, tickRate, label
}

func (m *Match) MatchJoinAttempt(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presence runtime.Presence, metadata map[string]string) (interface{}, bool, string) {
	acceptUser := true

	return state, acceptUser, ""
}

func (m *Match) MatchJoin(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	mState, _ := state.(*MatchState)

	for _, p := range presences {
		mState.presences[p.GetUserId()] = p
	}

	return mState
}

func (m *Match) MatchLeave(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, presences []runtime.Presence) interface{} {
	mState, _ := state.(*MatchState)

	for _, p := range presences {
		delete(mState.presences, p.GetUserId())
	}

	return mState
}

func (m *Match) MatchLoop(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, messages []runtime.MatchData) interface{} {
	mState, _ := state.(*MatchState)

	for _, presence := range mState.presences {
		logger.Info("Presence %v named %v", presence.GetUserId(), presence.GetUsername())
	}

	for _, message := range messages {
		logger.Info("Received %v from %v", string(message.GetData()), message.GetUserId())
		reliable := true
		dispatcher.BroadcastMessage(1, message.GetData(), []runtime.Presence{message}, nil, reliable)
	}

	return mState
}

func (m *Match) MatchTerminate(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, graceSeconds int) interface{} {
	message := "Server shutting down in " + strconv.Itoa(graceSeconds) + " seconds."
	reliable := true
	dispatcher.BroadcastMessage(2, []byte(message), []runtime.Presence{}, nil, reliable)

	return state
}

func (m *Match) MatchSignal(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, dispatcher runtime.MatchDispatcher, tick int64, state interface{}, data string) (interface{}, string) {
	return state, "signal received: " + data
}
