package model

import (
	"github.com/PlagueCat-Miao/goipfs-lab511/constdef"
	"time"
)

type ClientInfo struct {
	Dhash            string              `form:"dhash" json:"dhash" xml:"dhash" binding:"required"`
	Status           constdef.UserStatus `form:"status" json:"status" xml:"status" binding:"required"`
	Ip               string              `form:"ip" json:"ip" xml:"ip" binding:"required"`
	Port             int                 `form:"port" json:"port" xml:"port" binding:"required"`
	Capacity         int64              `form:"capacity" json:"capacity" xml:"capacity" binding:"required"`
	Remain           int64              `form:"remain" json:"remain" xml:"remain" binding:"required"`
	LastPingPongTime time.Time           `json:"lastpingpongtime,omitempty"`
}