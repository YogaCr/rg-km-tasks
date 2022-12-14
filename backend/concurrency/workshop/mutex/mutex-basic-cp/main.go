package main

import "sync"

//gunakan channel untuk memberpaiki masalah race condition!
func counter(output chan<- int) {
	// TODO: answer here
	mtx := &sync.Mutex{}
	var wg sync.WaitGroup
	count := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// TODO: answer here
			mtx.Lock()
			//kirim 1 ke channel
			count++
			// TODO: answer here
			mtx.Unlock()
		}()
	}
	wg.Wait()
	//mengubah nilai count menggunakan data dari channel
	output <- count
}
