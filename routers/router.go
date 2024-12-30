package routers

import (
	"fullstack-beego-app/controllers"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {

	beego.Router("/home", &controllers.MainController{}, "get:GetHome")

	// Authentication
	beego.Router("/", &controllers.AuthController{}, "get:LoginPage")
	beego.Router("/register", &controllers.AuthController{}, "get:RegisterPage")
	beego.Router("/login", &controllers.AuthController{}, "post:Login")
	beego.Router("/logout", &controllers.AuthController{}, "get:Logout")
	beego.Router("/register", &controllers.AuthController{}, "post:Register")
	beego.Router("/validate-session", &controllers.AuthController{}, "get:ValidateSession")

	// Product routes
	beego.Router("/product", &controllers.ProductController{}, "post:CreateProduct")
	beego.Router("/product/:id", &controllers.ProductController{}, "put:UpdateProduct;delete:DeleteProduct")
	beego.Router("/products", &controllers.ProductController{}, "get:GetAllProducts")

	// Stock routes
	beego.Router("/stock", &controllers.StockController{}, "post:AddStock")
	beego.Router("/stock/:id", &controllers.StockController{}, "put:UpdateStock;delete:DeleteStock")
	beego.Router("/stocks", &controllers.StockController{}, "get:GetAllStocks")

	// User routes
	beego.Router("/user", &controllers.UserController{}, "post:CreateUser")
	beego.Router("/user/:id", &controllers.UserController{}, "put:UpdateUser;delete:DeleteUser")
	beego.Router("/users", &controllers.UserController{}, "get:GetAllUsers")

	// Warehouse routes
	beego.Router("/warehouse", &controllers.WarehouseController{}, "post:CreateWarehouse")
	beego.Router("/warehouse/:id", &controllers.WarehouseController{}, "put:UpdateWarehouse;delete:DeleteWarehouse")
	beego.Router("/warehouses", &controllers.WarehouseController{}, "get:GetAllWarehouses")

	// StockHistory routes
	beego.Router("/stock-history", &controllers.StockHistoryController{}, "post:CreateStockHistory")
	beego.Router("/stock-history/:id", &controllers.StockHistoryController{}, "get:GetStockHistory;delete:DeleteStockHistory")
	beego.Router("/stock-histories", &controllers.StockHistoryController{}, "get:GetAllStockHistories")
}
