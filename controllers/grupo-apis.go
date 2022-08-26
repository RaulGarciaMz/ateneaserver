package controllers

import (
	"net/http"
	"strconv"

	"github.com/RaulGarciaMz/ateneaserver/models"
	"github.com/gin-gonic/gin"
)

// ActualizaGrupo godoc
// @Summary Actualiza datos de un grupo
// @Description Actualiza el registro del grupo indicado en el body de la petición
// @Tags grupo
// @Accept  json
// @Produce  json
// @Param equipo body models.GrupoParam true "Estructura del grupo a actualizar"
// @Success 200 {string} string
// @Failure 400
// @Failure 409 {string} string
// @Failure 417 {string} string
// @Failure 500
// @Router /grupo [put]
func (e *Atenea) ActualizaGrupo() gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyPost := &models.GrupoParam{}
		err := c.BindJSON(bodyPost)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ActualizaGrupo(bodyPost.Id, bodyPost.Nombre)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		switch single {
		case "Éxito":
			c.JSON(http.StatusCreated, single)
		case "Nombre existente":
			c.JSON(http.StatusConflict, single)
		default:
			c.JSON(http.StatusExpectationFailed, single)
		}
	}
}

// AltaGrupo godoc
// @Summary Crea un nuevo registro de grupo
// @Description Agrega un nuevo grupo
// @Tags grupo
// @Accept  json
// @Produce  json
// @Param equipo body models.GrupoParam true "Estructura del grupo a agregar"
// @Success 201 {string} string
// @Failure 400
// @Failure 409 {string} string
// @Failure 417 {string} string
// @Failure 500
// @Router /grupo [post]
func (e *Atenea) AltaGrupo() gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyPost := &models.GrupoParam{}
		err := c.BindJSON(bodyPost)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.AltaGrupo(bodyPost.Nombre)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		switch single {
		/* 		case "Éxito":
		c.JSON(http.StatusCreated, single) */
		case "Existente":
			c.JSON(http.StatusConflict, single)
		default:
			c.JSON(http.StatusCreated, single)
		}
	}
}

// BajaGrupo godoc
// @Summary Elimina grupos
// @Description Borra el grupo correspondiente al valor del identificador (id) indicado.
// @Tags grupo
// @Accept  json
// @Produce  json
// @Param id path string true "Identificador (id) del grupo a eliminar"
// @Success 200 {string} string
// @Failure 400
// @Failure 409 {string} string
// @Failure 417 {string} string
// @Failure 500
// @Router /grupo/id/{id} [delete]
func (e *Atenea) BajaGrupo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_usr, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.BajaGrupo(int32(id_usr))
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

// ListaGrupos godoc
// @Summary Obtiene la lista de grupos
// @Description Obtiene la lista de grupos
// @Tags grupo
// @Produce  json
// @Success 200 {array} models.Grupo
// @Failure 500
// @Router /grupo [get]
func (e *Atenea) ListaGrupos() gin.HandlerFunc {
	return func(c *gin.Context) {

		single, err := e.db.ListaGrupos()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, single)
	}
}

// ListaEquipoGrupos godoc
// @Summary Obtiene la lista de grupos por equipo
// @Description Obtiene la lista de grupos correspondiente al valor del identificador del equipo (id) indicado en el parámetro
// @Tags grupo
// @Produce  json
// @Param id path string true "Identificador (id) del grupo"
// @Success 200 {array} models.Grupo
// @Failure 500
// @Router /grupo/equipo/{id} [get]
func (e *Atenea) ListaEquipoGrupos() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_usr, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ListaEquipoGrupos(id_usr)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, single)
	}
}
