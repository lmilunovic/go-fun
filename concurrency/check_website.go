package concurrency

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {

	result := map[string]bool{}

	for _, url := range urls {
		go func(u string) {
			result[url] = wc(url)
		}(url)
	}

	return result
}
