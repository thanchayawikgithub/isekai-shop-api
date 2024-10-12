package services

import (
	adminModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/admin/models"
	playerModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/player/models"
)

type Oauth2Service interface {
	PlayerAccountCreating(playerCreatingReq playerModels.PlayerCreatingReq) error
	AdminAccountCreating(adminCreatingReq adminModels.AdminCreatingReq) error
	IsPlayer(playerID string) bool
	IsAdmin(adminID string) bool
}
