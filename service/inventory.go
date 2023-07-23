package service

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

type Inventory struct {
	Id           int64  `json:"id" gorm:"primary_key"`
	SerialNumber string `json:"serial_number" gorm:"unique_index:idx_serial_number_brand_model"`
	Brand        string `json:"brand" gorm:"unique_index:idx_serial_number_brand_model"`
	Model        string `json:"model" gorm:"unique_index:idx_serial_number_brand_model"`
	Status       string `json:"status"`
	DateBought   string `json:"date_bought"`
}

type InventoryOrm struct {
	db   *gorm.DB
	data *Inventory
}

func NewInventoryOrm(db *gorm.DB, data *Inventory) InventoryInt {
	return &InventoryOrm{
		db:   db,
		data: data,
	}
}

type InventoryInt interface {
	GetAll() ([]Inventory, error)
	GetByWhere() ([]Inventory, error)
	Create() error
	Update(id int64) error
	Delete() error
}

func (dm *InventoryOrm) GetAll() ([]Inventory, error) {
	var appList []Inventory
	err := dm.db.Model(dm.data).Find(&appList).Error
	return appList, err
}

func (dm *InventoryOrm) GetByWhere() ([]Inventory, error) {
	var appList []Inventory

	where := ""
	if dm.data.SerialNumber != "" {
		where += fmt.Sprintf("serial_number=\"%s\" and ", dm.data.SerialNumber)
	}
	if dm.data.Brand != "" {
		where += fmt.Sprintf("brand=\"%s\" and ", dm.data.Brand)
	}
	if dm.data.Model != "" {
		where += fmt.Sprintf("model=\"%s\" and ", dm.data.Model)
	}
	if dm.data.Status != "" {
		where += fmt.Sprintf("status=\"%s\" and ", dm.data.Status)
	}
	if dm.data.DateBought != "" {
		where += fmt.Sprintf("date_bought=\"%s\" and ", dm.data.DateBought)
	}

	if len(where) > 3 && where[len(where)-4:] == "and " {
		where = where[:len(where)-4]
	}
	log.Printf("where:%v", where)
	err := dm.db.Table("inventories").Where(where).Find(&appList).Error
	return appList, err
}

func (dm *InventoryOrm) Create() error {
	err := dm.db.Model(dm.data).Create(dm.data).Error
	return err
}

func (dm *InventoryOrm) Update(id int64) error {
	dm.data.Id = 0
	err := dm.db.Model(dm.data).Where("id = ?", id).Updates(dm.data).Error
	return err
}

func (dm *InventoryOrm) Delete() error {
	dm.data.Id = 0
	err := dm.db.Model(dm.data).Where("serial_number = ? and brand = ? and model = ?", dm.data.SerialNumber, dm.data.Brand, dm.data.Model).Delete(dm.data).Error
	return err
}
