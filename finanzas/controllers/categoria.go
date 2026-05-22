package controllers

import (
	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

type CategoriaController struct {
	beego.Controller
}

func (c *CategoriaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *CategoriaController) Post() {
	handlePost(&c.Controller, models.AddCategoria)
}

func (c *CategoriaController) GetOne() {
	handleGetOne(&c.Controller, models.GetCategoriaById)
}

func (c *CategoriaController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllCategoria)
}

func (c *CategoriaController) Put() {
	id, valid := pathID(&c.Controller)
	if !valid {
		return
	}
	handlePut(&c.Controller, &models.Categoria{Id: id}, models.UpdateCategoriaById)
}

func (c *CategoriaController) Delete() {
	handleDelete(&c.Controller, models.DeleteCategoria)
}
