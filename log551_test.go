package log551_test
import (
	"github.com/go51/log551"
	"testing"
)

func TestNew(t *testing.T) {

	log1 := log551.New(nil)
	log2 := log551.New(nil)

	if log1 == nil {
		t.Error("インスタンス生成に失敗しました。")
	}

	if log2 == nil {
		t.Error("インスタンス生成に失敗しました。")
	}

	t.Logf("log1: [%p] %#v", &log1, log1)
	t.Logf("log2: [%p] %#v", &log2, log2)

	if &log1 == &log2 {
		t.Error("インスタンス生成に失敗しました。")
	}
}

func BenchmarkNew(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = log551.New(nil)
	}
}

func TestDebugEachTime(t *testing.T) {
	conf := log551.Config{
		Debug:log551.ConfigDetail{
			EachTime:true,
			Cache:false,
		},
		Information:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Warning:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Error:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Critical:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Path:"/var/log/gorai/test.log",
	}

	l := log551.New(&conf)
	l.Open()
	defer l.Close()

	l.Debug("test string")
	l.Debugf("%s %s", "test", "string")
	l.Debugln("test string")
	l.Information("test string")
	l.Informationf("%s %s", "test", "string")
	l.Informationln("test string")
	l.Warning("test string")
	l.Warningf("%s %s", "test", "string")
	l.Warningln("test string")
	l.Error("test string")
	l.Errorf("%s %s", "test", "string")
	l.Errorln("test string")
	l.Critical("test string")
	l.Criticalf("%s %s", "test", "string")
	l.Criticalln("test string")
}

func TestInformationEachTime(t *testing.T) {
	conf := log551.Config{
		Debug:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Information:log551.ConfigDetail{
			EachTime:true,
			Cache:false,
		},
		Warning:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Error:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Critical:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Path:"/var/log/gorai/test.log",
	}

	l := log551.New(&conf)
	l.Open()
	defer l.Close()

	l.Debug("test string")
	l.Debugf("%s %s", "test", "string")
	l.Debugln("test string")
	l.Information("test string")
	l.Informationf("%s %s", "test", "string")
	l.Informationln("test string")
	l.Warning("test string")
	l.Warningf("%s %s", "test", "string")
	l.Warningln("test string")
	l.Error("test string")
	l.Errorf("%s %s", "test", "string")
	l.Errorln("test string")
	l.Critical("test string")
	l.Criticalf("%s %s", "test", "string")
	l.Criticalln("test string")
}

func TestWarningEachTime(t *testing.T) {
	conf := log551.Config{
		Debug:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Information:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Warning:log551.ConfigDetail{
			EachTime:true,
			Cache:false,
		},
		Error:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Critical:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Path:"/var/log/gorai/test.log",
	}

	l := log551.New(&conf)
	l.Open()
	defer l.Close()

	l.Debug("test string")
	l.Debugf("%s %s", "test", "string")
	l.Debugln("test string")
	l.Information("test string")
	l.Informationf("%s %s", "test", "string")
	l.Informationln("test string")
	l.Warning("test string")
	l.Warningf("%s %s", "test", "string")
	l.Warningln("test string")
	l.Error("test string")
	l.Errorf("%s %s", "test", "string")
	l.Errorln("test string")
	l.Critical("test string")
	l.Criticalf("%s %s", "test", "string")
	l.Criticalln("test string")
}

func TestErrorEachTime(t *testing.T) {
	conf := log551.Config{
		Debug:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Information:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Warning:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Error:log551.ConfigDetail{
			EachTime:true,
			Cache:false,
		},
		Critical:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Path:"/var/log/gorai/test.log",
	}

	l := log551.New(&conf)
	l.Open()
	defer l.Close()

	l.Debug("test string")
	l.Debugf("%s %s", "test", "string")
	l.Debugln("test string")
	l.Information("test string")
	l.Informationf("%s %s", "test", "string")
	l.Informationln("test string")
	l.Warning("test string")
	l.Warningf("%s %s", "test", "string")
	l.Warningln("test string")
	l.Error("test string")
	l.Errorf("%s %s", "test", "string")
	l.Errorln("test string")
	l.Critical("test string")
	l.Criticalf("%s %s", "test", "string")
	l.Criticalln("test string")
}

