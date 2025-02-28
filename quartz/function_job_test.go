package quartz_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/reugn/go-quartz/quartz"
)

func TestFunctionJob(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var n = 2
	funcJob1 := quartz.NewFunctionJob(func(_ context.Context) (string, error) {
		n += 2
		return "fired1", nil
	})

	funcJob2 := quartz.NewFunctionJob(func(_ context.Context) (int, error) {
		n += 2
		return 42, nil
	})

	sched := quartz.NewStdScheduler()
	sched.Start(ctx)
	sched.ScheduleJob(ctx, funcJob1, quartz.NewRunOnceTrigger(time.Millisecond*300))
	sched.ScheduleJob(ctx, funcJob2, quartz.NewRunOnceTrigger(time.Millisecond*800))
	time.Sleep(time.Second)
	sched.Clear()
	sched.Stop()

	assertEqual(t, funcJob1.JobStatus, quartz.OK)
	assertNotEqual(t, funcJob1.Result, nil)
	assertEqual(t, *funcJob1.Result, "fired1")

	assertEqual(t, funcJob2.JobStatus, quartz.OK)
	assertNotEqual(t, funcJob2.Result, nil)
	assertEqual(t, *funcJob2.Result, 42)

	assertEqual(t, n, 6)
}

func TestNewFunctionJobWithDescAndKey(t *testing.T) {
	jobDesc := "test job"

	funcJob1 := quartz.NewFunctionJobWithDesc(jobDesc, func(_ context.Context) (string, error) {
		return "fired1", nil
	})

	funcJob2 := quartz.NewFunctionJobWithDesc(jobDesc, func(_ context.Context) (string, error) {
		return "fired2", nil
	})

	assertEqual(t, funcJob1.Description(), jobDesc)
	assertEqual(t, funcJob1.Key(), funcJob1.Key())
	assertNotEqual(t, funcJob1.Key(), funcJob2.Key())
}

func TestFunctionJobRespectsContext(t *testing.T) {
	var n int
	funcJob2 := quartz.NewFunctionJob(func(ctx context.Context) (bool, error) {
		timer := time.NewTimer(time.Hour)
		defer timer.Stop()
		select {
		case <-ctx.Done():
			n--
			return false, ctx.Err()
		case <-timer.C:
			n++
			return true, nil
		}
	})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sig := make(chan struct{})
	go func() { defer close(sig); funcJob2.Execute(ctx) }()

	if n != 0 {
		t.Fatal("job should not have run yet")
	}
	cancel()
	<-sig

	if n != -1 {
		t.Fatal("job side effect should have reflected cancelation:", n)
	}
	if !errors.Is(funcJob2.Error, context.Canceled) {
		t.Fatal("unexpected error function", funcJob2.Error)
	}
	if funcJob2.Result != nil {
		t.Fatal("errored jobs should not return values")
	}
}
