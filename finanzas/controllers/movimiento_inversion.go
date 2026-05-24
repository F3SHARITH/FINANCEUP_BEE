package controllers

import (
	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

type MovimientoInversionController struct {
	beego.Controller
}

func (c *MovimientoInversionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create MovimientoInversion
// @Param	body	body	models.MovimientoInversion	true	"body for MovimientoInversion content"
// @Success 201 {int} models.MovimientoInversion
// @Failure 403 body is empty
// @router / [post]
func (c *MovimientoInversionController) Post() {
	handlePost(&c.Controller, models.AddMovimientoInversion)
}

// GetOne ...
// @Title Get One
// @Description get MovimientoInversion by id
// @Param	id	path	string	true	"The key for staticblock"
// @Success 200 {object} models.MovimientoInversion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MovimientoInversionController) GetOne() {
	handleGetOne(&c.Controller, models.GetMovimientoInversionById)
}

// GetAll ...
// @Title Get All
// @Description get MovimientoInversion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.MovimientoInversion
// @Failure 403
// @router / [get]
func (c *MovimientoInversionController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllMovimientoInversion)
}

// Put ...
// @Title Put
// @Description update the MovimientoInversion
// @Param	id	path	string	true	"The id you want to update"
// @Param	body	body	models.MovimientoInversion	true	"body for MovimientoInversion content"
// @Success 200 {object} models.MovimientoInversion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MovimientoInversionController) Put() {
	handlePut(&c.Controller, &models.MovimientoInversion{Id: pathID(&c.Controller)}, models.UpdateMovimientoInversionById)
}

// Delete ...
// @Title Delete
// @Description delete the MovimientoInversion
// @Param	id	path	string	true	"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MovimientoInversionController) Delete() {
	handleDelete(&c.Controller, models.DeleteMovimientoInversion)
}
