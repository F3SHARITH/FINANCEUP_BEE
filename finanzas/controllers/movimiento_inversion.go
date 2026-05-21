package controllers

import (
	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

type MovimientoInversionController struct {
	beego.Controller
}

func (c *MovimientoInversionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *MovimientoInversionController) Post() {
	handlePost(&c.Controller, models.AddMovimientoInversion)
}

func (c *MovimientoInversionController) GetOne() {
	handleGetOne(&c.Controller, models.GetMovimientoInversionById)
}

func (c *MovimientoInversionController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllMovimientoInversion)
}

func (c *MovimientoInversionController) Put() {
	handlePut(&c.Controller, &models.MovimientoInversion{Id: pathID(&c.Controller)}, models.UpdateMovimientoInversionById)
}

func (c *MovimientoInversionController) Delete() {
	handleDelete(&c.Controller, models.DeleteMovimientoInversion)
}
