package mytime

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

const (
	ts      = "1633603737443"
	format  = "2006-01-02T15:04:05.999Z"
	iso     = "2021-10-07T10:48:57.443Z"
	iso_cst = "2021-10-07T18:48:57.443Z"
)

func TestTsToStr(t *testing.T) {
	s := TsToStr(ts, format)
	assert.True(t, s == iso)
}

func TestTsToStrCST(t *testing.T) {
	s := TsToStrCST(ts, format)
	assert.True(t, s == iso_cst)
}

func TestTsToISO(t *testing.T) {
	s := TsToISO(ts)
	assert.True(t, s == iso)
}

func TestTsToISOCST(t *testing.T) {
	s := TsToISOCST(ts)
	assert.True(t, s == iso_cst)
}

// ts ->已测试 iso ->库函数 time
func TestTsToTime(t *testing.T) {
	tt, err := time.Parse("2006-01-02T15:04:05.999Z", iso)
	if err != nil {
		t.Log("ERROR: time.Parse().")
		return
	}
	iso1 := TsToISO(ts)
	time1 := ISOToTime(iso1) // 包装的库函数，不再测试
	assert.True(t, tt == time1)
}

func TestTsToTimeCST(t *testing.T) {
	tt, err := time.Parse("2006-01-02T15:04:05.999Z", iso_cst)
	if err != nil {
		t.Log("ERROR: time.Parse().")
		return
	}
	iso1 := TsToISOCST(ts)
	time1 := ISOToTime(iso1)
	assert.True(t, tt == time1)
}

func TestISOToTs(t *testing.T) {
	ts1 := ISOToTs(iso)
	assert.True(t, ts1 == ts)
}

func TestISOCSTToTs(t *testing.T) {
	ts1 := ISOCSTToTs(iso_cst)
	assert.True(t, ts1 == ts)
}

// 测试 iso to iso_cst: 2021-10-07T10:48:57.443Z  -> 2021-10-07T18:48:57.443Z
func TestISOToISOCST(t *testing.T) {
	iso_cst1 := ISOToISOCST(iso)
	assert.True(t, iso_cst1 == iso_cst)
}

// 测试 iso_cst to iso: 2021-10-07T18:48:57.443Z  -> 2021-10-07T10:48:57.443Z
func TestISOCSTToISO(t *testing.T) {
	iso1 := ISOCSTToISO(iso_cst)
	assert.True(t, iso1 == iso)
}

func TestTimeToTs(t *testing.T) {
	tt, err := time.Parse("2006-01-02T15:04:05.999Z", iso)
	if err != nil {
		t.Log("ERROR: time.Parse().")
		return
	}
	ts1 := TimeToTs(tt)
	assert.True(t, ts1 == ts)
}

func TestTimeToTsCST(t *testing.T) {
	tt, err := time.Parse("2006-01-02T15:04:05.999Z", iso_cst)
	if err != nil {
		t.Log("ERROR: time.Parse().")
		return
	}
	ts1 := TimeCSTToTs(tt)
	assert.True(t, ts1 == ts)
}

func TestTimeToISO(t *testing.T) {
	tt, err := time.Parse("2006-01-02T15:04:05.999Z", iso)
	if err != nil {
		t.Log("ERROR: time.Parse().")
		return
	}
	iso1 := TimeToISO(tt)
	assert.True(t, iso1 == iso)
}

func TestTimeToISOCST(t *testing.T) {
	tt, err := time.Parse("2006-01-02T15:04:05.999Z", iso_cst)
	if err != nil {
		t.Log("ERROR: time.Parse().")
		return
	}
	iso1 := TimeToISOCST(tt)
	assert.True(t, iso1 == iso_cst)
}

// 测试 ISONow() 获取当前时间 iso
func TestISONow(t *testing.T) {
	iso1 := ISONow()
	t.Log(iso1)
}

// 测试 ISONowCST() 获取当前时间 iso_cst
func TestISONowCST(t *testing.T) {
	iso_cst1 := ISONowCST()
	t.Log(iso_cst1)
}

// 测试 ConvertDate（）： yyyymmdd 转为 yyyy-mm-ddT09:30:00Z
func TestConvertDate(t *testing.T) {
	s := "20240401"
	right_s := "2024-04-01T09:30:00Z"
	assert.True(t, ConvertDate(s) == right_s)
}
