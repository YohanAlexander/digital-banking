package transfers

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// BeforeCreate hook do gorm para gerar uuid no automigrate do DB
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
