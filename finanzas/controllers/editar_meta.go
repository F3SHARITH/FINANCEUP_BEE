package controllers

import (
	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

// EditarMetaController operations for EditarMeta
type EditarMetaController struct {
	beego.Controller
}

// URLMapping ...
func (c *EditarMetaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create EditarMeta
// @Param	body		body 	models.EditarMeta	true		"body for EditarMeta content"
// @Success 201 {int} models.EditarMeta
// @Failure 403 body is empty
// @router / [post]
func (c *EditarMetaController) Post() {
	handlePost(&c.Controller, models.AddEditarMeta)
}

// GetOne ...
// @Title Get One
// @Description get EditarMeta by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.EditarMeta
// @Failure 403 :id is empty
// @router /:id [get]
func (c *EditarMetaController) GetOne() {
	handleGetOne(&c.Controller, models.GetEditarMetaById)
}

// GetAll ...
// @Title Get All
// @Description get EditarMeta
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.EditarMeta
// @Failure 403
// @router / [get]
func (c *EditarMetaController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllEditarMeta)
}

// Put ...
// @Title Put
// @Description update the EditarMeta
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.EditarMeta	true		"body for EditarMeta content"
// @Success 200 {object} models.EditarMeta
// @Failure 403 :id is not int
// @router /:id [put]
func (c *EditarMetaController) Put() {
	handlePut(&c.Controller, &models.EditarMeta{Id: pathID(&c.Controller)}, models.UpdateEditarMetaById)
}

// Delete ...
// @Title Delete
// @Description delete the EditarMeta
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *EditarMetaController) Delete() {
	handleDelete(&c.Controller, models.DeleteEditarMeta)
}
