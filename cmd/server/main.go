package server

import (
	"context"
	"github.com/11DingKing/digital-humanities-go/internal/config"
	"github.com/11DingKing/digital-humanities-go/internal/httpapi"
	"github.com/11DingKing/digital-humanities-go/internal/service"
	"github.com/11DingKing/digital-humanities-go/internal/storage"
	"github.com/11DingKing/digital-humanities-go/internal/worker"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Main() {
	cfg := config.Load()
	ctx := context.Background()
	db, e := storage.Open(ctx, cfg.DatabasePath)
	if e != nil {
		log.Fatal(e)
	}
	defer db.Close()
	if e = storage.Migrate(ctx, db); e != nil {
		log.Fatal(e)
	}
	svc := service.New(db)
	srv := &http.Server{Addr: cfg.Addr, Handler: httpapi.New(svc).Mux, ReadHeaderTimeout: 5 * time.Second}
	w := worker.New(db, time.Duration(cfg.WorkerIntervalSeconds)*time.Second)
	wctx, cancel := context.WithCancel(ctx)
	defer cancel()
	go w.Run(wctx)
	go func() {
		if e := srv.ListenAndServe(); e != nil && e != http.ErrServerClosed {
			log.Print(e)
		}
	}()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	w.Stop()
	sh, cc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cc()
	_ = srv.Shutdown(sh)
}
