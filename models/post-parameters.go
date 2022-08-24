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
	Id          int32  `json:"id"`
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
