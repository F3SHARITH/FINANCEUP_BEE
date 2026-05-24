package main

import (
	_ "educacion/routers"
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	//cargar .env
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}

	err = beego.LoadAppConfig("ini", "conf/app.conf")
	if err != nil {
		fmt.Println("Error loading app config:", err)
		panic(err)
	}

	pgUser, _ := beego.AppConfig.String("PGuser")
	pgPass, _ := beego.AppConfig.String("PGpass")
	pgHost, _ := beego.AppConfig.String("PGhost")
	pgPort, _ := beego.AppConfig.String("PGport")
	pgDb, _ := beego.AppConfig.String("PGdb")
	pgSchema, _ := beego.AppConfig.String("PGschema")

	fmt.Printf("PostgreSQL connection string: postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s\n", pgUser, pgPass, pgHost, pgPort, pgDb, pgSchema)
	orm.RegisterDataBase(
		"default",
		"postgres",
		"postgres://"+
			pgUser+":"+
			pgPass+"@"+
			pgHost+":"+
			pgPort+"/"+
			pgDb+
			"?sslmode=disable&search_path="+
			pgSchema,
	)

	beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
	}
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders: []string{"Origin", "x-requested-with",
			"content-type",
			"accept",
			"origin",
			"authorization",
			"x-csrftoken"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	beego.Run()
}
