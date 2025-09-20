package repository

import (
	"gorm.io/gorm"
)

// CREATE Transaction
func CreateTransaction(db *gorm.DB, tx *Transaction) error {
	return db.Create(tx).Error
}

// GET All Transactions
func GetTransactions(db *gorm.DB) ([]Transaction, error) {
	var transactions []Transaction
	err := db.Preload("Analysis.Alerts").Find(&transactions).Error
	return transactions, err
}

// GET Transaction by ID
func GetTransactionByID(db *gorm.DB, id string) (Transaction, error) {
	var tx Transaction
	err := db.Preload("Analysis.Alerts").First(&tx, "transaction_id = ?", id).Error
	return tx, err
}

// UPDATE Transaction
func UpdateTransaction(db *gorm.DB, id string, updates map[string]interface{}) error {
	return db.Model(&Transaction{}).Where("transaction_id = ?", id).Updates(updates).Error
}

// DELETE Transaction
func DeleteTransaction(db *gorm.DB, id string) error {
	return db.Delete(&Transaction{}, "transaction_id = ?", id).Error
}
