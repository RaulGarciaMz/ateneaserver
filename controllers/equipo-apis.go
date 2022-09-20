package controllers

import (
	"net/http"
	"strconv"

	"github.com/RaulGarciaMz/ateneaserver/models"
	"github.com/gin-gonic/gin"
)

// AltaEquipo godoc
// @Summary Crea un nuevo registro de equipo
// @Description Agrega un nueva equipo
// @Tags equipo
// @Accept  json
// @Produce  json
// @Param equipo body models.EquipoParam true "Estructura del equipo a agregar"
// @Success 201 {string} string
// @Failure 400
// @Failure 409 {string} string
// @Failure 417 {string} string
// @Failure 500
// @Router /equipo [post]
func (e *Atenea) AltaEquipo() gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyPost := &models.EquipoParam{}
		err := c.BindJSON(bodyPost)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.AltaEquipo(bodyPost.Nombre, bodyPost.Ip, bodyPost.Descripcion, bodyPost.Puerto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		switch single {
		case "Éxito":
			c.JSON(http.StatusCreated, single)
		case "Existente":
			c.JSON(http.StatusConflict, single)
		default:
			c.JSON(http.StatusExpectationFailed, single)
		}
	}
}

// FiltroEquipo godoc
// @Summary Cambia el valor de la bandera de filtro de equipos
// @Description cambia el valor de la bandera de filtro del equipo indicado
// @Tags equipo
// @Accept  json
// @Produce  json
// @Param equipo body models.EquipoParam true "Estructura del equipo a agregar"
// @Success 201 {string} string
// @Failure 400
// @Failure 409 {string} string
// @Failure 417 {string} string
// @Failure 500
// @Router /equipo/filtro [post]
func (e *Atenea) FiltroEquipo() gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyPost := &models.FiltroEquipoParam{}
		err := c.BindJSON(bodyPost)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.FiltroEquipo(bodyPost.Id, bodyPost.Filtro)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		switch single {
		case "Corrupto":
			c.JSON(http.StatusConflict, single)

		default:
			c.JSON(http.StatusOK, single)
		}
	}
}

// ActualizaEquipo godoc
// @Summary Actualiza datos de un equipo
// @Description Actualiza el registro del equipo indicado en el body de la petición
// @Tags equipo
// @Accept  json
// @Produce  json
// @Param equipo body models.EquipoParam true "Estructura del equipo a actualizar"
// @Success 200 {string} string
// @Failure 400
// @Failure 409 {string} string
// @Failure 417 {string} string
// @Failure 500
// @Router /equipo [put]
func (e *Atenea) ActualizaEquipo() gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyPost := &models.EquipoParam{}
		err := c.BindJSON(bodyPost)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ActualizaEquipo(bodyPost.Id, bodyPost.Nombre, bodyPost.Ip, bodyPost.Descripcion, bodyPost.Puerto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		switch single {
		case "Éxito":
			c.JSON(http.StatusOK, single)
		case "Nombre o IP existentes":
			c.JSON(http.StatusConflict, single)
		default:
			c.JSON(http.StatusExpectationFailed, single)
		}
	}
}

// BajaEquipo godoc
// @Summary Elimina equipos
// @Description Borra el equipo correspondiente al valor del identificador (id) indicado.
// @Tags equipo
// @Accept  json
// @Produce  json
// @Param id path string true "Identificador (id) del equipo a eliminar"
// @Success 200 {string} string
// @Failure 400
// @Failure 409 {string} string
// @Failure 417 {string} string
// @Failure 500
// @Router /equipo/id/{id} [delete]
func (e *Atenea) BajaEquipo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_equ, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.BajaEquipo(int32(id_equ))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		switch single {
		case "Éxito":
			c.JSON(http.StatusOK, single)
		case "Existente":
			c.JSON(http.StatusConflict, single)
		default:
			c.JSON(http.StatusExpectationFailed, single)
		}
	}
}

// ListaEquipos godoc
// @Summary Obtiene la lista de equipos
// @Description Obtiene la lista de equipos
// @Tags equipo
// @Produce  json
// @Success 200 {array} models.Equipo
// @Failure 500
// @Router /equipo [get]
func (e *Atenea) ListaEquipos() gin.HandlerFunc {
	return func(c *gin.Context) {

		single, err := e.db.ListaEquipos()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, single)
	}
}
