package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pawanpaudel93/go-mux-restapi/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func InitDatabase() {
	dsn := "host=localhost user=gorm password=gorm dbname=gorm sslmode=disable"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.Resource{})
}

func GetResources(w http.ResponseWriter, r *http.Request) {
	var resources []models.Resource
	db.Find(&resources)
	json.NewEncoder(w).Encode(&resources)
}

func GetResource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var resource models.Resource
	db.First(&resource, params["id"])
	json.NewEncoder(w).Encode(&resource)
}

func CreateResource(w http.ResponseWriter, r *http.Request) {
	var resource models.Resource
	json.NewDecoder(r.Body).Decode(&resource)
	db.Create(&resource)
	json.NewEncoder(w).Encode(&resource)
}

func UpdateResource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var resource models.Resource
	var updateData map[string]interface{}
	json.NewDecoder(r.Body).Decode(&updateData)
	db.First(&resource, params["id"])
	db.Model(&resource).Updates(updateData)
	json.NewEncoder(w).Encode(&resource)
}

func DeleteResource(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var resource models.Resource
	db.First(&resource, params["id"])
	db.Delete(&resource)

	var resources []models.Resource
	db.Find(&resources)
	json.NewEncoder(w).Encode(&resources)
}
