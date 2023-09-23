package transaction

import (
	"gorm.io/gorm"
)

type Repository interface {
	Save(transaction Transaction) (Transaction, error)
	Update(transaction Transaction) (Transaction, error)
	Delete(ID int) error
	FindByID(ID int) (Transaction, error)
	FindAllByWeekID(weekID int) ([]Transaction, error)
	FindAll() ([]Transaction, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) Delete(ID int) error {
	var transaction Transaction
	err := r.db.Where("id = ?", ID).Delete(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) FindByID(ID int) (Transaction, error) {
	var transaction Transaction
	err := r.db.Where("id = ?", ID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) FindAllByWeekID(weekID int) ([]Transaction, error) {
	var transaction []Transaction
	err := r.db.Where("minggu_ke = ?", weekID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *repository) FindAll() ([]Transaction, error) {
	var transaction []Transaction
	err := r.db.Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}
