package gdata

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Gfile represents a gene file
type Gfile struct {
	ID                 bson.ObjectId   `bson:"_id"`
	Filetype           string          `json:"filetype"`
	Filename           string          `json:"filename"`
	Filesize           uint32          `json:"filesize"`
	Path               string          `json:"path"`
	SourceFileIDs      []bson.ObjectId `json:"sourceFileIds"`
	CreationTime       time.Time       `json:"creationTime"`
	ModificationTime   time.Time       `json:"modificationTime"`
	TaskID             string          `json:"taskId"`
	SequencingType     string          `json:"sequencingType"`
	Probe              string          `json:"probe"`
	Sequencer          string          `json:"sequencer"`
	Platform           string          `json:"platform"`
	ReferenceGenome    string          `json:"referenceGenome"`
	DuplicationRemoved string          `json:"duplicationRemoved"`
	Bed                string          `json:"bed"`
	Commented          string          `json:"commented"`
}

//Gfiles is an array of Gfile
type Gfiles []Gfile
