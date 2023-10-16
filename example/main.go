package main

// import (
// 	"context"
// 	"database/sql"
// 	"fmt"
// 	"time"

// 	"example.com/go-project/src"

// 	"github.com/heroiclabs/nakama-common/runtime"
// )

// var (
// 	errInternalError  = runtime.NewError("internal server error", 13) // INTERNAL
// 	errMarshal        = runtime.NewError("cannot marshal type", 13)   // INTERNAL
// 	errNoInputAllowed = runtime.NewError("no input allowed", 3)       // INVALID_ARGUMENT
// 	errNoUserIdFound  = runtime.NewError("no user ID in context", 3)  // INVALID_ARGUMENT
// 	errUnmarshal      = runtime.NewError("cannot unmarshal type", 13) // INTERNAL
// )

// const (
// 	rpcIdRewards   = "rewards"
// 	rpcIdFindMatch = "find_match"
// 	rpcMyFunc      = "MyFunc"
// 	rpcMyFunc2     = "MyFunc2"
// )

// func rpcMyFunch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
// 	fmt.Println("xxxxxxxxxxxxxxx   sssssssssssssssssssss")
// 	// fmt.Println("xxxxxxxxxxxxxxxxxxxxx")
// 	userID := "32b9a76c-e10b-4765-8bbb-48f32c8fd569"
// 	changeset := map[string]int64{
// 		"coins": 10, // Add 10 coins to the user's wallet.
// 		"gems":  -5, // Remove 5 gems from the user's wallet.
// 	}
// 	metadata := map[string]interface{}{
// 		"game_result": "won",
// 	}
// 	updated, previous, err := nk.WalletUpdate(ctx, userID, changeset, metadata, true)

// 	fmt.Println(updated)
// 	fmt.Println(previous)
// 	fmt.Println(err)

// 	return "", nil
// }

// func rpcMyFunch2(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
// 	fmt.Println("xxxxxxxxxxxxxxx   sssssssssssssssssssss")
// 	// fmt.Println("xxxxxxxxxxxxxxxxxxxxx")
// 	userID := "93b917a8-f7e0-41c3-aaee-332360c29b9d"
// 	changeset := map[string]int64{
// 		"coins": 10, // Add 10 coins to the user's wallet.
// 		"gems":  5,  // Remove 5 gems from the user's wallet.
// 	}
// 	metadata := map[string]interface{}{
// 		"game_result": "won",
// 	}
// 	updated, previous, err := nk.WalletUpdate(ctx, userID, changeset, metadata, true)

// 	fmt.Println(updated)
// 	fmt.Println(previous)
// 	fmt.Println(err)

// 	return "", nil
// }

// func MyAccessSessionVars(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) error {
// 	vars, ok := ctx.Value(runtime.RUNTIME_CTX_VARS).(map[string]string)

// 	if !ok {
// 		logger.Info("User session does not contain any key-value pairs set")
// 		return nil
// 	}

// 	logger.Info("User session contains key-value pairs set by both the client and the before authentication hook: %v", vars)
// 	return nil
// }

// // noinspection GoUnusedExportedFunction
// func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
// 	fmt.Println("0------------------------------")
// 	initStart := time.Now()

// 	fmt.Println("2------------------------------")
// 	// if err := initializer.RegisterRpc(rpcMyFunc, rpcMyFunch); err != nil {
// 	// 	return err
// 	// }

// 	var systemId string
// 	if env, ok := ctx.Value(runtime.RUNTIME_CTX_ENV).(map[string]string); ok {
// 		systemId = env["SYSTEM_ID"]
// 	}

// 	_, err := db.ExecContext(ctx, `
// 	INSERT INTO users (id, username)
// 	VALUES ($1, $2)
// 	ON CONFLICT (id) DO NOTHING
// 	  `, systemId, "sysmtem_id")
// 	if err != nil {
// 		logger.Error("Error: %s", err.Error())
// 	}

// 	// register
// 	if err := initializer.RegisterRpc(rpcMyFunc2, rpcMyFunch2); err != nil {
// 		return err
// 	}

// 	if err := initializer.RegisterMatch("lobby", func(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) (runtime.Match, error) {
// 		return &src.Match{}, nil
// 	}); err != nil {
// 		logger.Error("unable to register: %v", err)
// 		return err
// 	}

// 	logger.Info("Plugin loaded in '%d' msec.", time.Now().Sub(initStart).Milliseconds())
// 	return nil
// }

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime/debug"

	"github.com/heroiclabs/nakama-common/runtime"
)

const apiBaseUrl = "https://pokeapi.co/api/v2"

// All Go modules must have a InitModule function with this exact signature.
func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	debug.PrintStack() // 打印调用堆栈信息
	fmt.Println("ii 1----------|", "")
	// Register the RPC function.
	if err := initializer.RegisterRpc("get_pokemon", GetPokemon); err != nil {
		logger.Error("Unable to register: %v", err)
		return err
	}

	return nil
}

func LookupPokemon(logger runtime.Logger, name string) (map[string]interface{}, error) {
	resp, err := http.Get(apiBaseUrl + "/pokemon/" + name)

	if err != nil {
		logger.Error("Failed request %v", err.Error())
		return nil, runtime.NewError("unable to retrieve api data", 13)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		logger.Error("Failed to read body %v", err.Error())
		return nil, runtime.NewError("failed to read body", 13)
	}

	if resp.StatusCode >= 400 {
		logger.Error("Failed request %v %v", resp.StatusCode, body)
		return nil, runtime.NewError("failed api request", 13)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, runtime.NewError("unable to unmarshal data", 13)
	}

	return result, nil
}

func GetPokemon(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	fmt.Println("qwewqe 1----------|", "")
	// We'll assume payload was sent as JSON and decode it.
	var input map[string]string
	if err := json.Unmarshal([]byte(payload), &input); err != nil {
		return "", runtime.NewError("unable to unmarshal payload", 13)
	}

	result, err := LookupPokemon(logger, input["PokemonName"])

	if err != nil {
		return "", runtime.NewError("unable to find pokemon", 5)
	}

	_, err2 := json.Marshal(result)

	if err2 != nil {
		return "", runtime.NewError("unable to marshal response", 13)
	}
	fmt.Println("q 1----------|", 1)
	return string("qqqqqqqqqqqqqqqqqqqqqqqq"), nil
}
