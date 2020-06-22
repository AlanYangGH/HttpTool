// 1. 全局logger
// 2. 支持日志分割
// 3. 接管os.StdOut
// 4. 可读性强的时间信息
// 5. Fatal或者Error打印调用栈
// 6. 显示文件和行号信息
// 7. 开发模式输出到os.Stdout

package internel

import (
	"os"

	"github.com/natefinch/lumberjack"

	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once       sync.Once
	logger     *zap.Logger
	globalUndo func()
	stdUndo    func()
)

func SetupLogger(filename string, level zapcore.Level) {
	once.Do(func() {
		caller := zap.AddCaller()
		development := zap.Development()
		stacktrace := zap.AddStacktrace(zap.ErrorLevel)

		core := initFileCore(filename, level)
		logger = zap.New(core, caller, development, stacktrace)

		// 替换zap的全局logger
		globalUndo = zap.ReplaceGlobals(logger)
		// 重定向log.Stdout
		stdUndo, _ = zap.RedirectStdLogAt(logger, zap.InfoLevel)
	})
}

func CloseLogger() {
	stdUndo()
	globalUndo()
	if err := logger.Sync(); err != nil {
		panic(err)
	}
}

func initFileCore(filename string, level zapcore.Level) zapcore.Core {
	writer := lumberjack.Logger{
		Filename:   filename, // 日志文件路径
		MaxSize:    2048,     // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: 100,      // 日志文件最多保存多少个备份
		MaxAge:     30,       // 文件最多保存多少天
		Compress:   false,    // 是否压缩
	}
	encoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	return zapcore.NewCore(encoder, zapcore.AddSync(&writer), zap.NewAtomicLevelAt(level))
}

func initConsoleCore(level zapcore.Level) zapcore.Core {
	// mutex保护os.Stdout
	syncer := zapcore.Lock(os.Stdout)

	// 默认选择development配置
	conf := zap.NewDevelopmentEncoderConfig()

	// 开发模式开启全路径，方便跳转
	conf.EncodeCaller = zapcore.FullCallerEncoder

	// 彩色终端
	conf.EncodeLevel = zapcore.CapitalColorLevelEncoder

	encoder := zapcore.NewConsoleEncoder(conf)
	return zapcore.NewCore(encoder, syncer, zap.NewAtomicLevelAt(level))
}
