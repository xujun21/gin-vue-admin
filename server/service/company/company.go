package company

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/company"
	companyReq "github.com/flipped-aurora/gin-vue-admin/server/model/company/request"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type CompanyService struct {
}

// CreateCompany 创建Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (compService *CompanyService) CreateCompany(comp company.Company) (err error) {
	err = global.GVA_DB.Create(&comp).Error
	return err
}

// DeleteCompany 删除Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (compService *CompanyService) DeleteCompany(comp company.Company) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&company.Company{}).Where("id = ?", comp.ID).Update("deleted_by", comp.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&comp).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCompanyByIds 批量删除Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (compService *CompanyService) DeleteCompanyByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&company.Company{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&company.Company{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCompany 更新Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (compService *CompanyService) UpdateCompany(comp company.Company) (err error) {
	err = global.GVA_DB.Save(&comp).Error
	return err
}

// GetCompany 根据id获取Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (compService *CompanyService) GetCompany(id uint) (comp company.Company, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&comp).Error
	return
}

// GetCompanyInfoList 分页获取Company记录
// Author [piexlmax](https://github.com/piexlmax)
func (compService *CompanyService) GetCompanyInfoList(info companyReq.CompanySearch) (list []company.Company, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&company.Company{})
	var comps []company.Company
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Company_name != "" {
		db = db.Where("company_name LIKE ?", "%"+info.Company_name+"%")
	}
	if info.Contact_name != "" {
		db = db.Where("contact_name LIKE ?", "%"+info.Contact_name+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}
	if info.Postcode != "" {
		db = db.Where("postcode LIKE ?", "%"+info.Postcode+"%")
	}
	if info.Note != "" {
		db = db.Where("note LIKE ?", "%"+info.Note+"%")
	}
	if info.Sage != "" {
		db = db.Where("sage LIKE ?", "%"+info.Sage+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["company_name"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("id desc")
	}
	err = db.Limit(limit).Offset(offset).Find(&comps).Error
	return comps, total, err
}

func (compService *CompanyService) ExportOrderExcel(info companyReq.CompanySearch, fileName string) (err error) {
	// 创建db
	db := global.GVA_DB.Model(&company.Company{})
	var comps []company.Company
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Company_name != "" {
		db = db.Where("company_name LIKE ?", "%"+info.Company_name+"%")
	}
	if info.Contact_name != "" {
		db = db.Where("contact_name LIKE ?", "%"+info.Contact_name+"%")
	}
	if info.Phone != "" {
		db = db.Where("phone LIKE ?", "%"+info.Phone+"%")
	}
	if info.Address != "" {
		db = db.Where("address LIKE ?", "%"+info.Address+"%")
	}
	if info.Postcode != "" {
		db = db.Where("postcode LIKE ?", "%"+info.Postcode+"%")
	}
	if info.Note != "" {
		db = db.Where("note LIKE ?", "%"+info.Note+"%")
	}
	if info.Sage != "" {
		db = db.Where("sage LIKE ?", "%"+info.Sage+"%")
	}

	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["company_name"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("id desc")
	}
	err = db.Find(&comps).Error
	if err != nil {
		return
	}

	// excel
	var xlsx *excelize.File = excelize.NewFile()
	var sheet = xlsx.GetSheetName(0)
	err = xlsx.SetSheetRow(sheet, "A1", &[]string{
		"Name",
		"Address",
		"Contact Person",
		"Contact Phone",
		"City",
		"Post Code",
		"Note",
	})
	if err != nil {
		return
	}

	for i, j := 0, len(comps); i < j; i++ {
		axis := fmt.Sprintf("A%d", i+2)

		err = xlsx.SetSheetRow(sheet, axis, &[]interface{}{
			comps[i].Company_name,
			comps[i].Address,
			comps[i].Contact_name,
			comps[i].Phone,
			comps[i].City,
			comps[i].Postcode,
			comps[i].Note,
		})
		if err != nil {
			return
		}
	}

	err = xlsx.SaveAs(fileName)
	return err

}
