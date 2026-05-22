package controllers

import (
	"educacion/models"

	beego "github.com/beego/beego/v2/server/web"
)

// LeccionController operations for Leccion
type LeccionController struct {
	beego.Controller
}

// URLMapping ...
func (c *LeccionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

// Post ...
// @Title Post
// @Description create Leccion
// @Param	body		body 	models.Leccion	true		"body for Leccion content"
// @Success 201 {int} models.Leccion
// @Failure 403 body is empty
// @router / [post]
func (c *LeccionController) Post() {
	handlePost(&c.Controller, models.AddLeccion)
}

// GetOne ...
// @Title Get One
// @Description get Leccion by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.Leccion
// @Failure 403 :id is empty
// @router /:id [get]
func (c *LeccionController) GetOne() {
	handleGetOne(&c.Controller, models.GetLeccionById)
}

// GetAll ...
// @Title Get All
// @Description get Leccion
// @Param	query	query	string	false	"Filter. e.g. col1:v1,col2:v2 ..."
// @Param	fields	query	string	false	"Fields returned. e.g. col1,col2 ..."
// @Param	sortby	query	string	false	"Sorted-by fields. e.g. col1,col2 ..."
// @Param	order	query	string	false	"Order corresponding to each sortby field, if single value, apply to all sortby fields. e.g. desc,asc ..."
// @Param	limit	query	string	false	"Limit the size of result set. Must be an integer"
// @Param	offset	query	string	false	"Start position of result set. Must be an integer"
// @Success 200 {object} models.Leccion
// @Failure 403
// @router / [get]
func (c *LeccionController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllLeccion)
}

// Put ...
// @Title Put
// @Description update the Leccion
// @Param	id		path 	string	true		"The id you want to update"
// @Param	body		body 	models.Leccion	true		"body for Leccion content"
// @Success 200 {object} models.Leccion
// @Failure 403 :id is not int
// @router /:id [put]
func (c *LeccionController) Put() {
	id, valid := pathID(&c.Controller)
	if !valid {
		return
	}
	handlePut(&c.Controller, &models.Leccion{Id: id}, models.UpdateLeccionById)
}

// Delete ...
// @Title Delete
// @Description delete the Leccion
// @Param	id		path 	string	true		"The id you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 id is empty
// @router /:id [delete]
func (c *LeccionController) Delete() {
	handleDelete(&c.Controller, models.DeleteLeccion)
}
