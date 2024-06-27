package models

import (
    "github.com/mengcuigoog/testgomod/utils"
)

//同步任务进度信息
type ThirdSyncTask struct {
        Id    uint32 `orm:"pk;auto" json:"id"`
        IdpId string
        //接口类型
        SyncDataType utils.DataType
        //上次同步时间
        LastSyncTime string
}

type IcbcRequestParam struct {
        Offset int64    `json:"offset"`
        Limit  int64    `json:"limit"`
        Time   []string `json:"time"`
        Key    string   `json:"key"`
}

type Result struct {
        error
}
