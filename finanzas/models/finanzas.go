package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
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
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetFinanzasById(id int) (v *Finanzas, err error) {
	o := orm.NewOrm()
	v = &Finanzas{Id: id}
	qs := o.QueryTable(new(Finanzas)).Filter("Id", id).RelatedSel()
	if err = qs.One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllFinanzas(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Finanzas)).RelatedSel()
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}

	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(order) == 1 {
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else if len(order) != 0 {
		return nil, errors.New("Error: unused 'order' fields")
	}

	var l []Finanzas
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

func UpdateFinanzasById(m *Finanzas) (err error) {
	o := orm.NewOrm()
	v := Finanzas{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func DeleteFinanzas(id int) (err error) {
	o := orm.NewOrm()
	v := Finanzas{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Finanzas{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
