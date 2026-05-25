package controllers

import (
	"encoding/json"
	"strconv"
	"strings"

	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

type TipoIngresoController struct {
	beego.Controller
}

func (c *TipoIngresoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create TipoIngreso
// @Param	body	body	models.TipoIngreso	true	"body for TipoIngreso content"
// @Success 201 {int} models.TipoIngreso
// @Failure 403 body is empty
// @router / [post]
func (c *TipoIngresoController) Post() {
	var v models.TipoIngreso
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := models.AddTipoIngreso(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = map[string]interface{}{"success": true, "status": 201, "Message": "Peticion exitosa", "data": v}
		} else {
			c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servidor Post: no se pudo crear el recurso", "error": err.Error()}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error al parsear JSON", "error": err.Error()}
	}
	c.ServeJSON()
}

// GetOne ...
// @Title Get One
// @Description get TipoIngreso by id
// @Param	id	path	string	true	"The key for staticblock"
// @Success 200 {object} models.TipoIngreso
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TipoIngresoController) GetOne() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v, err := models.GetTipoIngresoById(id)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servidor GetOne: recurso no encontrado", "error": err.Error()}
	} else {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Peticion exitosa", "data": v}
	}
	c.ServeJSON()
}

// GetAll ...
// @Title Get All
// @Description get TipoIngreso
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TipoIngreso
// @Failure 403
// @router / [get]
func (c *TipoIngresoController) GetAll() {
	var fields []string
	var sortby []string
	var order []string
	query := make(map[string]string)
	var limit int64 = 10
	var offset int64

	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error: invalid query key/value pair"}
				c.ServeJSON()
				return
			}
			query[kv[0]] = kv[1]
		}
	}

	l, err := models.GetAllTipoIngreso(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servidor GetAll", "error": err.Error()}
	} else {
		if l == nil {
			l = []interface{}{}
		}
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Peticion exitosa", "data": l}
	}
	c.ServeJSON()
}

// Put ...
// @Title Put
// @Description update the TipoIngreso
// @Param	id	path	string	true	"The id you want to update"
// @Param	body	body	models.TipoIngreso	true	"body for TipoIngreso content"
// @Success 200 {object} models.TipoIngreso
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TipoIngresoController) Put() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	v := models.TipoIngreso{Id: id}
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if err := models.UpdateTipoIngresoById(&v); err == nil {
			c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "Peticion exitosa", "data": v}
		} else {
			c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servidor Put: no se pudo actualizar el recurso", "error": err.Error()}
		}
	} else {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error al parsear JSON", "error": err.Error()}
	}
	c.ServeJSON()
}

// Delete ...
// @Title Delete
// @Description delete the TipoIngreso
// @Param	id	path	string	true	"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TipoIngresoController) Delete() {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := models.DeleteTipoIngreso(id); err == nil {
		c.Data["json"] = map[string]interface{}{"success": true, "status": 200, "Message": "dato eliminado", "data": id}
	} else {
		c.Data["json"] = map[string]interface{}{"success": false, "status": 400, "Message": "Error en el servidor Delete", "error": err.Error()}
	}
	c.ServeJSON()
}
