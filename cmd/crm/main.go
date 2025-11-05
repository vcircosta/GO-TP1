package main

import (
	"github.com/vcircosta/GO-TP1/internal/app"
	"github.com/vcircosta/GO-TP1/internal/storage"
)

func main() {
	app.Run(storage.NewMemoryStore())
}
