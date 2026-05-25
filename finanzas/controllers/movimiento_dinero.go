package controllers

import (
	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

// MovimientoDineroController operations for MovimientoDinero
type MovimientoDineroController struct {
	beego.Controller
}

// URLMapping ...
func (c *MovimientoDineroController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create MovimientoDinero
// @Param	body		body 	models.MovimientoDinero	true		"body for MovimientoDinero content"
// @Success 201 {int} models.MovimientoDinero
// @Failure 403 body is empty
// @router / [post]
func (c *MovimientoDineroController) Post() {
	handlePost(&c.Controller, models.AddMovimientoDinero)
}

// GetOne ...
// @Title Get One
// @Description get MovimientoDinero by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.MovimientoDinero
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MovimientoDineroController) GetOne() {
	handleGetOne(&c.Controller, models.GetMovimientoDineroById)
}

// GetAll ...
// @Title Get All
// @Description get MovimientoDinero
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.MovimientoDinero
// @Failure 403
// @router / [get]
func (c *MovimientoDineroController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllMovimientoDinero)
}

// Put ...
// @Title Put
// @Description update the MovimientoDinero
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.MovimientoDinero	true		"body for MovimientoDinero content"
// @Success 200 {object} models.MovimientoDinero
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MovimientoDineroController) Put() {
	handlePut(&c.Controller, &models.MovimientoDinero{Id: pathID(&c.Controller)}, models.UpdateMovimientoDineroById)
}

// Delete ...
// @Title Delete
// @Description delete the MovimientoDinero
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MovimientoDineroController) Delete() {
	handleDelete(&c.Controller, models.DeleteMovimientoDinero)
}
