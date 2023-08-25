package models

const (
	ORDER_STATUS_PROCESSING = "processing"
	ORDER_STATUS_COMPLETED  = "completed"
	ORDER_STATUS_CANCELED   = "canceled"
)

type Order struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	UserID      uint   `json:"user_id"`
	Status      string `gorm:"not null;default:processing" json:"status"`
	Description string `gorm:"not null" json:"description"`

	User User `json:"-"`
}

func CreateOrder(userId uint, description string) (id uint, err error) {
	order := Order{
		UserID:      userId,
		Description: description,
	}
	err = db.Create(&order).Error
	id = order.ID
	return
}

func GetOrderById(id uint) (order Order, err error) {
	err = db.First(&order, id).Error
	return
}

func GetAllOrders() (orders []Order, err error) {
	err = db.Find(&orders).Error
	return
}

func UpdateOrder(id uint, description, status string) int64 {
	return db.Where(id).Updates(Order{Description: description, Status: status}).RowsAffected
}
