package models

import "errors"

const (
	REFUND_STATUS_APPENDING = "appending"
	REFUND_STATUS_APPROVED  = "approved"
	REFUND_STATUS_CANCEL    = "cancel"

	REFUND_ERROR_BADORDER = "Order is not completed"
)

type Refund struct {
	ID      uint   `gorm:"primaryKey" json:"id"`
	OrderID uint   `gorm:"not null;unique" json:"order_id"`
	Status  string `gorm:"not null;default:appending" json:"status"`

	Order Order `json:"-"`
}

func CreateRefund(orderId uint) (id uint, err error) {
	order, err := GetOrderById(orderId)
	if err != nil {
		return 0, err
	}
	if order.Status != ORDER_STATUS_COMPLETED {
		return 0, errors.New(REFUND_ERROR_BADORDER)
	}

	Refund := Refund{
		OrderID: orderId,
		Status:  REFUND_STATUS_APPENDING,
	}
	err = db.Create(&Refund).Error
	id = Refund.ID
	return
}

func GetRefundById(id uint) (Refund Refund, err error) {
	err = db.First(&Refund, id).Error
	return
}

func GetAllRefunds() (Refunds []Refund, err error) {
	err = db.Find(&Refunds).Error
	return
}

func UpdateRefund(id uint, status string) int64 {
	return db.Where(id).Updates(Refund{Status: status}).RowsAffected
}

func DeleteRefund(id uint) int64 {
	return db.Delete(&Refund{}, id).RowsAffected
}
