package repository

import (
	"context"
	"errors"
)

// ErrNotFound is returned when a requested record does not exist.
var ErrNotFound = errors.New("not found")

// SubscriptionRepository is the read interface used by the service.
type SubscriptionRepository interface {
	FindByID(ctx context.Context, id string) (*SubscriptionRow, error)
}

// PlanRepository is the read interface used by the service.
type PlanRepository interface {
	FindByID(ctx context.Context, id string) (*PlanRow, error)
}

// StatementQuery defines the parameters for listing statements.
type StatementQuery struct {
	SubscriptionID string
	Kind           string
	Status         string
	StartAfter     string
	EndBefore      string
	PageSize       int
	Page           int
}

// StatementRepository is the read interface used by the service.
type StatementRepository interface {
	FindByID(ctx context.Context, id string) (*StatementRow, error)
	ListByCustomerID(ctx context.Context, customerID string, q StatementQuery) ([]*StatementRow, int, error)
}
