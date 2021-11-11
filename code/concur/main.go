package concurrency

type WebsiteChecker func(string) bool
type result struct {
    string
    bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
    results := make(map[string]bool)
    resultChannel := make(chan result)

    for _, url := range urls {
        go func(u string) {
            /*  
                results[u] = wc(u) -> X : Race Condition

                resolve race condition!
                URL을 반복할 때 map에 직접 기록하지 않고 각 호출에 대한 결과 구조를 결과 channel로 보냄.
            */
            resultChannel <- result{u, wc(u)}
        }(url)
    }

    for i := 0; i < len(urls); i++ {
        /*
            channel 데이터를 변수에 담기
        */
        r := <-resultChannel
        results[r.string] = r.bool
    }

    return results
}