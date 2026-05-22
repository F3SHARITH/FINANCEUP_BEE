package controllers

import (
	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

type MovimientoMetaController struct {
	beego.Controller
}

func (c *MovimientoMetaController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *MovimientoMetaController) Post() {
	handlePost(&c.Controller, models.AddMovimientoMeta)
}

func (c *MovimientoMetaController) GetOne() {
	handleGetOne(&c.Controller, models.GetMovimientoMetaById)
}

func (c *MovimientoMetaController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllMovimientoMeta)
}

func (c *MovimientoMetaController) Put() {
	id, valid := pathID(&c.Controller)
	if !valid {
		return
	}
	handlePut(&c.Controller, &models.MovimientoMeta{Id: id}, models.UpdateMovimientoMetaById)
}

func (c *MovimientoMetaController) Delete() {
	handleDelete(&c.Controller, models.DeleteMovimientoMeta)
}
