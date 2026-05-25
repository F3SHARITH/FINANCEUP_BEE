package controllers

import (
	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

// TipoInversionController operations for TipoInversion
type TipoInversionController struct {
	beego.Controller
}

// URLMapping ...
func (c *TipoInversionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create TipoInversion
// @Param	body		body 	models.TipoInversion	true		"body for TipoInversion content"
// @Success 201 {int} models.TipoInversion
// @Failure 403 body is empty
// @router / [post]
func (c *TipoInversionController) Post() {
	handlePost(&c.Controller, models.AddTipoInversion)
}

// GetOne ...
// @Title Get One
// @Description get TipoInversion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.TipoInversion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TipoInversionController) GetOne() {
	handleGetOne(&c.Controller, models.GetTipoInversionById)
}

// GetAll ...
// @Title Get All
// @Description get TipoInversion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TipoInversion
// @Failure 403
// @router / [get]
func (c *TipoInversionController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllTipoInversion)
}

// Put ...
// @Title Put
// @Description update the TipoInversion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TipoInversion	true		"body for TipoInversion content"
// @Success 200 {object} models.TipoInversion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TipoInversionController) Put() {
	handlePut(&c.Controller, &models.TipoInversion{Id: pathID(&c.Controller)}, models.UpdateTipoInversionById)
}

// Delete ...
// @Title Delete
// @Description delete the TipoInversion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TipoInversionController) Delete() {
	handleDelete(&c.Controller, models.DeleteTipoInversion)
}
