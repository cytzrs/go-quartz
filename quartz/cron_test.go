package quartz_test

import (
	"fmt"
	"strconv"
	"testing"
	"time"

	"github.com/gorhill/cronexpr"
	"github.com/reugn/go-quartz/quartz"
)

func TestCronExpression1(t *testing.T) {
	prev := time.Date(2023, 4, 22, 12, 00, 00, 00, time.UTC).UnixNano()
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("10/20 15 14 5-10 * ? *")
	cronTrigger.Description()
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 50)
	}
	assertEqual(t, result, "Sun Jul 9 14:15:30 2023")
}

func TestCronExpression2(t *testing.T) {
	prev := time.Date(2023, 4, 22, 12, 00, 00, 00, time.UTC).UnixNano()
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("* 5,7,9 14-16 * * ? *")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 50)
	}
	assertEqual(t, result, "Sat Apr 22 14:05:49 2023")
}

func TestCronExpression3(t *testing.T) {
	prev := time.Date(2023, 4, 22, 12, 00, 00, 00, time.UTC).UnixNano()
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("* 5,7,9 14/2 ? * WED,Sat *")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 50)
	}
	assertEqual(t, result, "Sat Apr 22 14:05:49 2023")
}

func TestCronExpression4(t *testing.T) {
	expression := "0 5,7 14 1 * Sun *"
	_, err := quartz.NewCronTrigger(expression)
	if err == nil {
		t.Fatalf("%s should fail", expression)
	}
}

func TestCronExpression5(t *testing.T) {
	prev := time.Date(2023, 4, 22, 12, 00, 00, 00, time.UTC).UnixNano()
	fmt.Println(time.Unix(0, prev).UTC())
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("* * * * * ? *")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 50)
	}
	assertEqual(t, result, "Sat Apr 22 12:00:50 2023")
}

func TestCronExpression6(t *testing.T) {
	prev := time.Date(2023, 4, 22, 12, 00, 00, 00, time.UTC).UnixNano()
	fmt.Println(time.Unix(0, prev).UTC())
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("* * 14/2 ? * mon/3 *")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 50)
	}
	assertEqual(t, result, "Mon Apr 24 14:00:49 2023")
}

func TestCronExpression7(t *testing.T) {
	prev := time.Date(2023, 4, 22, 12, 00, 00, 00, time.UTC).UnixNano()
	fmt.Println(time.Unix(0, prev).UTC())
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("* 5-9 14/2 ? * 1-3 *")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 50)
	}
	assertEqual(t, result, "Sun Apr 23 14:05:49 2023")
}

func TestCronExpression8(t *testing.T) {
	prev := time.Date(2023, 4, 22, 12, 00, 00, 00, time.UTC).UnixNano()
	fmt.Println(time.Unix(0, prev).UTC())
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("*/3 */51 */12 */2 */4 ? *")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 50)
	}
	assertEqual(t, result, "Mon May 1 12:00:27 2023")
}

func TestCronExpression9(t *testing.T) {
	prev := time.Date(2023, 4, 22, 12, 00, 00, 00, time.UTC).UnixNano()
	fmt.Println(time.Unix(0, prev).UTC())
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("*/15 * * ? * 1-7")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 50)
	}
	assertEqual(t, result, "Sat Apr 22 12:12:30 2023")
}

func TestCronExpression10(t *testing.T) {
	prev := time.Date(2023, 4, 22, 12, 00, 00, 00, time.UTC).UnixNano()
	fmt.Println(time.Unix(0, prev).UTC())
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("10,20 10,20 10,20 10,20 6,12 ?")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 50)
	}
	assertEqual(t, result, "Tue Dec 10 10:10:20 2024")
}

func TestCronExpression11(t *testing.T) {
	prev := time.Date(2023, 4, 22, 12, 00, 00, 00, time.UTC).UnixNano()
	fmt.Println(time.Unix(0, prev).UTC())
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("10,20 10,20 10,20 ? 6,12 3,6")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 50)
	}
	assertEqual(t, result, "Fri Jun 23 10:10:20 2023")
}

