package repository

import "context"

// MockSubscriptionRepo is an in-memory SubscriptionRepository for testing.
type MockSubscriptionRepo struct {
	records map[string]*SubscriptionRow
}

// NewMockSubscriptionRepo creates a MockSubscriptionRepo pre-populated with the given rows.
func NewMockSubscriptionRepo(rows ...*SubscriptionRow) *MockSubscriptionRepo {
	m := &MockSubscriptionRepo{records: make(map[string]*SubscriptionRow)}
	for _, r := range rows {
		m.records[r.ID] = r
	}
	return m
}

// FindByID returns the SubscriptionRow with the given ID, or ErrNotFound.
func (m *MockSubscriptionRepo) FindByID(_ context.Context, id string) (*SubscriptionRow, error) {
	row, ok := m.records[id]
	if !ok {
		return nil, ErrNotFound
	}
	return row, nil
}

// MockPlanRepo is an in-memory PlanRepository for testing.
type MockPlanRepo struct {
	records map[string]*PlanRow
}

// NewMockPlanRepo creates a MockPlanRepo pre-populated with the given rows.
func NewMockPlanRepo(rows ...*PlanRow) *MockPlanRepo {
	m := &MockPlanRepo{records: make(map[string]*PlanRow)}
	for _, r := range rows {
		m.records[r.ID] = r
	}
	return m
}

// FindByID returns the PlanRow with the given ID, or ErrNotFound.
func (m *MockPlanRepo) FindByID(_ context.Context, id string) (*PlanRow, error) {
	row, ok := m.records[id]
	if !ok {
		return nil, ErrNotFound
	}
	return row, nil
}

// MockStatementRepo is an in-memory StatementRepository for testing.
type MockStatementRepo struct {
	records map[string]*StatementRow
}

// NewMockStatementRepo creates a MockStatementRepo pre-populated with the given rows.
func NewMockStatementRepo(rows ...*StatementRow) *MockStatementRepo {
	m := &MockStatementRepo{records: make(map[string]*StatementRow)}
	for _, r := range rows {
		m.records[r.ID] = r
	}
	return m
}

// FindByID returns the StatementRow with the given ID, or ErrNotFound.
func (m *MockStatementRepo) FindByID(_ context.Context, id string) (*StatementRow, error) {
	row, ok := m.records[id]
	if !ok {
		return nil, ErrNotFound
	}
	return row, nil
}

// ListByCustomerID returns all StatementRows for the given customer, filtered by the query parameters.
func (m *MockStatementRepo) ListByCustomerID(_ context.Context, customerID string, q StatementQuery) ([]*StatementRow, int, error) {
	var statements []*StatementRow
	for _, r := range m.records {
		if r.DeletedAt != nil {
			continue
		}

		if r.CustomerID == customerID {
			statements = append(statements, r)
		}
	}

	if statements == nil {
		return nil, 0, nil
	}

	var filtered []*StatementRow
	for _, r := range statements {
		if r.SubscriptionID != q.SubscriptionID && q.SubscriptionID != "" {
			continue
		}

		if r.Kind != q.Kind && q.Kind != "" {
			continue
		}

		if r.Status != q.Status && q.Status != "" {
			continue
		}

		if r.PeriodStart < q.StartAfter && q.StartAfter != "" {
			continue
		}

		if r.PeriodEnd > q.EndBefore && q.EndBefore != "" {
			continue
		}

		filtered = append(filtered, r)
	}

	if q.PageSize <= 0 || q.Page <= 0 {
		q.PageSize = 10
		q.Page = 1
	}

	offsetEnd := q.Page * q.PageSize
	offsetStart := offsetEnd - q.PageSize
	if offsetStart >= len(filtered) {
		return nil, len(filtered), nil
	}

	if offsetEnd > len(filtered) {
		return filtered[offsetStart:], len(filtered), nil
	}

	final := filtered[offsetStart:offsetEnd]

	return final, len(filtered), nil
}
