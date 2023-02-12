package model

import (
	"strconv"
	"time"
)

type EventSubject struct {
	Type            string `json:"type" bson:"type"`
	FullName        string `json:"full_name" bson:"full_name"`
	RoleName        string `json:"role_name" bson:"role_name"`
	ApplicationList string `json:"application_list" bson:"application_list"`
	IP              string `json:"ip" bson:"ip"`
	UID             string `json:"uid" bson:"uid"`
}

func GetSubjectFromUser(u *User) *EventSubject {
	if u != nil {
		return &EventSubject{
			Type:            "用户",
			FullName:        u.Username,
			RoleName:        "",
			ApplicationList: "",
			UID:             strconv.FormatUint(uint64(u.ID), 10),
		}
	}
	return &EventSubject{Type: "用户"}

}

func GetSystemSubject() *EventSubject {
	return &EventSubject{
		Type:     "系统",
		FullName: "系统",
	}
}

type EventTarget struct {
	Uuid string `json:"uuid" bson:"uuid"`
	Name string `json:"name" bson:"name"`
}

type Log struct {
	ID          string        `json:"id" bson:"_id,omitempty"`
	Tag         int           `json:"tag" bson:"tag"`
	CreatedAt   time.Time     `json:"created_at" bson:"created_at"`
	Type        string        `json:"type" bson:"type"`
	Subject     *EventSubject `json:"subject" bson:"subject"`
	Target      *EventTarget  `json:"target" bson:"target"`
	Result      *bool         `json:"result" bson:"result"`
	Reason      string        `json:"reason" bson:"reason"`
	Description string        `json:"description" bson:"description"`
	Detail      any           `json:"-"`
}
