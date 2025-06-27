package repository

import (
	"hotel/internal/dto"
	"hotel/internal/model"

	"github.com/thoas/go-funk"
)


type ReturnError struct {
	Code int 
	Error error
}

type GuestRepository struct {
	
}

func NewGuestRepository() *GuestRepository {
	return &GuestRepository{
		
	}
}

func (g *GuestRepository) CreateGuest(input dto.CreateGuestDTO) error {
	guest := model.Guest{
		Name:  input.Name,
		Email: input.Email,
		Phone: input.Phone,
	}
	if err := DB.Create(&guest).Error; err != nil {
		return err
	}
	return nil
}

func (g *GuestRepository) FindAllGuests() ([]dto.FindGuestDTO, error) {
	var guests []model.Guest
	if err := DB.Find(&guests).Preload("Reservations").Error; err != nil {
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

func (g *GuestRepository) FindOneGuest(id string) (*dto.FindGuestDTO, error) {
	var guest model.Guest
	if err := DB.First(&guest, "id = ?", id).Preload("Reservations").Error; err != nil {
		return nil, err
	}
	mappedReservations := funk.Map(guest.Reservations, func(reservation model.Reservation) dto.FindReservationDTO {
		return  dto.FindReservationDTO{
			ID: reservation.ID,
			RoomID: reservation.RoomID,
			GuestID: reservation.GuestID,
			Guest: nil,
			CheckIn: reservation.CheckIn,
			CheckOut: reservation.CheckOut,
			TotalPrice: reservation.TotalPrice,
		}
	}).([]dto.FindReservationDTO)
	guestDTO := dto.FindGuestDTO{
		ID: guest.ID,
		Name: guest.Name,
		Email: guest.Email,
		Phone: guest.Phone,
		Reservations: mappedReservations,
	}
	return  &guestDTO, nil
}

func (g *GuestRepository) UpdateGuest(id string, input dto.UpdateGuestDTO) *ReturnError {
	_, err := g.FindOneGuest(id)
	if err != nil {
		return &ReturnError{
			Code: 404,
			Error: err,
		}
	}
	fields := make(map[string]interface{})
	if input.Email != nil {
		fields["Email"] = *input.Email
	}
	if input.Name != nil {
		fields["Name"] = *input.Name
	}
	if input.Phone != nil {
		fields["Phone"] = *input.Phone
	}
	if err := DB.Model(&model.Guest{}).Where("id = ?", id).Updates(&fields).Error; err != nil {
		return &ReturnError{
			Code: 500,
			Error: err,
		}
	}
	return nil
}
func (g *GuestRepository) DeleteGuest(id string) *ReturnError {
	_, err := g.FindOneGuest(id)
	if err != nil {
		return &ReturnError{
			Code: 404,
			Error: err,
		}
	}
	if err := DB.Delete(&model.Guest{}, "id = ?", id).Error; err != nil {
		return &ReturnError{
			Code: 500,
			Error: err,
		}
	}
	return  nil
}