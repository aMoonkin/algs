package main

import (
	"errors"
	"net/http"
)

func session(w http.ResponseWriter, r *http.Request) (sess data.Session, err error) {
	Cookie, err := r.Cookie("_cookit")
	if err == nil {
		sess = data.Session{Uuid: cookie.Value}
		if ok, _ := sess.Check(); !ok {
			err = errors.New("Invalid session")
		}
	}
	return
}
