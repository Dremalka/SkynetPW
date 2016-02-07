package main

import (
	"strings"
	"errors"
	"github.com/jbrodriguez/mlog"
	"net/url"
	"fmt"
	"net/http"
	"bytes"
	"strconv"
	"io/ioutil"
	"encoding/xml"
)

// getUidAndToken Функция передает на сервер mail.ru логин и пароль и получает uid и token для авторизации на игровом сервере
func getUidAndToken(login, password string) ([]byte, []byte, []byte, error) {
	mlog.Trace("Функция: getUidAndToken")
	var err error

	// проверить корректность логина
	split := strings.Split(login, "@")
	if len(split) < 2 {
		err = errors.New("Неправильный логин. Должен быть указан e-mail.")
		mlog.Error(err)
		return nil, nil, nil, err
	}

	domain := split[1]
	mailDomains := []string{"mail.ru", "inbox.ru", "bk.ru", "list.ru"}
	foundMailDomain := false
	for _, v := range mailDomains {
		if v == domain {
			foundMailDomain = true
			break
		}
	}

	var uid, uid2, token string
	if foundMailDomain {
		mlog.Trace("Логин от mail.ru")
		// TODO обработка ситуации, когда e-mail заведен в mail.ru
		err = errors.New("Отсутствует обработка получения uid и токена по email-у от mail.ru")
		mlog.Error(err)
		return nil, nil, nil, err
	} else {
		apiURL := "http://authdl.mail.ru"
		resource := "/sz.php"
		data := url.Values{}
		params := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?><AutoLogin ProjectId="61" SubProjectId="0" ShardId="0" Username="%s" Password="%s"/>`, login, password)
		u, _ := url.ParseRequestURI(apiURL)
		u.Path = resource
		urlStr := fmt.Sprintf("%v", u)

		client := &http.Client{}
		r, _ := http.NewRequest("POST", urlStr, bytes.NewBufferString(params)) // <-- URL-encoded payload
		r.Header.Add("User-Agent", "Downloader/4260")
		r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
		r.Header.Add("Accept-Encoding", "identity")
		resp, _ := client.Do(r)

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, nil, nil, err
		}

		type AuthPers struct {
			XMLName xml.Name `xml:"AutoLogin"`
			UID2    string   `xml:"PersId,attr"`
			Token   string   `xml:"Key,attr"`
		}
		var q AuthPers
		xml.Unmarshal(body, &q)
		uid2 = q.UID2
		token = q.Token

		//
		params = fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?><PersList ProjectId="61" SubProjectId="0" ShardId="0" Username="%s" Password="%s"/>`, login, password)
		r, _ = http.NewRequest("POST", urlStr, bytes.NewBufferString(params)) // <-- URL-encoded payload
		r.Header.Add("User-Agent", "Downloader/4260")
		r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
		r.Header.Add("Accept-Encoding", "identity")
		resp, _ = client.Do(r)

		body, err = ioutil.ReadAll(resp.Body)
		//fmt.Println(string(body))

		type Pers struct {
			XMLName xml.Name `xml:"Pers"`
			ID      string   `xml:"Id,attr"`
			Title   string   `xml:"Title,attr"`
			Cli     string   `xml:"Cli,attr"`
		}

		type PersList struct {
			XMLName  xml.Name `xml:"PersList"`
			PersID   string   `xml:"PersId,attr"`
			PersList []Pers   `xml:"Pers"`
		}

		var q1 PersList
		xml.Unmarshal(body, &q1)
		if len(q1.PersList) == 0 {
			err = errors.New("У учетной записи нет игровых аккаунтов.")
			mlog.Error(err)
			return nil, nil, nil, err
		}
		uid = q1.PersList[0].ID
	}
	return []byte(uid), []byte(uid2), []byte(token), nil
}