func TestCronExpression12(t *testing.T) {
	prev := time.Date(2023, 4, 22, 12, 00, 00, 00, time.UTC).UnixNano()
	fmt.Println(time.Unix(0, prev).UTC())
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("0 0 0 ? 4,6 SAT,MON")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 50)
	}
	assertEqual(t, result, "Sat Apr 18 00:00:00 2026")
}

func TestCronExpression13(t *testing.T) {
	prev := time.Date(2023, 4, 22, 12, 00, 00, 00, time.UTC).UnixNano()
	fmt.Println(time.Unix(0, prev).UTC())
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("0 0 0 29 2 ?")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 5)
	}
	assertEqual(t, result, "Wed Feb 29 00:00:00 2040")
}

func TestCronExpression14(t *testing.T) {
	prev := time.Date(2023, 4, 22, 12, 00, 00, 00, time.UTC).UnixNano()
	fmt.Println(time.Unix(0, prev).UTC())
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("0 0 0 1 5 ? 2023/2")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 10)
	}
	assertEqual(t, result, "Wed May 1 00:00:00 2041")
}

func TestCronExpressionWithLoc(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	prev := time.Date(2023, 4, 29, 12, 00, 00, 00, loc).UnixNano()
	fmt.Println(time.Unix(0, prev).UTC())
	result := ""
	if err != nil {
		t.Fatal(err)
	}
	cronTrigger, err := quartz.NewCronTriggerWithLoc("* 5 22-23 * * Sun *", loc)
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 50)
	}
	assertEqual(t, result, "Mon May 1 02:05:49 2023") // Result comparison is in UTC time
}

func TestCronExpressionWithLoc2(t *testing.T) {
	loc, err := time.LoadLocation("America/New_York")
	prev := time.Date(2023, 4, 29, 12, 00, 00, 00, loc).UnixNano()
	fmt.Println(time.Unix(0, prev).UTC())
	result := ""
	if err != nil {
		t.Fatal(err)
	}
	cronTrigger, err := quartz.NewCronTriggerWithLoc("0 0 10 * * Sun *", loc)
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 50)
	}
	assertEqual(t, result, "Sun Apr 7 14:00:00 2024")
}

func TestCronDaysOfWeek(t *testing.T) {
	daysOfWeek := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}
	expected := []string{
		"Sun Apr 21 00:00:00 2019",
		"Mon Apr 22 00:00:00 2019",
		"Tue Apr 23 00:00:00 2019",
		"Wed Apr 24 00:00:00 2019",
		"Thu Apr 18 00:00:00 2019",
		"Fri Apr 19 00:00:00 2019",
		"Sat Apr 20 00:00:00 2019",
	}

	for i := 0; i < len(daysOfWeek); i++ {
		cronDayOfWeek(t, daysOfWeek[i], expected[i])
		cronDayOfWeek(t, strconv.Itoa(i+1), expected[i])
	}
}

func cronDayOfWeek(t *testing.T, dayOfWeek, expected string) {
	prev := int64(1555524000000000000) // Wed Apr 17 18:00:00 2019
	expression := fmt.Sprintf("0 0 0 * * %s", dayOfWeek)
	cronTrigger, err := quartz.NewCronTrigger(expression)
	if err != nil {
		t.Fatal(err)
	} else {
		nextFireTime, err := cronTrigger.NextFireTime(prev)
		if err != nil {
			t.Fatal(err)
		} else {
			assertEqual(t, time.Unix(nextFireTime/int64(time.Second), 0).UTC().Format(readDateLayout),
				expected)
		}
	}
}

func TestCronYearly(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("@yearly")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 100)
	}
	assertEqual(t, result, "Sun Jan 1 00:00:00 2119")
}

func TestCronMonthly(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("@monthly")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 100)
	}
	assertEqual(t, result, "Sun Aug 1 00:00:00 2027")
}

func TestCronWeekly(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("@weekly")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 100)
	}
	assertEqual(t, result, "Sun Mar 14 00:00:00 2021")
}

func TestCronDaily(t *testing.T) {
	prev := int64(1555351200000000000)
	fmt.Println(time.Unix(0, prev).UTC())
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("@daily")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 1000)
	}
	assertEqual(t, result, "Sun Jan 9 00:00:00 2022")
}

