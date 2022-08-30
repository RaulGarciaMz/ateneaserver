package database

import (
	"github.com/RaulGarciaMz/ateneaserver/models"
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type AteneaRepo struct {
	db *gorm.DB
}

func NewAteneaRepo(db *gorm.DB) *AteneaRepo {
	return &AteneaRepo{db: db}
}

func (e *AteneaRepo) AltaEquipo(nombre, ip, descripcion string) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.alta_equipo (?, ?, ?) as resultado", nombre, ip, descripcion).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) FiltroEquipo(id int32, filtro bool) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.filtro_equipo (?, ?) as resultado", id, filtro).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) ActualizaEquipo(id int32, nombre, ip, descripcion string) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.actualiza_equipo(?,?,?,?) as resultado", id, nombre, ip, descripcion).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) ValidaUsuario(usuario, password string) (*models.Usuario, error) {

	single := models.Usuario{}
	err := e.db.Raw("SELECT * from  admin.valida_usuario (?, ?)", usuario, password).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return &single, err
}

func (e *AteneaRepo) ActualizaGrupo(id int32, nombre string) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.actualiza_grupo(?,?) as resultado", id, nombre).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) ActualizaPassword(usuario, password string) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.actualiza_password(?,?) as resultado", usuario, password).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) ActualizaUsuario(id, id_rol int32, nombre, e_mail, cel string) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.actualiza_usuario(?,?,?,?,?) as resultado", id, id_rol, nombre, e_mail, cel).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) AltaGrupo(nombre string) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.alta_grupo(?) as resultado", nombre).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) AltaRol(nombre string) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.alta_rol(?) as resultado", nombre).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) AltaUsuario(id_rol int32, usuario, pass, nombre, e_mail, cel string) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.alta_usuario(?,?,?,?,?,?) as resultado", id_rol, usuario, pass, nombre, e_mail, cel).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) BajaEquipo(id int32) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.baja_equipo(?) as resultado", id).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) BajaGrupo(id int32) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.baja_grupo(?) as resultado", id).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) BajaRol(id int32) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.baja_rol(?) as resultado", id).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) BajaUsuario(id int32) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.baja_usuario(?) as resultado", id).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) EliminaGrupoEquipo(id_grupo, id_equipo int32) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.elimina_grupo_equipo(?,?) as resultado", id_grupo, id_equipo).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) IntegraGrupoEquipo(id_grupo, id_equipo int32) (string, error) {

	single := models.FunctionResp{}
	err := e.db.Raw("SELECT * FROM admin.integra_grupo_equipo(?,?) as resultado", id_grupo, id_equipo).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, nil
}

