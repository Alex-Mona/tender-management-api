package models

import "time"

// BidStatus определяет статус предложения.
type BidStatus string

const (
    BidCreated   BidStatus = "Created"
    BidPublished BidStatus = "Published"
    BidCanceled  BidStatus = "Canceled"
    BidApproved  BidStatus = "Approved"
    BidRejected  BidStatus = "Rejected"
)

// BidAuthorType определяет тип автора предложения.
type BidAuthorType string

const (
    Organization BidAuthorType = "Organization"
    User         BidAuthorType = "User"
)

// Bid представляет предложение в тендере.
type Bid struct {
    ID             string       `json:"id" gorm:"primary_key"`
    Name           string       `json:"name"` // Добавлено для имени предложения
    Description    string       `json:"description"` // Добавлено для описания предложения
    TenderID       string       `json:"tenderId"`
    Amount         float64      `json:"amount"` // Добавлено для суммы предложения
    CreatorUsername string     `json:"creatorUsername"`
    AuthorType     BidAuthorType `json:"authorType"` // Добавлено для типа автора
    AuthorID       string       `json:"authorId"` // Добавлено для ID автора
    Status         BidStatus    `json:"status"`
    CreatedAt      time.Time    `json:"createdAt"`
    Version        int          `json:"version"` // Добавлено для версии предложения
}

// BidReview представляет отзыв на предложение.
type BidReview struct {
    ID          string    `json:"id" gorm:"primary_key"`
    BidID       string    `json:"bidId"` // Связь с предложением
    Description string    `json:"description"`
    CreatedAt   time.Time `json:"createdAt"`
}