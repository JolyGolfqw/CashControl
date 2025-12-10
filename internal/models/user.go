package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email     string `gorm:"uniqueIndex;not null" json:"email"`    // Электронная почта пользователя
	Username  string `gorm:"uniqueIndex;not null" json:"username"` // Имя пользователя
	Password  string `gorm:"not null" json:"-"`                    // Хешированный пароль пользователя

	// Связи
	Expenses          []Expense          `gorm:"foreignKey:UserID" json:"-"` // Все расходы пользователя
	Categories        []Category         `gorm:"foreignKey:UserID" json:"-"` // Все категории пользователя
	Budgets           []Budget           `gorm:"foreignKey:UserID" json:"-"` // Все бюджеты пользователя
	RecurringExpenses []RecurringExpense `gorm:"foreignKey:UserID" json:"-"` // Все регулярные расходы пользователя
	ActivityHistory   []ActivityHistory  `gorm:"foreignKey:UserID" json:"-"` // Вся история действий пользователя
}

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`    // Электронная почта для регистрации
	Username string `json:"username" binding:"required,min=3"` // Имя пользователя для регистрации
	Password string `json:"password" binding:"required,min=6"` // Пароль для регистрации
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"` // Электронная почта для входа
	Password string `json:"password" binding:"required"`    // Пароль для входа
}

type LoginResponse struct {
	Token string `json:"token"` // JWT токен для аутентификации
	User  *User  `json:"user"`  // Информация о пользователе
}
