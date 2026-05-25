package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/client/orm"
	beego "github.com/beego/beego/v2/server/web"
)

type apiResponse map[string]interface{}

func ok(c *beego.Controller, message string, data interface{}) {
	c.Ctx.Output.SetStatus(http.StatusOK)
	c.Data["json"] = apiResponse{"success": true, "status": http.StatusOK, "message": message, "data": data}
	c.ServeJSON()
}

func created(c *beego.Controller, data interface{}) {
	c.Ctx.Output.SetStatus(http.StatusCreated)
	c.Data["json"] = apiResponse{"success": true, "status": http.StatusCreated, "message": "Peticion exitosa", "data": data}
	c.ServeJSON()
}

func fail(c *beego.Controller, status int, message string, err error) {
	body := apiResponse{"success": false, "status": status, "message": message}
	if err != nil {
		body["error"] = err.Error()
	}
	c.Ctx.Output.SetStatus(status)
	c.Data["json"] = body
	c.ServeJSON()
}

func statusFromError(err error) int {
	if err == nil {
		return http.StatusOK
	}
	msg := strings.ToLower(err.Error())
	if errors.Is(err, orm.ErrNoRows) || strings.Contains(msg, "no row found") {
		return http.StatusNotFound
	}
	if strings.Contains(msg, "violates foreign key") || strings.Contains(msg, "foreign key constraint") {
		return http.StatusConflict
	}
	return http.StatusBadRequest
}

func handlePost[T any](c *beego.Controller, add func(*T) (int64, error)) {
	var v T
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &v); err == nil {
		if _, err := add(&v); err == nil {
			c.Ctx.Output.SetStatus(201)
			c.Data["json"] = v
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func handleGetOne[T any](c *beego.Controller, get func(int) (*T, error)) {
	id, valid := pathID(c)
	if !valid {
		return
	}
	v, err := get(id)
	if err != nil {
		fail(c, statusFromError(err), "No se encontro el registro solicitado", err)
		return
	}
	ok(c, "Peticion exitosa", v)
}

func handleGetAll(c *beego.Controller, getAll func(map[string]string, []string, []string, []string, int64, int64) ([]interface{}, error)) {
	var fields []string
	var sortby []string
	var order []string
	query := make(map[string]string)
	var limit int64 = 10
	var offset int64

	if v := c.GetString("fields"); v != "" {
		fields = strings.Split(v, ",")
	}
	if v, err := c.GetInt64("limit"); err == nil {
		limit = v
	}
	if v, err := c.GetInt64("offset"); err == nil {
		offset = v
	}
	if v := c.GetString("sortby"); v != "" {
		sortby = strings.Split(v, ",")
	}
	if v := c.GetString("order"); v != "" {
		order = strings.Split(v, ",")
	}
	if v := c.GetString("query"); v != "" {
		for _, cond := range strings.Split(v, ",") {
			kv := strings.SplitN(cond, ":", 2)
			if len(kv) != 2 {
				fail(c, http.StatusBadRequest, "El parametro query debe tener formato campo:valor", nil)
				return
			}
			query[kv[0]] = kv[1]
		}
	}

	l, err := getAll(query, fields, sortby, order, offset, limit)
	if err != nil {
		c.Data["json"] = err.Error()
	} else {
		c.Data["json"] = l
	}
	c.ServeJSON()
}

func handlePut[T any](c *beego.Controller, v *T, update func(*T) error) {
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, v); err == nil {
		if err := update(v); err == nil {
			c.Data["json"] = "OK"
		} else {
			c.Data["json"] = err.Error()
		}
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}

func pathID(c *beego.Controller) (int, bool) {
	idStr := c.Ctx.Input.Param(":id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		fail(c, http.StatusBadRequest, "El id debe ser un numero entero positivo", err)
		return 0, false
	}
	return id, true
}

func handleDelete(c *beego.Controller, del func(int) error) {
	idStr := c.Ctx.Input.Param(":id")
	id, _ := strconv.Atoi(idStr)
	if err := del(id); err == nil {
		c.Data["json"] = "OK"
	} else {
		c.Data["json"] = err.Error()
	}
	c.ServeJSON()
}
