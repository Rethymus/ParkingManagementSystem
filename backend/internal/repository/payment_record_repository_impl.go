package repository

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/zccccc01/ParkingManagementSystem/backend/internal/models"
)

type PaymentRecordRepositoryImpl struct {
	DB *gorm.DB
}

func NewPaymentRecordRepository(db *gorm.DB) PaymentRecordRepository {
	return &PaymentRecordRepositoryImpl{DB: db}
}

func (r *PaymentRecordRepositoryImpl) Create(payment *models.PaymentRecord) (bool, error) {
	result := r.DB.Create(&payment)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r *PaymentRecordRepositoryImpl) GetAmountByRecordID(id int) (float64, error) {
	var payment models.PaymentRecord
	result := r.DB.First(&payment, "RecordID=?", id)
	if result.Error != nil {
		return 0, result.Error
	}
	return payment.Amount, nil
}

func (r *PaymentRecordRepositoryImpl) GetAmountByReservationID(id int) (float64, error) {
	var payment models.PaymentRecord
	result := r.DB.Find(&payment, "ReservationID=?", id)
	if result.Error != nil {
		return 0, result.Error
	}
	return payment.Amount, nil
}

func (r *PaymentRecordRepositoryImpl) GetPaymentMethodByPaymentID(id int) (string, error) {
	var payment models.PaymentRecord
	result := r.DB.First(&payment, "PaymentID=?", id)
	if result.Error != nil {
		return "", result.Error
	}
	return payment.PaymentMethod, nil
}

func (r *PaymentRecordRepositoryImpl) GetPaymentTimeStampByPaymentID(id int) (time.Time, error) {
	var payment models.PaymentRecord
	result := r.DB.First(&payment, "PaymentID=?", id)
	if result.Error != nil {
		return time.Time{}, result.Error
	}
	return payment.PaymentTimestamp, nil
}

func (r *PaymentRecordRepositoryImpl) GetPaymentStatusByReservationID(id int) (string, error) {
	var payment models.PaymentRecord
	result := r.DB.Select("PaymentTimestamp").Find(&payment, "ReservationID = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "NoRecord", nil
		}
		return "", result.Error
	}

	if payment.PaymentTimestamp.Before(time.Now()) {
		return "PAY", nil
	}
	return "NoPAY", nil
}

func (r *PaymentRecordRepositoryImpl) GetPaymentStatusByRecordID(id int) (string, error) {
	var payment models.PaymentRecord
	result := r.DB.Select("PaymentTimestamp").Find(&payment, "RecordID = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return "NoRecord", nil
		}
		return "", result.Error
	}

	if payment.PaymentTimestamp.Before(time.Now()) {
		return "PAY", nil
	}
	return "NoPAY", nil
}

func (r *PaymentRecordRepositoryImpl) GetPaymentFeeByPlateNumber(plateNumber string) ([]float64, error) {
	/**
	select Amount from paymentrecord where RecordID in (
		select RecordID from parkingrecord where RecordID in (
			select VehicleID from vehicle where PlateNumber = ?
		)
	)
	*/
	type Amount struct {
		Amount float64 `gorm:"column:Amount"`
	}
	var amounts []Amount
	var fees []float64
	query := `
	SELECT Amount 
	FROM paymentrecord 
	WHERE RecordID IN (
		SELECT RecordID 
		FROM parkingrecord 
		WHERE RecordID IN (
			SELECT VehicleID 
			FROM vehicle 
			WHERE PlateNumber = ?
		)
	)
	`
	result := r.DB.Raw(query, plateNumber).Scan(&amounts)
	if result.Error != nil {
		return nil, result.Error
	}
	for _, amount := range amounts {
		fees = append(fees, amount.Amount)
	}
	return fees, nil
}
