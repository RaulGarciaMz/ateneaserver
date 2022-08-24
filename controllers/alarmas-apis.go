package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// ListaAlarmaCriticalEquipo godoc
// @Summary Obtiene la lista de alarmas criticas de equipos
// @Description Obtiene la lista de alarmas criticas correspondiente al identificador del equipo indicado
// @Tags alarma
// @Produce  json
// @Param id path string true "Identificador (id) del equipo"
// @Success 200 {array} models.ListaAlarma
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /alarma/critica/equipo/{id} [get]
func (e *Atenea) ListaAlarmaCriticalEquipo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_eq, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ListaAlarmaCriticalEquipo(int32(id_eq))
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

// ListaAlarmaCriticalGrupo godoc
// @Summary Obtiene la lista de alarmas criticas por grupos
// @Description Obtiene la lista de alarmas criticas correspondiente al identificador del grupo indicado
// @Tags alarma
// @Produce  json
// @Param id path string true "Identificador (id) del grupo"
// @Success 200 {array} models.ListaAlarma
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /alarma/critica/grupo/{id} [get]
func (e *Atenea) ListaAlarmaCriticalGrupo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_eq, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ListaAlarmaCriticalGrupo(int32(id_eq))
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

// ListaAlarmaGrupo godoc
// @Summary Obtiene la lista de alarmas por grupos
// @Description Obtiene la lista de alarmas correspondiente al identificador del grupo indicado
// @Tags alarma
// @Produce  json
// @Param id path string true "Identificador (id) del grupo"
// @Success 200 {array} models.ListaAlarma
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /alarma/grupo/{id} [get]
func (e *Atenea) ListaAlarmaGrupo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_eq, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ListaAlarmaGrupo(int32(id_eq))
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

// ListaAlarmaMajorEquipo godoc
// @Summary Obtiene la lista de alarmas major por equipos
// @Description Obtiene la lista de alarmas major correspondiente al identificador del equipo indicado
// @Tags alarma
// @Produce  json
// @Param id path string true "Identificador (id) del equipo"
// @Success 200 {array} models.ListaAlarma
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /alarma/major/equipo/{id} [get]
func (e *Atenea) ListaAlarmaMajorEquipo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_eq, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ListaAlarmaMajorEquipo(int32(id_eq))
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

// ListaAlarmaMajorGrupo godoc
// @Summary Obtiene la lista de alarmas major por grupos
// @Description Obtiene la lista de alarmas major correspondiente al identificador del grupo indicado
// @Tags alarma
// @Produce  json
// @Param id path string true "Identificador (id) del grupo"
// @Success 200 {array} models.ListaAlarma
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /alarma/major/grupo/{id} [get]
func (e *Atenea) ListaAlarmaMajorGrupo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_eq, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ListaAlarmaMajorGrupo(int32(id_eq))
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

// ListaAlarmaNotReachableEquipo godoc
// @Summary Obtiene la lista de alarmas not-reachable por equipos
// @Description Obtiene la lista de alarmas not-reachable correspondiente al identificador del equipo indicado
// @Tags alarma
// @Produce  json
// @Param id path string true "Identificador (id) del equipo"
// @Success 200 {array} models.ListaAlarma
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /alarma/not-reachable/equipo/{id} [get]
func (e *Atenea) ListaAlarmaNotReachableEquipo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_eq, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ListaAlarmaNotReachableEquipo(int32(id_eq))
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

// ListaAlarmaNotReachableGrupo godoc
// @Summary Obtiene la lista de alarmas not-reachable por grupos
// @Description Obtiene la lista de alarmas not-reachable correspondiente al identificador del grupo indicado
// @Tags alarma
// @Produce  json
// @Param id path string true "Identificador (id) del grupo"
// @Success 200 {array} models.ListaAlarma
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /alarma/not-reachable/grupo/{id} [get]
func (e *Atenea) ListaAlarmaNotReachableGrupo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_eq, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ListaAlarmaNotReachableGrupo(int32(id_eq))
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

// ListaAlarmaNotifyEquipo godoc
// @Summary Obtiene la lista de alarmas notify por equipos
// @Description Obtiene la lista de alarmas notify correspondiente al identificador del equipo indicado
// @Tags alarma
// @Produce  json
// @Param id path string true "Identificador (id) del equipo"
// @Success 200 {array} models.ListaAlarma
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /alarma/notify/equipo/{id} [get]
func (e *Atenea) ListaAlarmaNotifyEquipo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_eq, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ListaAlarmaNotifyEquipo(int32(id_eq))
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

// ListaAlarmaNotifyGrupo godoc
// @Summary Obtiene la lista de alarmas notify por grupos
// @Description Obtiene la lista de alarmas notify correspondiente al identificador del grupo indicado
// @Tags alarma
// @Produce  json
// @Param id path string true "Identificador (id) del grupo"
// @Success 200 {array} models.ListaAlarma
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /alarma/notify/grupo/{id} [get]
func (e *Atenea) ListaAlarmaNotifyGrupo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_eq, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ListaAlarmaNotifyGrupo(int32(id_eq))
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

// ListaAlarmaWarningEquipo godoc
// @Summary Obtiene la lista de alarmas warning por equipos
// @Description Obtiene la lista de alarmas warning correspondiente al identificador del equipo indicado
// @Tags alarma
// @Produce  json
// @Param id path string true "Identificador (id) del equipo"
// @Success 200 {array} models.ListaAlarma
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /alarma/warning/equipo/{id} [get]
func (e *Atenea) ListaAlarmaWarningEquipo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_eq, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ListaAlarmaWarningEquipo(int32(id_eq))
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

// ListaAlarmaWarningGrupo godoc
// @Summary Obtiene la lista de alarmas warning por grupos
// @Description Obtiene la lista de alarmas warning correspondiente al identificador del grupo indicado
// @Tags alarma
// @Produce  json
// @Param id path string true "Identificador (id) del grupo"
// @Success 200 {array} models.ListaAlarma
// @Failure 400
// @Failure 404
// @Failure 500
// @Router /alarma/warning/grupo/{id} [get]
func (e *Atenea) ListaAlarmaWarningGrupo() gin.HandlerFunc {
	return func(c *gin.Context) {

		id := c.Param("id")

		id_eq, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		single, err := e.db.ListaAlarmaWarningGrupo(int32(id_eq))
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
