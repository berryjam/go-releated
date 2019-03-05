package main

import (
	"time"
	"fmt"
)

var Funnels = make(map[string]*Funnel)

type Funnel struct {
	Capacity    float64
	LeakingRate float64
	LeftQuota   float64
	LeakingTime int64
}

func (f *Funnel) MakeSpace() {
	deltaTime := time.Now().Unix() - f.LeakingTime
	deltaQuota := f.LeakingRate * float64(deltaTime)
	if deltaQuota < 1 {
		return
	}

	if deltaQuota+f.LeftQuota >= f.Capacity {
		f.LeftQuota = f.Capacity
	} else {
		f.LeftQuota += deltaQuota
	}
}

func (f *Funnel) Watering(quota float64) bool {
	f.MakeSpace()
	if f.LeftQuota >= quota {
		f.LeftQuota -= quota
		return true
	} else {
		return false
	}
}

func IsActionAllowed(uid, action string, capacity float64, leakingRate float64) bool {
	key := fmt.Sprintf("%v:%v", uid, action)
	if _, ok := Funnels[key]; !ok {
		Funnels[key] = &Funnel{Capacity: capacity, LeftQuota: capacity, LeakingRate: leakingRate, LeakingTime: time.Now().Unix()}
	}
	return Funnels[key].Watering(1)
}

func main() {
	for i := 0; i < 20; i++ {
		fmt.Printf("%+v\n", IsActionAllowed("laoqian", "reply", 15, 0.5))
	}
}
