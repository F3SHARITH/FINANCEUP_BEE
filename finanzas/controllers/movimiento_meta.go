package controllers

import (
	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

type MovimientoMetaController struct {
	beego.Controller
}

func (c *MovimientoMetaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create MovimientoMeta
// @Param	body	body	models.MovimientoMeta	true	"body for MovimientoMeta content"
// @Success 201 {int} models.MovimientoMeta
// @Failure 403 body is empty
// @router / [post]
func (c *MovimientoMetaController) Post() {
	handlePost(&c.Controller, models.AddMovimientoMeta)
}

// GetOne ...
// @Title Get One
// @Description get MovimientoMeta by id
// @Param	id	path	string	true	"The key for staticblock"
// @Success 200 {object} models.MovimientoMeta
// @Failure 403 :id is empty
// @router /:id [get]
func (c *MovimientoMetaController) GetOne() {
	handleGetOne(&c.Controller, models.GetMovimientoMetaById)
}

// GetAll ...
// @Title Get All
// @Description get MovimientoMeta
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.MovimientoMeta
// @Failure 403
// @router / [get]
func (c *MovimientoMetaController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllMovimientoMeta)
}

// Put ...
// @Title Put
// @Description update the MovimientoMeta
// @Param	id	path	string	true	"The id you want to update"
// @Param	body	body	models.MovimientoMeta	true	"body for MovimientoMeta content"
// @Success 200 {object} models.MovimientoMeta
// @Failure 403 :id is not int
// @router /:id [put]
func (c *MovimientoMetaController) Put() {
	handlePut(&c.Controller, &models.MovimientoMeta{Id: pathID(&c.Controller)}, models.UpdateMovimientoMetaById)
}

// Delete ...
// @Title Delete
// @Description delete the MovimientoMeta
// @Param	id	path	string	true	"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *MovimientoMetaController) Delete() {
	handleDelete(&c.Controller, models.DeleteMovimientoMeta)
}
