package structs

import "time"

type PhantomPayUserData struct {
	UserId    string    `gorm:"primaryKey;type:varchar(100)"`
	Username  string    `gorm:"type:varchar(100);not null"`
	Email     string    `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password  string    `gorm:"type:text;not null"`
	CreatedAt time.Time 
	UpdatedAt time.Time 
}
