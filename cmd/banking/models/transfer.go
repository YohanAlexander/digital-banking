package models

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yohanalexander/desafio-banking-go/pkg/app"
	"gorm.io/gorm"
)

// BeforeCreate hook do gorm para gerar uuid no create
func (t *Transfer) BeforeCreate(tx *gorm.DB) (err error) {
	t.ID = uuid.New()
	return
}

// Transfer modelo para transferências do usuário
type Transfer struct {
	gorm.Model           `json:"-"`
	ID                   uuid.UUID `json:"id" gorm:"type:uuid"`
	AccountOriginID      uuid.UUID `json:"account_origin_id" validate:"required"`
	AccountDestinationID uuid.UUID `json:"account_destination_id" validate:"required"`
	Amount               float64   `json:"amount" validate:"required"`
	CreatedAt            time.Time `json:"created_at"`
}

// Create realiza uma transferência entre contas
func (t *Transfer) Create(app *app.App) error {

	// inicia o modo de transaction
	tx := app.DB.Client.Begin()

	// verifica se a conta de destino existe
	if err := t.CheckDestinationAccount(app); err != nil {
		// caso não encontre faz rollback
		tx.Rollback()
		return err
	}

	// verifica se a conta de origem tem saldo suficiente
	if err := t.CheckOriginBalance(app); err != nil {
		// caso não tenha saldo faz rollback
		tx.Rollback()
		return err
	}

	// cria o struct transfer no DB
	if err := tx.Create(&Transfer{
		ID:                   t.ID,
		AccountOriginID:      t.AccountOriginID,
		AccountDestinationID: t.AccountDestinationID,
		Amount:               t.Amount,
		CreatedAt:            t.CreatedAt,
	}); err.Error != nil {
		// caso ocorra erro faz rollback
		tx.Rollback()
		return err.Error
	}

	// atualiza os saldos da conta de origem e destino
	if err := t.BalanceAccounts(app); err != nil {
		// caso ocorra erro faz rollback
		tx.Rollback()
		return err
	}

	// transferência sem erros é comitada
	tx.Commit()

	// retorna erro nulo
	return nil

}

// CheckDestinationAccount verifica se a conta de destino existe
func (t *Transfer) CheckDestinationAccount(app *app.App) error {

	// captura a conta de destino no banco
	a := &Account{}
	result := app.DB.Client.First(&a, &t.AccountDestinationID)
	return result.Error

}

// CheckOriginBalance verifica se a conta de origem tem saldo suficiente
func (t *Transfer) CheckOriginBalance(app *app.App) error {

	// captura a conta de origem no banco
	a := &Account{}
	result := app.DB.Client.First(&a, &t.AccountOriginID)
	if result.Error != nil {
		return result.Error
	}

	// caso não tenha saldo suficiente retorna erro adequado
	if (t.Amount - a.Balance) < 0 {
		return errors.New("Saldo da conta insuficiente")
	}

	// caso tenha saldo suficiente retorna erro nulo
	return nil

}

// BalanceAccounts atualiza os saldos da conta de origem e destino
func (t *Transfer) BalanceAccounts(app *app.App) error {

	// inicia o modo de transaction
	tx := app.DB.Client.Begin()

	// captura a conta de origem no DB
	origem := &Account{}
	if err := app.DB.Client.First(&origem, &t.AccountOriginID); err.Error != nil {
		tx.Rollback()
		return err.Error
	}

	// captura a conta de destino no DB
	destino := &Account{}
	if err := app.DB.Client.First(&destino, &t.AccountDestinationID); err.Error != nil {
		tx.Rollback()
		return err.Error
	}

	// atualiza o saldo da conta de origem
	origem.Balance = origem.Balance - t.Amount
	if err := tx.Save(&origem); err != nil {
		tx.Rollback()
		return err.Error
	}

	// atualiza o saldo da conta de destino
	destino.Balance = destino.Balance + t.Amount
	if err := tx.Save(&destino); err != nil {
		tx.Rollback()
		return err.Error
	}

	// balanceamento sem erros é comitado
	tx.Commit()

	// retorna erro nulo
	return nil

}
