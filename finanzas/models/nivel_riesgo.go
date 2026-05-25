package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
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
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

func GetNivelRiesgoById(id int) (v *NivelRiesgo, err error) {
	o := orm.NewOrm()
	v = &NivelRiesgo{Id: id}
	qs := o.QueryTable(new(NivelRiesgo)).Filter("Id", id).RelatedSel()
	if err = qs.One(v); err == nil {
		return v, nil
	}
	return nil, err
}

func GetAllNivelRiesgo(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(NivelRiesgo)).RelatedSel()
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

	var l []NivelRiesgo
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

func UpdateNivelRiesgoById(m *NivelRiesgo) (err error) {
	o := orm.NewOrm()
	v := NivelRiesgo{Id: m.Id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

func DeleteNivelRiesgo(id int) (err error) {
	o := orm.NewOrm()
	v := NivelRiesgo{Id: id}
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&NivelRiesgo{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
