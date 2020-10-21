package controllers

import (
	"context"
	"fmt"
	"git.qix.sx/hackathon/01-go/tank-rush/hackathon-2020/models"
	"github.com/gin-gonic/gin"
	redisCache "github.com/go-redis/cache/v8"
	"gorm.io/gorm"
	"log"
	"strconv"
)

func Put_info(c *gin.Context, db *gorm.DB) {
	db.AutoMigrate(&models.Review{})
	var product models.Product
	db.First(&product, "code = ?", c.Query("asin"))
	InsertQuery := &models.Review{Asin: c.Query("asin"), Title: c.Query("title"), Content: c.Query("content"), ProdId: product.ID}
	result := db.Create(&InsertQuery)
	if result != nil {
		c.JSON(200, gin.H{
			"Review ID":    InsertQuery.ID,
			"Review ASIN":  InsertQuery.Asin,
			"Review Title": InsertQuery.Title,
			"Review Body": InsertQuery.Content,
		})
	}
}

func Get_info(c *gin.Context, db *gorm.DB, cache *redisCache.Cache) {
	product_id := c.Param("product_id")
	token := c.Query("token")
	var reviews_page string = "1"
	var per_page string = "2"
	if len(c.Query("reviews_page")) > 0 {
		reviews_page = c.Query("reviews_page")
	}
	if len(c.Query("per_page")) > 0 {
		per_page = c.Query("per_page")
	}
	limit, err := strconv.Atoi(per_page)
	if err != nil {
		log.Fatal(err)
	}
	offset, err := strconv.Atoi(reviews_page)
	var product models.Product
	var reviews []models.Review

	productKey := fmt.Sprintf("product_%s", product_id)
	reviewsKey := fmt.Sprintf("reviews_%s", product_id)
	if err := cache.Get(context.TODO(), productKey, &product); err != nil {
		err := db.First(&product, "code = ?", product_id).Error
		if err != nil {
			c.JSON(404, gin.H{
				"status":  false,
				"message": "Product not found",
			})
		}

	} //TODO: Add error handling for product getter
	if err := cache.Get(context.TODO(), reviewsKey, &reviews); err != nil {
		db.Limit(limit).Offset(offset*limit-limit).Find(&reviews, "asin = ?", product.Code) //Todo checks offest*limit, to prevent OFFSET 1 on default, and OFFSET larger that rows.affected
	} //TODO: Add error handling for reviews getter

	c.JSON(200, gin.H{
		"Product Title":  product.Title,
		"Product ASIN":   product.Code,
		"Security token": token,
		"Reviews":        reviews,
	})
}