func (e *AteneaRepo) ListaAlarmaCriticalEquipo(id int32) ([]*models.ListaAlarma, error) {

	single := []*models.ListaAlarma{}
	err := e.db.Raw("SELECT * FROM admin.lista_alarmas_critical_equipo(?)", id).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaAlarmaCriticalGrupo(id int32) ([]*models.ListaAlarma, error) {

	single := []*models.ListaAlarma{}
	err := e.db.Raw("SELECT * FROM  admin.lista_alarmas_critical_grupo(?)", id).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaAlarmaGrupo(id int32) ([]*models.ListaAlarma, error) {

	single := []*models.ListaAlarma{}
	err := e.db.Raw("SELECT * FROM admin.lista_alarmas_grupo(?)", id).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaAlarmaMajorEquipo(id int32) ([]*models.ListaAlarma, error) {

	single := []*models.ListaAlarma{}
	err := e.db.Raw("SELECT * FROM admin.lista_alarmas_major_equipo(?)", id).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaAlarmaMajorGrupo(id int32) ([]*models.ListaAlarma, error) {

	single := []*models.ListaAlarma{}
	err := e.db.Raw("SELECT * FROM admin.lista_alarmas_major_grupo(?)", id).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaAlarmaNotReachableEquipo(id int32) ([]*models.ListaAlarma, error) {

	single := []*models.ListaAlarma{}
	err := e.db.Raw("SELECT * FROM admin.lista_alarmas_not_reachable_equipo(?)", id).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaAlarmaNotReachableGrupo(id int32) ([]*models.ListaAlarma, error) {

	single := []*models.ListaAlarma{}
	err := e.db.Raw("SELECT * FROM admin.lista_alarmas_not_reachable_grupo(?)", id).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaAlarmaNotifyEquipo(id int32) ([]*models.ListaAlarma, error) {

	single := []*models.ListaAlarma{}
	err := e.db.Raw("SELECT * FROM admin.lista_alarmas_notify_equipo(?)", id).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaAlarmaNotifyGrupo(id int32) ([]*models.ListaAlarma, error) {

	single := []*models.ListaAlarma{}
	err := e.db.Raw("SELECT * FROM admin.lista_alarmas_notify_grupo(?)", id).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaAlarmaWarningEquipo(id int32) ([]*models.ListaAlarma, error) {

	single := []*models.ListaAlarma{}
	err := e.db.Raw("SELECT * FROM admin.lista_alarmas_warning_equipo(?)", id).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaAlarmaWarningGrupo(id int32) ([]*models.ListaAlarma, error) {

	single := []*models.ListaAlarma{}
	err := e.db.Raw("SELECT * FROM admin.lista_alarmas_warning_grupo(?)", id).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) MonitoreoEquipo(id int32) ([]*models.MonitoreoEquipo, error) {

	single := []*models.MonitoreoEquipo{}
	err := e.db.Raw("SELECT * FROM admin.monitoreo_equipo(?)", id).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) MonitoreoEquipos() ([]*models.MonitoreoEquipo, error) {
	single := []*models.MonitoreoEquipo{}
	err := e.db.Raw("SELECT * FROM admin.monitoreo_equipos()").Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) MonitoreoGrupo(id int32) ([]*models.MonitoreoGrupo, error) {

	single := []*models.MonitoreoGrupo{}
	err := e.db.Raw("SELECT * FROM admin.monitoreo_grupo(?)", id).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) MonitoreoGrupos() ([]models.MonitoreoGrupo, error) {
	single := []models.MonitoreoGrupo{}
	err := e.db.Raw("SELECT * FROM admin.monitoreo_grupos()").Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaUsuarios() ([]models.Usuario, error) {
	single := []models.Usuario{}
	err := e.db.Raw("SELECT * FROM admin.lista_usuarios()").Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaRoles() ([]models.Rol, error) {
	single := []models.Rol{}
	err := e.db.Raw("SELECT * FROM admin.lista_roles()").Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaRolesById(id int32) ([]models.Rol, error) {
	single := []models.Rol{}
	err := e.db.Raw("SELECT * FROM admin.lista_roles_id(?)", id).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaGrupos() ([]models.Grupo, error) {
	single := []models.Grupo{}
	err := e.db.Raw("SELECT * FROM admin.lista_grupos()").Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaEquipos() ([]models.Equipo, error) {
	single := []models.Equipo{}
	err := e.db.Raw("SELECT * FROM admin.lista_equipos()").Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaGruposEquipos() ([]models.GrupoEquipo, error) {
	single := []models.GrupoEquipo{}
	err := e.db.Raw("SELECT * FROM admin.lista_grupos_equipos()").Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) IntegraGrupoEquipoLista(id int32, eq []int32) (string, error) {
	single := models.FunctionResp{}

	err := e.db.Raw("SELECT * FROM admin.integra_grupo_equipo_lista(?, ?) as resultado", id, pq.Array(eq)).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, err
}

func (e *AteneaRepo) ProcesaAlarmasLista(als *models.ListaAlarmasParam) (string, error) {
	single := models.FunctionResp{}

	err := e.db.Raw("SELECT * FROM admin.procesa_alarmas_lista(?, ?,?,?,?,?,?,?,?,?,?,?,?) as resultado",
		als.Id, pq.Array(als.MsgIds), pq.Array(als.MsgSlots), pq.Array(als.MsgPorts),
		pq.Array(als.MsgTexts), pq.Array(als.MsgSourcesNames), pq.Array(als.MsgSeveryties), pq.Array(als.MsgInstances),
		pq.Array(als.MsgSetTimes), pq.Array(als.MsgCardIds), pq.Array(als.TimestampInicios), als.DateServer,
		als.TimeServer).Scan(&single).Error
	if err != nil {
		return "", err
	}

	return *single.Resultado, err
}

func (e *AteneaRepo) TotalAlarmas() (*models.TotalAlarma, error) {
	single := models.TotalAlarma{}

	err := e.db.Raw("SELECT * FROM admin.total_alarmas()").Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return &single, err
}

func (e *AteneaRepo) TotalEquiposAlarmas() (*models.TotalAlarma, error) {
	single := models.TotalAlarma{}

	err := e.db.Raw("SELECT * FROM admin.total_equipos_alarmas()").Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return &single, err
}

func (e *AteneaRepo) GraficoPrincipal() ([]models.GraficoPrincipal, error) {
	single := []models.GraficoPrincipal{}

	err := e.db.Raw("SELECT * FROM admin.grafico_principal()").Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) TotalNoAlcanzable() (*int32, error) {
	single := models.FunctionRespInteger{}

	err := e.db.Raw("SELECT * FROM admin.total_no_alcanzable() as resultado").Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single.Resultado, err
}

func (e *AteneaRepo) TopEquiposAlarmados() ([]models.TopEquipoAlarmado, error) {
	single := []models.TopEquipoAlarmado{}

	err := e.db.Raw("SELECT * FROM admin.top_equipos_alarmados()").Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) TopOcurrenciaTipoAlarmas() ([]models.TopEquipoAlarmado, error) {
	single := []models.TopEquipoAlarmado{}

	err := e.db.Raw("SELECT * FROM admin.top_ocurrencia_tipo_alarmas()").Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) TopNoAlcanzable() ([]models.TopEquipoAlarmado, error) {
	single := []models.TopEquipoAlarmado{}

	err := e.db.Raw("SELECT * FROM admin.top_no_alcanzable()").Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}

func (e *AteneaRepo) ListaEquipoGrupos(equipo int) ([]models.Grupo, error) {
	single := []models.Grupo{}
	err := e.db.Raw("SELECT * FROM admin.lista_equipo_grupos(?)", equipo).Scan(&single).Error
	if err != nil {
		return nil, err
	}

	return single, err
}
