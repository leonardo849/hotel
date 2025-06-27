package service

import (
	"hotel/internal/dto"
	"hotel/internal/repository"
	"hotel/internal/validator"
)

type ReservationService struct {
	reservationRepository *repository.ReservationRepository
}

func NewReservationService(reservationRepository *repository.ReservationRepository) *ReservationService {
	return  &ReservationService{
		reservationRepository: reservationRepository,
	}
}

func (r *ReservationService) CreateReservation(input dto.CreateReservationDTO) (status int, message string) {
	if err := validator.Validate.Struct(input); err != nil {
		return 400, err.Error()
	}
	err := r.reservationRepository.CreateReservation(input)
	if err != nil {
		return 500, "internal server error"
	}
	return 201, "reservation was created"
}