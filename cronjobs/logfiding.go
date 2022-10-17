package cronjobs

import (
	"fmt"
	"kubecontroller/kcontrollers"
	"kubecontroller/models"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/robfig/cron/v3"
)

func RunCronjob() {
	c := cron.New()
	c.AddFunc("*/1 * * * *",
		func() {
			log.Println("UploadTraffix")
			fmt.Println("UploadTraffix")
			fmt.Println(os.Getenv("INGRESSNAMESPACE"))
			logs, err := kcontrollers.GetPodLogs(os.Getenv("INGRESSNAMESPACE"), os.Getenv("INGRESSPODNAME"))
			if err != nil {
				log.Println(err)
			}
			//logs := "10.200.9.115 - - [17/Oct/2022:10:40:30 +0000] \"POST /adminapi/reportsservice/logs/Notification HTTP/2.0\" 200 428 \"https://csdpdev-admin.d21.co.in/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36\" 877 0.303 [default-reportservice-80] [] 10.200.9.121:5006 450 0.300 200 04f6fe5b641fe5fe94f1b21e0393c0f0\n10.200.9.115 - - [17/Oct/2022:10:40:33 +0000] \"OPTIONS /adminapi/accessservice/rightspolicy/GetAccess HTTP/2.0\" 204 0 \"https://csdpdev-admin.d21.co.in/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36\" 56 0.002 [default-accessservice-80] [] 10.200.9.98:5002 0 0.000 204 06b2300652a3d1809ef8c65e29ba3abc\n10.200.9.115 - - [17/Oct/2022:10:40:33 +0000] \"OPTIONS /adminapi/packsservice/cycle/get HTTP/2.0\" 204 0 \"http://localhost:4200/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36\" 47 0.001 [default-packservice-80] [] 10.200.9.171:5004 0 0.000 204 1c04783ec1dff026aa7bacd4f5118ae2\n10.200.9.115 - - [17/Oct/2022:10:40:33 +0000] \"POST /adminapi/accessservice/rightspolicy/GetAccess HTTP/2.0\" 200 452 \"https://csdpdev-admin.d21.co.in/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36\" 198 0.011 [default-accessservice-80] [] 10.200.9.98:5002 474 0.012 200 d2fe3e63451c700db05c383a41d558d6\n10.200.9.115 - - [17/Oct/2022:10:40:33 +0000] \"POST /adminapi/packsservice/cycle/get HTTP/2.0\" 200 1820 \"http://localhost:4200/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36\" 449 0.013 [default-packservice-80] [] 10.200.9.171:5004 1877 0.012 200 205a8b471d718261e94d4df9b97ac277\n10.200.9.115 - - [17/Oct/2022:10:40:34 +0000] \"OPTIONS /adminapi/accessservice/user/GetUserStatusDashboard HTTP/2.0\" 204 0 \"https://csdpdev-admin.d21.co.in/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36\" 60 0.002 [default-accessservice-80] [] 10.200.9.98:5002 0 0.004 204 a2744ca9deaacf67882349cd47f71674\n10.200.9.115 - - [17/Oct/2022:10:40:34 +0000] \"OPTIONS /adminapi/assetsservice/masters/userroles HTTP/2.0\" 204 0 \"https://csdpdev-admin.d21.co.in/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36\" 52 0.004 [default-assetsservice-80] [] 10.200.9.104:5003 0 0.004 204 e1eeea92581f98ea239c13754e315d22\n10.200.9.115 - - [17/Oct/2022:10:40:34 +0000] \"OPTIONS /adminapi/accessservice/user/search HTTP/2.0\" 204 0 \"https://csdpdev-admin.d21.co.in/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36\" 48 0.002 [default-accessservice-80] [] 10.200.9.98:5002 0 0.004 204 a19e1f8ca78e9bb70a214dc8f29358c1\n10.200.9.115 - - [17/Oct/2022:10:40:34 +0000] \"POST /adminapi/assetsservice/masters/userroles HTTP/2.0\" 200 455 \"https://csdpdev-admin.d21.co.in/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36\" 206 0.005 [default-assetsservice-80] [] 10.200.9.104:5003 477 0.004 200 bcc34134ace2d265e851d3cbdd9fff6d\n10.200.9.115 - - [17/Oct/2022:10:40:34 +0000] \"POST /adminapi/accessservice/user/GetUserStatusDashboard HTTP/2.0\" 200 487 \"https://csdpdev-admin.d21.co.in/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36\" 243 0.010 [default-accessservice-80] [] 10.200.9.98:5002 504 0.012 200 7c5d1ef38abe351b9e554d4fcd19dec7\n10.200.9.115 - - [17/Oct/2022:10:40:34 +0000] \"POST /adminapi/accessservice/user/search HTTP/2.0\" 200 619 \"https://csdpdev-admin.d21.co.in/\" \"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/106.0.0.0 Safari/537.36\" 336 0.029 [default-accessservice-80] [] 10.200.9.98:5002 641 0.028 200 eec5981e01b20619096533356e0c5946\nW1017 10:40:43.647693       7 controller.go:1016] Error obtaining Endpoints for Service \"default/stable-grafana\": no object matching key \"default/stable-grafana\" in local store\n"
			split := strings.Split(logs, "\n")
			//fmt.Println(len(split))
			for _, line := range split {
				//fmt.Println(index)
				//fmt.Println(len(strings.Split(line, " ")))
				//fmt.Println(line)
				if len(strings.Split(line, " ")) == 32 {
					var traffic models.Traffic
					data := strings.Split(line, " ")
					//fmt.Println(data[3][1:])
					date := dateconvert(data[3][1:])
					timetaken, err := strconv.ParseFloat(data[29], 8)
					if err != nil {
						log.Println(err)
					}
					//fmt.Println(date)
					if result := models.DB.Where("Log = ?", line).First(&traffic).RowsAffected; result == 0 {
						tf := models.Traffic{IP: data[0], TimeTaken: timetaken, Date: date, Method: data[5][1:], Path: data[6], StatusCode: data[8], Log: line}
						models.DB.Create(&tf)
					}
				}
			}

		})
	log.Println("Start cronjob")
	c.Start()
}

