package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type NivelRiesgo struct {
	Id                int       `orm:"column(id_nivel_riesgo);pk;auto"`
	Nombre            string    `orm:"column(nombre)"`
	Activo            bool      `orm:"column(activo)"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(timestamp without time zone);auto_now_add"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone);auto_now"`
}

func (t *NivelRiesgo) TableName() string {
	return "nivel_riesgo"
}

func init() {
	orm.RegisterModel(new(NivelRiesgo))
}

func AddNivelRiesgo(m *NivelRiesgo) (id int64, err error) {
	return addRecord(m)
}

func GetNivelRiesgoById(id int) (v *NivelRiesgo, err error) {
	o := orm.NewOrm()
	v = &NivelRiesgo{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllNivelRiesgo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	return getAllRecords(new(NivelRiesgo), query, fields, sortby, order, offset, limit)
}

func UpdateNivelRiesgoById(m *NivelRiesgo) (err error) {
	return updateRecord(m, &NivelRiesgo{Id: m.Id})
}

func DeleteNivelRiesgo(id int) (err error) {
	return deleteRecord(&NivelRiesgo{Id: id})
}
