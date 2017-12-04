package concurrent

import (
	"sync"
	"net/http"
	"io/ioutil"
	"time"
)

// 测试高并发访问获取账号接口
// url: 测试连接
// total: 总访问数
// concurrency: 并发数
func TestConcurrent(url string, total int, concurrency int) (totalTime int64, result []string, err error) {
	sem := make(chan bool, concurrency)
	result = make([]string, 0)
	var wt sync.WaitGroup
	begin := time.Now().Unix()
	for i := 1; i <= total; i ++ {
		sem <- true
		wt.Add(1)
		go func(index int) {
			resp, err := http.Get(url)
			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)
			result = append(result, string(body))
			if err != nil {
				return
			}
			<-sem
			wt.Done()
		}(i)
	}

	wt.Wait()
	end := time.Now().Unix()
	return end - begin, result, nil
}