func UploadTrafficLog() {
	log.Println("UploadTraffix")
	fmt.Println("UploadTraffix")
	fmt.Println(os.Getenv("INGRESSNAMESPACE"))
	logs, err := kcontrollers.GetPodLogs(os.Getenv("INGRESSNAMESPACE"), os.Getenv("INGRESSPODNAME"))
	if err != nil {
		log.Println(err)
	}
	split := strings.Split(logs, "\n")
	for index, line := range split {
		if len(strings.Split(line, " ")) != 0 && len(split) != index+1 {
			var traffic models.Traffic
			data := strings.Split(line, " ")
			if Is_valid_ip(data[0]) {
				date, err := time.Parse(data[3][1:], data[3][1:])
				if err != nil {
					log.Println(err)
				}
				if result := models.DB.Where("Log = ?", line).First(&traffic).RowsAffected; result == 0 {
					tf := models.Traffic{IP: data[0], Date: date, Method: data[5][1:], Path: data[6], StatusCode: data[8], Log: line}
					models.DB.Create(&tf)
				}

			}

		}
	}

}

func dateconvert(datet string) time.Time {
	time_spl := strings.Split(datet, ":")
	date_spl := strings.Split(time_spl[0], "/")
	var date_month time.Month
	if date_spl[1] == "Oct" {
		date_month = time.October
	}
	if date_spl[1] == "Jan" {
		date_month = time.January
	}
	if date_spl[1] == "Feb" {
		date_month = time.February
	}
	if date_spl[1] == "Mar" {
		date_month = time.March
	}
	if date_spl[1] == "Apr" {
		date_month = time.April
	}
	if date_spl[1] == "Jun" {
		date_month = time.June
	}
	if date_spl[1] == "Jul" {
		date_month = time.July
	}
	if date_spl[1] == "Aug" {
		date_month = time.August
	}
	if date_spl[1] == "May" {
		date_month = time.May
	}
	if date_spl[1] == "Sep" {
		date_month = time.September
	}
	if date_spl[1] == "Oct" {
		date_month = time.October
	}
	if date_spl[1] == "Nov" {
		date_month = time.November
	}
	if date_spl[1] == "Dec" {
		date_month = time.December
	}

	s_year, err := strconv.Atoi(date_spl[2])
	if err != nil {
		fmt.Println(err)
	}
	s_day, err := strconv.Atoi(date_spl[0])
	if err != nil {
		fmt.Println(err)
	}
	s_hours, err := strconv.Atoi(time_spl[1])
	if err != nil {
		fmt.Println(err)
	}
	s_minute, err := strconv.Atoi(time_spl[2])
	if err != nil {
		fmt.Println(err)
	}
	s_second, err := strconv.Atoi(time_spl[3])
	if err != nil {
		fmt.Println(err)
	}
	t := time.Date(s_year, date_month, s_day, s_hours, s_minute, s_second, 0, time.UTC)
	return t
}

func Is_valid_ip(ip string) bool {
	re, _ := regexp.Compile(`^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$`)
	if re.MatchString(ip) {
		return true
	}
	return false
}
