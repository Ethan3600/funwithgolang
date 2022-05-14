package cpu

import "echoapp/service"

func GetCpuIntensiveWork(times int) []int {
	return service.CPU(times)
}
