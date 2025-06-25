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