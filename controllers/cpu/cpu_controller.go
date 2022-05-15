package cpu

import "github.com/Ethan3600/funwithgolang/service"

func GetCpuIntensiveWork(times int) []int {
	return service.CPU(times)
}
