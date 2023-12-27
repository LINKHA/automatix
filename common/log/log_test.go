package log_test

import (
	"testing"

	"github.com/LINKHA/automatix/common/log"
)

func TestStdlog(t *testing.T) {

	//测试 默认debug输出
	log.Debug("zinx debug content1")
	log.Debug("zinx debug content2")

	log.Debugf(" zinx debug a = %d\n", 10)

	//设置log标记位，加上长文件名称 和 微秒 标记
	log.ResetFlags(log.BitDate | log.BitLongFile | log.BitLevel)
	log.Info("zinx info content")

	//设置日志前缀，主要标记当前日志模块
	log.SetPrefix("MODULE")
	log.Error("zinx error content")

	//添加标记位
	log.AddFlag(log.BitShortFile | log.BitTime)
	log.Stack(" Zinx Stack! ")

	//设置日志写入文件
	log.SetLogFile("./log", "testfile.log")
	log.Debug("===> zinx debug content ~~666")
	log.Debug("===> zinx debug content ~~888")
	log.Error("===> zinx Error!!!! ~~~555~~~")

	//调试隔离级别
	log.Debug("=================================>")
	//1.debug
	log.SetLogLevel(log.LogInfo)
	log.Debug("===> 调试Debug：debug不应该出现")
	log.Info("===> 调试Debug：info应该出现")
	log.Warn("===> 调试Debug：warn应该出现")
	log.Error("===> 调试Debug：error应该出现")
	//2.info
	log.SetLogLevel(log.LogWarn)
	log.Debug("===> 调试Info：debug不应该出现")
	log.Info("===> 调试Info：info不应该出现")
	log.Warn("===> 调试Info：warn应该出现")
	log.Error("===> 调试Info：error应该出现")
	//3.warn
	log.SetLogLevel(log.LogError)
	log.Debug("===> 调试Warn：debug不应该出现")
	log.Info("===> 调试Warn：info不应该出现")
	log.Warn("===> 调试Warn：warn不应该出现")
	log.Error("===> 调试Warn：error应该出现")
	//4.error
	log.SetLogLevel(log.LogPanic)
	log.Debug("===> 调试Error：debug不应该出现")
	log.Info("===> 调试Error：info不应该出现")
	log.Warn("===> 调试Error：warn不应该出现")
	log.Error("===> 调试Error：error不应该出现")
}

func Testlogger(t *testing.T) {
}
