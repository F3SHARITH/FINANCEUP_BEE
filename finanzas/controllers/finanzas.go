package controllers

import (
	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

type FinanzasController struct {
	beego.Controller
}

func (c *FinanzasController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *FinanzasController) Post() {
	handlePost(&c.Controller, models.AddFinanzas)
}

func (c *FinanzasController) GetOne() {
	handleGetOne(&c.Controller, models.GetFinanzasById)
}

func (c *FinanzasController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllFinanzas)
}

func (c *FinanzasController) Put() {
	handlePut(&c.Controller, &models.Finanzas{Id: pathID(&c.Controller)}, models.UpdateFinanzasById)
}

func (c *FinanzasController) Delete() {
	handleDelete(&c.Controller, models.DeleteFinanzas)
}
