package models

import (
	"errors"
	"fmt"
	"time"

	"Base1.com/go-backend/configs"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Almacen struct {
	gorm.Model
	Nombre   string `gorm:"type:varchar(50)"`
	Cantidad int    `gorm:"null"`
}

type AccionesRH struct {
	gorm.Model
	Id               string    `gorm:"type:varchar(12);primaryKey;not null"`
	Fecha            time.Time `gorm:"null"`
	Empleado         string    `gorm:"type:varchar(10);not null"`
	TipoNovedad      string    `gorm:"type:varchar(3);not null"`
	FechaEfectividad time.Time `gorm:"null"`
	SueldoAnterior   float64   `gorm:"null"`
	Sueldo           float64   `gorm:"null"`
	Observaciones    string    `gorm:"type:text;not null"`
	Fecha_salida     time.Time `gorm:"null"`
	Fecha_entrada    time.Time `gorm:"null"`
	Dia              int       `gorm:"null"`
	Id_regreso       string    `gorm:"type:varchar(12);not null"`
}

type Usuario struct {
	gorm.Model
	Usuario string `gorm:"type:varchar(30) not null"`  //varchar(30)) NOT NULL,
	Clave   string `gorm:"type:varchar(200) not null"` //varchar(200)) NOT NULL,
	Dire    string `gorm:"type:varchar(40) null"`      //varchar(40)) NULL,,
	Tele    string `gorm:"type:varchar(20) null"`      //varchar(20)) NULL,,
	Correo  string `gorm:"type:varchar(80) null"`      //varchar(80)) NULL,,
}

// Antes de guardar un nuevo usuario o actualizar uno existente,
// puedes llamar a CreatePassword para encriptar la contraseña.

func (a *AccionesRH) BeforeCreate(tx *gorm.DB) (err error) {
	a.Id = uuid.New().String()
	return
}

type Transa01 struct {
	Empresa       int       `gorm:"null"` //int NULL,
	Fecha         time.Time `gorm:"null"`
	Tipo          string    `gorm:"type:varchar(2) null"`  //varchar(2)) NULL,,
	Numero        string    `gorm:"type:varchar(12) null"` //varchar(12)) NULL,,
	Documento     string    `gorm:"type:varchar(12) null"` //varchar(12)) NULL,,
	Documento1    string    `gorm:"type:varchar(12) null"` //varchar(12)) NULL,,
	Turno1        string    `gorm:"type:varchar(1) null"`  //varchar(1)) NULL,,
	Ncf           string    `gorm:"type:varchar(19) null"` //varchar(19)) NULL,,
	Cliente       string    `gorm:"type:varchar(15) null"` //varchar(15)) NULL,,
	Suplidor      string    `gorm:"type:varchar(15) null"` //varchar(15)) NULL,,
	Termino       string    `gorm:"type:varchar(3) null"`  //varchar(3)) NULL,,
	Dia           int       `gorm:"null"`                  //int NULL,
	Dia1          int       `gorm:"null"`                  //int NULL,
	Vence         time.Time `gorm:"null"`
	Digitado      time.Time `gorm:"null"`
	Itbis         float64   `gorm:"null"` //decimal(18, 2) NULL,
	Fechadesc     time.Time `gorm:"null"`
	Otros         float64   `gorm:"null"`                   //decimal(18, 2) NULL,
	Flete         float64   `gorm:"null"`                   //decimal(18, 2) NULL,
	Descuen       float64   `gorm:"null"`                   //decimal(18, 2) NULL,
	Nograva       float64   `gorm:"null"`                   //decimal(18, 2) NULL,
	Grava         float64   `gorm:"null"`                   //decimal(18, 2) NULL,
	Monto         float64   `gorm:"null"`                   //decimal(18, 2) NULL,
	Total         float64   `gorm:"null"`                   //decimal(18, 2) NULL,
	Balance       float64   `gorm:"null"`                   //decimal(18, 2) NULL,
	Nombre        string    `gorm:"type:varchar(50) null"`  //varchar(50)) NULL,,
	Vendedor      string    `gorm:"type:varchar(3) null"`   //varchar(3)) NULL,,
	Vendedor1     string    `gorm:"type:varchar(3) null"`   //varchar(3)) NULL,,
	Venname       string    `gorm:"type:varchar(50) null"`  //varchar(50)) NULL,,
	Zona          string    `gorm:"type:varchar(3) null"`   //varchar(3)) NULL,,
	Pedido        string    `gorm:"type:varchar(12) null"`  //varchar(12)) NULL,,
	Conduce       string    `gorm:"type:varchar(12) null"`  //varchar(12)) NULL,,
	Observa       string    `gorm:"type:varchar(254) null"` //varchar(254)) NULL,,
	Validado      string    `gorm:"type:varchar(1) null"`   //varchar(1)) NULL,,
	Posteado      string    `gorm:"type:varchar(1) null"`   //varchar(1)) NULL,,
	Efectivo      float64   `gorm:"null"`                   //decimal (18, 4)  NULL,
	Cheque        float64   `gorm:"null"`                   //decimal (18, 4)  NULL,
	Tarjeta       float64   `gorm:"null"`                   //decimal (18, 4)  NULL,
	Gastos        float64   `gorm:"null"`                   //numeric (18, 2) NULL,
	Nulo          string    `gorm:"type:varchar(1) null"`   //varchar(1)) NULL,,
	ABIERTO       string    `gorm:"type:varchar(1) null"`   //varchar(1)) NULL,,
	ENVIADO       float64   `gorm:"null"`                   //decimal(18, 2) NULL,
	Despacho      string    `gorm:"type:varchar(1) null"`   //varchar(1)) NULL,,
	Provincia     string    `gorm:"type:varchar(3) null"`   //varchar(3)) NULL,,
	Almacen       string    `gorm:"type:varchar(2) null"`   //varchar(2)) NULL,,
	POSTEO        float64   `gorm:"null"`                   //decimal (18, 0) null,
	Socio         string    `gorm:"type:varchar(1) null"`   //varchar(1)) NULL,,
	Creado        string    `gorm:"type:varchar(20) null"`  //varchar(20)) NULL,,
	Dsc           string    `gorm:"type:varchar(1) null"`   //varchar(1)) NULL,,
	Prnc          string    `gorm:"type:varchar(1) null"`   //varchar(1)) NULL,,
	Prn           string    `gorm:"type:varchar(1) null"`   //varchar(1)) NULL,,
	Observa1      string    `gorm:"type:varchar(254) null"` //varchar(254)) NULL,,
	Usuario       string    `gorm:"type:varchar(30) null"`  //varchar(30)) NULL,,
	Chofer        string    `gorm:"type:varchar(10) null"`  //varchar(10)) NULL,,
	Promotor      string    `gorm:"type:varchar(3) null"`   //varchar(3)) NULL,,
	Cajero        string    `gorm:"type:varchar(20) null"`  //varchar(20)) NULL,,
	Descrip1      string    `gorm:"type:varchar(50) null"`  //varchar(50)) NULL,,
	Descrip2      string    `gorm:"type:varchar(50) null"`  //varchar(50)) NULL,,
	Turno         string    `gorm:"type:varchar(1) null"`   //varchar(1)) NULL,,
	Almacen1      string    `gorm:"type:varchar(2) null"`   //varchar(2)) NULL,,
	Fecham        string    `gorm:"type:varchar(23) null"`  //varchar(23)) NULL,
	Fechavalidado string    `gorm:"type:varchar(26) null"`  //varchar(26)) NULL,,
	Dire          string    `gorm:"type:varchar(50) null"`  //varchar(50)) NULL,,
	Ciudad        string    `gorm:"type:varchar(30) null"`  //varchar(30)) NULL,,
	Tele          string    `gorm:"type:varchar(30) null"`  //varchar(30)) NULL,,
	Tasa1         float64   `gorm:"null"`                   //decimal(18, 2) NULL,
	Masitbis      string    `gorm:"type:varchar(1) null"`   //varchar(1)) NULL,,
	Cedula        string    `gorm:"type:varchar(15) null"`  //varchar(15)) NULL,,
	Comision      float64   `gorm:"null"`                   //decimal (18, 0) null,
	Ncf1          string    `gorm:"type:varchar(19) null"`  //varchar(19)) NULL,,
	Egreso        float64   `gorm:"null"`                   //decimal (19, 2)  NULL,
	Factor        float64   `gorm:"null"`                   //decimal(18, 2) NULL,
	Procesar      string    `gorm:"type:varchar(1) null"`   //varchar(1)) NULL,,
	Observa2      string    `gorm:"type:varchar(254) null"` //varchar(254)) NULL,,
	Anulado       string    `gorm:"type:varchar(50) null"`  //varchar(50)) NULL,,
	Conduce2      string    `gorm:"type:varchar(12) null"`  //varchar(12)) NULL,,
	Conduce3      string    `gorm:"type:varchar(12) null"`  //varchar(12)) NULL,,
	Conduce4      string    `gorm:"type:varchar(12) null"`  //varchar(12)) NULL,,
	Lectura1      float64   `gorm:"null"`                   //decimal(18, 2) NULL,
	Lectura2      float64   `gorm:"null"`                   //decimal(18, 2) NULL,
	Lectura3      float64   `gorm:"null"`                   //decimal(18, 2) NULL,
	Ncftipo       string    `gorm:"type:varchar(2) null"`   //varchar(2)) NULL,,
	Generancf     string    `gorm:"type:varchar(12) null"`  //varchar(12)) NULL,,
	Diferencia    float64   `gorm:"null"`                   //numeric (18, 2) NULL,
	Devolucion    float64   `gorm:"null"`                   //numeric (18, 2) NULL,
	Modificado    string    `gorm:"type:varchar(40) null"`  //varchar(40)) NULL,,
	Devuelta      float64   `gorm:"null"`                   //numeric (18, 2) NULL,
	Tcredito      string    `gorm:"type:varchar(1) null"`   //varchar(1)) NULL,,
	Sucursal1     string    `gorm:"type:varchar(2) null"`   //varchar(2)) NULL,,
	Cnttransa     float64   `gorm:"null"`                   //numeric (18, 0) NULL,
	Fechancf      time.Time `gorm:"null"`
	Contacto      string    `gorm:"type:varchar(50) null"` //varchar(50)) NULL,,
	Idsucursal    string    `gorm:"type:varchar(5) null"`  //varchar(5)) NULL,,
	Tasa          float64   `gorm:"null"`                  //decimal (18, 6)  NULL,
	Clasifica2    string    `gorm:"type:varchar(3) null"`  //varchar(3)) NULL,,
	E_ncfcode     string    `gorm:"type:varchar(20) null"` //varchar(20)) NULL,,
	Transferencia float64   `gorm:"null"`                  //decimal (18, 4)  NULL,
	Notacredito   float64   `gorm:"null"`                  //decimal (18, 4)  NULL,
	Notacredito1  string    `gorm:"type:varchar(12) null"` //varchar(12)) NULL,,
	NivelPrecio   int       `gorm:"null"`                  //int NULL,
	Itbisencosto  string    `gorm:"type:varchar(1) null"`  //varchar(1)) NULL,,
	Reimpresion   int       `gorm:"null"`                  //int NULL,
	FactTmp       string    `gorm:"type:varchar(12) null"` //varchar(12)) NULL,,
	AntDocumento  string    `gorm:"type:varchar(12) null"` //varchar(12)) NULL,,
	Arancel       float64   `gorm:"null"`                  //decimal (12, 6)  NULL,
	Gasto1        float64   `gorm:"null"`                  //decimal (12, 6)  NULL,
	Autorizado    string    `gorm:"type:varchar(20) null"` //varchar(20)) NULL,
}

// modelo transaciones query1

func GetTransaccionesVentas(startDate, endDate time.Time) ([]Transa01, error) {
	configs.ConnectToDB()
	defer configs.CloseDB()

	rows, err := configs.DB.Raw("SELECT * FROM Transa01 WHERE Fecha BETWEEN ? AND ?", startDate, endDate).Rows()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []Transa01
	for rows.Next() {
		var transaction Transa01
		if err := rows.Scan(&transaction.Empresa, &transaction.Fecha, &transaction.Tipo, &transaction.Numero, &transaction.Documento, &transaction.Documento1, &transaction.Turno1, &transaction.Ncf, &transaction.Cliente, &transaction.Suplidor, &transaction.Termino, &transaction.Dia, &transaction.Dia1, &transaction.Vence, &transaction.Digitado, &transaction.Itbis, &transaction.Fechadesc, &transaction.Otros, &transaction.Flete, &transaction.Descuen, &transaction.Nograva, &transaction.Grava, &transaction.Monto, &transaction.Total, &transaction.Balance, &transaction.Nombre, &transaction.Vendedor, &transaction.Vendedor1, &transaction.Venname, &transaction.Zona, &transaction.Pedido, &transaction.Conduce, &transaction.Observa, &transaction.Validado, &transaction.Posteado, &transaction.Efectivo, &transaction.Cheque, &transaction.Tarjeta, &transaction.Gastos, &transaction.Nulo, &transaction.ABIERTO, &transaction.ENVIADO, &transaction.Despacho, &transaction.Provincia, &transaction.Almacen, &transaction.POSTEO, &transaction.Socio, &transaction.Creado, &transaction.Dsc, &transaction.Prnc, &transaction.Prn, &transaction.Observa1, &transaction.Usuario, &transaction.Chofer, &transaction.Promotor, &transaction.Cajero, &transaction.Descrip1, &transaction.Descrip2, &transaction.Turno, &transaction.Almacen1, &transaction.Fecham, &transaction.Fechavalidado, &transaction.Dire, &transaction.Ciudad, &transaction.Tele, &transaction.Tasa1, &transaction.Masitbis, &transaction.Cedula, &transaction.Comision, &transaction.Ncf1, &transaction.Egreso, &transaction.Factor, &transaction.Procesar, &transaction.Observa2, &transaction.Anulado, &transaction.Conduce2, &transaction.Conduce3, &transaction.Conduce4, &transaction.Lectura1, &transaction.Lectura2, &transaction.Lectura3, &transaction.Ncftipo, &transaction.Generancf, &transaction.Diferencia, &transaction.Devolucion, &transaction.Modificado, &transaction.Devuelta, &transaction.Tcredito, &transaction.Sucursal1, &transaction.Cnttransa, &transaction.Fechancf, &transaction.Contacto, &transaction.Idsucursal, &transaction.Tasa, &transaction.Clasifica2, &transaction.E_ncfcode, &transaction.Transferencia, &transaction.Notacredito, &transaction.Notacredito1, &transaction.NivelPrecio, &transaction.Itbisencosto, &transaction.Reimpresion, &transaction.FactTmp, &transaction.AntDocumento, &transaction.Arancel, &transaction.Gasto1, &transaction.Autorizado /*... otros campos ...*/); err != nil {
			return nil, fmt.Errorf("error al escanear filas: %w", err)

		}
		transactions = append(transactions, transaction)
	}

	if len(transactions) == 0 {
		return nil, errors.New("no hay transacciones válidas en esta fecha")
	}

	return transactions, nil
}

/*func Main() {
	startDate := time.Date(2021, 1, 1, 0, 0, 0, 0, time.UTC)
	endDate := time.Date(2021, 12, 31, 0, 0, 0, 0, time.UTC)

	transactions, err := GetTransaccionesVentas(startDate, endDate)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Transacciones de ventas:", transactions)
}*/
