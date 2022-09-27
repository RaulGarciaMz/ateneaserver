package routes

import (
	"github.com/RaulGarciaMz/ateneaserver/controllers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AteneaRoute struct {
}

func (p AteneaRoute) CreaRuta(router *gin.RouterGroup, dbConn *gorm.DB, relativePath string, withOptionsHandler gin.HandlerFunc) {

	controller := controllers.AteneaController(dbConn)

	grupo := router.Group(relativePath, withOptionsHandler)
	{
		//equipos
		grupo.GET("/equipo", controller.ListaEquipos())
		grupo.POST("/equipo", controller.AltaEquipo())
		grupo.POST("/equipo/filtro", controller.FiltroEquipo())
		grupo.POST("/equipo/alcanzable", controller.EquipoAlcanzable())
		grupo.PUT("/equipo", controller.ActualizaEquipo())
		grupo.DELETE("/equipo/id/:id", controller.BajaEquipo())

		//usuarios
		grupo.GET("/usuario", controller.ListaUsuarios())
		grupo.POST("/usuario", controller.AltaUsuario())
		grupo.PUT("/usuario", controller.ActualizaUsuario())
		grupo.DELETE("/usuario/id/:id", controller.BajaUsuario())
		grupo.POST("/usuario/valida", controller.ValidaUsuario())
		grupo.PUT("/usuario/actpass", controller.ActualizaPassword())

		//Grupos
		grupo.GET("/grupo", controller.ListaGrupos())
		grupo.GET("/grupo/equipo/:id", controller.ListaEquipoGrupos())
		grupo.POST("/grupo", controller.AltaGrupo())
		grupo.PUT("/grupo", controller.ActualizaGrupo())
		grupo.DELETE("/grupo/id/:id", controller.BajaGrupo())

		//Roles
		grupo.GET("/rol", controller.ListaRoles())
		grupo.GET("/rol/:id", controller.ListaRolesById())
		grupo.POST("/rol", controller.AltaRol())
		grupo.DELETE("/rol/id/:id", controller.BajaRol())

		//Grupos-equipos
		grupo.GET("/grupo-equipo", controller.ListaGruposEquipos())
		grupo.POST("/grupo-equipo", controller.IntegraGrupoEquipo())
		grupo.POST("/grupo-equipo/lista", controller.IntegraGrupoEquipoLista())
		grupo.DELETE("/grupo-equipo/:id_grupo/:id_equipo", controller.EliminaGrupoEquipo())

		//alarmas
		grupo.GET("/alarma/critica/equipo/:id", controller.ListaAlarmaCriticalEquipo())
		grupo.GET("/alarma/critica/grupo/:id", controller.ListaAlarmaCriticalGrupo())
		grupo.GET("/alarma/grupo/:id", controller.ListaAlarmaGrupo())
		grupo.GET("/alarma/major/equipo/:id", controller.ListaAlarmaMajorEquipo())
		grupo.GET("/alarma/major/grupo/:id", controller.ListaAlarmaMajorGrupo())
		grupo.GET("/alarma/not-reachable/equipo/:id", controller.ListaAlarmaNotReachableEquipo())
		grupo.GET("/alarma/not-reachable/grupo/:id", controller.ListaAlarmaNotReachableGrupo())
		grupo.GET("/alarma/notify/equipo/:id", controller.ListaAlarmaNotifyEquipo())
		grupo.GET("/alarma/notify/grupo/:id", controller.ListaAlarmaNotifyGrupo())
		grupo.GET("/alarma/warning/equipo/:id", controller.ListaAlarmaWarningEquipo())
		grupo.GET("/alarma/warning/grupo/:id", controller.ListaAlarmaWarningGrupo())
		grupo.POST("/alarma/procesa", controller.ProcesaAlarmasLista())

		grupo.GET("/alarma/total", controller.TotalAlarmas())
		grupo.GET("/alarma/equipo/total", controller.TotalEquiposAlarmas())
		grupo.GET("/alarma/resumen", controller.GraficoPrincipal())
		grupo.GET("/alarma/not-reachable", controller.TotalNoAlcanzable())
		grupo.GET("/alarma/top/equipo", controller.TopEquiposAlarmados())
		grupo.GET("/alarma/top/ocurrencia", controller.TopOcurrenciaTipoAlarmas())
		grupo.GET("/alarma/top/not-reachable", controller.TopNoAlcanzable())

		//Monitoreo
		grupo.GET("/monitoreo/equipo/:id", controller.MonitoreoEquipo())
		grupo.GET("/monitoreo/equipos", controller.MonitoreoEquipos())
		grupo.GET("/monitoreo/grupos", controller.MonitoreoGrupos())
		grupo.GET("/monitoreo/grupo/:id", controller.MonitoreoGrupo())

		grupo.OPTIONS("/*id", withOptionsHandler)
	}
}
