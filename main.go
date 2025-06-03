package main

import (
	"context"
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os/exec"
	"runtime"
	"sync"
	"time"
)

const port = "8080"

var ctx, cancel = context.WithCancel(context.Background())

//go:embed dist
var content embed.FS

type CountMsg struct {
	Count int
}

// var in_data struct {
// 	Action string `json:"action"`
// 	Port   string `json:"port"`
// 	Baud   int    `json:"baud"`
// }

func api_count(w http.ResponseWriter, r *http.Request) {
	log.Println(r.RemoteAddr, r.RequestURI)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	m := CountMsg{Count: count}
	json.NewEncoder(w).Encode(m)
	count++
}

func api_shutdown(w http.ResponseWriter, r *http.Request) {
	log.Println(r.Host, r.RequestURI)
}

var count = 0

func main() {
	sm := http.NewServeMux()
	static, _ := fs.Sub(content, "dist") //чтобы получать доступ к файлам без указания корневой директории
	sm.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RemoteAddr, r.RequestURI)
		w.Header().Set("Cache-Control", "no-store")
		http.FileServer(http.FS(static)).ServeHTTP(w, r)
	})
	sm.HandleFunc("GET /api/count", api_count)
	sm.HandleFunc("GET /api/shutdown", api_shutdown)

	httpServ := &http.Server{
		Addr:           ":" + port,
		Handler:        sm,
		ReadTimeout:    20 * time.Second,
		WriteTimeout:   20 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		if err := httpServ.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
		wg.Done()
	}()

	go func() {
		<-ctx.Done()
		log.Println("Shutdown")
		httpServ.Shutdown(context.Background())
	}()

	startBrowser("http://localhost:" + port)

	wg.Wait()
}

func startBrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
