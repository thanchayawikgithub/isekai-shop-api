package services

import (
	"github.com/thanchayawikgithub/isekai-shop-api/internal/entities"
	adminModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/admin/models"
	adminRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/admin/repositories"
	playerModels "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/player/models"
	playerRepositories "github.com/thanchayawikgithub/isekai-shop-api/internal/modules/player/repositories"
)

type googleOAuth2Service struct {
	playerRepo playerRepositories.PlayerRepository
	adminRepo  adminRepositories.AdminRepository
}

func NewGoogleOAuth2Service(playerRepo playerRepositories.PlayerRepository, adminRepo adminRepositories.AdminRepository) Oauth2Service {
	return &googleOAuth2Service{playerRepo, adminRepo}
}

func (s *googleOAuth2Service) PlayerAccountCreating(playerCreatingReq *playerModels.PlayerCreatingReq) error {
	if s.IsPlayer(playerCreatingReq.ID) {
		return nil
	}

	playerEntity := &entities.Player{
		ID:     playerCreatingReq.ID,
		Name:   playerCreatingReq.Name,
		Email:  playerCreatingReq.Email,
		Avatar: playerCreatingReq.Avatar,
	}

	_, err := s.playerRepo.Creating(playerEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *googleOAuth2Service) AdminAccountCreating(adminCreatingReq *adminModels.AdminCreatingReq) error {
	if s.IsAdmin(adminCreatingReq.ID) {
		return nil
	}

	adminEntity := &entities.Admin{
		ID:     adminCreatingReq.ID,
		Name:   adminCreatingReq.Name,
		Email:  adminCreatingReq.Email,
		Avatar: adminCreatingReq.Avatar,
	}

	_, err := s.adminRepo.Creating(adminEntity)
	if err != nil {
		return err
	}

	return nil
}

func (s *googleOAuth2Service) IsPlayer(playerID string) bool {
	player, err := s.playerRepo.FindByID(playerID)
	if err != nil {
		return false
	}

	return player != nil
}

func (s *googleOAuth2Service) IsAdmin(adminID string) bool {
	admin, err := s.adminRepo.FindByID(adminID)
	if err != nil {
		return false
	}

	return admin != nil
}
