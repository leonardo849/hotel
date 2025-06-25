package integration_test

import (
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

func TestFindAllRooms(t *testing.T) {
	e := newExpect(t)
	response := e.GET("/room/all").
	Expect(). 
	Status(200).JSON().Array()

	if response.Length().Raw() == 0 {
		t.Errorf("length is less than 1")
	}

	firstRoom := response.Value(0).Object()

	firstRoom.Value("number").Number().Gt(0)
	firstRoom.Value("type").String().NotEmpty()
	firstRoom.Value("price_per_night").Number().Gt(0)

	
}