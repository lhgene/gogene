package gdata

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//Gfile represents a gene file
type Gfile struct {
	ID                 bson.ObjectId   `json:"_id" bson:"_id"`
	Filetype           string          `json:"filetype,omitempty" bson:"filetype,omitempty"`
	Filename           string          `json:"filename,omitempty" bson:"filename,omitempty"`
	Filesize           uint32          `json:"filesize,omitempty" bson:"filesize,omitempty"`
	Path               string          `json:"path,omitempty" bson:"path,omitempty"`
	IP                 string          `json:"ip,omitempty" bson:"ip,omitempty"`
	SourceFileIDs      []bson.ObjectId `json:"sourceFileIds,omitempty" bson:"sourceFileIds,omitempty"`
	CreationTime       time.Time       `json:"creationTime" bson:"creationTime"`
	ModificationTime   time.Time       `json:"modificationTime" bson:"modificationTime"`
	TaskID             string          `json:"taskId,omitempty" bson:"taskId,omitempty"`
	SequencingType     string          `json:"sequencingType,omitempty" bson:"sequencingType,omitempty"`
	Probe              string          `json:"probe,omitempty" bson:"probe,omitempty"`
	Sequencer          string          `json:"sequencer,omitempty" bson:"sequencer,omitempty"`
	Platform           string          `json:"platform,omitempty" bson:"platform,omitempty"`
	ReferenceGenome    string          `json:"referenceGenome,omitempty" bson:"referenceGenome,omitempty"`
	DuplicationRemoved bool            `json:"duplicationRemoved,omitempty" bson:"duplicationRemoved,omitempty"`
	Bed                string          `json:"bed,omitempty" bson:"bed,omitempty"`
	Commented          bool            `json:"commented,omitempty" bson:"commented,omitempty"`
}

//Gfiles is an array of Gfile
// type Gfiles []Gfile

//GfileResults represents Gfile search result
type GfileResults struct {
	Ret    uint32  `json:"ret"`
	Msg    string  `json:"msg"`
	Gfiles []Gfile `json:"data"`
}
