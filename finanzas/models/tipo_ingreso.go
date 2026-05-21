package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type TipoIngreso struct {
	Id                   int               `orm:"column(id_tipo_ingreso);pk;auto"`
	IdMovimientoDinero   *MovimientoDinero `orm:"column(id_movimiento_dinero);rel(fk);null"`
	NombreMovimientoPago string            `orm:"column(nombre_movimiento_pago)"`
	Descripcion          string            `orm:"column(descripcion);null"`
	Activo               bool              `orm:"column(activo)"`
	FechaCreacion        time.Time         `orm:"column(fecha_creacion);type(timestamp without time zone);auto_now_add"`
	FechaModificacion    time.Time         `orm:"column(fecha_modificacion);type(timestamp without time zone);auto_now"`
}

func (t *TipoIngreso) TableName() string {
	return "tipo_ingreso"
}

func init() {
	orm.RegisterModel(new(TipoIngreso))
}

func AddTipoIngreso(m *TipoIngreso) (id int64, err error) {
	return addRecord(m)
}

func GetTipoIngresoById(id int) (v *TipoIngreso, err error) {
	o := orm.NewOrm()
	v = &TipoIngreso{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllTipoIngreso(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	return getAllRecords(new(TipoIngreso), query, fields, sortby, order, offset, limit)
}

func UpdateTipoIngresoById(m *TipoIngreso) (err error) {
	return updateRecord(m, &TipoIngreso{Id: m.Id})
}

func DeleteTipoIngreso(id int) (err error) {
	return deleteRecord(&TipoIngreso{Id: id})
}
