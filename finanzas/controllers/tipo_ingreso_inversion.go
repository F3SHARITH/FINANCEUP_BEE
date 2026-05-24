package controllers

import (
	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

type TipoIngresoInversionController struct {
	beego.Controller
}

func (c *TipoIngresoInversionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create TipoIngresoInversion
// @Param	body	body	models.TipoIngresoInversion	true	"body for TipoIngresoInversion content"
// @Success 201 {int} models.TipoIngresoInversion
// @Failure 403 body is empty
// @router / [post]
func (c *TipoIngresoInversionController) Post() {
	handlePost(&c.Controller, models.AddTipoIngresoInversion)
}

// GetOne ...
// @Title Get One
// @Description get TipoIngresoInversion by id
// @Param	id	path	string	true	"The key for staticblock"
// @Success 200 {object} models.TipoIngresoInversion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TipoIngresoInversionController) GetOne() {
	handleGetOne(&c.Controller, models.GetTipoIngresoInversionById)
}

// GetAll ...
// @Title Get All
// @Description get TipoIngresoInversion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TipoIngresoInversion
// @Failure 403
// @router / [get]
func (c *TipoIngresoInversionController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllTipoIngresoInversion)
}

// Put ...
// @Title Put
// @Description update the TipoIngresoInversion
// @Param	id	path	string	true	"The id you want to update"
// @Param	body	body	models.TipoIngresoInversion	true	"body for TipoIngresoInversion content"
// @Success 200 {object} models.TipoIngresoInversion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TipoIngresoInversionController) Put() {
	handlePut(&c.Controller, &models.TipoIngresoInversion{Id: pathID(&c.Controller)}, models.UpdateTipoIngresoInversionById)
}

// Delete ...
// @Title Delete
// @Description delete the TipoIngresoInversion
// @Param	id	path	string	true	"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TipoIngresoInversionController) Delete() {
	handleDelete(&c.Controller, models.DeleteTipoIngresoInversion)
}
