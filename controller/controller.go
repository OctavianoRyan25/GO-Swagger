package controller

import (
	"OctavianoRyan25/GOSwagger/database"
	"OctavianoRyan25/GOSwagger/model"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateCar godoc
// @Summary Post Details for a given id
// @Description Post detail of car corresponding to the input ID
// @Tags cars
// @Accept json
// @Produce json
// @Param model.Car body model.Car true "create car"
// @Success 200 {object} model.Car
// @Router /cars [post]
func CreateCar(ctx *gin.Context) {
	var db = database.GetDB()
	var input model.Car

	if err := ctx.ShouldBindJSON(&input); err != nil{
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	carInput := model.Car{
		Merk	: input.Merk,
		Type	: input.Type,
		Pemilik	: input.Pemilik,
	}
	db.Create(&carInput)

	ctx.JSON(http.StatusCreated, gin.H{
		"car": carInput,
	})
}

// GetAll godoc
// @Summary Get All Object
// @Description Get Detail All object
// @Tags cars
// @Accept json
// @Produce json
// @Success 200 {object} model.Car
// @Router /cars [get]
func GetAllCar(ctx *gin.Context) {
	var db = database.GetDB()
	var cars []model.Car

	err := db.Find(&cars).Error

	if err != nil {
		fmt.Println("Error getting car datas : ", err)
	}

	ctx.JSON(http.StatusFound, gin.H{
		"car": cars,
	})
}

// GetCarByID godoc
// @Summary Get Car for a given ID
// @Description get detail of car corresponding to the input ID
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200 {object} model.Car
// @Router /car/{id} [get]
func GetCarByID(ctx *gin.Context) {
	var db = database.GetDB()
	var car model.Car

	err := db.First(&car, "ID = ?", ctx.Param("id")).Error

	if err != nil {
		fmt.Println("Error getting car datas : ", err)
	}

	ctx.JSON(http.StatusFound, gin.H{
		"car": car,
	})
}

// UpdateCarByID godoc
// @Summary Update Car for a given ID
// @Description Update detail of car corresponding to the input ID
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200 {object} model.Car
// @Router /car/{id} [patch]
func UpdateCarByID(ctx *gin.Context) {
	var db = database.GetDB()
	var car model.Car

	err := db.First(&car, "ID = ?", ctx.Param("id")).Error

	if err != nil {
		fmt.Println("Error getting car datas : ", err)
	}

	var input model.Car
	if err := ctx.ShouldBindJSON(&input);err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error " : err.Error(),
		})
		return
	}

}

// DeleteCArByID godoc
// @Summary Delete Car for a given ID
// @Description Delete detail of car corresponding to the input ID
// @Tags cars
// @Accept json
// @Produce json
// @Param Id path int true "ID of the car"
// @Success 200 {object} model.Car
// @Router /cars/{id} [delete]
func DeleteCarByID(ctx *gin.Context) {
	var db = database.GetDB()
	var carDelete model.Car

	err := db.First(&carDelete, "ID = ?", ctx.Param("id")).Error

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "record not found"})
	}

	db.Delete(&carDelete)
	ctx.JSON(http.StatusOK, gin.H{"Message": "Succesfully delete data"})

}