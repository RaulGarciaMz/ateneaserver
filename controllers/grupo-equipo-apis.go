package controllers

import (
	"net/http"
	"strconv"

	"github.com/RaulGarciaMz/ateneaserver/models"
	"github.com/gin-gonic/gin"
)

// EliminaGrupoEquipo godoc
// @Summary Elimina combinaciones de grupos - equipos
// @Description Borra la combinación grupo - equipo correspondiente a los valores de los identificadores (id-equipo, id_grupo) indicado.
// @Tags grupo-equipo
// @Accept  json
// @Produce  json
// @Param id_grupo path string true "Identificador del grupo"
// @Param id_equipo path string true "Identificador del equipo"
// @Success 200 {string} string
// @Failure 400
// @Failure 409 {string} string
// @Failure 417 {string} string
// @Failure 500
// @Router /grupo-equipo/{id_grupo}/{id_equipo} [delete]
func (e *Atenea) EliminaGrupoEquipo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id_g := c.Param("id_grupo")
		id_e := c.Param("id_equipo")

		id_gpo, err := strconv.Atoi(id_g)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		id_equ, err := strconv.Atoi(id_e)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.EliminaGrupoEquipo(int32(id_gpo), int32(id_equ))
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

// ListaGruposEquipos godoc
// @Summary Obtiene la lista de grupos-equipos
// @Description Obtiene la lista de grupos-equipos
// @Tags grupo-equipo
// @Produce  json
// @Success 200 {array} models.GrupoEquipo
// @Failure 500
// @Router /grupo-equipo [get]
func (e *Atenea) ListaGruposEquipos() gin.HandlerFunc {
	return func(c *gin.Context) {

		single, err := e.db.ListaGruposEquipos()
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, single)
	}
}

// IntegraGrupoEquipo godoc
// @Summary Crea un nuevo registro de grupo-equipo
// @Description Agrega un nuevo grupo-equipo
// @Tags grupo-equipo
// @Accept  json
// @Produce  json
// @Param equipo body models.GrupoEquipoParam true "Estructura del grupo-equipo a agregar"
// @Success 201 {string} string
// @Failure 400
// @Failure 409 {string} string
// @Failure 417 {string} string
// @Failure 500
// @Router /grupo-equipo [post]
func (e *Atenea) IntegraGrupoEquipo() gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyPost := &models.GrupoEquipoParam{}
		err := c.BindJSON(bodyPost)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.IntegraGrupoEquipo(bodyPost.IdGrupo, bodyPost.IdEquipo)
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

// IntegraGrupoEquipoLista godoc
// @Summary Crea un nuevo registro de grupo-equipo
// @Description Agrega un nuevo grupo-equipo
// @Tags grupo-equipo
// @Accept  json
// @Produce  json
// @Param equipo body models.GrupoEquipoIntegraParam true "Estructura del grupo-equipo a agregar"
// @Success 201 {string} string
// @Failure 400
// @Failure 409 {string} string
// @Failure 417 {string} string
// @Failure 500
// @Router /grupo-equipo/lista [post]
func (e *Atenea) IntegraGrupoEquipoLista() gin.HandlerFunc {
	return func(c *gin.Context) {

		bodyPost := &models.GrupoEquipoIntegraParam{}
		err := c.BindJSON(bodyPost)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.IntegraGrupoEquipoLista(bodyPost.IdGrupo, bodyPost.Equipos)
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
