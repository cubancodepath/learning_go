package main

import (
	"context"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

type Level string

const (
	Debug Level = "debug"
	Info  Level = "info"
)

func Log(ctx context.Context, level Level, message string) {
	var inLevel Level
	inLevel = getLog(ctx)
	if level == Debug && inLevel == Debug {
		fmt.Println(message)
	}
	if level == Info && (inLevel == Info || inLevel == Debug) {
		fmt.Println(message)
	}
}

func storeLog(ctx context.Context, level Level) context.Context {
	return context.WithValue(ctx, "logLevel", level)
}

func getLog(ctx context.Context) Level {
	return ctx.Value("logLevel").(Level)
}

func getLoggin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		logLevel := string(Info)
		logLevel = req.URL.Query().Get("log_level")

		if logLevel != string(Debug) && logLevel != string(Info) {
			w.Write([]byte("log_level_invalid"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		level := Level(logLevel)
		ctx := storeLog(req.Context(), level)
		next.ServeHTTP(w, req.WithContext(ctx))
	})
}

func main() {
	ctx, cancelFunc := context.WithTimeout(context.Background(), time.Second*2)
	defer cancelFunc()
	wg := &sync.WaitGroup{}
	wg.Add(1)

	sum := 0
	i := 0
	go func() {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				cancelFunc()
				return
			default:
				num := rand.Intn(99_9999)
				sum += num
				i++
				if num == 1234 {
					cancelFunc()
					return
				}
			}
		}
	}()

	<-ctx.Done()
	wg.Wait()
	fmt.Printf("the sum is %d after %d iterations\n", sum, i)
}
