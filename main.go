package main

import (
	"github.com/astaxie/beego"
	"kitty/app/action/job"
)

func main() {


	beego.Router("/jobinfo/",&job.JobInfo{})
}
