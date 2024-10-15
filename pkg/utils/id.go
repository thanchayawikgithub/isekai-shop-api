package utils

import (
	"github.com/labstack/echo/v4"
	adminExceptions "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/admin/exceptions"
	playerExceptions "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/player/exceptions"
)

func GetReqAdminD(ctx echo.Context) (string, error) {
	if adminID, ok := ctx.Get("adminID").(string); !ok || adminID == "" {
		return "", &adminExceptions.AdminNotFound{AdminID: "Unknown"}
	} else {
		return adminID, nil
	}
}

func GetReqPlayerID(ctx echo.Context) (string, error) {
	if playerID, ok := ctx.Get("playerID").(string); !ok || playerID == "" {
		return "", &playerExceptions.PlayerNotFound{PlayerID: "Unknown"}
	} else {
		return playerID, nil
	}
}
