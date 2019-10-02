package main

type HealthCheck []func() error

func NewHealthCheck() HealthCheck {
	return make(HealthCheck, 0)
}

// 添加心跳检查条目
func (health HealthCheck) AddHealthCheck(job func() error) HealthCheck {
	return append(health, job)
}

// 检查心跳
func (health HealthCheck) Check() []error {
	errors := make([]error, 0)

	for _, h := range health {
		err := h()

		if err != nil {
			errors = append(errors, err)
		}
	}

	return errors
}
