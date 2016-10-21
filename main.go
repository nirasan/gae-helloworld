package gae_helloworld

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"time"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine"
	"strconv"
)

func CreateHandler() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/entry/create", EntryCreateHandler)
	r.HandleFunc("/entry/show", EntryShowHandler)
	r.HandleFunc("/entry/update", EntryUpdateHandler)
	r.HandleFunc("/entry/delete", EntryDeleteHandler)
	return r
}

type Entry struct {
	Title string
	Body string
	CreatedTime time.Time
}

func EntryCreateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	key := datastore.NewIncompleteKey(ctx, "Entry", nil)
	e := &Entry{
		Title: r.FormValue("t"),
		Body: r.FormValue("b"),
	}
	if key, err := datastore.Put(ctx, key, e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintln(w, "entry created: IntID=", key.IntID())
	}
}

func EntryShowHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
	key := datastore.NewKey(ctx, "Entry", "", id, nil)
	var e Entry
	if err := datastore.Get(ctx, key, &e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintln(w, "entry show: title=", e.Title, " body=", e.Body)
	}
}

func EntryUpdateHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
	key := datastore.NewKey(ctx, "Entry", "", id, nil)
	var e Entry
	if err := datastore.Get(ctx, key, &e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	e.Title = r.FormValue("t")
	e.Body = r.FormValue("b")
	if key, err := datastore.Put(ctx, key, &e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintln(w, "entyr updated: IntID=", key.IntID())
	}
}

func EntryDeleteHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
	key := datastore.NewKey(ctx, "Entry", "", id, nil)
	if err := datastore.Delete(ctx, key); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintln(w, "entyr deleted: IntID=", key.IntID())
	}
}