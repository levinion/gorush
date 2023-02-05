package log

import "go.uber.org/zap"

// 初始化logger
func InitLogger() {
	logger, err := zap.NewDevelopment(zap.WithCaller(false))
	if err != nil {
		panic(err)
	}
	defer logger.Sync()
	//替换全局logger，使用zap.L()调用
	zap.ReplaceGlobals(logger)
}
