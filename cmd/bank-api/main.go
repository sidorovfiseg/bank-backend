package main

import (
	"bank-backend/internal/di"
	"context"
	"log/slog"
	"net/http"
)




func main() {
	ctx := context.Background()
	container := di.New(ctx)
	defer container.CloseConnection()
	slog.Info("Starting server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		slog.Error("Unable to start server", err)
	} 
}