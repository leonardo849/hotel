package service

import (
	"hotel/internal/dto"
	"hotel/internal/logger"
	"hotel/internal/repository"
	"hotel/internal/validator"

	"go.uber.org/zap"
)

type GuestService struct {
	guestRepository *repository.GuestRepository
}

func NewGuestService(guestRepository *repository.GuestRepository) *GuestService {
	return  &GuestService{
		guestRepository: guestRepository,
	}
}

func (g *GuestService) CreateGuest(input dto.CreateGuestDTO) (status int, message string) {
	if err := validator.Validate.Struct(input); err != nil {
		logger.ZapLogger.Error(
			"validation error in create guest",
			zap.Error(err),
			zap.String("function", "create guest"),
		)
		return 400, err.Error()
	}
	if err := g.guestRepository.CreateGuest(input); err != nil {
		logger.ZapLogger.Error(
			"internal error in create guest",
			zap.Error(err),
			zap.String("function", "create guest"),
		)
		return 500, err.Error()
	}
	return 201, "guest was created"
}

func (g *GuestService) FindAllGuests() (status int, message interface{}) {
	var guests []dto.FindGuestDTO
	var err error
	if guests, err = g.guestRepository.FindAllGuests(); err != nil {
		logger.ZapLogger.Error(
			"internal error in create guest",
			zap.Error(err),
			zap.String("function", "find all guests"),
		)
		return 500, err.Error()
	}
	return 200, guests
}