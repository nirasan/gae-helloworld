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
	r.HandleFunc("/entry/put", EntryPutHandler)
	r.HandleFunc("/entry/get", EntryGetHandler)
	return r
}

type Entry struct {
	Title string
	Body string
	CreatedTime time.Time
}

func EntryPutHandler(w http.ResponseWriter, r *http.Request) {
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
		fmt.Fprintln(w, "put entry: IntID=", key.IntID())
	}
}

func EntryGetHandler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)
	id, _ := strconv.ParseInt(r.FormValue("id"), 10, 64)
	key := datastore.NewKey(ctx, "Entry", "", id, nil)
	var e Entry
	if err := datastore.Get(ctx, key, &e); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	} else {
		fmt.Fprintln(w, "title: ", e.Title, " body: ", e.Body)
	}
}
