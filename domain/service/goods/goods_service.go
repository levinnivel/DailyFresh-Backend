package service

import (
	"net/http"
	"strconv"

	Auth "DailyFresh-Backend/authentication"
	Model "DailyFresh-Backend/domain/model/goods"
	Repo "DailyFresh-Backend/domain/repository/goods"
	Response "DailyFresh-Backend/response"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func GetGoods(c *gin.Context) {
	id := c.Query("id")
	goods := Repo.GetGoods(id)

	var responses Response.Response
	if goods != nil {
		responses.Message = "Get Goods success"
		responses.Status = 200
		responses.Data = goods
	} else {
		responses.Message = "Get Goods failed"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func GetGoodsBySeller(c *gin.Context) {
	id := c.Query("seller_id")
	goods := Repo.GetGoodsBySeller(id)

	var responses Response.Response
	if goods != nil {
		responses.Message = "Get Goods By Seller success"
		responses.Status = 200
		responses.Data = goods
	} else {
		responses.Message = "Get Goods By Seller failed"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func GetGoodsByCategory(c *gin.Context) {
	id := c.Query("category")
	goods := Repo.GetGoodsByCategory(id)

	var responses Response.Response
	if goods != nil {
		responses.Message = "Get Goods By Category success"
		responses.Status = 200
		responses.Data = goods
	} else {
		responses.Message = "Get Goods By Category failed"
		responses.Status = 400
	}
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
}

func PostGoods(c *gin.Context) {
	authStatus := Auth.Authenticate(c, "seller")

	if authStatus {
		name := c.PostForm("name")
		price, _ := strconv.Atoi(c.PostForm("price"))
		description := c.PostForm("description")
		category := c.PostForm("category")
		stock, _ := strconv.Atoi(c.PostForm("stock"))
		image := c.PostForm("image")
		seller_id, _ := strconv.Atoi(c.Query("seller_id"))

		var Goods Model.Goods
		Goods.Name = name
		Goods.Price = price
		Goods.Description = description
		Goods.Category = category
		Goods.Stock = stock
		Goods.Image = image
		Goods.SellerID = seller_id

		SuccessPost := Repo.PostGoods(Goods)

		var responses Response.Response
		if SuccessPost {
			responses.Message = "Success Post Goods"
			responses.Status = 200
		} else {
			responses.Message = "Failed Post Goods"
			responses.Status = 400
		}

		c.Header("Content-Type", "application/json")
		c.JSON(http.StatusOK, responses)
	}
}

func UpdateGoods(c *gin.Context) {
	// authStatus := Auth.Authenticate(c, "customer")

	// if authStatus {
	name := c.PostForm("name")
	price, _ := strconv.Atoi(c.PostForm("price"))
	description := c.PostForm("description")
	category := c.PostForm("category")
	stock, _ := strconv.Atoi(c.PostForm("stock"))
	image := c.PostForm("image")
	id := c.PostForm("id")

	Goods := Repo.GetGoods(id)[0]

	if name != "" {
		Goods.Name = name
	}

	if price != 0 {
		Goods.Price = price
	}

	if description != "" {
		Goods.Description = description
	}

	if category != "" {
		Goods.Category = category
	}

	if stock != 0 {
		Goods.Stock = stock
	}

	if image != "" {
		Goods.Image = image
	}

	SuccessPost := Repo.UpdateGoods(Goods)

	var responses Response.Response
	if SuccessPost {
		responses.Message = "Success Update Goods"
		responses.Status = 200
	} else {
		responses.Message = "Failed Update Goods"
		responses.Status = 400
	}

	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, responses)
	// }
}
