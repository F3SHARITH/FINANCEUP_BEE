package controllers

import (
	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

type TipoIngresoController struct {
	beego.Controller
}

func (c *TipoIngresoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *TipoIngresoController) Post() {
	handlePost(&c.Controller, models.AddTipoIngreso)
}

func (c *TipoIngresoController) GetOne() {
	handleGetOne(&c.Controller, models.GetTipoIngresoById)
}

func (c *TipoIngresoController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllTipoIngreso)
}

func (c *TipoIngresoController) Put() {
	handlePut(&c.Controller, &models.TipoIngreso{Id: pathID(&c.Controller)}, models.UpdateTipoIngresoById)
}

func (c *TipoIngresoController) Delete() {
	handleDelete(&c.Controller, models.DeleteTipoIngreso)
}
