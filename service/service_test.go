package service

import (
	"errors"
	"testing"
)

// --- Mock-реализации для тестов ---

type mockRepo struct {
	createOrderCalled bool
	createOrderError  error
}

func (m *mockRepo) CreateOrder(customer string, products []string, total float64) error {
	m.createOrderCalled = true
	return m.createOrderError
}

type mockNotifier struct {
	notifyCalled bool
	notifyTo     string
	notifyMsg    string
	notifyError  error
	sendCalled   bool
	sendTo       string
	sendMsg      string
}

func (m *mockNotifier) Send(to string, message string) error {
	m.sendCalled = true
	m.sendTo = to
	m.sendMsg = message
	return m.notifyError
}

func (m *mockNotifier) Notify(to string, message string) error {
	m.notifyCalled = true
	m.notifyTo = to
	m.notifyMsg = message
	m.sendCalled = true
	m.sendTo = to
	m.sendMsg = message
	return m.notifyError
}

// --- Тесты ---

func TestCreateOrder_Success(t *testing.T) {
	mockRepo := &mockRepo{}
	mockNotifier := &mockNotifier{}

	svc := NewOrderService(mockRepo, mockNotifier)

	err := svc.CreateOrder("Иван", []string{"apple", "banana"}, 10.5)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !mockRepo.createOrderCalled {
		t.Error("expected repo.CreateOrder to be called")
	}

	if !mockNotifier.notifyCalled {
		t.Error("expected notifier.Notify to be called")
	}

	expectedMsg := "Заказ на сумму 10.50 успешно создан"
	if mockNotifier.notifyMsg != expectedMsg {
		t.Errorf("expected notify message %q, got %q", expectedMsg, mockNotifier.notifyMsg)
	}

	if mockNotifier.notifyTo != "Иван" {
		t.Errorf("expected notify to 'Иван', got %q", mockNotifier.notifyTo)
	}
}

func TestCreateOrder_RepoError(t *testing.T) {
	expectedErr := errors.New("database error")
	mockRepo := &mockRepo{createOrderError: expectedErr}
	mockNotifier := &mockNotifier{}

	svc := NewOrderService(mockRepo, mockNotifier)

	err := svc.CreateOrder("Мария", []string{"orange"}, 25.0)
	if err != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	if !mockRepo.createOrderCalled {
		t.Error("expected repo.CreateOrder to be called")
	}

	if mockNotifier.notifyCalled {
		t.Error("expected notifier.Notify NOT to be called when repo fails")
	}
}

func TestCreateOrder_NotifierError(t *testing.T) {
	expectedErr := errors.New("notification error")
	mockRepo := &mockRepo{}
	mockNotifier := &mockNotifier{notifyError: expectedErr}

	svc := NewOrderService(mockRepo, mockNotifier)

	err := svc.CreateOrder("Олег", []string{"coffee"}, 18.0)
	if err != expectedErr {
		t.Errorf("expected error %v, got %v", expectedErr, err)
	}

	if !mockRepo.createOrderCalled {
		t.Error("expected repo.CreateOrder to be called")
	}

	if !mockNotifier.notifyCalled {
		t.Error("expected notifier.Notify to be called")
	}
}

func TestCreateOrder_EmptyProducts(t *testing.T) {
	mockRepo := &mockRepo{}
	mockNotifier := &mockNotifier{}

	svc := NewOrderService(mockRepo, mockNotifier)

	err := svc.CreateOrder("Тест", []string{}, 0.0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if !mockRepo.createOrderCalled {
		t.Error("expected repo.CreateOrder to be called")
	}
}

func TestCreateOrder_ZeroTotal(t *testing.T) {
	mockRepo := &mockRepo{}
	mockNotifier := &mockNotifier{}

	svc := NewOrderService(mockRepo, mockNotifier)

	err := svc.CreateOrder("Бесплатный заказ", []string{"item"}, 0.0)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	expectedMsg := "Заказ на сумму 0.00 успешно создан"
	if mockNotifier.notifyMsg != expectedMsg {
		t.Errorf("expected notify message %q, got %q", expectedMsg, mockNotifier.notifyMsg)
	}
}

func TestCreateOrder_NotifyDelegatesToSend(t *testing.T) {
	mockRepo := &mockRepo{}
	mockNotifier := &mockNotifier{}

	svc := NewOrderService(mockRepo, mockNotifier)

	err := svc.CreateOrder("User", []string{"product"}, 100.0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Notify вызывает Send (проверяем, что Send тоже был вызван)
	if !mockNotifier.sendCalled {
		t.Error("expected notifier.Send to be called (Notify delegates to Send)")
	}
}

func TestNewOrderService(t *testing.T) {
	mockRepo := &mockRepo{}
	mockNotifier := &mockNotifier{}

	svc := NewOrderService(mockRepo, mockNotifier)

	if svc == nil {
		t.Fatal("expected non-nil service")
	}
}
