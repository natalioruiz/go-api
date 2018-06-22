package album

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

// Repository ...
type Repository struct{}

// Constants
const (
	SERVER  = "127.0.0.1:27017"
	DBNAME  = "musicstore"
	DOCNAME = "albums"
)

// GetAlbums returns the list of albums
func (r Repository) GetAlbums() Albums {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to connecto to MongoDB")
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	results := Albums{}
	if err := c.Find(nil).All(&results); err != nil {
		fmt.Println("Failed to get results", err)
	}
	return results
}

// GetAlbum returns a single album
func (r Repository) GetAlbum(id string) (*Album, error) {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to connecto to MongoDB")
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	album := Album{}
	if err := c.FindId(bson.ObjectIdHex(id)).One(&album); err != nil {
		return nil, err
	}
	return &album, nil
}

// AddAlbum adds an album to the database
func (r Repository) AddAlbum(album Album) (*Album, error) {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to connecto to MongoDB")
	}
	defer session.Close()

	album.ID = bson.NewObjectId()
	if err := session.DB(DBNAME).C(DOCNAME).Insert(album); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return &album, nil
}

// UpdateAlbum updates an existing album
func (r Repository) UpdateAlbum(album Album) (*Album, error) {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to connecto to MongoDB")
	}
	defer session.Close()
	if err := session.DB(DBNAME).C(DOCNAME).UpdateId(album.ID, album); err != nil {
		return nil, err
	}
	return &album, nil
}

// DeleteAlbum deletes an existing album
func (r Repository) DeleteAlbum(id string) bool {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		fmt.Println("Failed to connecto to MongoDB")
	}
	defer session.Close()
	c := session.DB(DBNAME).C(DOCNAME)
	if err := c.RemoveId(bson.ObjectIdHex(id)); err != nil {
		return false
	}
	return true
}
