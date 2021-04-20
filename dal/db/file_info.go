package db

import (
	"fmt"
	"github.com/PlagueCat-Miao/goipfs-lab511/model"
	"github.com/jinzhu/gorm"
	json "github.com/json-iterator/go"
	"strings"
)

var FileInfoDB *gorm.DB

type IIPFSFileInfoDB interface {
	CreateInfo(fileInfo *model.FileInfo) error
	GetInfoByFhash(fhash string) (*model.FileInfo, error)
	GetInfoList(offset, limit int64, queryFactor map[string]string) ([]*model.FileInfo, int64, error)
	ModifyInfo(queryFactor map[string]string, updateMap map[string]interface{}) error
	OwnerIncrByFhash(fhash string, newfileInfo *model.FileInfo) error
}

type IPFSFileInfoDB struct{}

func NewIPFSFileInfoDB() IIPFSFileInfoDB {
	return &IPFSFileInfoDB{}
}

func (i *IPFSFileInfoDB) CreateInfo(fileInfo *model.FileInfo) error {
	query := FileInfoDB.Debug()
	query = query.Create(fileInfo)
	return query.Error
}

func (i *IPFSFileInfoDB) GetInfoByFhash(fhash string) (*model.FileInfo, error) {
	ansFileInfo := &model.FileInfo{}
	query := FileInfoDB.Debug()
	query = query.Where("fhash = ?", fhash).First(ansFileInfo)
	return ansFileInfo, query.Error
}

// limit == 0代表没有输出限制
func (i *IPFSFileInfoDB) GetInfoList(offset, limit int64, queryFactor map[string]string) ([]*model.FileInfo, int64, error) {
	ret := []*model.FileInfo{}
	var total int64
	if limit < 0 || offset < 0 {
		return nil, 0, fmt.Errorf(" limit: %v || offset: %v invalid ", limit, offset)
	} else if limit == 0 && len(queryFactor) == 0 { //输出全部内容，存在风险，禁止
		return nil, 0, fmt.Errorf(" need queryFactor ")
	}
	//查询条件
	query := FileInfoDB.Debug()
	queryMap := queryFactorParse(queryFactor)
	for k, v := range queryMap {
		query = query.Where(k, v)
	}
	// 获取条件下的count
	query = query.Model(&ret).Count(&total)
	if query.Error != nil {
		return nil, 0, query.Error
	} else if total == 0 { //没有查到结果
		return nil, 0, nil
	} else if total <= offset {
		return nil, 0, fmt.Errorf(" offset > total:Cross the border,offset:%v total:%v  ", offset, total)
	}
	// 获取数据
	if limit != 0 {
		query = query.Limit(limit)
	}
	err := query.Offset(offset).Find(&ret).Error
	return ret, total, err
}

func (i *IPFSFileInfoDB) OwnerIncrByFhash(fhash string, newfileInfo *model.FileInfo) error {
	fileInfo := &model.FileInfo{}
	txquery := FileInfoDB.Debug().Begin()
	//查询条件
	var count int64
	txquery = txquery.Where("fhash = ?", fhash).Model(&model.FileInfo{}).Count(&count)
	if count == 0 {
		txquery = txquery.Create(newfileInfo)
		err := txquery.Error
		if err != nil {
			err = fmt.Errorf("[Create-err]:%v , newfileInfo:%+v", err, &newfileInfo)
			txquery.Rollback()
		} else {
			txquery.Commit()
		}
		return err
	}
	//添加新的owners
	txquery = txquery.Where("fhash = ?", fhash).First(&fileInfo)
	OwnersMap, _ := fileInfo.OwnersUnmarshal()
	newOwnersMap, _ := newfileInfo.OwnersUnmarshal()
	mergeInfo := map[string]model.ClientInfo{}
	for key, val := range OwnersMap {
		mergeInfo[key] = val
	}
	for key, val := range newOwnersMap { //去重
		mergeInfo[key] = val
	}
	mergeByte, _ := json.Marshal(mergeInfo)
	err := txquery.Updates(map[string]interface{}{"owners": string(mergeByte)}).Error
	if err != nil {
		err = fmt.Errorf("[Updates-err]:%v ,fileInfo:%+v mergeByte:%v", err, &fileInfo, string(mergeByte))
		txquery.Rollback()
	} else {
		txquery.Commit()
	}
	return err

}

func (i *IPFSFileInfoDB) ModifyInfo(queryFactor map[string]string, updateMap map[string]interface{}) error {
	fileInfo := &model.FileInfo{}
	txquery := FileInfoDB.Debug().Begin()
	txquery = txquery.Model(fileInfo)
	//查询条件
	queryMap := queryFactorParse(queryFactor)
	for k, v := range queryMap {
		txquery = txquery.Where(k, v)
	}
	err := txquery.Updates(updateMap).Error
	if err != nil {
		txquery.Rollback()
	} else {
		txquery.Commit()
	}
	return err
}

func queryFactorParse(queryFactor map[string]string) map[string]interface{} {
	ret := make(map[string]interface{})
	for key, str := range queryFactor {
		switch key {
		case "create_time", "update_time":
			splits := strings.Split(str, ";")
			ret[fmt.Sprintf("%+v >=  FROM_UNIXTIME(?)", key)] = splits[0]
			ret[fmt.Sprintf("%+v < FROM_UNIXTIME(?)", key)] = splits[1]
		case "fhash", "title", "uploader":
			ret[fmt.Sprintf("%v = ?", key)] = str
		default:
			ret[fmt.Sprintf("%+v like ?", key)] = "%" + str + "%"
		}
	}
	return ret
}
