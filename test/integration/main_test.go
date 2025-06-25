package integration_test

import (
	"hotel/config"
	"hotel/internal/logger"
	"hotel/internal/model"
	"hotel/internal/repository"
	"hotel/internal/router"
	"hotel/internal/validator"
	"log"
	"net/http"
	"os"
	"testing"

	"github.com/gavv/httpexpect/v2"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

var app *fiber.App


type fiberRoundTripper struct {
	app *fiber.App
}

var client = &http.Client{
	Transport: fiberRoundTripper{app: app},
}



func (rt fiberRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	return rt.app.Test(req)
}

var DB *gorm.DB

func TestMain(m *testing.M) {
	err := config.SetupEnvVar()
	if err != nil {
		log.Panic(err.Error())
	}
	if err = logger.StartLogger(); err != nil {
		log.Panic(err.Error())
	}
	db, err := repository.ConnectToDatabase()
	if err != nil {
		log.Panic(err.Error())
	}
	validator.StartValidate()
	
	DB = db
	app = router.SetupApp()
	sqldb, err := db.DB()
	if err != nil {
		log.Panic(err.Error())
	}
	
	cleanDatabase(db)
	code := m.Run()
	sqldb.Close()
	os.Exit(code)
}

func newExpect(t *testing.T) *httpexpect.Expect {
	client := &http.Client{
		Transport: fiberRoundTripper{app: app},
	}
	return httpexpect.WithConfig(httpexpect.Config{
		BaseURL:  "http://localhost",
		Client:   client,
		Reporter: httpexpect.NewRequireReporter(t),
	})
}

func cleanDatabase(db *gorm.DB) {
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Reservation{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Guest{})
	db.Session(&gorm.Session{AllowGlobalUpdate: true}).Delete(&model.Room{})
}