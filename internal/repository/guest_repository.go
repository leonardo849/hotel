package repository

import (
	"hotel/internal/dto"
	"hotel/internal/model"

	"github.com/thoas/go-funk"
	"gorm.io/gorm"
)

type GuestRepository struct {
	guestModel *gorm.DB
}

func NewGuestRepository(guestModel *gorm.DB) *GuestRepository {
	return &GuestRepository{
		guestModel: guestModel,
	}
}

func (g *GuestRepository) CreateGuest(input dto.CreateGuestDTO) error {
	guest := model.Guest{
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}
	if err := g.guestModel.Create(&guest).Error; err != nil {
		return err
	}
	return nil
}

func (g *GuestRepository) FindAllGuests() ([]dto.FindGuestDTO, error) {
	var guests []model.Guest
	if err := g.guestModel.Find(&guests).Preload("Reservations").Error; err != nil {
		return nil, err
	}
	mapped := funk.Map(guests, func(guest model.Guest) dto.FindGuestDTO {
		reservationsMapped := funk.Map(guest.Reservations, func(reservation model.Reservation) dto.FindReservationDTO {
			return dto.FindReservationDTO{
				ID: reservation.ID,
				RoomID: reservation.RoomID,
				GuestID: reservation.GuestID,
				Guest: nil,
				CheckIn: reservation.CheckIn,
				CheckOut: reservation.CheckOut,
				TotalPrice: reservation.TotalPrice,
			}
		}).([]dto.FindReservationDTO)
		return dto.FindGuestDTO{
			ID:        guest.ID,
			Name:      guest.Name,
			Email:     guest.Email,
			Phone:     guest.Phone,
			Reservations: reservationsMapped,
			CreatedAt: guest.CreatedAt,
			UpdatedAt: guest.UpdatedAt,
		}
	}).([]dto.FindGuestDTO)
	return mapped, nil

}
