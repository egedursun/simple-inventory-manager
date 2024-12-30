package middlewares

import (
	"fullstack-beego-app/models"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"log"
	"time"
)

func LoggingMiddleware(ctx *context.Context) {
	start := time.Now()
	log.Printf("Started %s %s", ctx.Request.Method, ctx.Request.RequestURI)
	ctx.ResponseWriter.Header().Set("X-Custom-Header", "Inventory Wizard")

	defer func() {
		log.Printf("Completed %s in %v", ctx.Request.RequestURI, time.Since(start))
	}()
}

func AuthMiddleware(ctx *context.Context) {
	allowedURLs := []string{
		"/",
		"/login",
		"/register",
		"/logout",
		"/validate-session",
	}

	for _, url := range allowedURLs {
		if ctx.Input.URL() == url {
			return
		}
	}

	if ctx.Input.Session("uid") == nil {
		ctx.Redirect(302, "/")
	}
}

func CORSMiddleware(ctx *context.Context) {
	ctx.Output.Header("Access-Control-Allow-Origin", "*")
	ctx.Output.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE")
	ctx.Output.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
	if ctx.Input.Method() == "OPTIONS" {
		ctx.ResponseWriter.WriteHeader(204)
	}

}

func CacheControlMiddleware(ctx *context.Context) {
	ctx.Output.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	ctx.Output.Header("Pragma", "no-cache")
	ctx.Output.Header("Expires", "0")
}

func MethodOverride(ctx *context.Context) {
	log.Printf("Original method: %s, URL: %s", ctx.Input.Method(), ctx.Input.URL())
	if ctx.Input.Method() == "POST" {
		overrideMethod := ctx.Input.Query("_method")
		log.Printf("Override method: %s", overrideMethod)
		if overrideMethod == "PUT" || overrideMethod == "DELETE" || overrideMethod == "PATCH" {
			ctx.Input.SetData("MethodOverride", overrideMethod)
			ctx.Request.Method = overrideMethod
			log.Printf("Overridden method: %s", ctx.Request.Method)
		}
	}
}

func AdminOnlyMiddleware(ctx *context.Context) {
	allowedURLs := []string{
		"/",
		"/login",
		"/register",
		"/logout",
		"/validate-session",
	}

	for _, url := range allowedURLs {
		if ctx.Input.URL() == url {
			return
		}
	}

	uid := ctx.Input.Session("uid")
	if uid == nil {
		log.Println("User not logged in; redirecting to login.")
		ctx.Redirect(303, "/login")
		return
	}

	o := orm.NewOrm()
	user := models.User{ID: uid.(int)}
	err := o.Read(&user)
	if err != nil {
		log.Printf("Error fetching user with ID %d: %v; redirecting to login.", uid, err)
		ctx.Redirect(303, "/login")
		return
	}

	if user.Role == "admin" {
		return
	}

	method := ctx.Input.Method()
	if method != "GET" {
		log.Printf("Non-admin user '%v' attempted %v operation; denying access.", user.Username, method)

		// Render the access denied page
		ctx.Output.Status = 403
		err = ctx.Output.Body([]byte(`
					<!DOCTYPE html>
					<html lang="en">
					<head>
						<meta charset="UTF-8">
						<meta name="viewport" content="width=device-width, initial-scale=1.0">
						<title>Access Denied</title>
						<link href="https://cdn.jsdelivr.net/npm/bootstrap@5.3.0/dist/css/bootstrap.min.css" rel="stylesheet">
					</head>
					<body>
						<div class="container text-center mt-5">
							<h1 class="text-danger">Access Denied</h1>
							<p>You do not have the necessary permissions to perform this action.</p>
							<p class="text-muted">User: ` + user.Username + `</p>
							<a href="/home" class="btn btn-primary">Go Back</a>
						</div>
					</body>
					</html>
				`))
		if err != nil {
			log.Printf("Error rendering access denied page: %v", err)
			return
		}
		return
	}
}
