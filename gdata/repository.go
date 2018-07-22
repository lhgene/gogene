package gdata

import (
	"fmt"
	"log"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//Repository ...
type Repository struct{}

// SERVER the DB server
const SERVER = "localhost:27017"

// DBNAME the name of the DB instance
const DBNAME = "gene"

// DOCNAME the name of the document
const DOCNAME = "gfile"

// GetGfiles returns the list of gfile
func (r Repository) GetGfiles() Gfiles {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to establish connection to Mongo server:", err)
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Gfiles{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to write results:", err)
	}
	return results
}

// AddGfile inserts an gfile in the DB
func (r Repository) AddGfile(gfile Gfile) bool {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	gfile.ID = bson.NewObjectId()
	var now = time.Now()
	gfile.CreationTime = now
	gfile.ModificationTime = now
	session.DB(DBNAME).C(DOCNAME).Insert(gfile)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

// // UpdateAlbum updates an Album in the DB (not used for now)
// func (r Repository) UpdateAlbum(album Album) bool {
// 	session, err := mgo.Dial(SERVER)
// 	defer session.Close()
// 	session.DB(DBNAME).C(DOCNAME).UpdateId(album.ID, album)
// 	if err != nil {
// 		log.Fatal(err)
// 		return false
// 	}
// 	return true
// }

// // DeleteAlbum deletes an Album (not used for now)
// func (r Repository) DeleteAlbum(id string) string {
// 	session, err := mgo.Dial(SERVER)
// 	defer session.Close()
// 	// Verify id is ObjectId, otherwise bail
// 	if !bson.IsObjectIdHex(id) {
// 		return "NOT FOUND"
// 	}
// 	// Grab id
// 	oid := bson.ObjectIdHex(id)
// 	// Remove user
// 	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
// 		log.Fatal(err)
// 		return "INTERNAL ERR"
// 	}
// 	// Write status
// 	return "OK"
// }
