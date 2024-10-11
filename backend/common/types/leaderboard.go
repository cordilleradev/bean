package types

type SupportedPeriods struct {
	FixedPeriods  []string       `json:"fixed_periods"`
	CustomPeriods *CustomPeriods `json:"custom_periods,omitempty"`
}

type CustomPeriods struct {
	MinDays int `json:"min_days"`
	MaxDays int `json:"max_days"`
}

func NewSupportedPeriods(fixedPeriods []string, customPeriods *CustomPeriods) SupportedPeriods {
	return SupportedPeriods{
		FixedPeriods:  fixedPeriods,
		CustomPeriods: customPeriods,
	}
}
