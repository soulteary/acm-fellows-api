package network_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/soulteary/acm-fellows-api/model/network"
)

func TestGetRemotePage(t *testing.T) {
	// 创建一个模拟的 HTTP 服务器
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello, world!"))
	}))
	defer ts.Close()

	// 调用 GetRemotePage 函数获取模拟服务器的内容
	got, err := network.GetRemotePage(ts.URL)
	if err != nil {
		t.Fatalf("GetRemotePage(%q) error: %v", ts.URL, err)
	}

	// 检查返回值是否符合预期
	want := []byte("Hello, world!")
	if string(got) != string(want) {
		t.Errorf("GetRemotePage(%q) = %q, want %q", ts.URL, got, want)
	}
}
