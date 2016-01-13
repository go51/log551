package log551_test
import (
	"github.com/go51/log551"
	"testing"
)

func TestNew(t *testing.T) {
	log1 := log551.New()
	log2 := log551.New()

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
		_ = log551.New()
	}
}