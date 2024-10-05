package common

import (
	"fmt"
)

func FormatTraffic(trafficBytes int64) (size string) {
	if trafficBytes < 1024 {
		return fmt.Sprintf("%.2f بایت", float64(trafficBytes)/float64(1))
	} else if trafficBytes < (1024 * 1024) {
		return fmt.Sprintf("%.2f کیلوبایت", float64(trafficBytes)/float64(1024))
	} else if trafficBytes < (1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f مگابایت", float64(trafficBytes)/float64(1024*1024))
	} else if trafficBytes < (1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f گیگابایت", float64(trafficBytes)/float64(1024*1024*1024))
	} else if trafficBytes < (1024 * 1024 * 1024 * 1024 * 1024) {
		return fmt.Sprintf("%.2f ترابایت", float64(trafficBytes)/float64(1024*1024*1024*1024))
	} else {
		return fmt.Sprintf("%.2f اگزابایت", float64(trafficBytes)/float64(1024*1024*1024*1024*1024))
	}
}
