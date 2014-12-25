/**
* Providing DB interaction functions for learning object.
 */

package models

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/jmcvetta/neoism"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"io"
	"log"
)

// Create node on graph and creeate the relationships
func CreateGraphNodeAndRelationships(m map[string]interface{}, id string) error {

	// Connect to the graph database server from remote
	db, err := neoism.Connect("http://54.187.83.59:7474/db/data")
	if err != nil {
		return err
	}
	// Create a node with a Cypher query
	res0 := []struct {
		N neoism.Node // Column "n" gets automagically unmarshalled into field N
	}{}
	cq0 := neoism.CypherQuery{
		Statement: "CREATE (n:LearningObject {tile: {tile},_id: {_id}}) RETURN n",
		// Use parameters instead of constructing a query string
		Parameters: neoism.Props{"tile": m["title"], "_id": id},
		Result:     &res0,
	}
	db.Cypher(&cq0)
	n1 := res0[0].N // Only one row of data returned
	n1.Db = db      // Must manually set Db with objects returned from Cypher query
	//
	// Create a relationship
	prerequisites := m["prerequisites"].([]interface{})

	// Create relationship if has
	res := []struct {
		N neoism.Node // Column "n" gets automagically unmarshalled into field N
	}{}
	for _, id := range prerequisites {
		cq := neoism.CypherQuery{
			Statement:  `MATCH (n {_id:{_id}}) RETURN n`,
			Parameters: neoism.Props{"_id": id},
			Result:     &res,
		}
		db.Cypher(&cq)
		fmt.Println(id)
		if len(res) != 0 {
			n2 := res[0].N
			n2.Db = db
			// Create relationship
			n2.Relate("prerequisite for", n1.Id(), neoism.Props{}) // Empty Props{} is okay
		}
	}

	return nil
}

// Save learning object to mongodb
func SaveLODocObj(m map[string]interface{}) (string, error) {
	// Connect to mongodb server from remote.
	session, err := mgo.Dial("adaptivelearner:81hocyupang@54.187.83.59/learningobjects")
	if err != nil {
		return "", err
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	// Use database learningobecjts and select learningobjects collection.
	c := session.DB("learningobjects").C("learningobjects")
	// Insert one learning object.
	uuid, err := newUUID()
	if err != nil {
		return "", err
	}
	m["_id"] = uuid
	err = c.Insert(m)
	if err != nil {
		return "", err
		log.Fatal(err)
	} else {
		return uuid, nil
	}

	return "", err
}

// newUUID generates a random UUID according to RFC 4122
func newUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

// Reture ids of learnning objects
func GetLearningObjectsIds() ([]byte, error) {
	// Connect to mongodb server from remote.
	session, err := mgo.Dial("adaptivelearner:81hocyupang@54.187.83.59/learningobjects")
	if err != nil {
		return nil, err
	}
	defer session.Close()
	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)
	// Use database learningobecjts and select learningobjects collection.
	c := session.DB("learningobjects").C("learningobjects")

	//query := c.Find(nil).Select(bson.M{"_id": 1},"title":1)
	var results []interface{}
	c.Find(bson.M{}).All(&results)

	js, err := json.Marshal(results)
	if err != nil {
		return nil, err
	}

	return js, nil

}
