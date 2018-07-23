package gdata

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

//Controller ...
type Controller struct {
	Repository Repository
}

// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {
	files := c.Repository.GetGfiles(r) // list of all gfiles
	log.Println(files)
	data, _ := json.Marshal(files)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// // CountAlbum /
// func (c *Controller) CountAlbum(w http.ResponseWriter, r *http.Request) {
// 	albums := c.Repository.GetAlbums() // list of all albums
// 	log.Println(albums)
// 	data, _ := json.Marshal(albums)
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(data)
// 	return
// }

// AddGfile POST /
func (c *Controller) AddGfile(w http.ResponseWriter, r *http.Request) {
	var gfile Gfile
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Panicln("Error AddGfile", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Panicln("Error AddGfile", err)
	}
	if err := json.Unmarshal(body, &gfile); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Panicln("Error AddGfile unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	success, newID := c.Repository.AddGfile(gfile) // adds the album to the DB
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	var buffer bytes.Buffer
	buffer.WriteString("{\"_id\":\"")
	buffer.WriteString(newID)
	buffer.WriteString("\"}")
	log.Println(buffer.String())
	w.Write(buffer.Bytes())
	return
}

// // UpdateAlbum PUT /
// func (c *Controller) UpdateAlbum(w http.ResponseWriter, r *http.Request) {
// 	var album Album
// 	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
// 	if err != nil {
// 		log.Fatalln("Error UpdateAlbum", err)
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	if err := r.Body.Close(); err != nil {
// 		log.Fatalln("Error AddaUpdateAlbumlbum", err)
// 	}
// 	if err := json.Unmarshal(body, &album); err != nil { // unmarshall body contents as a type Candidate
// 		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 		w.WriteHeader(422) // unprocessable entity
// 		if err := json.NewEncoder(w).Encode(err); err != nil {
// 			log.Fatalln("Error UpdateAlbum unmarshalling data", err)
// 			w.WriteHeader(http.StatusInternalServerError)
// 			return
// 		}
// 	}
// 	success := c.Repository.UpdateAlbum(album) // updates the album in the DB
// 	if !success {
// 		w.WriteHeader(http.StatusInternalServerError)
// 		return
// 	}
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.WriteHeader(http.StatusOK)
// 	return
// }

// DeleteGfile DELETE /
func (c *Controller) DeleteGfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]                                    // param id
	if err := c.Repository.DeleteGfile(id); err != "" { // delete a gfile by id
		if strings.Contains(err, "404") {
			w.WriteHeader(http.StatusNotFound)
		} else if strings.Contains(err, "500") {
			w.WriteHeader(http.StatusInternalServerError)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}
