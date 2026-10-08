package model

import "testing"

func TestConstants_OrderStatusPending(t *testing.T) {
	if OrderStatusPending != "pending" {
		t.Errorf("expected OrderStatusPending = 'pending', got %q", OrderStatusPending)
	}
}

func TestConstants_NotificationFormats(t *testing.T) {
	tests := []struct {
		name   string
		format string
		to     string
		msg    string
	}{
		{"EmailNotifyFormat", EmailNotifyFormat, "user@example.com", "Hello"},
		{"SMSNotifyFormat", SMSNotifyFormat, "+79001234567", "Hello"},
		{"TelegramNotifyFormat", TelegramNotifyFormat, "@user", "Hello"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.format
			if got == "" {
				t.Errorf("expected non-empty format for %s", tt.name)
			}
		})
	}
}

func TestConstants_OrderCreatedMessage(t *testing.T) {
	if OrderCreatedMessage == "" {
		t.Error("OrderCreatedMessage should not be empty")
	}
}

func TestEmailSender_Send(t *testing.T) {
	sender := NewEmailSender()
	err := sender.Send("test@example.com", "Hello")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestEmailSender_Notify(t *testing.T) {
	sender := NewEmailSender()
	err := sender.Notify("test@example.com", "Order confirmed")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSMSSender_Send(t *testing.T) {
	sender := NewSMSSender()
	err := sender.Send("+79001234567", "Hello")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestSMSSender_Notify(t *testing.T) {
	sender := NewSMSSender()
	err := sender.Notify("+79001234567", "Order confirmed")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestTelegramNotifier_Send(t *testing.T) {
	notifier := NewTelegramNotifier()
	err := notifier.Send("@user", "Hello")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestTelegramNotifier_Notify(t *testing.T) {
	notifier := NewTelegramNotifier()
	err := notifier.Notify("@user", "Order confirmed")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestOrderStruct(t *testing.T) {
	order := Order{
		ID:       1,
		Customer: "Иван",
		Products: "apple,banana",
		Total:    10.5,
		Status:   OrderStatusPending,
	}

	if order.Customer != "Иван" {
		t.Errorf("expected Customer = 'Иван', got %q", order.Customer)
	}
	if order.Status != OrderStatusPending {
		t.Errorf("expected Status = %q, got %q", OrderStatusPending, order.Status)
	}
	if order.Total != 10.5 {
		t.Errorf("expected Total = 10.5, got %f", order.Total)
	}
}
