package controllers

import (
	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

type TipoIngresoInversionController struct {
	beego.Controller
}

func (c *TipoIngresoInversionController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *TipoIngresoInversionController) Post() {
	handlePost(&c.Controller, models.AddTipoIngresoInversion)
}

func (c *TipoIngresoInversionController) GetOne() {
	handleGetOne(&c.Controller, models.GetTipoIngresoInversionById)
}

func (c *TipoIngresoInversionController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllTipoIngresoInversion)
}

func (c *TipoIngresoInversionController) Put() {
	handlePut(&c.Controller, &models.TipoIngresoInversion{Id: pathID(&c.Controller)}, models.UpdateTipoIngresoInversionById)
}

func (c *TipoIngresoInversionController) Delete() {
	handleDelete(&c.Controller, models.DeleteTipoIngresoInversion)
}
