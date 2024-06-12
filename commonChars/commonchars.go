package commonChars

import "sync"

// Given a string array words, return an array of all characters that show up in all strings within the words (including duplicates).
//You may return the answer in any order.

func commonChars(words []string) []string {
	var wg sync.WaitGroup
	ch := make(chan map[rune]int, len(words))

	for _, word := range words {
		wg.Add(1)
		go func(w string) {
			defer wg.Done()
			m := make(map[rune]int)
			for _, c := range w {
				m[c]++
			}
			ch <- m
		}(word)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	chars := <-ch
	for m := range ch {
		for k := range chars {
			if v, ok := m[k]; ok {
				if v < chars[k] {
					chars[k] = v
				}
			} else {
				delete(chars, k)
			}
		}
	}

	res := []string{}
	for k, v := range chars {
		for i := 0; i < v; i++ {
			res = append(res, string(k))
		}
	}

	return res
}
