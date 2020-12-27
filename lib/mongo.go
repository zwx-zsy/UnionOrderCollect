package lib

import (
	"gopkg.in/mgo.v2"
	"log"
	"time"
)

const (
	MaxCon       = 300
)

type CronSession struct {
	*mgo.Session
	Csession *mgo.Session
}
func (c *CronSession)NewClient() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{ServerConf.DBConf.Host+":"+ServerConf.DBConf.Port},
		Timeout:  60 * time.Second,
		Database: ServerConf.DBConf.DatabaseName,
		Source:   ServerConf.DBConf.AuthDBName,
		Username: ServerConf.DBConf.User,
		Password: ServerConf.DBConf.Password,
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		log.Fatalf("CreateSession failed:%\n", err)
	}
	//设置连接池的大小
	session.SetPoolLimit(MaxCon)
	defer session.Close()
	c.Csession = session.Clone()
}

func (c *CronSession)Collection(collection string) *mgo.Collection {
	return c.Csession.DB(ServerConf.DBConf.AuthDBName).C(collection)
}

