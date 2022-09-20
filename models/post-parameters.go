package models

type UsuarioParam struct {
	Usuario string `json:"usuario"`
	Nombre  string `json:"nombre"`
	Passwd  string `json:"password"`
	EMail   string `json:"email"`
	Cel     string `json:"cel"`
	Id      int32  `json:"id"`
	IdRol   int32  `json:"id_rol"`
}

type UsuarioAuthParam struct {
	Usuario string `json:"usuario"`
	Passwd  string `json:"password"`
}

type EquipoParam struct {
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Ip          string `json:"ip"`
	Puerto      int32  `json:"puerto"`
	Id          int32  `json:"id"`
}

type FiltroEquipoParam struct {
	Id     int32 `json:"id"`
	Filtro bool  `json:"filtro"`
}

type GrupoParam struct {
	Nombre string `json:"nombre"`
	Id     int32  `json:"id"`
}

type RolParam struct {
	Nombre string `json:"nombre"`
	Id     int32  `json:"id"`
}

type GrupoEquipoParam struct {
	IdGrupo  int32 `json:"id_grupo"`
	IdEquipo int32 `json:"id_equipo"`
}

type GrupoEquipoIntegraParam struct {
	IdGrupo int32   `json:"id_grupo"`
	Equipos []int32 `json:"equipos"`
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
