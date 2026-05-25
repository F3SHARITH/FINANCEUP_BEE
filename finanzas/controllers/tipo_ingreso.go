package controllers

import (
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
	handlePost(&c.Controller, models.AddTipoIngreso)
}

// GetOne ...
// @Title Get One
// @Description get TipoIngreso by id
// @Param	id	path	string	true	"The key for staticblock"
// @Success 200 {object} models.TipoIngreso
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TipoIngresoController) GetOne() {
	handleGetOne(&c.Controller, models.GetTipoIngresoById)
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
	handleGetAll(&c.Controller, models.GetAllTipoIngreso)
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
	handlePut(&c.Controller, &models.TipoIngreso{Id: pathID(&c.Controller)}, models.UpdateTipoIngresoById)
}

// Delete ...
// @Title Delete
// @Description delete the TipoIngreso
// @Param	id	path	string	true	"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TipoIngresoController) Delete() {
	handleDelete(&c.Controller, models.DeleteTipoIngreso)
}
