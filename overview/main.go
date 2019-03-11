package main

import (
	"github.com/kataras/iris"
)

type User struct {
	Username  string `json:"username"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	City      string `json:"city"`
	Age       int    `json:"age"`
}

func main() {
	app := iris.New()

	app.RegisterView(iris.HTML("./views", ".html").Reload(true))

	app.OnErrorCode(iris.StatusInternalServerError, func(ctx iris.Context) {
		errMessage := ctx.Values().GetString("error")
		if errMessage != "" {
			ctx.Writef("Internal Server error: %s", errMessage)
			return
		}

		ctx.Writef("Unexpected internal server error")
	})

	app.Use(func(ctx iris.Context) {
		ctx.Application().Logger().Infof("begin request for path: %s", ctx.Path())
		ctx.Next()
	})

	app.Subdomain("mysubdomain.").Post("/decode", func(ctx iris.Context) {

	})

	app.Post("/decode", func(ctx iris.Context) {
		var user User
		ctx.ReadJSON(&user)
		ctx.Writef("%s, %s is %d years old and comes from %s", user.Firstname, user.Lastname, user.Age, user.City)
	})

	app.Get("/encode", func(ctx iris.Context) {
		doe := User{
			Username:  "xiesansi",
			Firstname: "Mingfei",
			Lastname:  "Ding",
			City:      "Hangzhou",
			Age:       23,
		}
		ctx.JSON(doe)
	})

	app.Get("/profile/{username:string}", profileByUsername)

	usersRoutes := app.Party("/users", logThisMiddleware)
	{
		usersRoutes.Get("/{id:int min(1)}", getUserById)
		usersRoutes.Post("/create", createUser)
	}

	app.Run(iris.Addr(":9123"), iris.WithCharset("UTF-8"), iris.WithoutServerError(iris.ErrServerClosed))
}

func logThisMiddleware(ctx iris.Context) {
	ctx.Application().Logger().Infof("Path: %s | IP: %s", ctx.Path(), ctx.RemoteAddr())
	ctx.Next()
}

func profileByUsername(ctx iris.Context) {
	username := ctx.Params().Get("username")
	ctx.ViewData("username", username)
	ctx.View("user/profile.html")

}

func getUserById(ctx iris.Context) {
	userID := ctx.Params().Get("id")
	user := User{
		Username: "username" + userID,
	}
	ctx.JSON(user)
}

func createUser(ctx iris.Context) {
	var user User
	err := ctx.ReadForm(&user)
	if err != nil {
		ctx.Values().Set("error", "creating user, read and parse form failed. "+err.Error())
		ctx.StatusCode(iris.StatusInternalServerError)
		return
	}
	// renders "./views/user/create_verification.html"
	// with {{ . }} equals to the User object, i.e {{ .Username }} , {{ .Firstname}} etc...
	ctx.ViewData("", user)
	ctx.View("user/create_verification.html")
}
