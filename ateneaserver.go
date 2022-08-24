package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RaulGarciaMz/ateneaserver/configuration"
	"github.com/RaulGarciaMz/ateneaserver/database"
	"github.com/RaulGarciaMz/ateneaserver/routes"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	/*
		 	_ "github.com/swaggo/files"
			_ "github.com/swaggo/gin-swagger"
	*/)

const executableID = "ATENEA"

var (
	Sha1ver        string // clave sha1 del commit usado para compilar el programa
	Branch         string //nombre de la rama usada para compilar el programa
	BuildTime      string // when the executable was built
	LastCommitTime string // when the last commit was
	Tag            string // nombre del último tag registrado en la rama
	flgVersion     bool
)

// @title Servidor de APIs para ATENEA
// @version 1.0
// @description Contiene las APIs para atender las peticiones de la IHM hacia la base de datos ATENEA

// @contact.name Soporte para las APIS: Raúl García
// @contact.url http://www.ineel.mx
// @contact.email rgarciaxxx@yahoo.com.mx

// @license.name Atenea Inc.
// @license.url http://www.atenea.mx/LICENSE-2.0.html

// @host IPServer:8085
// @BasePath /atenea/admin
func main() {
	var err error
	var dbConn *gorm.DB

	version := generaVersion()
	parseCmdLineFlags(version)

	confFile, err := os.Open("ateneaconf.json")
	if err != nil {
		panic(err)
	}
	defer confFile.Close()
	conf, err := ioutil.ReadAll(confFile)
	if err != nil {
		panic(err)
	}

	myConf := configuration.Configuration{}
	err = json.Unmarshal(conf, &myConf)
	if err != nil {
		panic(err)
	}

	dbConn, err = database.GetConnection(myConf.Database.Username, myConf.Database.Password, myConf.Database.Dbname, myConf.Database.Host)
	if err != nil {
		//redislog.Fatal("error obteniendo conexión a BD: ", err)
		panic(err)
	}

	router := gin.New()
	//router.Use(gin.LoggerWithWriter(redislog.GlobalWriter()), gin.Recovery())

	confApi := router.Group("atenea")
	withOptionsHandler := WithOptions(http.MethodGet, http.MethodDelete, http.MethodPut, http.MethodPost)

	/// agrega rutas
	routes.AteneaRoute{}.CreaRuta(confApi, dbConn, "admin", withOptionsHandler)

	confApi = router.Group("api")
	withOptionsVer := WithOptions(http.MethodGet)
	publishGroup := confApi.Group("/version-api", withOptionsVer)
	{
		publishGroup.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, version)
		})
		publishGroup.OPTIONS("/*id", withOptionsVer)
	}

	err = router.Run(myConf.Server.Port)
	if err != nil {
		//redislog.Fatal("Error atando servidor")
		panic(err)
	}
}

func parseCmdLineFlags(version string) {
	flag.BoolVar(&flgVersion, "version", false, "si true, imprime la versión y termina el programa")
	flag.Parse()

	if flgVersion {
		fmt.Println(version)
		os.Exit(0)
	}
}

func generaVersion() string {
	var version string

	tUnix, err := strconv.ParseInt(LastCommitTime, 10, 64)
	if err != nil {
		version = "Compilado el: " + BuildTime + " Rama: " + Branch + " Commit (sha1): " + Sha1ver + " Tag: " + Tag
	}
	timeT := time.Unix(tUnix, 0)
	version = "Compilado el: " + BuildTime +
		" Rama: " + Branch +
		" Commit (sha1): " + Sha1ver +
		" Fecha Commit: " + timeT.Format("02-01-2006 15:04:00") +
		" Tag: " + Tag

	return version
}

func WithOptions(methods ...string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", strings.Join(methods, ", "))
		if len(c.Request.Header["Access-Control-Request-Headers"]) > 0 {
			var headersList []string
			for _, header := range c.Request.Header[http.CanonicalHeaderKey("Access-Control-Request-Headers")] {
				headersList = append(headersList, header)
			}
			c.Writer.Header().Set("Access-Control-Allow-Headers", strings.Join(headersList, ", "))
		}

		if c.Request.Method != http.MethodOptions {
			c.Next()
		}
	}
}
