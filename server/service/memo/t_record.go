package memo

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/memo"
	memoReq "github.com/flipped-aurora/gin-vue-admin/server/model/memo/request"
	"gorm.io/gorm"
)

type TRecordService struct{}

// CreateTRecord 创建备忘录记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (tRecordService *TRecordService) CreateTRecord(tRecord *memo.TRecord) (err error) {
	err = global.GVA_DB.Create(tRecord).Error
	return err
}

// DeleteTRecord 删除备忘录记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (tRecordService *TRecordService) DeleteTRecord(ID string, userID uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&memo.TRecord{}).Where("id = ?", ID).Update("deleted_by", userID).Error; err != nil {
			return err
		}
		if err = tx.Delete(&memo.TRecord{}, "id = ?", ID).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteTRecordByIds 批量删除备忘录记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (tRecordService *TRecordService) DeleteTRecordByIds(IDs []string, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&memo.TRecord{}).Where("id in ?", IDs).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", IDs).Delete(&memo.TRecord{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateTRecord 更新备忘录记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (tRecordService *TRecordService) UpdateTRecord(tRecord memo.TRecord) (err error) {
	err = global.GVA_DB.Model(&memo.TRecord{}).Where("id = ?", tRecord.ID).Updates(&tRecord).Error
	return err
}

// GetTRecord 根据ID获取备忘录记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (tRecordService *TRecordService) GetTRecord(ID string) (tRecord memo.TRecord, err error) {
	err = global.GVA_DB.Where("id = ?", ID).First(&tRecord).Error
	return
}

// GetTRecordInfoList 分页获取备忘录记录记录
// Author [piexlmax](https://github.com/piexlmax)
func (tRecordService *TRecordService) GetTRecordInfoList(info memoReq.TRecordSearch) (list []memo.TRecord, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&memo.TRecord{})
	var tRecords []memo.TRecord
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.RecordName != "" {
		db = db.Where("record_name LIKE ?", "%"+info.RecordName+"%")
	}
	if info.CreatedBy != nil {
		db = db.Where("created_by = ?", info.CreatedBy)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	if limit != 0 {
		db = db.Limit(limit).Offset(offset)
	}

	err = db.Find(&tRecords).Error
	return tRecords, total, err
}
