package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Finanzas struct {
	Id                 int               `orm:"column(id_finanzas);pk;auto"`
	IdUsuario          int               `orm:"column(id_usuario)"`
	IdMovimientoDinero *MovimientoDinero `orm:"column(id_movimiento_dinero);rel(fk);null"`
	IdCategoria        *Categoria        `orm:"column(id_categoria);rel(fk);null"`
	MontoPresupuesto   float64           `orm:"column(monto_presupuesto);null"`
	Gasto              float64           `orm:"column(gasto);null"`
	Disponible         float64           `orm:"column(disponible);null"`
	Fecha              time.Time         `orm:"column(fecha);type(date);null"`
	Activo             bool              `orm:"column(activo)"`
	FechaCreacion      time.Time         `orm:"column(fecha_creacion);type(timestamp without time zone);auto_now_add"`
	FechaModificacion  time.Time         `orm:"column(fecha_modificacion);type(timestamp without time zone);auto_now"`
}

func (t *Finanzas) TableName() string {
	return "finanzas"
}

func init() {
	orm.RegisterModel(new(Finanzas))
}

func AddFinanzas(m *Finanzas) (id int64, err error) {
	return addRecord(m)
}

func GetFinanzasById(id int) (v *Finanzas, err error) {
	o := orm.NewOrm()
	v = &Finanzas{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllFinanzas(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	return getAllRecords(new(Finanzas), query, fields, sortby, order, offset, limit)
}

func UpdateFinanzasById(m *Finanzas) (err error) {
	return updateRecord(m, &Finanzas{Id: m.Id})
}

func DeleteFinanzas(id int) (err error) {
	return deleteRecord(&Finanzas{Id: id})
}
