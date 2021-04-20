package model

import (
	json "github.com/json-iterator/go"
	"time"
)

type FileInfo struct {
	Fhash         string    `gorm:"column:fhash;primary_key"` // 主键
	Title         string    `gorm:"column:title"`
	Owners        string    `gorm:"column:owners"`
	Uploader      string    `gorm:"column:uploader"`
	Size          int64     `gorm:"column:size"`
	AuthorityCode int       `gorm:"column:authority_code"`
	Note          string    `gorm:"column:note"`
	CreateTime    time.Time `gorm:"-"`
	UpdateTime    time.Time `gorm:"-"`
	DeletedAt     time.Time `gorm:"-"`
}

func (f *FileInfo) OwnersMarshal(Owners map[string]ClientInfo) {
	OwnersByte, _ := json.Marshal(Owners)
	f.Owners = string(OwnersByte)
}

func (f *FileInfo) OwnersUnmarshal() (map[string]ClientInfo, error) {
	ret := map[string]ClientInfo{}
	err := json.Unmarshal([]byte(f.Owners), &ret)
	return ret, err
}
