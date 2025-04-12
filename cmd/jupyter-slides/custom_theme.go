package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"sync"
)

type CustomThemesFlag []string

func (f *CustomThemesFlag) String() string {
	return strings.Join(*f, ", ")
}
func (f *CustomThemesFlag) Set(value string) error {
	*f = append(*f, value)
	return nil
}

func (f CustomThemesFlag) ToArray() ([]map[string]any, error) {
	var (
		wg      sync.WaitGroup
		items   []map[string]any
		errChan = make(chan error, len(f))
	)
	for _, path := range f {
		wg.Add(1)
		go func(path string) {
			defer wg.Done()
			ct, err := os.ReadFile(path)
			if err != nil {
				errChan <- err
				return
			}
            var data map[string]any
            if err := json.Unmarshal(ct, &data); err != nil {
                errChan <- fmt.Errorf("JSON parse error: %v", err)
				return
            }
			items = append(items, data)
		}(path)
	}
	wg.Wait()
    close(errChan)

	for err := range errChan {
		if err != nil {
			return nil, err
		}
	}
	return items, nil
}
