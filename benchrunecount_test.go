package benchrunecount

import (
	"testing"
	"unicode/utf8"
)

func BenchmarkRunCountInString16ascii(b *testing.B) {
	str := "abcdefghijklmnop"
	for i := 0; i < b.N; i++ {
		_ = utf8.RuneCountInString(str)
	}
}

func BenchmarkRunCountInString16multi(b *testing.B) {
	str := "日本語を数える速さを確認します。"
	for i := 0; i < b.N; i++ {
		_ = utf8.RuneCountInString(str)
	}
}

func BenchmarkRunCountInString32ascii(b *testing.B) {
	str := "abcdefghijklmnopqrstuvwxyz012345"
	for i := 0; i < b.N; i++ {
		_ = utf8.RuneCountInString(str)
	}
}

func BenchmarkRunCountInString32multi(b *testing.B) {
	str := "日本語を数える速さを確認します。どれぐらい差があるのでしょうか。"
	for i := 0; i < b.N; i++ {
		_ = utf8.RuneCountInString(str)
	}
}

func BenchmarkCastToRuneArray16ascii(b *testing.B) {
	str := "abcdefghijklmnop"
	for i := 0; i < b.N; i++ {
		_ = len([]rune(str))
	}
}

func BenchmarkCastToRuneArray16multi(b *testing.B) {
	str := "日本語を数える速さを確認します。"
	for i := 0; i < b.N; i++ {
		_ = len([]rune(str))
	}
}

func BenchmarkCastToRuneArray32ascii(b *testing.B) {
	str := "abcdefghijklmnopqrstuvwxyz012345"
	for i := 0; i < b.N; i++ {
		_ = len([]rune(str))
	}
}

func BenchmarkCastToRuneArray32multi(b *testing.B) {
	str := "日本語を数える速さを確認します。どれぐらい差があるのでしょうか。"
	for i := 0; i < b.N; i++ {
		_ = len([]rune(str))
	}
}
