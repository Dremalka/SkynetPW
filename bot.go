package main

import (
	"github.com/jbrodriguez/mlog"
	"errors"
	"fmt"
	"strconv"
)

type bot struct {
	ID int
	Name string
}

var mbot map[int]bot

func init() {
	// при необходимости иницилизировать map
	if mbot == nil {
		mbot = make(map[int]bot)
	}
}

// newBot Создать нового бота и запомнить его в map
func newBot() (int, error) {
	var err error
	mlog.Trace("Функция: newBot")

	// найти новый незанятый id
	id := 0
	mlog.Trace("Функция: newBot. Поиск нового незанятого id")
	for i := 1; i < 10000; i++ {
		_, ok := mbot[i]
		if !ok {
			id = i
			break
		}
	}
	if id == 0 {	// превышен индекс
		err = errors.New("Превышено количество индексов ботов. id > 999")
		mlog.Error(err)
	}else{
		b := bot{}	// инициализировать бота
		b.ID = id
		b.Name = fmt.Sprintf("Бот %s", strconv.Itoa(id))
		mlog.Trace("Функция: newBot. Новый бот.", b)
		mbot[id] = b
	}

	return id, err
}

// delBot Найти бота по идентификатору и удалить
func delBot(id int) error {
	mlog.Trace("Функция: delBot")
	_, ok := mbot[id]
	if ok {
		delete(mbot, id)
		mlog.Trace("Функция: delBot. Бот с id = %d удален.", id)
	}
	return nil
}

// infBot Служебная структура для обмена данными с веб-интерфейсом
type infBot struct {
	ID int
	Name string
}

// lstBot Функция подготавливает и возвращает массив с информацией по текущим ботам
func lstBot() ([]infBot, error) {
	mlog.Trace("Функция: lstBot")
	list := make([]infBot, len(mbot))
	i := 0
	for id, bot := range mbot {
		inf := infBot{}
		inf.ID = id
		inf.Name = bot.Name
		list[i] =inf
		i++
	}
	return list, nil
}

// updBot Функция обновляет информацию бота по указанному идентификатору
func updBot(id int, inf infBot) error {
	mlog.Trace("Функция: updBot")
	var err error
	bot, ok := mbot[id]
	if !ok {
		err = errors.New("В mbot нет бота с указанным индексом.")
		mlog.Error(err)
		return err
	}
	bot.Name = inf.Name
	mbot[id] = bot
	mlog.Trace("Функция: updBot. Обновлена информация бота с id = %d.", id)
	return nil
}