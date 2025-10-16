package testutil

import (
	"context"

	domain "github.com/flexprice/flexprice/internal/domain/pdf"
	"github.com/flexprice/flexprice/internal/logger"
	"github.com/flexprice/flexprice/internal/pdf"
	"github.com/flexprice/flexprice/internal/typst"
	"github.com/stretchr/testify/mock"
)

var _ pdf.Generator = (*MockPDFGenerator)(nil)

type MockPDFGenerator struct {
	logger *logger.Logger
	typst  typst.Compiler
	mock.Mock
}

// RenderInvoicePdf implements pdf.Generator.
func (m *MockPDFGenerator) RenderInvoicePdf(ctx context.Context, data *domain.InvoiceData) ([]byte, error) {
	args := m.Called(ctx, data)
	return args.Get(0).([]byte), args.Error(1)
}

func NewMockPDFGenerator(logger *logger.Logger) pdf.Generator {
	return &MockPDFGenerator{
		logger: logger,
		typst:  typst.DefaultCompiler(logger),
	}
}
