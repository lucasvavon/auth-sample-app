package models

import (
	"database/sql/driver"
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	BudgetID    int       `json:"budget_id" gorm:"not null"`
	Amount      float32   `json:"amount" gorm:"not null"`
	Type        TransType `json:"trans_type" gorm:"type:trans_type;not null"`
	Recurring   bool      `gorm:"default:false"`
	Frequency   Frequency `json:"frequency" gorm:"type:frequency;not null"`
	Description string    `json:"description"`
	CategoryID  int       `json:"categorie_id" gorm:"not null"`
}

type TransType string

const (
	Income  TransType = "income"
	Expense TransType = "expense"
)

func (t *TransType) Scan(value interface{}) error {
	*t = TransType(value.([]byte))
	return nil
}

func (t TransType) Value() (driver.Value, error) {
	return string(t), nil
}

type Frequency string

const (
	Annual  Frequency = "annual"
	Monthly Frequency = "monthly"
	Weekly  Frequency = "weekly"
	Daily   Frequency = "daily"
)

func (p *Frequency) Scan(value interface{}) error {
	*p = Frequency(value.([]byte))
	return nil
}

func (p Frequency) Value() (driver.Value, error) {
	return string(p), nil
}
