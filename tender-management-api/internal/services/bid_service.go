package services

import (
    "errors"
    "strconv"
    "tender-management-api/internal/db"
    "tender-management-api/internal/models"
    "github.com/google/uuid"
)

// GetBids получает список предложений для указанного тендера с пагинацией.
func GetBids(tenderID, limit, offset string) ([]models.Bid, error) {
    // Преобразование limit и offset в целые числа
    lim, off := 10, 0
    var err error

    if limit != "" {
        lim, err = strconv.Atoi(limit)
        if err != nil {
            return nil, errors.New("invalid limit parameter")
        }
    }

    if offset != "" {
        off, err = strconv.Atoi(offset)
        if err != nil {
            return nil, errors.New("invalid offset parameter")
        }
    }

    // Получение предложений для указанного тендера
    var bids []models.Bid
    if err := db.DB.Where("tender_id = ?", tenderID).Offset(off).Limit(lim).Find(&bids).Error; err != nil {
        return nil, err
    }

    return bids, nil
}

// CreateBid создает новое предложение.
func CreateBid(bid *models.Bid) (*models.Bid, error) {
    // Генерация UUID для поля ID
    bid.ID = uuid.New().String()

    // Валидация данных
    if err := validateBid(bid); err != nil {
        return nil, err
    }

    // Сохранение предложения в базу данных
    if err := db.DB.Create(bid).Error; err != nil {
        return nil, err
    }

    return bid, nil
}

// GetBid получает предложение по его ID.
func GetBid(bidID string) (*models.Bid, error) {
    var bid models.Bid
    if err := db.DB.Where("id = ?", bidID).First(&bid).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return nil, errors.New("bid not found")
        }
        return nil, err
    }
    return &bid, nil
}

// UpdateBid обновляет предложение.
func UpdateBid(bidID string, updates map[string]interface{}) (*models.Bid, error) {
    var bid models.Bid

    // Найти предложение по ID
    if err := db.DB.Where("id = ?", bidID).First(&bid).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return nil, errors.New("bid not found")
        }
        return nil, err
    }

    // Обновление полей предложения
    if err := db.DB.Model(&bid).Updates(updates).Error; err != nil {
        return nil, err
    }

    return &bid, nil
}

// DeleteBid удаляет предложение по его ID.
func DeleteBid(bidID string) error {
    if err := db.DB.Where("id = ?", bidID).Delete(&models.Bid{}).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return errors.New("bid not found")
        }
        return err
    }
    return nil
}

// validateBid проверяет корректность данных предложения.
func validateBid(bid *models.Bid) error {
    if bid.Name == "" {
        return errors.New("name cannot be empty")
    }
    if bid.Description == "" {
        return errors.New("description cannot be empty")
    }
    if bid.TenderID == "" {
        return errors.New("tenderId cannot be empty")
    }
    // Добавьте дополнительные проверки по необходимости
    return nil
}
