/**
 * @Author Mr.LiuQH
 * @Description TODO
 * @Date 2021/2/3 3:03 下午
 **/
package main

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type DeFaultFieldHook struct {

}

func (d *DeFaultFieldHook) Fire( entry *logrus.Entry) error  {
	entry.Data["append"] = time.Now().Unix()
	return nil
}

func (d *DeFaultFieldHook) Levels()  []logrus.Level {
	return logrus.AllLevels
}


func main() {
	logrus.AddHook(&DeFaultFieldHook{})
	// 设置日志级别
	logrus.SetLevel(logrus.DebugLevel)
	// 将日志输出到控制台打印出来
	logrus.SetOutput(os.Stdout)
	// 设置为true则显示日志在代码什么位置打印的
	logrus.SetReportCaller(true)
	//  设置日志以json格式输出，默认以text格式输出
	logrus.SetFormatter(&logrus.JSONFormatter{})
	logrus.Trace("这是Trace,日志信息")
	logrus.Debug("这是Debug,日志信息")
	logrus.Info("这是Info,日志信息")
	logrus.Error("这是Error,日志信息")
	logrus.Warn("这是Warn,日志信息")
	// log之后会调用调用 os.Exit(1)
	//logrus.Fatal("这是Fatal,日志信息")
	//// log之后会Painc()
	//logrus.Panic("这是Panic,日志信息")

	logrus.WithFields(logrus.Fields{
		"uerName":"zhangsan",
		"age":28,
		"pice":500.12,
		"likes":[]string{"游戏","旅游"},
	}).Info("请求入参信息~")

}