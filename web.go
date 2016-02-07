package main

import (
	"github.com/jbrodriguez/mlog"
	"github.com/labstack/echo"
	mid "github.com/labstack/echo/middleware"
	"net/http"
	"strconv"
	"fmt"
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
		e.Put("/api/bot/:id", updateBot)   // обновить бота
		e.Put("/api/bot/:id/connect", connectToServer)   // подключить бота к серверу
		e.Put("/api/bot/:id/disconnect", disconnectFromServer)   // отключить бота от сервера
		e.Delete("/api/bot/:id", deleteBot) // удалить бота

		e.Run(":8080")
	}()
}

// listBot сформировать и выдать в формате json список текущих ботов
func listBot(c *echo.Context) error {
	mlog.Trace("Функция: listBot")
	var err error
	list, err := lstBot()	// запросить список ботов
	mlog.Trace(fmt.Sprintf("Функция: listBot. list = %v, err = %v", list, err))
	if err != nil {
		mlog.Error(err)
		return c.JSON(http.StatusConflict, nil)
	}
	return c.JSON(http.StatusOK, list)
}

// createBot Создать бота и вернуть присвоенный ему ID
func createBot(c *echo.Context) error {
	mlog.Trace("Функция: createBot")
	id, err := newBot()
	if err != nil {
		mlog.Trace("Функция newBot вернула ошибку.")
		return c.String(http.StatusOK, "0")
	}
	mlog.Trace("Функция: createBot. id =", id)
	return c.String(http.StatusOK, strconv.Itoa(id))
}

// deleteBot Получить из контехта ID и передать функции удаления. По завершению уведомить, что бот удален
func deleteBot(c *echo.Context) error {
	mlog.Trace("Функция: deleteBot")
	id, err := strconv.Atoi(c.Param("id"))	// преобразовать в int
	if err != nil {
		mlog.Error(err)
	}else{
		err = delBot(id)	// передать в функцию удаления
		if err != nil {
			mlog.Error(err)
		}
	}
	return c.String(http.StatusOK, "ok")	// сообщить, что бот удален
}

func updateBot(c *echo.Context) error {
	mlog.Trace("Функция: updateBot")
	id, err := strconv.Atoi(c.Param("id"))	// преобразовать в int
	if err != nil {
		mlog.Error(err)
	}else{
		inf := make(map[string]interface{})
		// TODO перед заполнением проверить наличие переданных значений. И заполнять только те, которые пришли с запросом
		inf["id"] = id
		inf["name"] = c.Form("name")
		inf["login"] = c.Form("login")
		inf["password"] = c.Form("password")

		mlog.Trace("Функция: updateBot. Получены данные ", inf)
		if err := updBot(id, inf); err != nil {
			mlog.Error(err)
		}
	}
	return c.String(http.StatusOK, "ok")	// сообщить, что бот изменен
}

func connectToServer(c *echo.Context) error {
	mlog.Trace("Функция: connectToServer")
	id, err := strconv.Atoi(c.Param("id"))	// преобразовать в int
	if err != nil {
		mlog.Error(err)
	}else{
		if err := connectBotToServer(id); err != nil {
			mlog.Error(err)
		}
	}
	return c.String(http.StatusOK, "ok")	// сообщить, что бот изменен
}

func disconnectFromServer(c *echo.Context) error {
	mlog.Trace("Функция: disconnectFromServer")
	id, err := strconv.Atoi(c.Param("id"))	// преобразовать в int
	if err != nil {
		mlog.Error(err)
	}else{
		if err := disconnectBotFromServer(id); err != nil {
			mlog.Error(err)
		}
	}
	return c.String(http.StatusOK, "ok")	// сообщить, что бот изменен
}