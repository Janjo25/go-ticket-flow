package domain

import "time"

type Ticket struct {
	ID           uint            `gorm:"primaryKey"`
	Category     string          `gorm:"not null"`
	Description  string          `gorm:"not null"`
	Priority     string          `gorm:"not null"`                     // Prioridades: Baja, Media, Alta, Crítica.
	Status       string          `gorm:"not null;default:'Pendiente'"` // Estados: Pendiente, En proceso, Resuelto.
	Title        string          `gorm:"not null"`
	UserID       uint            `gorm:"not null"`     // Identificador del empleado que creó el ticket.
	TechnicianID *uint           `gorm:"default:null"` // Identificador del técnico asignado (puede ser nulo).
	Comments     []Comment       `gorm:"foreignKey:TicketID"`
	History      []TicketHistory `gorm:"foreignKey:TicketID"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Comment representa un comentario en un ticket.
type Comment struct {
	ID        uint   `gorm:"primaryKey"`
	TicketID  uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"`
	Content   string `gorm:"not null"`
	CreatedAt time.Time
}

// TicketHistory representa un cambio de estado en un ticket.
type TicketHistory struct {
	ID        uint   `gorm:"primaryKey"`
	TicketID  uint   `gorm:"not null"`
	UserID    uint   `gorm:"not null"` // Quién hizo el cambio.
	NewStatus string `gorm:"not null"`
	OldStatus string `gorm:"not null"`
	ChangedAt time.Time
}
