package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type MovimientoInversion struct {
	Id                int       `orm:"column(id_movimiento_dinero);pk;auto"`
	Nombre            string    `orm:"column(nombre)"`
	Monto             float64   `orm:"column(monto)"`
	EsIngreso         bool      `orm:"column(es_ingreso)"`
	Activo            bool      `orm:"column(activo)"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(timestamp without time zone);auto_now_add"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone);auto_now"`
}

func (t *MovimientoInversion) TableName() string {
	return "movimiento_inversion"
}

func init() {
	orm.RegisterModel(new(MovimientoInversion))
}

func AddMovimientoInversion(m *MovimientoInversion) (id int64, err error) {
	return addRecord(m)
}

func GetMovimientoInversionById(id int) (v *MovimientoInversion, err error) {
	o := orm.NewOrm()
	v = &MovimientoInversion{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllMovimientoInversion(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	return getAllRecords(new(MovimientoInversion), query, fields, sortby, order, offset, limit)
}

func UpdateMovimientoInversionById(m *MovimientoInversion) (err error) {
	return updateRecord(m, &MovimientoInversion{Id: m.Id})
}

func DeleteMovimientoInversion(id int) (err error) {
	return deleteRecord(&MovimientoInversion{Id: id})
}
