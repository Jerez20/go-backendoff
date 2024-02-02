package controllers

import (
	"net/http"
	"time"

	"Base1.com/go-backend/configs"
	"Base1.com/go-backend/models"
	"github.com/gin-gonic/gin"
)

type Transa01RequestBody struct {
	Empresa       int       `json:"empresa"` //int NULL,
	Fecha         time.Time `json:"fecha" binding:"required"`
	Tipo          string    `json:"tipo"`       //varchar(2)) NULL,,
	Numero        string    `json:"numero"`     //varchar(12)) NULL,,
	Documento     string    `json:"documento"`  //varchar(12)) NULL,,
	Documento1    string    `json:"documento1"` //varchar(12)) NULL,,
	Turno1        string    `json:"turno1"`     //varchar(1)) NULL,,
	Ncf           string    `json:"ncf"`        //varchar(19)) NULL,,
	Cliente       string    `json:"cliente"`    //varchar(15)) NULL,,
	Suplidor      string    `json:"suplidor"`   //varchar(15)) NULL,,
	Termino       string    `json:"termino"`    //varchar(3)) NULL,,
	Dia           int       `json:"dia"`        //int NULL,
	Dia1          int       `json:"dia1"`       //int NULL,
	Vence         time.Time `json:"vence"`
	Digitado      time.Time `json:"digitado"`
	Itbis         float64   `json:"itebis"` //decimal(18, 2) NULL,
	Fechadesc     time.Time `json:"fechadesc"`
	Otros         float64   `json:"otros"`         //decimal(18, 2) NULL,
	Flete         float64   `json:"flete"`         //decimal(18, 2) NULL,
	Descuen       float64   `json:"descuen"`       //decimal(18, 2) NULL,
	Nograva       float64   `json:"nograva"`       //decimal(18, 2) NULL,
	Grava         float64   `json:"grava"`         //decimal(18, 2) NULL,
	Monto         float64   `json:"monto"`         //decimal(18, 2) NULL,
	Total         float64   `json:"total"`         //decimal(18, 2) NULL,
	Balance       float64   `json:"balance"`       //decimal(18, 2) NULL,
	Nombre        string    `json:"nombre"`        //varchar(50)) NULL,,
	Vendedor      string    `json:"vendedor"`      //varchar(3)) NULL,,
	Vendedor1     string    `json:"vendedor1"`     //varchar(3)) NULL,,
	Venname       string    `json:"venname"`       //varchar(50)) NULL,,
	Zona          string    `json:"zona"`          //varchar(3)) NULL,,
	Pedido        string    `json:"pedido"`        //varchar(12)) NULL,,
	Conduce       string    `json:"conduce"`       //varchar(12)) NULL,,
	Observa       string    `json:"observa"`       //varchar(254)) NULL,,
	Validado      string    `json:"validado"`      //varchar(1)) NULL,,
	Posteado      string    `json:"posteado"`      //varchar(1)) NULL,,
	Efectivo      float64   `json:"efectivo"`      //decimal (18, 4)  NULL,
	Cheque        float64   `json:"cheque"`        //decimal (18, 4)  NULL,
	Tarjeta       float64   `json:"tarjeta"`       //decimal (18, 4)  NULL,
	Gastos        float64   `json:"gastos"`        //numeric (18, 2) NULL,
	Nulo          string    `json:"nulo"`          //varchar(1)) NULL,,
	ABIERTO       string    `json:"abierto"`       //varchar(1)) NULL,,
	ENVIADO       float64   `json:"enviado"`       //decimal(18, 2) NULL,
	Despacho      string    `json:"despacho"`      //varchar(1)) NULL,,
	Provincia     string    `json:"provincia"`     //varchar(3)) NULL,,
	Almacen       string    `json:"almacen"`       //varchar(2)) NULL,,
	POSTEO        float64   `json:"posteo"`        //decimal (18, 0) null,
	Socio         string    `json:"socio"`         //varchar(1)) NULL,,
	Creado        string    `json:"creado"`        //varchar(20)) NULL,,
	Dsc           string    `json:"dsc"`           //varchar(1)) NULL,,
	Prnc          string    `json:"prnc"`          //varchar(1)) NULL,,
	Prn           string    `json:"prn"`           //varchar(1)) NULL,,
	Observa1      string    `json:"observa1"`      //varchar(254)) NULL,,
	Usuario       string    `json:"usuario"`       //varchar(30)) NULL,,
	Chofer        string    `json:"chofer"`        //varchar(10)) NULL,,
	Promotor      string    `json:"promotor"`      //varchar(3)) NULL,,
	Cajero        string    `json:"cajero"`        //varchar(20)) NULL,,
	Descrip1      string    `json:"descrip1"`      //varchar(50)) NULL,,
	Descrip2      string    `json:"descript2"`     //varchar(50)) NULL,,
	Turno         string    `json:"turno"`         //varchar(1)) NULL,,
	Almacen1      string    `json:"almacen1"`      //varchar(2)) NULL,,
	Fecham        string    `json:"fecham"`        //varchar(23)) NULL,
	Fechavalidado string    `json:"fechavalidado"` //varchar(26)) NULL,,
	Dire          string    `json:"dire"`          //varchar(50)) NULL,,
	Ciudad        string    `json:"ciudad"`        //varchar(30)) NULL,,
	Tele          string    `json:"tele"`          //varchar(30)) NULL,,
	Tasa1         float64   `json:"tasa1"`         //decimal(18, 2) NULL,
	Masitbis      string    `json:"masitibis"`     //varchar(1)) NULL,,
	Cedula        string    `json:"cedula"`        //varchar(15)) NULL,,
	Comision      float64   `json:"comision"`      //decimal (18, 0) null,
	Ncf1          string    `json:"ncf1"`          //varchar(19)) NULL,,
	Egreso        float64   `json:"egreso"`        //decimal (19, 2)  NULL,
	Factor        float64   `json:"factor"`        //decimal(18, 2) NULL,
	Procesar      string    `json:"procesar"`      //varchar(1)) NULL,,
	Observa2      string    `json:"observa2"`      //varchar(254)) NULL,,
	Anulado       string    `json:"anulado"`       //varchar(50)) NULL,,
	Conduce2      string    `json:"conduce2"`      //varchar(12)) NULL,,
	Conduce3      string    `json:"conduce3"`      //varchar(12)) NULL,,
	Conduce4      string    `json:"conduce4"`      //varchar(12)) NULL,,
	Lectura1      float64   `json:"lectura1"`      //decimal(18, 2) NULL,
	Lectura2      float64   `json:"lectura2"`      //decimal(18, 2) NULL,
	Lectura3      float64   `json:"lectura3"`      //decimal(18, 2) NULL,
	Ncftipo       string    `json:"ncftipo"`       //varchar(2)) NULL,,
	Generancf     string    `json:"generancf"`     //varchar(12)) NULL,,
	Diferencia    float64   `json:"diferencia"`    //numeric (18, 2) NULL,
	Devolucion    float64   `json:"devolucion"`    //numeric (18, 2) NULL,
	Modificado    string    `json:"modificado"`    //varchar(40)) NULL,,
	Devuelta      float64   `json:"devuelta"`      //numeric (18, 2) NULL,
	Tcredito      string    `json:"tcredito"`      //varchar(1)) NULL,,
	Sucursal1     string    `json:"sucursal1"`     //varchar(2)) NULL,,
	Cnttransa     float64   `json:"cnttransa"`     //numeric (18, 0) NULL,
	Fechancf      time.Time `json:"fechancf"`
	Contacto      string    `json:"contacto"`      //varchar(50)) NULL,,
	Idsucursal    string    `json:"idsucursal"`    //varchar(5)) NULL,,
	Tasa          float64   `json:"tasa"`          //decimal (18, 6)  NULL,
	Clasifica2    string    `json:"calific2"`      //varchar(3)) NULL,,
	E_ncfcode     string    `json:"e_ncfcode"`     //varchar(20)) NULL,,
	Transferencia float64   `json:"transferencia"` //decimal (18, 4)  NULL,
	Notacredito   float64   `json:"notacredito"`   //decimal (18, 4)  NULL,
	Notacredito1  string    `json:"notacredito1"`  //varchar(12)) NULL,,
	NivelPrecio   int       `json:"nivelprecio"`   //int NULL,
	Itbisencosto  string    `json:"itebisencosto"` //varchar(1)) NULL,,
	Reimpresion   int       `json:"reimpresion"`   //int NULL,
	FactTmp       string    `json:"factTmp"`       //varchar(12)) NULL,,
	AntDocumento  string    `json:"antDocumento"`  //varchar(12)) NULL,,
	Arancel       float64   `json:"arancel"`       //decimal (12, 6)  NULL,
	Gasto1        float64   `json:"gasto1"`        //decimal (12, 6)  NULL,
	Autorizado    string    `json:"autorizado"`    //varchar(20)) NULL,
}

func Transaget(c *gin.Context) {
	var Transas []models.Transa01
	configs.DB.Find(&Transas)
	c.JSON(200, &Transas)
}

// transaciones model query 1
func GetTransaccionesVentas(c *gin.Context) {
	textValue := c.Query("start_date")
	startDate, err := time.Parse("20060102", textValue)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al parsear la fecha de inicio: " + err.Error()})
		return
	}

	textValue = c.Query("end_date")
	endDate, err := time.Parse("20060102", textValue)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Error al parsear la fecha de fin: " + err.Error()})
		return
	}

	transacciones, err := models.GetTransaccionesVentas(startDate, endDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener las transacciones de ventas"})
		return
	}

	c.JSON(http.StatusOK, transacciones)
}