func TestCriticalEachTime(t *testing.T) {
	conf := log551.Config{
		Debug:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Information:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Warning:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Error:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Critical:log551.ConfigDetail{
			EachTime:true,
			Cache:false,
		},
		Path:"/var/log/gorai/test.log",
	}

	l := log551.New(&conf)
	l.Open()
	defer l.Close()

	l.Debug("test string")
	l.Debugf("%s %s", "test", "string")
	l.Debugln("test string")
	l.Information("test string")
	l.Informationf("%s %s", "test", "string")
	l.Informationln("test string")
	l.Warning("test string")
	l.Warningf("%s %s", "test", "string")
	l.Warningln("test string")
	l.Error("test string")
	l.Errorf("%s %s", "test", "string")
	l.Errorln("test string")
	l.Critical("test string")
	l.Criticalf("%s %s", "test", "string")
	l.Criticalln("test string")
}

func TestDebugCache(t *testing.T) {
	conf := log551.Config{
		Debug:log551.ConfigDetail{
			EachTime:false,
			Cache:true,
		},
		Information:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Warning:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Error:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Critical:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Path:"/var/log/gorai/test.log",
	}

	l := log551.New(&conf)
	l.Open()
	defer l.Close()

	l.Debug("test string")
	l.Debugf("%s %s", "test", "string")
	l.Debugln("test string")
	l.Information("test string")
	l.Informationf("%s %s", "test", "string")
	l.Informationln("test string")
	l.Warning("test string")
	l.Warningf("%s %s", "test", "string")
	l.Warningln("test string")
	l.Error("test string")
	l.Errorf("%s %s", "test", "string")
	l.Errorln("test string")
	l.Critical("test string")
	l.Criticalf("%s %s", "test", "string")
	l.Criticalln("test string")
}

func TestInformationCache(t *testing.T) {
	conf := log551.Config{
		Debug:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Information:log551.ConfigDetail{
			EachTime:false,
			Cache:true,
		},
		Warning:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Error:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Critical:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Path:"/var/log/gorai/test.log",
	}

	l := log551.New(&conf)
	l.Open()
	defer l.Close()

	l.Debug("test string")
	l.Debugf("%s %s", "test", "string")
	l.Debugln("test string")
	l.Information("test string")
	l.Informationf("%s %s", "test", "string")
	l.Informationln("test string")
	l.Warning("test string")
	l.Warningf("%s %s", "test", "string")
	l.Warningln("test string")
	l.Error("test string")
	l.Errorf("%s %s", "test", "string")
	l.Errorln("test string")
	l.Critical("test string")
	l.Criticalf("%s %s", "test", "string")
	l.Criticalln("test string")
}

func TestWarningCache(t *testing.T) {
	conf := log551.Config{
		Debug:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Information:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Warning:log551.ConfigDetail{
			EachTime:false,
			Cache:true,
		},
		Error:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Critical:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Path:"/var/log/gorai/test.log",
	}

	l := log551.New(&conf)
	l.Open()
	defer l.Close()

	l.Debug("test string")
	l.Debugf("%s %s", "test", "string")
	l.Debugln("test string")
	l.Information("test string")
	l.Informationf("%s %s", "test", "string")
	l.Informationln("test string")
	l.Warning("test string")
	l.Warningf("%s %s", "test", "string")
	l.Warningln("test string")
	l.Error("test string")
	l.Errorf("%s %s", "test", "string")
	l.Errorln("test string")
	l.Critical("test string")
	l.Criticalf("%s %s", "test", "string")
	l.Criticalln("test string")
}

