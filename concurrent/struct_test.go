package concurrent

import (
	"testing"
	"fmt"
)

func TestHighConcurrent(t *testing.T) {

	for i := 1; i < 100; i ++ {
		url := fmt.Sprintf("http://10.8.210.235:8099/v1/device/getAccount/?device_id=%v&game_id=9019&seq=1&strategy=1&tag=9019&type=1", i)
		totalTime, result, err := TestConcurrent(url, 100, 100)
		fmt.Printf("totalTime:%v, result:%v, err:%v\n ", totalTime, result, err)
	}
}
