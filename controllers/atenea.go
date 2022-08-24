package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/RaulGarciaMz/ateneaserver/database"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

const AteneaLogKey = "atenea"

type Atenea struct {
	db *database.AteneaRepo
}

func AteneaController(db *gorm.DB) *Atenea {

	return &Atenea{
		db: database.NewAteneaRepo(db),
	}
}

// MonitoreoEquipo godoc
// @Summary Obtiene el monitoreo por equipos
// @Description Obtiene el monitoreo del equipo correspondiente al identificador del equipo indicado
// @Tags monitoreo
// @Produce  json
// @Param id path string true "Identificador (id) del equipo"
// @Success 200 {array} models.MonitoreoEquipo
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /monitoreo/equipo/{id} [get]
func (e *Atenea) MonitoreoEquipo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_eq, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.MonitoreoEquipo(int32(id_eq))
		if err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, err)
				return
			}

			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, single)
	}
}

// MonitoreoEquipos godoc
// @Summary Obtiene el monitoreo por equipos
// @Description Obtiene el monitoreo de todos los equipos
// @Tags monitoreo
// @Produce  json
// @Success 200 {array} models.MonitoreoEquipo
// @Failure 404
// @Failure 500
// @Router /monitoreo/equipo [get]
func (e *Atenea) MonitoreoEquipos() gin.HandlerFunc {
	return func(c *gin.Context) {

		single, err := e.db.MonitoreoEquipos()
		if err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, err)
				return
			}

			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, single)
	}
}

// MonitoreoGrupos godoc
// @Summary Obtiene los datos de monitoreo de los grupos
// @Description Obtiene los datos de monitoreo de los grupos
// @Tags monitoreo
// @Produce  json
// @Success 200 {array} models.MonitoreoGrupo
// @Failure 404
// @Failure 500
// @Router /monitoreo/grupos [get]
func (e *Atenea) MonitoreoGrupos() gin.HandlerFunc {
	return func(c *gin.Context) {

		single, err := e.db.MonitoreoGrupos()
		if err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, err)
				return
			}

			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, single)
	}
}

// MonitoreoGrupo godoc
// @Summary Obtiene el monitoreo por grupos
// @Description Obtiene el monitoreo del equipo correspondiente al identificador del grupo indicado
// @Tags monitoreo
// @Produce  json
// @Param id path string true "Identificador (id) del grupo"
// @Success 200 {array} models.MonitoreoEquipo
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /monitoreo/grupo/{id} [get]
func (e *Atenea) MonitoreoGrupo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_eq, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.MonitoreoEquipo(int32(id_eq))
		if err != nil {

			if errors.Is(err, gorm.ErrRecordNotFound) {
				c.JSON(http.StatusNotFound, err)
				return
			}

			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, single)
	}
}
