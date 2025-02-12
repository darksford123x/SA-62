package main

import (
	"context"
	"log"

	"github.com/darksford123x/repairinvoice-record/controllers"
	_ "github.com/darksford123x/repairinvoice-record/docs"
	"github.com/darksford123x/repairinvoice-record/ent"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Devices struct {
	Device []Device
}

type Device struct {
	Dname string
	id    int
}

type Symptoms struct {
	Symptom []Symptom
}

type Symptom struct {
	Syname string
}

type StatusRs struct {
	StatusR []StatusR
}

type StatusR struct {
	Sname string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:ent.db?&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")
	controllers.NewDeviceController(v1, client)
	controllers.NewStatusRController(v1, client)
	controllers.NewSymptomController(v1, client)
	controllers.NewRepairInvoiceController(v1, client)

	// Set Statusrs Data
	statusrs := StatusRs{
		StatusR: []StatusR{
			StatusR{"In process "},
			StatusR{"done"},
		},
	}
	for _, s := range statusrs.StatusR {
		client.StatusR.
			Create().
			SetSname(s.Sname).
			Save(context.Background())
	}

	// Set Devices Data
	devices := Devices{
		Device: []Device{
			Device{"Acer Predator", 1},
			Device{"Mi Air", 1},
			Device{"Huawei Y800", 1},
			Device{"Lenovo NotePad", 2},
			Device{"Samsung ok10", 2},
			Device{"MacBook Pro 2020", 2},
		},
	}
	for _, d := range devices.Device {
		client.Device.
			Create().
			SetDname(d.Dname).
			Save(context.Background())
	}

	// Set Symptoms Data
	symptoms := Symptoms{
		Symptom: []Symptom{
			Symptom{"Damaged Hardware"},
			Symptom{"Damaged Software"},
		},
	}
	for _, sy := range symptoms.Symptom {
		client.Symptom.
			Create().
			SetSyname(sy.Syname).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
