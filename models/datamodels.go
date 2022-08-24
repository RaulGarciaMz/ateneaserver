package models

import "time"

type FunctionResp struct {
	Resultado *string `gorm:"column:resultado" json:"resultado"`
}

type MonitoreoGrupo struct {
	Grupo        *string `gorm:"column:grupo" json:"grupo"`
	IdGrupo      *int32  `gorm:"column:id_grupo" json:"id_grupo"`
	TotEquipos   *int32  `gorm:"column:tot_equipos" json:"tot_equipos"`
	Critical     *int    `gorm:"column:critical" json:"critical"`
	Major        *int    `gorm:"column:major" json:"major"`
	Warning      *int    `gorm:"column:warning_" json:"warning"`
	Notify       *int    `gorm:"column:notify_" json:"notify"`
	NotReachable *int    `gorm:"column:not_reachable" json:"not_reachable"`
}

type Usuario struct {
	Usuario *string `gorm:"column:usuario" json:"usuario"`
	Pass    *string `gorm:"column:pass" json:"pass"`
	Nombre  *string `gorm:"column:nombre" json:"nombre"`
	EMail   *string `gorm:"column:e_mail" json:"email"`
	Cel     *string `gorm:"column:cel" json:"cel"`
	Id      *int32  `gorm:"column:id" json:"id"`
	IdRol   *int32  `gorm:"column:id_rol" json:"id_rol"`
}

type ListaAlarma struct {
	Equipo        string    `json:"equipo" gorm:"column:equipo"`
	IdEquipo      int32     `json:"id_equipo" gorm:"column:id_equipo"`
	MsgId         int32     `json:"msg_id" gorm:"column:msg_id"`
	MsgSlot       int       `json:"msg_slot" gorm:"column:msg_slot"`
	MsgPort       int       `json:"msg_port" gorm:"column:msg_port"`
	MsgText       string    `json:"msg_text" gorm:"column:msg_text"`
	MsgSourceName string    `json:"msg_source_name" gorm:"column:msg_source_name"`
	MsgSeverity   string    `json:"msg_severity" gorm:"column:msg_severity"`
	MsgInstance   int       `json:"msg_instance" gorm:"column:msg_instance"`
	MsgSetTime    time.Time `json:"msg_set_time" gorm:"column:msg_set_time"`
	MsgCardId     int32     `json:"msg_card_id" gorm:"column:msg_card_id"`
}

type MonitoreoEquipo struct {
	Equipo       string `json:"equipo" gorm:"column:equipo"`
	Ip           string `json:"ip" gorm:"column:ip"`
	IdEquipo     int32  `json:"id_equipo" gorm:"column:id_equipo"`
	Critical     int    `json:"critical" gorm:"column:critical"`
	Major        int    `json:"major" gorm:"column:major"`
	Warning      int    `json:"warning" gorm:"column:warning_"`
	Notify       int    `json:"notify" gorm:"column:notify_"`
	NotReachable int    `json:"not_reachable" gorm:"column:not_reachable"`
}

type Rol struct {
	Id     *int32  `gorm:"column:id" json:"id"`
	Nombre *string `gorm:"column:nombre" json:"nombre"`
}

type Grupo struct {
	Id     *int32  `gorm:"column:id" json:"id"`
	Nombre *string `gorm:"column:nombre" json:"nombre"`
}

type Equipo struct {
	Id          *int32  `gorm:"column:id" json:"id"`
	Nombre      *string `gorm:"column:nombre" json:"nombre"`
	Ip          *string `gorm:"column:ip" json:"ip"`
	Descripcion *string `gorm:"column:descripcion" json:"descripcion"`
}

type GrupoEquipo struct {
	IdGrupo  *int32 `gorm:"column:id_grupo" json:"id_grupo"`
	IdEquipo *int32 `gorm:"column:id_equipo" json:"id_equipo"`
}

type ListaAlarmasParam struct {
	Id               int32
	MsgIds           []int64
	MsgSlots         []int32
	MsgPorts         []*int64
	MsgTexts         []string
	MsgSourcesNames  []string
	MsgSeveryties    []string
	MsgInstances     []int32
	MsgSetTimes      []string
	MsgCardIds       []int32
	TimestampInicios []string
	DateServer       string
	TimeServer       string
}
