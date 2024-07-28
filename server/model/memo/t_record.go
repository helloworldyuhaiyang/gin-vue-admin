// 自动生成模板TRecord
package memo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// 备忘录记录 结构体  TRecord
type TRecord struct {
	global.GVA_MODEL
	RecordName  string `json:"recordName" form:"recordName" gorm:"column:record_name;comment:记录名字;size:191;"`        //记录名字
	ResourceUrl string `json:"resourceUrl" form:"resourceUrl" gorm:"column:resource_url;comment:第三方存储的地址;size:191;"` //第三方存储的地址
	Location    string `json:"location" form:"location" gorm:"column:location;comment:记录位置;size:191;"`               //记录位置
	RSize       *int   `json:"rSize" form:"rSize" gorm:"column:r_size;comment:记录文件大小;"`                              //记录文件大小
	CreatedBy   uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy   uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy   uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName 备忘录记录 TRecord自定义表名 t_record
func (TRecord) TableName() string {
	return "t_record"
}
