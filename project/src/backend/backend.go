package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-gorm/postgres"
	_ "github.com/lib/pq"
	"gorm.io/gorm"
)

// Dados de acesso à base de dados
const (
	host     = "database"
	port     = "5432"
	user     = "test"
	password = "passw0rd"
	dbname   = "apidb"
)

// Tabelas da base de dados

type Worker struct {
	gorm.Model
	Worker_ID int    "gorm:primaryKey"
	Username  string "gorm:unique"
	password  string
	Name      string "gorm:unique"
	Admin     bool
}

type Zone struct {
	gorm.Model
	Zone_ID   int    "gorm:primaryKey"
	Name      string "gorm:unique"
	Latitude  float32
	Longitude float32
	Limit     int
}

type Access struct { // Lista de Zonas a que o worker pode aceder
	gorm.Model
	IdWorker int "gorm:\"foreignKey:Worker_ID\""
	IdZone   int "gorm:\"foreignKey:Zone_ID\""
}

type WorkersZones struct { // Lista de workers a trabalhar em cada zona
	gorm.Model
	IdWorker int "gorm:\"foreignKey:Worker_ID\""
	IdZone   int "gorm:\"foreignKey:Zone_ID\""
}

var db *gorm.DB

// Ligação à base de dados e criação de tabelas
func init() {

	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(
		&Worker{},
		&Zone{},
		&Access{},
		&WorkersZones{},
	)

	fmt.Println("Successfully connected!")
}

// Inicialização das rotas
func main() {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true           //maybe reduce to localhost only
	config.AllowCredentials = true          //
	config.AddAllowHeaders("authorization") //maybe not needed

	router := gin.Default()
	router.Use(cors.New(config))
	router.POST("/login", login)

	/* 	routerZones := router.Group("/zones")
		{
			//routerZones.GET("/", listZones)
			//routerZones.POST("/:id", addWorkerToZone)
			//routerZones.DELETE("/:id", removeWorkerFromZone)
		}

		routerbackOffice := router.Group("/admin")
		{
			routerbackOffice.GET("/users", adminListUsers)
			routerbackOffice.POST("/users", adminCreateUser)
			routerbackOffice.PUT("/users/:id", adminManageUser)
			routerbackOffice.DELETE("/users/:id", adminRemoveUser)

	    routerbackOffice.GET("/zones", adminListZones)
			routerbackOffice.POST("/zones", adminCreateZone)
			routerbackOffice.DELETE("/zones/:id", adminRemoveZone) // feito

		} */

	router.Run(":8081")
}

/////////////////////////////////////// CONFIRMAR //////////////////////////////////////

type loginData struct {
	user  Worker
	token string
}

// Login
func login(c *gin.Context) {
	var worker Worker
	username := c.PostForm("username")
	password := c.PostForm("password")
	db.Find(&worker, "Username = ?", username)
	if worker.password == password { //usar hashes
		data := loginData{user: worker, token: "123456789"}
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
			"data":   data,
		})
	}
}

// Adicionar Zona
func adminCreateZone(c *gin.Context) {
	zoneid, _ := strconv.Atoi(c.PostForm("Zone_ID"))
	latitude, _ := strconv.ParseFloat(c.PostForm("latitude"), 32)
	longitude, _ := strconv.ParseFloat(c.PostForm("longitude"), 32)
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	zone := Zone{Zone_ID: zoneid, Name: c.PostForm("name"), Latitude: float32(latitude), Longitude: float32(longitude), Limit: limit}
	db.Save(&zone)
	c.JSON(http.StatusCreated, gin.H{"status": http.StatusCreated, "message": "Todo item created successfully!", "resourceId": zone.ID})
}

// Remover Zona
func adminRemoveZone(c *gin.Context) {
	var zone Zone
	zoneID := c.Param("id")

	db.Delete(&zone, "id = ?", zoneID)

	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "message": "Zone deleted successfully!"})
}
