// generated-from:104b79afd9df75dea2d8a9a296e587b968f198fd376a8947a11c0b3c142f3b4e DO NOT REMOVE, DO UPDATE

package test

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/moov-io/base/log"
	"github.com/moov-io/base/stime"
	"github.com/stretchr/testify/require"

	"github.com/moov-io/fincen/pkg/service"
)

func Ctx() context.Context {
	return context.Background()
}

type TestEnvironment struct {
	T          *testing.T
	Assert     *require.Assertions
	StaticTime stime.StaticTimeService
	AccountID  string

	service.Environment
}

func NewEnvironment(t *testing.T, router *mux.Router) *TestEnvironment {
	testEnv := &TestEnvironment{}

	testEnv.T = t
	testEnv.PublicRouter = router
	testEnv.Assert = require.New(t)
	testEnv.Logger = log.NewDefaultLogger()
	testEnv.StaticTime = stime.NewStaticTimeService()
	testEnv.TimeService = testEnv.StaticTime
	testEnv.AccountID = uuid.NewString()

	cfg, err := service.LoadConfig(testEnv.Logger)
	if err != nil {
		t.Fatal(err)
	}
	testEnv.Config = cfg

	_, err = service.NewEnvironment(&testEnv.Environment)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(testEnv.Shutdown)

	return testEnv
}
