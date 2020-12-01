package accounts

import (
	"time"

	"github.com/google/uuid"
	"github.com/yohanalexander/desafio-banking-go/cmd/banking/models/transfers"
	"github.com/yohanalexander/desafio-banking-go/pkg/app"
	"github.com/yohanalexander/desafio-banking-go/pkg/secret"
	"gorm.io/gorm"
)

// BeforeCreate hook do gorm para gerar uuid no create
func (a *Account) BeforeCreate(tx *gorm.DB) (err error) {
	a.ID = uuid.New()
	a.Secret, err = secret.HashPassword(a.Secret)
	if err != nil {
		return err
	}
	return
}

// Account modelo para conta do usuário
type Account struct {
	gorm.Model `json:"-"`
	ID         uuid.UUID            `json:"id" gorm:"type:uuid"`
	Name       string               `json:"name" validate:"required"`
	CPF        string               `gorm:"unique" json:"cpf" validate:"required"`
	Secret     string               `json:"secret" validate:"required"`
	Balance    float64              `json:"balance" validate:"required"`
	CreatedAt  time.Time            `json:"created_at"`
	Transfers  []transfers.Transfer `json:"-" gorm:"foreignKey:ID"`
}

// Create cria uma conta de usuário
func (a *Account) Create(app *app.App) *gorm.DB {

	result := app.DB.Client.Create(&Account{
		ID:        a.ID,
		Name:      a.Name,
		CPF:       a.CPF,
		Secret:    a.Secret,
		Balance:   a.Balance,
		CreatedAt: a.CreatedAt,
		Transfers: a.Transfers,
	})

	return result

}
