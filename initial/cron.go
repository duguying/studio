package initial

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/toolbox"
	"github.com/gogather/com"
	"github.com/gogather/com/log"
	"github.com/gogather/statistics"
)

func InitCron() {
	tk := toolbox.NewTask("statistic", "0 0 * * * *", githubStat)

	if beego.AppConfig.String("runmode") == "dev" {
		err := tk.Run()
		if err != nil {
			log.Warnln("[Run Task Failed]")
			log.Warnln(err)
		}
	}

	toolbox.AddTask("statistic", tk)
	toolbox.StartTask()
}

func githubStat() error {
	if !com.FileExist("static") {
		com.Mkdir("static")
	} else {
		if com.FileExist("static/upload") {
			com.Mkdir("static/upload")
		}
	}

	token := beego.AppConfig.String("github_token")
	user := beego.AppConfig.String("github_user")

	json, err := statistics.GetRepos(user, token)

	if err != nil {
		return err
	}

	stat := beego.AppConfig.String("github_statistics")

	err = com.WriteFile(stat, json)

	return err
}
