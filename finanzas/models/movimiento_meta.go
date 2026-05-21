package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type MovimientoMeta struct {
	Id                int       `orm:"column(id_movimiento_dinero);pk;auto"`
	Nombre            string    `orm:"column(nombre)"`
	Monto             float64   `orm:"column(monto)"`
	EsIngreso         bool      `orm:"column(es_ingreso)"`
	Activo            bool      `orm:"column(activo)"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(timestamp without time zone);auto_now_add"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone);auto_now"`
}

func (t *MovimientoMeta) TableName() string {
	return "movimiento_meta"
}

func init() {
	orm.RegisterModel(new(MovimientoMeta))
}

func AddMovimientoMeta(m *MovimientoMeta) (id int64, err error) {
	return addRecord(m)
}

func GetMovimientoMetaById(id int) (v *MovimientoMeta, err error) {
	o := orm.NewOrm()
	v = &MovimientoMeta{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllMovimientoMeta(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	return getAllRecords(new(MovimientoMeta), query, fields, sortby, order, offset, limit)
}

func UpdateMovimientoMetaById(m *MovimientoMeta) (err error) {
	return updateRecord(m, &MovimientoMeta{Id: m.Id})
}

func DeleteMovimientoMeta(id int) (err error) {
	return deleteRecord(&MovimientoMeta{Id: id})
}
