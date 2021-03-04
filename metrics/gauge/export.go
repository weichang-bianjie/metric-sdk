package gauge

import "github.com/weichang-bianjie/metric-sdk/types"

type Client interface {
	RegisterMetricInfo(name string, help string, constLabels map[string]string) types.Gauge
	NewMetricFamilyScrap(name, help string, groupingLabels map[string]string, value float64) *types.GobbableMetricFamily
}
