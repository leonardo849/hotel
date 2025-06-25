package dto


type CreateRoomDTO struct {
	Number uint `validate:"required,number,gt=0" json:"number"`
	Type string `json:"type" validate:"required,typeofroom"`
	PricePerNight float64   `json:"price_per_night" validate:"required,gt=0,number"`
}