package models

import (
	"time"
)

type Traffic struct {
	IP         string    `json:"IP" gorm:"size:500;not null"`
	Date       time.Time `json:"Date" gorm:"not null"`
	Method     string    `json:"Method" gorm:"size:500;not null"`
	Path       string    `json:"Path" gorm:"size:500;not null"`
	StatusCode string    `json:"StatusCode" gorm:"size:500;not null"`
	TimeTaken  float64   `json:"TimeTaken" gorm:"type:decimal(7,6);default:0.000"`
	Log        string    `json:"Log" gorm:"size:5000;not null"`
}
