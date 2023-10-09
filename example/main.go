package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/heroiclabs/nakama-common/runtime"
)

var (
	errInternalError  = runtime.NewError("internal server error", 13) // INTERNAL
	errMarshal        = runtime.NewError("cannot marshal type", 13)   // INTERNAL
	errNoInputAllowed = runtime.NewError("no input allowed", 3)       // INVALID_ARGUMENT
	errNoUserIdFound  = runtime.NewError("no user ID in context", 3)  // INVALID_ARGUMENT
	errUnmarshal      = runtime.NewError("cannot unmarshal type", 13) // INTERNAL
)

const (
	rpcIdRewards   = "rewards"
	rpcIdFindMatch = "find_match"
	rpcMyFunc      = "MyFunc"
	rpcMyFunc2     = "MyFunc2"
)

func rpcMyFunch(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	fmt.Println("xxxxxxxxxxxxxxx   sssssssssssssssssssss")
	// fmt.Println("xxxxxxxxxxxxxxxxxxxxx")
	userID := "32b9a76c-e10b-4765-8bbb-48f32c8fd569"
	changeset := map[string]int64{
		"coins": 10, // Add 10 coins to the user's wallet.
		"gems":  -5, // Remove 5 gems from the user's wallet.
	}
	metadata := map[string]interface{}{
		"game_result": "won",
	}
	updated, previous, err := nk.WalletUpdate(ctx, userID, changeset, metadata, true)

	fmt.Println(updated)
	fmt.Println(previous)
	fmt.Println(err)

	return "", nil
}

func rpcMyFunch2(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	fmt.Println("xxxxxxxxxxxxxxx   sssssssssssssssssssss")
	// fmt.Println("xxxxxxxxxxxxxxxxxxxxx")
	userID := "93b917a8-f7e0-41c3-aaee-332360c29b9d"
	changeset := map[string]int64{
		"coins": 10, // Add 10 coins to the user's wallet.
		"gems":  5,  // Remove 5 gems from the user's wallet.
	}
	metadata := map[string]interface{}{
		"game_result": "won",
	}
	updated, previous, err := nk.WalletUpdate(ctx, userID, changeset, metadata, true)

	fmt.Println(updated)
	fmt.Println(previous)
	fmt.Println(err)

	return "", nil
}

func MyAccessSessionVars(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule) error {
	vars, ok := ctx.Value(runtime.RUNTIME_CTX_VARS).(map[string]string)

	if !ok {
		logger.Info("User session does not contain any key-value pairs set")
		return nil
	}

	logger.Info("User session contains key-value pairs set by both the client and the before authentication hook: %v", vars)
	return nil
}

// noinspection GoUnusedExportedFunction
func InitModule(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, initializer runtime.Initializer) error {
	fmt.Println("0------------------------------")
	initStart := time.Now()

	fmt.Println("2------------------------------")
	// if err := initializer.RegisterRpc(rpcMyFunc, rpcMyFunch); err != nil {
	// 	return err
	// }

	if err := initializer.RegisterRpc(rpcMyFunc2, rpcMyFunch2); err != nil {
		return err
	}

	logger.Info("Plugin loaded in '%d' msec.", time.Now().Sub(initStart).Milliseconds())
	return nil
}
