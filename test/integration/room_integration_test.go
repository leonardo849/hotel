package integration_test

import (
	"hotel/internal/model"
	"testing"
)

func TestCreateRoom(t *testing.T) {
	e := newExpect(t)

	e.POST("/room/create").
	WithJSON(map[string]interface{}{
		"number": 40,
		"type": "suite",
		"price_per_night": 40, 
	}).
	Expect(). 
	Status(201)

}



var idInFindAllRooms string
func TestFindAllRooms(t *testing.T) {
	e := newExpect(t)
	response := e.GET("/room/all").
	Expect(). 
	Status(200).JSON().Array()

	if response.Length().Raw() == 0 {
		t.Errorf("length is less than 1")
	}

	firstRoom := response.Value(0).Object()
	idInFindAllRooms = firstRoom.Value("id").String().Raw()
	firstRoom.Value("number").Number().Gt(0)
	firstRoom.Value("type").String().NotEmpty()
	firstRoom.Value("price_per_night").Number().Gt(0)
}


func TestFindOneRoom(t *testing.T) {
	e := newExpect(t)
	var rooms []model.Room
	if err := DB.Find(&rooms).Error; err != nil {
		t.Errorf("error in find rooms %v", err.Error())
	}
	id := rooms[0].ID.String()
	url := "/room/one/" + string(id)
	response := e.GET(url).
	Expect(). 
	Status(200).JSON().
	Object()
	
	idString := response.Value("id").String().NotEmpty().Raw()
	if idString != idInFindAllRooms {
		t.Errorf("id is wrong")
	}
	response.Value("number").Number().Gt(0)
	response.Value("type").String().NotEmpty()
}

func TestUpdateRoom(t *testing.T) {
	url := "/room/update/" + string(idInFindAllRooms)
	e := newExpect(t)
	e.PUT(url).
	WithJSON(map[string]interface{}{
		"number": 50,
	}).
	Expect(). 
	Status(200)
}

func TestDeleteRoom(t *testing.T) {
	url := "/room/delete/" + string(idInFindAllRooms)
	e := newExpect(t)
	e.DELETE(url).
	Expect(). 
	Status(200)
}