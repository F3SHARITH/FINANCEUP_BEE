package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type TipoIngresoInversion struct {
	Id                   int                  `orm:"column(id_tipo_ingreso);pk;auto"`
	IdMovimientoDinero   *MovimientoInversion `orm:"column(id_movimiento_dinero);rel(fk);null"`
	NombreMovimientoPago string               `orm:"column(nombre_movimiento_pago)"`
	Descripcion          string               `orm:"column(descripcion);null"`
	Activo               bool                 `orm:"column(activo)"`
	FechaCreacion        time.Time            `orm:"column(fecha_creacion);type(timestamp without time zone);auto_now_add"`
	FechaModificacion    time.Time            `orm:"column(fecha_modificacion);type(timestamp without time zone);auto_now"`
}

func (t *TipoIngresoInversion) TableName() string {
	return "tipo_ingreso_inversion"
}

func init() {
	orm.RegisterModel(new(TipoIngresoInversion))
}

func AddTipoIngresoInversion(m *TipoIngresoInversion) (id int64, err error) {
	return addRecord(m)
}

func GetTipoIngresoInversionById(id int) (v *TipoIngresoInversion, err error) {
	o := orm.NewOrm()
	v = &TipoIngresoInversion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllTipoIngresoInversion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	return getAllRecords(new(TipoIngresoInversion), query, fields, sortby, order, offset, limit)
}

func UpdateTipoIngresoInversionById(m *TipoIngresoInversion) (err error) {
	return updateRecord(m, &TipoIngresoInversion{Id: m.Id})
}

func DeleteTipoIngresoInversion(id int) (err error) {
	return deleteRecord(&TipoIngresoInversion{Id: id})
}
