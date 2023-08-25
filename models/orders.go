package models

type Order struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	UserID      uint   `json:"user_id"`
	Status      string `gorm:"not null;default:processing" json:"status"`
	Description string `gorm:"not null" json:"description"`

	User User `json:"-"`
}

func CreateOrder(userId uint, description string) (uint, error) {
	order := Order{
		UserID:      userId,
		Description: description,
	}
	err := db.Create(&order).Error
	return order.ID, err
}
