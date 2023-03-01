package test_helper

import (
	"context"
	"github.com/golang/mock/gomock"
	"testing"
)

func RunTestParallel(name string, t *testing.T, f func(t *testing.T)) {
	t.Helper()
	t.Run(name, func(t *testing.T) {
		t.Parallel()
		f(t)
	})
}

func RunTestWithGoMock(t *testing.T, f func(ctrl *gomock.Controller)) {
	ctx, cancel := context.WithCancel(context.Background())
	ctrl, ctx := gomock.WithContext(ctx, t)
	defer ctrl.Finish()

	go func() {
		f(ctrl)
		cancel()
	}()

	<-ctx.Done()
}
