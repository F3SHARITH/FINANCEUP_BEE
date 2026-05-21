package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

type Categoria struct {
	Id                int       `orm:"column(id_categoria);pk;auto"`
	Nombre            string    `orm:"column(nombre)"`
	Descripcion       string    `orm:"column(descripcion);null"`
	Activo            bool      `orm:"column(activo)"`
	FechaCreacion     time.Time `orm:"column(fecha_creacion);type(timestamp without time zone);auto_now_add"`
	FechaModificacion time.Time `orm:"column(fecha_modificacion);type(timestamp without time zone);auto_now"`
}

func (t *Categoria) TableName() string {
	return "categoria"
}

func init() {
	orm.RegisterModel(new(Categoria))
}

func AddCategoria(m *Categoria) (id int64, err error) {
	return addRecord(m)
}

func GetCategoriaById(id int) (v *Categoria, err error) {
	o := orm.NewOrm()
	v = &Categoria{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllCategoria(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	return getAllRecords(new(Categoria), query, fields, sortby, order, offset, limit)
}

func UpdateCategoriaById(m *Categoria) (err error) {
	return updateRecord(m, &Categoria{Id: m.Id})
}

func DeleteCategoria(id int) (err error) {
	return deleteRecord(&Categoria{Id: id})
}
