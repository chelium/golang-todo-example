package db

import (
	mgo "gopkg.in/mgo.v2"
)

var mgoSession *mgo.Session
var mongoConnStr = "mongodb:27017"

// GetMongoSession creates a new session if mgoSession is nil i.e there is no active mongo session.
// If there is an active mongo session it will return a Clone
func GetMongoSession() (*mgo.Session, error) {
	if mgoSession == nil {
		var err error
		mgoSession, err = mgo.Dial(mongoConnStr)
		if err != nil {
			return nil, err
		}
	}
	return mgoSession.Clone(), nil
}
