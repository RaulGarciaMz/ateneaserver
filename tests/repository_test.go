package tests

import (
	"fmt"
	"testing"

	"github.com/RaulGarciaMz/ateneaserver/database"
	"github.com/RaulGarciaMz/ateneaserver/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// equipos
func TestRepo_AltaEquipo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.AltaEquipo("ultimoRGM34", "1.0.0.0", "prueba de alta de equipo")
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ActualizaEquipo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ActualizaEquipo(1, "pato", "10.4.5.6", "actualizaciónde equipo")
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_BajaEquipo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.BajaEquipo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

// usuarios
func TestRepo_AltaUsuario(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.AltaUsuario(1, "nuevo", "bbb", "el ultimo", "nvo@gmail.com", "7777777777")
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ActualizaUsuario(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ActualizaUsuario(1, 1, "pass", "pato@gmail.com", "7777777777")
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_BajaUsuario(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.BajaUsuario(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ValidaUsuario(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ValidaUsuario("vcb", "bbb")
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ActualizaPassword(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ActualizaPassword("vcb", "bbb")
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

// Grupos
func TestRepo_AltaGrupo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.AltaGrupo("nuevoGrupo")
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ActualizaGrupo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ActualizaGrupo(1, "vcb")
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_BajaGrupo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.BajaGrupo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

// Roles
func TestRepo_AltaRol(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.AltaRol("rolvcb")
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_BajaRol(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.BajaRol(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

// Grupos-equipos
func TestRepo_IntegraGrupoEquipo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.IntegraGrupoEquipo(1, 1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_EliminaGrupoEquipo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.EliminaGrupoEquipo(1, 1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

// alarmas
func TestRepo_ListaAlarmaCriticalEquipo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ListaAlarmaCriticalEquipo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ListaAlarmaCriticalGrupo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ListaAlarmaCriticalGrupo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ListaAlarmaGrupo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ListaAlarmaGrupo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ListaAlarmaMajorEquipo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ListaAlarmaMajorEquipo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ListaAlarmaMajorGrupo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ListaAlarmaMajorGrupo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ListaAlarmaNotReachableEquipo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ListaAlarmaNotReachableEquipo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ListaAlarmaNotReachableGrupo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ListaAlarmaNotReachableGrupo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ListaAlarmaNotifyEquipo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ListaAlarmaNotifyEquipo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ListaAlarmaNotifyGrupo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ListaAlarmaNotifyGrupo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ListaAlarmaWarningEquipo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ListaAlarmaWarningEquipo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ListaAlarmaWarningGrupo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ListaAlarmaWarningGrupo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

// monitoreo
func TestRepo_MonitoreoEquipos(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.MonitoreoEquipos()
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_MonitoreoEquipo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.MonitoreoEquipo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_MonitoreoGrupos(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.MonitoreoGrupos()
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_MonitoreoGrupo(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.MonitoreoGrupo(1)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_ProcesaAlarmasLista(t *testing.T) {

	var err error

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	v := int64(0)

	alarmas := models.ListaAlarmasParam{
		Id:               1,
		MsgIds:           []int64{1234567, 7654321},
		MsgSlots:         []int32{7, 22},
		MsgPorts:         []*int64{&v, nil},
		MsgTexts:         []string{"No Connection PRUEBA1", "Ouput bitrate PRUEBA2"},
		MsgSourcesNames:  []string{"Stream", "Demodulation"},
		MsgSeveryties:    []string{"CRITICAL", "MAJOR"},
		MsgInstances:     []int32{3, 0},
		MsgSetTimes:      []string{"2022-08-19 14:11:02+00", "2022-08-19 14:11:15"},
		MsgCardIds:       []int32{73, 17},
		TimestampInicios: []string{"2022-08-21 21:26:21+00", "2022-08-21 21:26:21+00"},
		DateServer:       "2022-08-19",
		TimeServer:       "19:17:22",
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ProcesaAlarmasLista(&alarmas)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

// utilerías para testing
func createConnection() (*gorm.DB, error) {
	uri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", "184.72.110.87", "admin", "atenea", "admin")

	db, err := gorm.Open(postgres.Open(uri), &gorm.Config{})
	return db, err
}
