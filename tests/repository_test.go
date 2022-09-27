package tests

import (
	"fmt"
	"testing"
	"time"

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

	b, err := repo.AltaEquipo("ultimoRGM34", "1.0.0.0", "prueba de alta de equipo", 20)
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

	b, err := repo.ActualizaEquipo(1, "pato", "10.4.5.6", "actualizaciónde equipo", 20)
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

func TestRepo_EquipoAlcanzable(t *testing.T) {

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.EquipoAlcanzable(28, true, time.Now())
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

	b, err := repo.ListaAlarmaGrupo(21)
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

	v0 := int64(0)
	v3 := int64(3)
	v16 := int64(16)
	v1005 := int64(1005)

	alarmas := models.ListaAlarmasParam{
		Id:               11,
		MsgIds:           []int64{4390932, 1048669, 3473417, 1048589, 1048589, 1048580, 3473418, 1048611, 4259950, 3539048, 3539045, 2883585, 1048580},
		MsgSlots:         []int32{1, 5, 8, 14, 0, 0, 8, 9, 16, 16, 16, 1, 5},
		MsgPorts:         []*int64{&v3, nil, nil, nil, nil, &v1005, nil, nil, &v0, &v0, &v16, &v0, &v0},
		MsgTexts:         []string{"No Lock", "Output bitrate exceeded on 230.2.3.6:6000", "Bulkdescrambler not authorised. All services blocked", "No signal", "No signal on port A", "239.255.1.4:1234: No bitrate on: Service 13005", "Latens controller connectivity failure", "Link down input", "Control port link down", "No Connection from IP 10.101.1.1, port 6002", "Control port link down", "PMT Missing: Service 13077: Service 13077 (From Hotbird Vert. Low)", "239.200.1.1:1234: No bitrate on: Service 13005"},
		MsgSourcesNames:  []string{"Demodulation", "Stream", "bulkdscr-latens", "Stream", "Stream", "Stream", "bulkdscr-latens", "Network", "Network", "EMMG/PDG", "Network", "PSI", "Stream"},
		MsgSeveryties:    []string{"CRITICAL", "CRITICAL", "CRITICAL", "CRITICAL", "CRITICAL", "CRITICAL", "MAJOR", "CRITICAL", "CRITICAL", "CRITICAL", "CRITICAL", "CRITICAL", "CRITICAL"},
		MsgInstances:     []int32{0, 11, 0, 0, 0, 0, 0, 0, 0, 3, 0, 13077, 0},
		MsgSetTimes:      []string{"2021-09-16 16:30:47", "2022-08-22 11:40:02", "2021-02-10 14:21:30", "2021-06-08 16:33:24", "2022-04-03 06:33:42", "2022-08-21 09:25:15", "2021-02-10 14:21:30", "2022-02-17 17:24:33", "2022-07-08 20:45:30", "2022-07-08 20:48:35", "2022-07-08 20:48:05", "2022-08-21 09:25:14", "2022-08-21 09:25:14"},
		MsgCardIds:       []int32{16, 13, 51, 14, 32, 37, 51, 90, 11, 11, 11, 16, 13},
		TimestampInicios: []string{"2020-10-01 20:56:52", "2022-08-12 11:15:22", "2019-07-23 16:38:18", "2020-03-15 21:02:06", "2021-11-03 01:02:42", "2022-08-10 06:45:48", "2019-07-23 16:38:18", "2021-08-05 22:44:24", "2022-05-15 05:26:18", "2022-05-15 05:32:28", "2022-05-15 05:31:28", "2022-08-10 06:45:46", "2022-08-10 06:45:46"},
		DateServer:       "2022-09-01",
		TimeServer:       "12:04:42",
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.ProcesaAlarmasLista(&alarmas)
	_ = b
	if err != nil {
		t.Error("No obtuvo datos por error")
	}

}

func TestRepo_TotalAlarmas(t *testing.T) {

	var err error

	cn, err := createConnection()
	if err != nil {
		t.Fatal("No se conectó")
	}

	repo := database.NewAteneaRepo(cn)

	b, err := repo.TotalAlarmas()
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
