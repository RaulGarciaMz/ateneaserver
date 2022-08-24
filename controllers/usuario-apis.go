package controllers

import (
	"net/http"
	"strconv"

	"github.com/RaulGarciaMz/ateneaserver/models"
	"github.com/gin-gonic/gin"
)

// ValidaUsuario godoc
// @Summary Valida credenciales de usuarios
// @Description Valida credenciales del usuario indicado en el body de la petición
// @Tags usuario
// @Accept  json
// @Produce  json
// @Param equipo body models.UsuarioAuthParam true "Estructura del usuario a agregar"
// @Success 200 {object} models.Usuario
// @Failure 400
// @Failure 401 {string} string
// @Failure 500
// @Router /usuario/valida [post]
func (e *Atenea) ValidaUsuario() gin.HandlerFunc {
	return func(c *gin.Context) {
		bodyPost := &models.UsuarioAuthParam{}
		err := c.BindJSON(bodyPost)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		user, err := e.db.ValidaUsuario(bodyPost.Usuario, bodyPost.Passwd)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		if user.Usuario == nil {
			c.JSON(http.StatusUnauthorized, err)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

// ActualizaPassword godoc
// @Summary Actualiza el password de usuarios
// @Description Actualiza el password del usuario indicado en el body de la petición
// @Tags usuario
// @Accept  json
// @Produce  json
// @Param equipo body models.UsuarioAuthParam true "Estructura del usuario y pwd a actualizar"
// @Success 200 {string} string
// @Failure 400
// @Failure 417 {string} string
// @Failure 500
// @Router /usuario/actpass [put]
func (e *Atenea) ActualizaPassword() gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyPost := &models.UsuarioAuthParam{}
		err := c.BindJSON(bodyPost)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ActualizaPassword(bodyPost.Usuario, bodyPost.Passwd)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		switch single {
		case "Éxito":
			c.JSON(http.StatusCreated, single)
		default:
			c.JSON(http.StatusExpectationFailed, single)
		}
	}
}

// ActualizaUsuario godoc
// @Summary Actualiza datos de usuarios
// @Description Actualiza el registro del usuario indicado en el body de la petición
// @Tags usuario
// @Accept  json
// @Produce  json
// @Param equipo body models.UsuarioParam true "Estructura del usuario a actualizar"
// @Success 200 {string} string
// @Failure 400
// @Failure 417 {string} string
// @Failure 500
// @Router /usuario [put]
func (e *Atenea) ActualizaUsuario() gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyPost := &models.UsuarioParam{}
		err := c.BindJSON(bodyPost)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ActualizaUsuario(bodyPost.Id, bodyPost.IdRol, bodyPost.Nombre, bodyPost.EMail, bodyPost.Cel)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		switch single {
		case "Éxito":
			c.JSON(http.StatusCreated, single)
		default:
			c.JSON(http.StatusExpectationFailed, single)
		}
	}
}

// AltaUsuario godoc
// @Summary Crea un nuevo registro de usuario
// @Description Agrega un nuevo usuario
// @Tags usuario
// @Accept  json
// @Produce  json
// @Param equipo body models.UsuarioParam true "Estructura del usuario a agregar"
// @Success 201 {string} string
// @Failure 400
// @Failure 409 {string} string
// @Failure 417 {string} string
// @Failure 500
// @Router /usuario [post]
func (e *Atenea) AltaUsuario() gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyPost := &models.UsuarioParam{}
		err := c.BindJSON(bodyPost)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.AltaUsuario(bodyPost.IdRol, bodyPost.Usuario, bodyPost.Passwd, bodyPost.Nombre, bodyPost.EMail, bodyPost.Cel)
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

// BajaUsuario godoc
// @Summary Elimina usuarios
// @Description Borra el usuario correspondiente al valor del identificador (id) indicado.
// @Tags usuario
// @Accept  json
// @Produce  json
// @Param id path string true "Identificador (id) del usuario a eliminar"
// @Success 200 {string} string
// @Failure 400
// @Failure 409 {string} string
// @Failure 417 {string} string
// @Failure 500
// @Router /usuario/id/{id} [delete]
func (e *Atenea) BajaUsuario() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_usr, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.BajaUsuario(int32(id_usr))
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

// ListaUsuarios godoc
// @Summary Obtiene la lista de usuarios
// @Description Obtiene la lista de usuarios
// @Tags usuarios
// @Produce  json
// @Success 200 {array} models.Usuario
// @Failure 500
// @Router /usuario [get]
func (e *Atenea) ListaUsuarios() gin.HandlerFunc {
	return func(c *gin.Context) {

		single, err := e.db.ListaUsuarios()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, single)
	}
}
