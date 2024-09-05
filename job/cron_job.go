package job

import (
	service "stocks/service"

	"github.com/robfig/cron"
)

func CheckAgain() {
	cr := cron.New()
	cr.AddFunc("@every 00h01m00s", service.VerifyAndSearchForStocks)
	cr.Start()
}
