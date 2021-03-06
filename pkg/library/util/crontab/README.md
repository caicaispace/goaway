```go
package main

import (
	"log"

	"comma/library/util/crontab"
)

var (
	cron = cronlib.New()
)

func main() {
	handleClean()
	go start()

	// cron already start, dynamic add job
	handleBackup()

	select {}
}

func start() {
	cron.Start()
	cron.Wait()
}

func handleClean() {
	job, err := cronlib.NewJobModel(
		"*/5 * * * * *",
		func() {
			pstdout("do clean gc action")
		},
	)
	if err != nil {
		panic(err.Error())
	}

	err = cron.Register("clean", job)
	if err != nil {
		panic(err.Error())
	}
}

func handleBackup() {
	job, err := cronlib.NewJobModel(
		"*/5 * * * * *",
		func() {
			pstdout("do backup action")
		},
	)
	if err != nil {
		panic(err.Error())
	}

	err = cron.DynamicRegister("backup", job)
	if err != nil {
		panic(err.Error())
	}
}

func pstdout(srv string) {
	log.Println(srv)
}
```

### set job attr

open async mode and try catch mode

```go
func run() error {
	cron := cronlib.New()

	// set async mode
	job, err = cronlib.NewJobModel(
		"0 * * * * *",
		func(),
		cronlib.AsyncMode(),
		cronlib.TryCatchMode(),
	)

	...
}
```

other method

```go
func run() error {
	cron := cronlib.New()

	// set async mode
	job, err = cronlib.NewJobModel(
		"0 * * * * *",
		func(),
	)

	...

	job.SetTryCatch(cronlib.OnMode)
	job.SetAsyncMode(cronlib.OnMode)

	...
}

```

### stop job

```go
cron := cronlib.New()
...
cron.StopService(srvName)
```

### update job

```go
spec := "*/3 * * * * *"
srv := "risk.scan.total.5s.to.3s"

job, _ := cronlib.NewJobModel(
	spec,
	func() {
		stdout(srv, spec)
	},
)

err := cron.UpdateJobModel(srv, job)
...
```

## Example

```go
package main

// test for crontab spec

import (
	"log"
	"time"

	"comma/library/util/crontab"
)

func main() {
	cron := cronlib.New()

	specList := map[string]string{
		"risk.scan.total.1s":       "*/1 * * * * *",
		"risk.scan.total.2s":       "*/2 * * * * *",
		"risk.scan.total.3s":       "*/3 * * * * *",
		"risk.scan.total.4s":       "*/4 * * * * *",
		"risk.scan.total.5s.to.3s": "*/5 * * * * *",
	}

	for srv, spec := range specList {
		tspec := spec // copy
		ssrv := srv   // copy
		job, err := cronlib.NewJobModel(
			spec,
			func() {
				stdout(ssrv, tspec)
			},
		)
		if err != nil {
			panic(err.Error())
		}

		err = cron.Register(srv, job)
		if err != nil {
			panic(err.Error())
		}
	}

	// update test
	time.AfterFunc(10*time.Second, func() {
		spec := "*/3 * * * * *"
		srv := "risk.scan.total.5s.to.3s"
		log.Println("reset 5s to 3s", srv)
		job, _ := cronlib.NewJobModel(
			spec,
			func() {
				stdout(srv, spec)
			},
		)
		cron.UpdateJobModel(srv, job)
		log.Println("reset finish", srv)

	})

	// kill test
	time.AfterFunc(3*time.Second, func() {

		srv := "risk.scan.total.1s"
		log.Println("stoping", srv)
		cron.StopService(srv)
		log.Println("stop finish", srv)

	})

	time.AfterFunc(11*time.Second, func() {

		srvPrefix := "risk"
		log.Println("stoping srv prefix", srvPrefix)
		cron.StopServicePrefix(srvPrefix)

	})

	cron.Start()
	cron.Wait()
}

func stdout(srv, spec string) {
	log.Println(srv, spec)
}

```