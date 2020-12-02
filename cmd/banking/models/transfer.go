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

// CreateTransfer realiza uma transferência entre contas
func (t *Transfer) CreateTransfer(app *app.App) error {

	// inicia o modo de transaction
	tx := app.DB.Client.Begin()

	// verifica se a conta de destino existe
	if err := t.checkDestinationAccount(app); err != nil {
		// caso não encontre faz rollback
		tx.Rollback()
		return err
	}

	// verifica se a conta de origem tem saldo suficiente
	if err := t.checkOriginBalance(app); err != nil {
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
		return errors.New("Erro na criação da transferência")
	}

	// atualiza o saldo da conta de origem
	if err := t.balanceOriginAccount(app); err != nil {
		// caso ocorra erro faz rollback
		tx.Rollback()
		return err
	}

	// atualiza o saldo da conta de destino
	if err := t.balanceDestinationAccount(app); err != nil {
		// caso ocorra erro faz rollback
		tx.Rollback()
		return err
	}

	// transferência sem erros é comitada
	tx.Commit()

	// caso sucesso retorna erro nulo
	return nil

}

// checkDestinationAccount verifica se a conta de destino existe
func (t *Transfer) checkDestinationAccount(app *app.App) error {

	if t.AccountDestinationID == t.AccountOriginID {
		return errors.New("Contas de transferência devem ser diferentes")
	}

	// captura a conta de destino no banco
	a := &Account{}
	if result := app.DB.Client.First(&a, &t.AccountDestinationID); result.Error != nil {
		return errors.New("Conta de destino não encontrada")
	}

	// retorna exista conta de destino retorna erro nulo
	return nil

}

// checkOriginBalance verifica se a conta de origem tem saldo suficiente
func (t *Transfer) checkOriginBalance(app *app.App) error {

	// captura a conta de origem no banco
	a := &Account{}
	if result := app.DB.Client.First(&a, &t.AccountOriginID); result.Error != nil {
		return errors.New("Conta de origem não encontrada")
	}

	// caso não tenha saldo suficiente retorna erro adequado
	if (a.Balance - t.Amount) < 0 {
		return errors.New("Saldo da conta insuficiente")
	}

	// caso tenha saldo suficiente retorna erro nulo
	return nil

}

// balanceOriginAccount atualiza o saldo da conta de origem
func (t *Transfer) balanceOriginAccount(app *app.App) error {

	// inicia o modo de transaction
	tx := app.DB.Client.Begin()

	// captura a conta de origem no DB
	origem := &Account{}
	if result := tx.First(&origem, &t.AccountOriginID); result.Error != nil {
		tx.Rollback()
		return errors.New("Conta de origem não encontrada")
	}

	// atualiza o saldo da conta de origem
	origem.Balance = origem.Balance - t.Amount
	if result := tx.Save(&origem); result.Error != nil {
		tx.Rollback()
		return errors.New("Erro ao atualizar saldo da conta de origem")
	}

	// atualização sem erros é comitada
	tx.Commit()

	// caso sucesso retorna erro nulo
	return nil

}

// balanceDestinationAccount atualiza o saldo da conta de destino
func (t *Transfer) balanceDestinationAccount(app *app.App) error {

	// inicia o modo de transaction
	tx := app.DB.Client.Begin()

	// captura a conta de destino no DB
	destino := &Account{}
	if result := tx.First(&destino, &t.AccountDestinationID); result.Error != nil {
		tx.Rollback()
		return errors.New("Conta de destino não encontrada")
	}

	// atualiza o saldo da conta de destino
	destino.Balance = destino.Balance + t.Amount
	if result := tx.Save(&destino); result.Error != nil {
		tx.Rollback()
		return errors.New("Erro ao atualizar saldo da conta de destino")
	}

	// atualização sem erros é comitada
	tx.Commit()

	// caso sucesso retorna erro nulo
	return nil

}
