package models

import (
	"time"
)

type Period string

const (
	Annual         = "annual"
	Monthly Period = "monthly"
	Weekly  Period = "weekly"
	Daily   Period = "daily"
)

type Budget struct {
	ID     int    `json:"id" gorm:"primaryKey"`
	Max    int    `json:"max" gorm:"not null"`
	Period Period `gorm:"type:enum('admin', 'moderator', 'user')"`
}

type Income struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	BudgetID    int       `json:"budget_id" gorm:"not null"`
	UserID      int       `json:"user_id" gorm:"not null"`
	Amount      float32   `json:"amount" gorm:"not null"`
	Description string    `json:"description" gorm:"not null"`
	CreatedAt   time.Time `json:"created_at"`
	Budgets     []Budget  `gorm:"foreignKey:BudgetID;constraint:OnDelete:CASCADE;"`
}

type Expense struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	BudgetID    int       `json:"budget_id" gorm:"not null"`
	CategoryID  int       `json:"category_id" gorm:"not null"`
	Amount      float32   `json:"amount" gorm:"not null"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	Budgets     []Budget  `gorm:"foreignKey:BudgetID;constraint:OnDelete:CASCADE;"`
}

type Category struct {
	ID          int    `json:"id" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"not null"`
	Logo        string `json:"logo" gorm:"not null"` // data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAABAAAAAQCAYAAAAf8/9hAAAABHNCSVQICAgIfAhkiAAAAAlwSFlzAAAAdgAAAHYBTnsmCAAAABl0RVh0U29mdHdhcmUAd3d3Lmlua3NjYXBlLm9yZ5vuPBoAAAF7SURBVDiNldM/SNVRFAfwz3sq/kEjSbOh4CElGBlJQSSNiRAFCm21JEgYSvQWBSeHiMBBcGoIWwVbmyTEbLMhFCGoSSsahMjsEUU23POr31L0Dlw4XM73e+453+/l/2MeEziAh9llsQqCQtSPo6YK3O8oox7NQQJqqyBoxrc4bX8juIKTqMMC3uTA56QdwBk04WsGrEM3ljCIu3gahXATRyNvxAiuZy+4jEvojMIOrGAWU/iAI9gOggoa0Era6lk8xzS+4LW0rH08wfnI8zEXWEW0ow/r+IlJ3I7Cq3jpH3IXsYd7+IFrOIhXOIXDeC95oIQhtOQJagJQwEccj7uR2EklSJvidZsYxTNcxGrmrjK6YrbW6PpO0nvHH7nKOIYX6MfjYjDP4HPk43gQ3bcC3IH7+IRdnMAtDOSN1CnJeSfGacTp6LQbYwriTJVK/lOUsIweSc4b0XENF/AWj4K4V/JQqZAjqMdYLHE/CmolOy9iGIfwXTLVJjZ+AX9GTzyDVsiyAAAAAElFTkSuQmCC
	Description string `json:"description"`
}

type Transaction struct {
	ID       int `json:"id" gorm:"primaryKey"`
	BudgetID int `json:"budget_id" gorm:"not null"`

	Amount      float32   `json:"amount" gorm:"not null"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	type: string (dépense, revenu)
}
