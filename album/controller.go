package album

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gopkg.in/mgo.v2/bson"
)

// Controller is the Album controller
type Controller struct {
	Repository Repository
}

// Index return a list of albums
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	albums := c.Repository.GetAlbums()
	log.Println(albums)
	data, _ := json.Marshal(albums)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// GetAlbum returns an album
func (c *Controller) GetAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	album, err := c.Repository.GetAlbum(vars["id"])
	if err != nil {
		log.Fatalln("Error GetAlbum", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(album)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// AddAlbum adds an album
func (c *Controller) AddAlbum(w http.ResponseWriter, r *http.Request) {
	var album Album
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error AddAlbum", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error AddAlbum", err)
	}
	if err := json.Unmarshal(body, &album); err != nil {
		log.Fatalln("Error AddAlbum", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	newAlbum, err := c.Repository.AddAlbum(album)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(newAlbum)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(data)
	return
}

// UpdateAlbum updates an album
func (c *Controller) UpdateAlbum(w http.ResponseWriter, r *http.Request) {
	var album Album
	vars := mux.Vars(r)
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error UpdateAlbum", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error UpdateAlbum", err)
	}
	if err := json.Unmarshal(body, &album); err != nil {
		log.Fatalln("Error UpdateAlbum", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	album.ID = bson.ObjectIdHex(vars["id"])
	newAlbum, err := c.Repository.UpdateAlbum(album)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	data, _ := json.Marshal(newAlbum)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// DeleteAlbum deletes an album
func (c *Controller) DeleteAlbum(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if success := c.Repository.DeleteAlbum(vars["id"]); !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}
