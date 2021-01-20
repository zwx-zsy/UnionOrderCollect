package lib

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

func UserTask() {
	var RedisClientObj RedisClient
	RedisClientObj.NewRedisClient()
	var NewClient CronSession
	NewClient.NewClient()
	CronUser(RedisClientObj.RClient, NewClient.Csession, time.Now().UTC())
}

func CronUser(r *redis.Client, s *mgo.Session, t time.Time) {
	defer r.Close()
	hasMore, err := UpdateUser(r, s)
	if err != nil {
		log.Fatal(err)
	}
	if hasMore {
		CronUser(r, s, t)
	}
}

func UpdateUser(r *redis.Client, s *mgo.Session) (hasMore bool, err error) {
	xMessageSliceCmd := r.XRange("newUser", "-", "+")
	fmt.Println("xMessageSliceCmd ====>:", xMessageSliceCmd)
	for _, message := range xMessageSliceCmd.Val() {
		updated, err := UpdateFansCount(s, message.Values["updateFansCount"])
		if err != nil {
			continue
		}
		if updated {
			//从stream中删除任务
			r.XDel("newUser", message.ID)
		}
	}
	xLen := r.XLen("newUser")
	if xLen.Val() > 0 {
		return true, err
	} else {
		return false, err
	}

}

type NewUserValue struct {
	MemberGradeId string `json:"memberGradeId"`
	Id            string `json:"Id"`
}

func UpdateFansCount(s *mgo.Session, userinfo interface{}) (updated bool, err error) {
	defer s.Close()
	var newUserValue NewUserValue
	marshal, err := json.Marshal(userinfo)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	err = json.Unmarshal(marshal, &newUserValue)
	if err != nil {
		log.Fatal(err)
		return false, err
	}
	collection := s.DB(ServerConf.DBConf.DatabaseName).C("JdUnion_User")
	CurrentUserObj := struct {
		InviterId string `json:"inviterId" bson:"inviterId"`
	}{}
	err = collection.FindId(bson.ObjectIdHex(newUserValue.Id)).One(&CurrentUserObj)
	if err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return false, err
	}
	if CurrentUserObj.InviterId == "" {
		fmt.Println(err)
		return false, err
	}
	err = collection.UpdateId(bson.ObjectIdHex(CurrentUserObj.InviterId), bson.M{"$inc": bson.M{"fansCount": 1}})
	if err != nil {
		fmt.Println(err)
		return false, err
	}
	return true, err
}
