package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"net/http"
	"os"
)

func main() {
	fmt.Println("Hello world")
	databaseTest()
}

// >docker run --name some-mysql -e MYSQL_ROOT_PASSWORD="12345678a" -e MYSQL_DATABASE="food_delivery" -p 3306:3306 -d mysql:latest
func databaseTest() {
	dsn := os.Getenv("MYSQL_CONNECTION")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Cann't connect to DB")
	}
	log.Println(db, err)
	//newRes := Restaurant{Name: "test", Address: "Nha 2 con sau"}
	//err = db.Create(&newRes).Error
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println("New id", newRes.Id)
	//var myRestaurant Restaurant
	//err = db.Where("id = ?", 3).First(&myRestaurant).Error
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(myRestaurant)
	//myRestaurant.Name = "Nha sau voi"
	//err = db.Where("id = ?", myRestaurant.Id).Updates(&myRestaurant).Error
	//if err != nil {
	//	log.Println(err)
	//}
	//log.Println(myRestaurant)
	//db.Table(Restaurant{}.TableName()).Where("id = ?", myRestaurant.Id).Delete(nil)
	r := gin.Default()
	r.GET("/hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})
	r.GET("/restaurant/:id", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "Hello",
		})
	})
	r.Run()
}

type Restaurant struct {
	Id      int    `json:"id" gorm:"column:id;"`
	Name    string `json:"name" gorm:"column:name"`
	Address string `json:"address" gorm:"column:address"`
}

// Dung de update cac truong
type RestaurantUpdate struct {
	Name    *string `json:"name" gorm:"column:name"`
	Address *string `json:"address" gorm:"column:address"`
}

func (res Restaurant) TableName() string {
	return "restaurants"
}

func (res RestaurantUpdate) TableName() string {
	return "restaurants"
}
