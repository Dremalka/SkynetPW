package main

import (
	"github.com/jbrodriguez/mlog"
	"errors"
)

type bot struct {
	ID int
}

var mbot map[int]bot

// newBot Создать нового бота и запомнить его в map
func newBot() (int, error) {
	var err error
	mlog.Trace("Функция: newBot")
	// при необходимости иницилизировать map
	if mbot == nil {
		mbot = make(map[int]bot)
	}

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
