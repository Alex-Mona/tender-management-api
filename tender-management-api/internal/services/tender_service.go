package services

import (
    "errors"
    "strconv"
    "tender-management-api/internal/db"
    "tender-management-api/internal/models"
    "github.com/google/uuid"
)

// GetTenders получает список тендеров с фильтрацией по типу услуг и пагинацией.
func GetTenders(serviceType, limit, offset string) ([]models.Tender, error) {
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

    // Создание запроса с учетом фильтрации по типу услуг
    query := db.DB.Model(&models.Tender{}).Offset(off).Limit(lim)
    if serviceType != "" {
        query = query.Where("service_type IN (?)", serviceType)
    }

    var tenders []models.Tender
    if err := query.Find(&tenders).Error; err != nil {
        return nil, err
    }

    return tenders, nil
}

// CreateTender создает новый тендер.
func CreateTender(tender *models.Tender) (*models.Tender, error) {
    // Генерация UUID для поля ID
    tender.ID = uuid.New().String()

    // Валидация данных
    if err := validateTender(tender); err != nil {
        return nil, err
    }

    // Сохранение тендера в базу данных
    if err := db.DB.Create(tender).Error; err != nil {
        return nil, err
    }

    return tender, nil
}

// GetUserTenders получает тендеры пользователя с пагинацией.
func GetUserTenders(username, limit, offset string) ([]models.Tender, error) {
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

    // Получение тендеров пользователя
    var tenders []models.Tender
    if err := db.DB.Where("creator_username = ?", username).Offset(off).Limit(lim).Find(&tenders).Error; err != nil {
        return nil, err
    }

    return tenders, nil
}

// GetTenderStatus получает статус тендера.
func GetTenderStatus(tenderID, username string) (models.TenderStatus, error) {
    var tender models.Tender

    if err := db.DB.Where("id = ? AND creator_username = ?", tenderID, username).First(&tender).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return "", errors.New("tender not found")
        }
        return "", err
    }

    return tender.Status, nil
}

// UpdateTenderStatus изменяет статус тендера.
func UpdateTenderStatus(tenderID, status, username string) (*models.Tender, error) {
    var tender models.Tender

    // Найти тендер по ID
    if err := db.DB.Where("id = ? AND creator_username = ?", tenderID, username).First(&tender).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return nil, errors.New("tender not found")
        }
        return nil, err
    }

    // Изменение статуса тендера
    tender.Status = models.TenderStatus(status)
    if err := db.DB.Save(&tender).Error; err != nil {
        return nil, err
    }

    return &tender, nil
}

// EditTender редактирует параметры тендера.
func EditTender(tenderID, username string, updates map[string]interface{}) (*models.Tender, error) {
    var tender models.Tender

    // Найти тендер по ID
    if err := db.DB.Where("id = ? AND creator_username = ?", tenderID, username).First(&tender).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return nil, errors.New("tender not found")
        }
        return nil, err
    }

    // Обновление полей тендера
    if err := db.DB.Model(&tender).Updates(updates).Error; err != nil {
        return nil, err
    }

    return &tender, nil
}

// RollbackTender откатывает параметры тендера к указанной версии.
func RollbackTender(tenderID string, version int, username string) (*models.Tender, error) {
    var tender models.Tender

    // Найти тендер по ID
    if err := db.DB.Where("id = ? AND creator_username = ?", tenderID, username).First(&tender).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return nil, errors.New("tender not found")
        }
        return nil, err
    }

    // Проверка версии
    if version >= tender.Version {
        return nil, errors.New("invalid version number")
    }

    // Откат к указанной версии (добавьте логику для получения предыдущих версий)
    // Здесь представлена упрощенная логика, которая требует адаптации под вашу структуру базы данных
    var oldTender models.Tender
    if err := db.DB.Where("id = ? AND version = ?", tenderID, version).First(&oldTender).Error; err != nil {
        if err == db.ErrRecordNotFound {
            return nil, errors.New("tender version not found")
        }
        return nil, err
    }

    // Обновление текущего тендера до старой версии
    oldTender.Version++
    if err := db.DB.Save(&oldTender).Error; err != nil {
        return nil, err
    }

    return &oldTender, nil
}

// validateTender проверяет корректность данных тендера.
func validateTender(tender *models.Tender) error {
    if tender.Name == "" {
        return errors.New("name cannot be empty")
    }
    if tender.Description == "" {
        return errors.New("description cannot be empty")
    }
    // Добавьте дополнительные проверки по необходимости
    return nil
}
