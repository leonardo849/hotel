package integration_test

import (
	"testing"
)

const (
	name string = "batman and robin are really cool"
	email string = "batman@gmail.com"
	phone string = "333-333-1434"
)

func TestCreateGuest(t *testing.T) {
	e := newExpect(t)
	e.POST("/guest/create"). 
	WithJSON(map[string]string{
		"name": name,
		"email": email,
		"phone": phone,
	}). 
	Expect(). 
	Status(201)

	
}



func TestCreateGuestHopeForError(t *testing.T) {
	e := newExpect(t)
	e.POST("/guest/create"). 
	WithJSON(map[string]string{
		"name": name,
		"email": email,
		"phone": "3233-3233-14324",
	}). 
	Expect(). 
	Status(400)
}

var idGuestInFindAllGuests string

func TestFindAllGuests(t *testing.T) {
	e := newExpect(t)
	response := e.GET("/guest/all").
	Expect(). 
	Status(200).JSON().Array()
	
	if response.Length().Raw() == 0 {
		t.Errorf("length is less than 1")
	}
	firstGuest := response.Value(0).Object()
	idGuestInFindAllGuests = firstGuest.Value("id").String().Raw()
	firstGuest.Value("name").String().IsEqual(name)
	firstGuest.Value("email").String().IsEqual(email)
	firstGuest.Value("phone").String().IsEqual(phone)
}

func UpdateGuest(t *testing.T) {
	url := "/update/" + idGuestInFindAllGuests
	e := newExpect(t)
	e.PUT(url).
	WithJSON(map[string]interface{}{
		"email": "batman1@gmail.com",
	}).
	Expect(). 
	Status(200)
}

func DeleteGuest(t *testing.T) {
	url := "/delete/" + idGuestInFindAllGuests
	e := newExpect(t)
	e.DELETE(url).
	Expect(). 
	Status(200)
}