func TestCronHourly(t *testing.T) {
	prev := int64(1555351200000000000)
	result := ""
	cronTrigger, err := quartz.NewCronTrigger("@hourly")
	if err != nil {
		t.Fatal(err)
	} else {
		result, _ = iterate(prev, cronTrigger, 1000)
	}
	assertEqual(t, result, "Mon May 27 10:00:00 2019")
}

var readDateLayout = "Mon Jan 2 15:04:05 2006"

func iterate(prev int64, cronTrigger *quartz.CronTrigger, iterations int) (string, error) {
	var err error
	for i := 0; i < iterations; i++ {
		prev, err = cronTrigger.NextFireTime(prev)
		// fmt.Println(time.Unix(prev/int64(time.Second), 0).UTC().Format(readDateLayout))
		if err != nil {
			fmt.Println(err)
			return "", err
		}
	}
	return time.Unix(prev/int64(time.Second), 0).UTC().Format(readDateLayout), nil
}

func TestCronExpressionError(t *testing.T) {
	tests := []string{
		"*/X * * * * *",
	}
	for _, test := range tests {
		t.Run(test, func(t *testing.T) {
			_, err := quartz.NewCronTrigger(test)
			if err == nil {
				t.Error("expected error")
			}
		})
	}
}

// 代表一个任务
type CronJob struct {
	id       string
	expr     *cronexpr.Expression
	nextTime time.Time // expr.Next(now)
	doWork   func(id string) string
}

var scheduleTable = make(map[string]*CronJob)
var now = time.Now()

func TestTain(t *testing.T) {
	// 定时任务字典， key: 任务的名字, value 任务对象

	// 定义定时任务以每 5s 执行一次
	// MustParse 如果遇到解析 contab 错误时会直接抛出 panic ，不会像 Parse 一样返回一个错误
	expr := cronexpr.MustParse("*/15 * * * * * *")
	cronJob := &CronJob{
		id:       "1",
		expr:     expr,
		nextTime: expr.Next(now),
		doWork: func(id string) string {
			fmt.Println("id: ", id)
			return id
		},
	}
	// 任务注册到调度表
	scheduleTable["job1"] = cronJob

	// 定义定时任务以每 3s 执行一次
	expr = cronexpr.MustParse("*/3 * * * * * *")
	cronJob = &CronJob{
		id:       "2",
		expr:     expr,
		nextTime: expr.Next(now),
		doWork: func(id string) string {
			fmt.Println("id2: ", id)
			return id
		},
	}
	// 任务注册到调度表
	scheduleTable["job2"] = cronJob
	register("job3")

	var i int = 0
	// 定时检查一下任务调度表
	for {
		now := time.Now()

		for jobName, cronJob := range scheduleTable {
			// 判断是否到期，当前时间等于定时任务的下次执行时间，或者当前时间大于任务的定时时间
			if now.Equal(cronJob.nextTime) || now.After(cronJob.nextTime) {
				// 启动一个协程, 执行这个任务
				go func(jobName string) {
					fmt.Println("执行:", jobName)
				}(jobName)

				i++
				if i == 5 {
					degister("job5")
				}

				if i == 3 {
					register("job5")
				}

				// 计算下一次调度时间
				cronJob.nextTime = cronJob.expr.Next(now)
				cronJob.doWork(cronJob.id)
				fmt.Println(jobName, "下次执行时间:", cronJob.nextTime)
			}
		}

		// 等待 1s，减少 CPU 消耗
		t := <-time.NewTimer(1 * time.Second).C
		fmt.Println(t)
	}

}

func register(key string) {
	// 定义定时任务以每 5s 执行一次
	// MustParse 如果遇到解析 contab 错误时会直接抛出 panic ，不会像 Parse 一样返回一个错误
	expr := cronexpr.MustParse("*/5 * * * * * *")
	cronJob := &CronJob{
		id:       key,
		expr:     expr,
		nextTime: expr.Next(now),
		doWork: func(id string) string {
			fmt.Println("id: ", id)
			return id
		},
	}
	// 任务注册到调度表
	scheduleTable[key] = cronJob
}

func degister(key string) {

	delete(scheduleTable, key)
	fmt.Println("fsdfsdf")
}
