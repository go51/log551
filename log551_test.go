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

func TestDebug(t *testing.T) {
	conf := log551.Config{
		Debug:true,
		Information:false,
		Warning:false,
		Error:false,
		Critical:false,
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

func TestInformation(t *testing.T) {
	conf := log551.Config{
		Debug:false,
		Information:true,
		Warning:false,
		Error:false,
		Critical:false,
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

func TestWarning(t *testing.T) {
	conf := log551.Config{
		Debug:false,
		Information:false,
		Warning:true,
		Error:false,
		Critical:false,
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

func TestError(t *testing.T) {
	conf := log551.Config{
		Debug:false,
		Information:false,
		Warning:false,
		Error:true,
		Critical:false,
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

func TestCritical(t *testing.T) {
	conf := log551.Config{
		Debug:false,
		Information:false,
		Warning:false,
		Error:false,
		Critical:true,
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

func BenchmarkDebug(b *testing.B) {
//	b.SkipNow()
	conf := log551.Config{
		Debug:true,
		Information:true,
		Warning:true,
		Error:true,
		Critical:true,
		Path:"/var/log/gorai/test.log",
	}

	l := log551.New(&conf)
	l.Open()
	defer l.Close()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
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

}