package main

import (
	"github.com/jbrodriguez/mlog"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
	"net/http"
)

func CreateWeb() <-chan struct{} {
	ch := make(<-chan struct{})
	router()
	return ch
}

func router() {
	go func() {
		e := echo.New()

		e.Use(mid.Logger()) // выводить лог
		//e.Use(mid.Recover())	// игнорировать ошибки при работе сервера

		//api
		e.Get("/api/bots", listBot)         // вывести json-список текущих ботов
		e.Post("/api/bots", createBot)      // создать нового бота
		e.Delete("/api/bot/:id", deleteBot) // удалить бота

		e.Run(":8080")
	}()
}

func listBot(c *echo.Context) error {
	mlog.Trace("Функция: listBot")
	return c.String(http.StatusOK, "ok")
}

func createBot(c *echo.Context) error {
	mlog.Trace("Функция: createBot")
	return c.String(http.StatusOK, "ok")
}

func deleteBot(c *echo.Context) error {
	mlog.Trace("Функция: deleteBot")
	return c.String(http.StatusOK, "ok")
}
