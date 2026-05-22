package controllers

import (
	"educacion/models"

	beego "github.com/beego/beego/v2/server/web"
)

// ProgresoEducativoController operations for ProgresoEducativo
type ProgresoEducativoController struct {
	beego.Controller
}

// URLMapping ...
func (c *ProgresoEducativoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create ProgresoEducativo
// @Param	body		body 	models.ProgresoEducativo	true		"body for ProgresoEducativo content"
// @Success 201 {int} models.ProgresoEducativo
// @Failure 403 body is empty
// @router / [post]
func (c *ProgresoEducativoController) Post() {
	handlePost(&c.Controller, models.AddProgresoEducativo)
}

// GetOne ...
// @Title Get One
// @Description get ProgresoEducativo by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.ProgresoEducativo
// @Failure 403 :id is empty
// @router /:id [get]
func (c *ProgresoEducativoController) GetOne() {
	handleGetOne(&c.Controller, models.GetProgresoEducativoById)
}

// GetAll ...
// @Title Get All
// @Description get ProgresoEducativo
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.ProgresoEducativo
// @Failure 403
// @router / [get]
func (c *ProgresoEducativoController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllProgresoEducativo)
}

// Put ...
// @Title Put
// @Description update the ProgresoEducativo
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.ProgresoEducativo	true		"body for ProgresoEducativo content"
// @Success 200 {object} models.ProgresoEducativo
// @Failure 403 :id is not int
// @router /:id [put]
func (c *ProgresoEducativoController) Put() {
	id, valid := pathID(&c.Controller)
	if !valid {
		return
	}
	handlePut(&c.Controller, &models.ProgresoEducativo{Id: id}, models.UpdateProgresoEducativoById)
}

// Delete ...
// @Title Delete
// @Description delete the ProgresoEducativo
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *ProgresoEducativoController) Delete() {
	handleDelete(&c.Controller, models.DeleteProgresoEducativo)
}
