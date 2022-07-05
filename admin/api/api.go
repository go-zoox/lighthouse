package api

import (
	"strconv"

	"github.com/go-zoox/zoox"
	"gorm.io/gorm"
)

func database() *gorm.DB {
	return GetDB()
}

func List() zoox.HandlerFunc {
	return func(ctx *zoox.Context) {
		type ListQuery struct {
			Page     string `json:"page,omitempty"`
			PageSize string `json:"page_size,omitempty"`
		}

		var query ListQuery
		if err := ctx.BindQuery(&query); err != nil {
			ctx.JSON(400, zoox.H{
				"code":    400,
				"message": err.Error(),
				"result":  nil,
			})
			return
		}

		page := 1
		pageSize := 10
		if query.Page != "" {
			page, _ = strconv.Atoi(query.Page)
		}
		if query.PageSize != "" {
			pageSize, _ = strconv.Atoi(query.PageSize)
		}

		total, dnss := DBList(page, pageSize)

		ctx.JSON(200, zoox.H{
			"code":    200,
			"message": "",
			"result": zoox.H{
				"total": total,
				"data":  dnss,
			},
		})
	}
}

func Create() zoox.HandlerFunc {
	return func(ctx *zoox.Context) {
		var dns DNS
		if err := ctx.BindJSON(&dns); err != nil {
			panic(err)
		}

		DBCreate(&DNS{
			Type:        dns.Type,
			Host:        dns.Host,
			Value:       dns.Value,
			IsActive:    true,
			Description: dns.Description,
		})

		ctx.Success(nil)
	}
}

func Update() zoox.HandlerFunc {
	return func(ctx *zoox.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}

		var dnsDTO DNS
		if err := ctx.BindJSON(&dnsDTO); err != nil {
			panic(err)
		}

		if dnsOrigin := DBRetrieve(id); dnsOrigin != nil {
			panic(err)
		} else {
			dnsOrigin.Host = dnsDTO.Host
			dnsOrigin.Value = dnsDTO.Value
			dnsOrigin.IsActive = dnsDTO.IsActive
			dnsOrigin.Description = dnsDTO.Description
			dnsOrigin.Type = dnsDTO.Type

			DBUpdate(dnsOrigin)
		}

		ctx.Success(nil)
	}
}

func Delete() zoox.HandlerFunc {
	return func(ctx *zoox.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			panic(err)
		}

		DBDelete(id)

		ctx.Success(nil)
	}
}