func TestErrorCache(t *testing.T) {
	conf := log551.Config{
		Debug:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Information:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Warning:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Error:log551.ConfigDetail{
			EachTime:false,
			Cache:true,
		},
		Critical:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Path:"/var/log/gorai/test.log",
	}

	l := log551.New(&conf)
	l.Open()
	defer l.Close()

	l.Debug("test string")
	l.Debugf("%s %s", "test", "string")
	l.Debugln("test string")
	l.Information("test string")
	l.Informationf("%s %s", "test", "string")
	l.Informationln("test string")
	l.Warning("test string")
	l.Warningf("%s %s", "test", "string")
	l.Warningln("test string")
	l.Error("test string")
	l.Errorf("%s %s", "test", "string")
	l.Errorln("test string")
	l.Critical("test string")
	l.Criticalf("%s %s", "test", "string")
	l.Criticalln("test string")
}

func TestCriticalCache(t *testing.T) {
	conf := log551.Config{
		Debug:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Information:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Warning:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Error:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Critical:log551.ConfigDetail{
			EachTime:false,
			Cache:true,
		},
		Path:"/var/log/gorai/test.log",
	}

	l := log551.New(&conf)
	l.Open()
	defer l.Close()

	l.Debug("test string")
	l.Debugf("%s %s", "test", "string")
	l.Debugln("test string")
	l.Information("test string")
	l.Informationf("%s %s", "test", "string")
	l.Informationln("test string")
	l.Warning("test string")
	l.Warningf("%s %s", "test", "string")
	l.Warningln("test string")
	l.Error("test string")
	l.Errorf("%s %s", "test", "string")
	l.Errorln("test string")
	l.Critical("test string")
	l.Criticalf("%s %s", "test", "string")
	l.Criticalln("test string")
}

func BenchmarkEachTime(b *testing.B) {
//	b.SkipNow()
	conf := log551.Config{
		Debug:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Information:log551.ConfigDetail{
			EachTime:true,
			Cache:false,
		},
		Warning:log551.ConfigDetail{
			EachTime:true,
			Cache:false,
		},
		Error:log551.ConfigDetail{
			EachTime:true,
			Cache:false,
		},
		Critical:log551.ConfigDetail{
			EachTime:true,
			Cache:false,
		},
		Path:"/var/log/gorai/test.log",
	}

	l := log551.New(&conf)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Open()
		l.Debug("test string")
		l.Debugf("%s %s", "test", "string")
		l.Debugln("test string")
		l.Information("test string")
		l.Informationf("%s %s", "test", "string")
		l.Informationln("test string")
		l.Warning("test string")
		l.Warningf("%s %s", "test", "string")
		l.Warningln("test string")
		l.Error("test string")
		l.Errorf("%s %s", "test", "string")
		l.Errorln("test string")
		l.Critical("test string")
		l.Criticalf("%s %s", "test", "string")
		l.Criticalln("test string")
		l.Close()
	}
}

func BenchmarkCache(b *testing.B) {
//	b.SkipNow()
	conf := log551.Config{
		Debug:log551.ConfigDetail{
			EachTime:false,
			Cache:false,
		},
		Information:log551.ConfigDetail{
			EachTime:false,
			Cache:true,
		},
		Warning:log551.ConfigDetail{
			EachTime:false,
			Cache:true,
		},
		Error:log551.ConfigDetail{
			EachTime:false,
			Cache:true,
		},
		Critical:log551.ConfigDetail{
			EachTime:false,
			Cache:true,
		},
		Path:"/var/log/gorai/test.log",
	}

	l := log551.New(&conf)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Open()

		l.Debug("test string")
		l.Debugf("%s %s", "test", "string")
		l.Debugln("test string")
		l.Information("test string")
		l.Informationf("%s %s", "test", "string")
		l.Informationln("test string")
		l.Warning("test string")
		l.Warningf("%s %s", "test", "string")
		l.Warningln("test string")
		l.Error("test string")
		l.Errorf("%s %s", "test", "string")
		l.Errorln("test string")
		l.Critical("test string")
		l.Criticalf("%s %s", "test", "string")
		l.Criticalln("test string")

		l.Close()
	}
}