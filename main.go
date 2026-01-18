package main

import (
	"fmt"
	"strconv"
	"time"
)

type PaymentInfo struct {
	ID          int
	Description string
	Usd         int
	Cancelled   bool
}

type PaymentModuleWithSlice struct {
	s []PaymentInfo
}

func (m *PaymentModuleWithSlice) AddInfo(info PaymentInfo) {
	if m.s == nil { // Защита от nil
		m.s = []PaymentInfo{}
	}
	m.s = append(m.s, info)
}

func (m *PaymentModuleWithSlice) GetInfo(id int) PaymentInfo {
	for _, info := range m.s {
		if info.ID == id {
			return info
		}
	}
	return PaymentInfo{}
}

//--------------------------------------------------------------------

type PaymentModuleWithMap struct {
	m map[int]PaymentInfo
}

func (m *PaymentModuleWithMap) AddInfo(info PaymentInfo) {
	m.m[info.ID] = info
}

func (m *PaymentModuleWithMap) GetInfo(id int) PaymentInfo {
	info, ok := m.m[id]
	if !ok {
		return PaymentInfo{}
	}
	return info
}

func main() {
	pSlice := &PaymentModuleWithSlice{}
	pMap := &PaymentModuleWithMap{
		m: make(map[int]PaymentInfo),
	}

	iterations := 100000 // 1 миллион вместо 100

	//-----------------------------------------------------------------

	before := time.Now()
	for i := 0; i < iterations; i++ {
		info := PaymentInfo{
			ID:          i,
			Description: strconv.Itoa(i),
		}
		pSlice.AddInfo(info)
	}
	fmt.Println("slice add:", time.Since(before))

	before = time.Now()
	for i := 0; i < iterations; i++ {
		info := PaymentInfo{
			ID:          i,
			Description: strconv.Itoa(i),
		}
		pMap.AddInfo(info)
	}
	fmt.Println("map add:", time.Since(before))

	//-------------------------------------------------------------------

	before = time.Now()
	for i := range iterations {
		info := pSlice.GetInfo(i)
		_ = info
	}
	fmt.Println("slice get:", time.Since(before))

	before = time.Now()
	for i := range iterations {
		info := pMap.GetInfo(i)
		_ = info
	}
	fmt.Println("map get:", time.Since(before))

}
