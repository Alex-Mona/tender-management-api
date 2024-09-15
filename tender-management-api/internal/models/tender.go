// package models

// import "time"

// // TenderStatus определяет статус тендера.
// type TenderStatus string

// const (
//     TenderCreated   TenderStatus = "Created"
//     TenderPublished TenderStatus = "Published"
//     TenderClosed    TenderStatus = "Closed"
// )

// // TenderServiceType определяет тип услуги тендера.
// type TenderServiceType string

// const (
//     Construction TenderServiceType = "Construction"
//     Delivery     TenderServiceType = "Delivery"
//     Manufacture  TenderServiceType = "Manufacture"
// )

// // Tender представляет тендер.
// type Tender struct {
//     ID             string             `json:"id" gorm:"primary_key"`
//     Name           string             `json:"name"`
//     Description    string             `json:"description"`
//     ServiceType    TenderServiceType  `json:"serviceType"`
//     Status         TenderStatus       `json:"status"`
//     OrganizationID string             `json:"organizationId"`
//     Version        int                `json:"version"`
//     CreatedAt      time.Time          `json:"createdAt"`
// }

package models

import "time"

// TenderStatus определяет статус тендера.
type TenderStatus string

const (
    TenderCreated   TenderStatus = "Created"
    TenderPublished TenderStatus = "Published"
    TenderClosed    TenderStatus = "Closed"
)

// TenderServiceType определяет тип услуги тендера.
type TenderServiceType string

const (
    Construction TenderServiceType = "Construction"
    Delivery     TenderServiceType = "Delivery"
    Manufacture  TenderServiceType = "Manufacture"
)

// Tender представляет тендер.
type Tender struct {
    ID             string             `json:"id" gorm:"primary_key"` // Строковое поле ID
    Name           string             `json:"name"`
    Description    string             `json:"description"`
    ServiceType    TenderServiceType  `json:"serviceType"`
    Status         TenderStatus       `json:"status"`
    OrganizationID string             `json:"organizationId"`
    Version        int                `json:"version"`
    CreatedAt      time.Time          `json:"createdAt"`
}
