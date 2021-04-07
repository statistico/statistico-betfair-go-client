package betfair

import (
	"context"
	"github.com/stretchr/testify/mock"
)

type MockClient struct {
	mock.Mock
}

func (m *MockClient) AccountFunds(ctx context.Context) (*AccountFunds, error) {
	args := m.Called(ctx)
	return args.Get(0).(*AccountFunds), args.Error(1)
}

func (m *MockClient) AccountDetails(ctx context.Context) (*AccountDetails, error) {
	args := m.Called(ctx)
	return args.Get(0).(*AccountDetails), args.Error(1)
}

func (m *MockClient) ListCompetitions(ctx context.Context, req ListCompetitionsRequest) ([]CompetitionResult, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]CompetitionResult), args.Error(1)
}

func (m *MockClient) ListEventTypes(ctx context.Context, req ListEventTypesRequest) ([]EventTypeResult, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]EventTypeResult), args.Error(1)
}

func (m *MockClient) ListEvents(ctx context.Context, req ListEventsRequest) ([]EventResult, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]EventResult), args.Error(1)
}

func (m *MockClient) ListMarketCatalogue(ctx context.Context, req ListMarketCatalogueRequest) ([]MarketCatalogue, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]MarketCatalogue), args.Error(1)
}

func (m *MockClient) ListMarketBook(ctx context.Context, req ListMarketBookRequest) ([]MarketBook, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]MarketBook), args.Error(1)
}

func (m *MockClient) ListRunnerBook(ctx context.Context, req ListRunnerBookRequest) ([]MarketBook, error) {
	args := m.Called(ctx, req)
	return args.Get(0).([]MarketBook), args.Error(1)
}

func (m *MockClient) PlaceOrder(ctx context.Context, req PlaceOrderRequest) (*PlaceExecutionReport, error) {
	args := m.Called(ctx, req)
	return args.Get(0).(*PlaceExecutionReport), args.Error(1)
}
