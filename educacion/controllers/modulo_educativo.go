package controllers

import (
	"educacion/models"

	beego "github.com/beego/beego/v2/server/web"
)

// ModuloEducativoController operations for ModuloEducativo
type ModuloEducativoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ModuloEducativoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ModuloEducativo
// @Param	body		body 	models.ModuloEducativo	true		"body for ModuloEducativo content"
// @Success 201 {int} models.ModuloEducativo
// @Failure 403 body is empty
// @router / [post]
func (c *ModuloEducativoController) Post() {
	handlePost(&c.Controller, models.AddModuloEducativo)
}

// GetOne ...
// @Title Get One
// @Description get ModuloEducativo by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ModuloEducativo
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ModuloEducativoController) GetOne() {
	handleGetOne(&c.Controller, models.GetModuloEducativoById)
}

// GetAll ...
// @Title Get All
// @Description get ModuloEducativo
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ModuloEducativo
// @Failure 403
// @router / [get]
func (c *ModuloEducativoController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllModuloEducativo)
}

// Put ...
// @Title Put
// @Description update the ModuloEducativo
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ModuloEducativo	true		"body for ModuloEducativo content"
// @Success 200 {object} models.ModuloEducativo
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ModuloEducativoController) Put() {
	id, valid := pathID(&c.Controller)
	if !valid {
		return
	}
	handlePut(&c.Controller, &models.ModuloEducativo{Id: id}, models.UpdateModuloEducativoById)
}

// Delete ...
// @Title Delete
// @Description delete the ModuloEducativo
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ModuloEducativoController) Delete() {
	handleDelete(&c.Controller, models.DeleteModuloEducativo)
}
