package controllers

import (
	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

// TipoIngresoMetaController operations for TipoIngresoMeta
type TipoIngresoMetaController struct {
	beego.Controller
}

// URLMapping ...
func (c *TipoIngresoMetaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create TipoIngresoMeta
// @Param	body		body 	models.TipoIngresoMeta	true		"body for TipoIngresoMeta content"
// @Success 201 {int} models.TipoIngresoMeta
// @Failure 403 body is empty
// @router / [post]
func (c *TipoIngresoMetaController) Post() {
	handlePost(&c.Controller, models.AddTipoIngresoMeta)
}

// GetOne ...
// @Title Get One
// @Description get TipoIngresoMeta by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.TipoIngresoMeta
// @Failure 403 :id is empty
// @router /:id [get]
func (c *TipoIngresoMetaController) GetOne() {
	handleGetOne(&c.Controller, models.GetTipoIngresoMetaById)
}

// GetAll ...
// @Title Get All
// @Description get TipoIngresoMeta
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.TipoIngresoMeta
// @Failure 403
// @router / [get]
func (c *TipoIngresoMetaController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllTipoIngresoMeta)
}

// Put ...
// @Title Put
// @Description update the TipoIngresoMeta
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.TipoIngresoMeta	true		"body for TipoIngresoMeta content"
// @Success 200 {object} models.TipoIngresoMeta
// @Failure 403 :id is not int
// @router /:id [put]
func (c *TipoIngresoMetaController) Put() {
	handlePut(&c.Controller, &models.TipoIngresoMeta{Id: pathID(&c.Controller)}, models.UpdateTipoIngresoMetaById)
}

// Delete ...
// @Title Delete
// @Description delete the TipoIngresoMeta
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *TipoIngresoMetaController) Delete() {
	handleDelete(&c.Controller, models.DeleteTipoIngresoMeta)
}
