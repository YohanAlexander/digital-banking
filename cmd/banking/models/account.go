package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yohanalexander/desafio-banking-go/pkg/app"
	"github.com/yohanalexander/desafio-banking-go/pkg/secret"
	"gorm.io/gorm"
)

// BeforeCreate hook do gorm para gerar uuid no create
func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	a.Secret, err = secret.HashPassword(a.Secret)
	if err != nil {
		return errors.New("Erro ao criptografar senha")
	}
	return
}

// Account modelo para conta do usuário
type Account struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID  `json:"id" gorm:"type:uuid"`
	Name       string     `json:"name" validate:"required"`
	CPF        string     `gorm:"unique" json:"cpf" validate:"required,len=11"`
	Secret     string     `json:"secret" validate:"required"`
	Balance    float64    `json:"balance" validate:"required"`
	CreatedAt  time.Time  `json:"created_at"`
	Transfers  []Transfer `json:"-" gorm:"foreignKey:AccountOriginID"`
}

// CreateAccount cria uma conta de usuário
func (a *Account) CreateAccount(app *app.App) (*Account, error) {

	account := &Account{
		ID:        a.ID,
		Name:      a.Name,
		CPF:       a.CPF,
		Secret:    a.Secret,
		Balance:   a.Balance,
		CreatedAt: a.CreatedAt,
		Transfers: a.Transfers,
	}

	result := app.DB.Client.Create(account)

	if result.Error != nil {
		return nil, errors.New("Erro na criação da conta")
	}

	return account, nil

}
