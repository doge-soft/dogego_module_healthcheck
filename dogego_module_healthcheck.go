package dogego_module_healthcheck

import "fmt"

type HealthCheck []func() error

func NewHealthCheck() HealthCheck {
	return make(HealthCheck, 0)
}

// 添加心跳检查条目
func (health HealthCheck) AddHealthCheck(job func() error) HealthCheck {
	return append(health, job)
}

// 检查心跳
func (health HealthCheck) Check() {
	panic_string := ""

	for _, h := range health {
		err := h()

		if err != nil {
			panic_string += fmt.Sprintf("%s\n", err.Error())
		}
	}

	panic(panic_string)
}
