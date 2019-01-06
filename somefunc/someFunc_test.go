// package単位でテストできる
// go test (テストしたいパッケージパス)でできる
package somefunc

import (
	"testing" // テストで使える関数・構造体が用意されているパッケージをimport
)

func TestSomeFuncSuccess(t *testing.T) {
	// someFunc()が正常に終わる引数を渡す。
	// resultやerrに異常のときの値が入っているか探して見つけたらテスト失敗
	result, err := someFunc("hoge")
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if result != 0 {
		t.Fatal("failed test")
	}
}

func TestSomeFuncFailed(t *testing.T) {
	// someFunc()がエラーを返す引数を渡す
	// resultやerrに正常なときの値が入っているか探して見つけたらテスト失敗
	result, err := someFunc("fuga")
	if err == nil {
		t.Fatal("failed test")
	}
	if result != 1 {
		t.Fatal("failed test")
	}
}
