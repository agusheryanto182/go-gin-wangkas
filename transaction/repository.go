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

type TransactionRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *TransactionRepository {
	return &TransactionRepository{db}
}

func (r *TransactionRepository) Save(transaction Transaction) (Transaction, error) {
	err := r.db.Create(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *TransactionRepository) Update(transaction Transaction) (Transaction, error) {
	err := r.db.Save(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *TransactionRepository) Delete(ID int) error {
	var transaction Transaction
	err := r.db.Where("id = ?", ID).Delete(&transaction).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *TransactionRepository) FindByID(ID int) (Transaction, error) {
	var transaction Transaction
	err := r.db.Where("id = ?", ID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *TransactionRepository) FindAllByWeekID(weekID int) ([]Transaction, error) {
	var transaction []Transaction
	err := r.db.Where("minggu_ke = ?", weekID).Find(&transaction).Error
	if err != nil {
		return transaction, err
	}
	return transaction, nil
}

func (r *TransactionRepository) FindAll() ([]Transaction, error) {
	var transactions []Transaction
	err := r.db.Find(&transactions).Error
	if err != nil {
		return transactions, err
	}
	return transactions, nil
}
