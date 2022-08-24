package controllers

import (
	"net/http"
	"strconv"

	"github.com/RaulGarciaMz/ateneaserver/models"
	"github.com/gin-gonic/gin"
)

// AltaRol godoc
// @Summary Crea un nuevo registro de rol
// @Description Agrega un nuevo rol
// @Tags rol
// @Accept  json
// @Produce  json
// @Param equipo body models.RolParam true "Estructura del rol a agregar"
// @Success 201 {string} string
// @Failure 400
// @Failure 409 {string} string
// @Failure 417 {string} string
// @Failure 500
// @Router /rol [post]
func (e *Atenea) AltaRol() gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyPost := &models.RolParam{}
		err := c.BindJSON(bodyPost)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.AltaRol(bodyPost.Nombre)
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

// BajaRol godoc
// @Summary Elimina roles
// @Description Borra el rol correspondiente al valor del identificador (id) indicado.
// @Tags rol
// @Accept  json
// @Produce  json
// @Param id path string true "Identificador (id) del rol a eliminar"
// @Success 200 {string} string
// @Failure 400
// @Failure 409 {string} string
// @Failure 417 {string} string
// @Failure 500
// @Router /rol/id/{id} [delete]
func (e *Atenea) BajaRol() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_usr, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.BajaRol(int32(id_usr))
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

// ListaRoles godoc
// @Summary Obtiene la lista de roles
// @Description Obtiene la lista de roles
// @Tags rol
// @Produce  json
// @Success 200 {array} models.Rol
// @Failure 500
// @Router /rol [get]
func (e *Atenea) ListaRoles() gin.HandlerFunc {
	return func(c *gin.Context) {

		single, err := e.db.ListaRoles()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, single)
	}
}

// ListaRolesById godoc
// @Summary Obtiene roles
// @Description Obtiene el rol correspondiente al identificador (id) indicado
// @Tags rol
// @Produce  json
// @Param id path string true "Identificador (id) del rol"
// @Success 200 {array} models.Rol
// @Failure 400
// @Failure 500
// @Router /rol/{id} [get]
func (e *Atenea) ListaRolesById() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_eq, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ListaRolesById(int32(id_eq))
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, single)
	}
}
