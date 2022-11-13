package notifications_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/gabrielsc1998/inventory-control-golang/internal/seedwork/domain/notifications"
)

func TestShouldAddErrors(t *testing.T) {
	notification := notifications.NewNotification()

	const ctx = "ctx"

	notification.AddError(ctx, "msg 1")
	notification.AddError(ctx, "msg 2")

	assert.True(t, notification.HasErrors())
	assert.Contains(t, notification.GetErrors(ctx), ctx+": msg 1, msg 2")

	const ctx2 = "ctx2"

	notification.AddError(ctx2, "msg 1")
	notification.AddError(ctx2, "msg 2")

	assert.Contains(t, notification.GetErrors(ctx2), ctx2+": msg 1, msg 2")
	assert.Equal(t, notification.GetErrors(), ctx+": msg 1, msg 2\n"+ctx2+": msg 1, msg 2\n")
}

func TestShouldReturnAnEmptyStringWhenTheErrorsAreEmtpy(t *testing.T) {
	notification := notifications.NewNotification()

	const ctx = "ctx"

	assert.False(t, notification.HasErrors())
	assert.Equal(t, notification.GetErrors(ctx), "")
}
