package gdata

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
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
func (r Repository) GetGfiles(req *http.Request) GfileResults {
	results := GfileResults{}
	results.Ret = 0
	session, err := mgo.Dial(SERVER)
	if err != nil {
		results.Ret = 1
		results.Msg = "Failed to establish connection to Mongo server:" + err.Error()
		results.Gfiles = make([]Gfile, 0)
		fmt.Println("Failed to establish connection to Mongo server:", err)
		return results
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	filter := bson.M{}
	req.ParseForm()
	fmt.Println("Form: ", req.Form)
	gf := Gfile{}
	st := reflect.TypeOf(gf)
	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		if jsonKey, ok := field.Tag.Lookup("json"); ok {
			if index := strings.IndexByte(jsonKey, ','); index > 0 {
				jsonKey = jsonKey[:index]
			}
			if jsonKey != "" {
				// fmt.Println("json key: ", jsonKey, " type: ", field.Type)
				if formval, haskey := req.Form[jsonKey]; haskey {
					switch field.Type.String() {
					case "string", "bool", "int", "uint32":
						filter[jsonKey] = formval[0]
					case "bson.ObjectId":
						if bson.IsObjectIdHex(formval[0]) {
							filter[jsonKey] = bson.ObjectIdHex(formval[0])
						}
					}
				}
			}
		}
	}
	fmt.Println("filter: ", filter)

	if err := c.Find(filter).All(&results.Gfiles); err != nil {
		results.Ret = 2
		results.Msg = err.Error()
		fmt.Println("Failed to load results:", err)
	}
	if results.Gfiles == nil {
		results.Gfiles = make([]Gfile, 0)
	}
	return results
}

// AddGfile inserts an gfile in the DB
func (r Repository) AddGfile(gfile Gfile) (bool, string) {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	gfile.ID = bson.NewObjectId()
	var now = time.Now()
	gfile.CreationTime = now
	gfile.ModificationTime = now
	session.DB(DBNAME).C(DOCNAME).Insert(gfile)
	if err != nil {
		log.Fatal(err)
		return false, ""
	}
	return true, gfile.ID.Hex()
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

// DeleteGfile deletes an Gfile (not used for now)
func (r Repository) DeleteGfile(id string) string {
	session, err := mgo.Dial(SERVER)
	defer session.Close()
	// Verify id is ObjectId, otherwise bail
	if !bson.IsObjectIdHex(id) {
		return "NOT FOUND"
	}
	// Grab id
	oid := bson.ObjectIdHex(id)
	// Remove gfile
	if err = session.DB(DBNAME).C(DOCNAME).RemoveId(oid); err != nil {
		log.Panic(err)
		return "INTERNAL ERR"
	}
	// Write status
	return "OK"
}
