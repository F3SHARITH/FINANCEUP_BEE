package controllers

import (
	"finanzas/models"

	beego "github.com/beego/beego/v2/server/web"
)

type NivelRiesgoController struct {
	beego.Controller
}

func (c *NivelRiesgoController) URLMapping() {
	c.Mapping("Post", c.Post)
	c.Mapping("GetOne", c.GetOne)
	c.Mapping("GetAll", c.GetAll)
	c.Mapping("Put", c.Put)
	c.Mapping("Delete", c.Delete)
}

func (c *NivelRiesgoController) Post() {
	handlePost(&c.Controller, models.AddNivelRiesgo)
}

func (c *NivelRiesgoController) GetOne() {
	handleGetOne(&c.Controller, models.GetNivelRiesgoById)
}

func (c *NivelRiesgoController) GetAll() {
	handleGetAll(&c.Controller, models.GetAllNivelRiesgo)
}

func (c *NivelRiesgoController) Put() {
	id, valid := pathID(&c.Controller)
	if !valid {
		return
	}
	handlePut(&c.Controller, &models.NivelRiesgo{Id: id}, models.UpdateNivelRiesgoById)
}

func (c *NivelRiesgoController) Delete() {
	handleDelete(&c.Controller, models.DeleteNivelRiesgo)
}
