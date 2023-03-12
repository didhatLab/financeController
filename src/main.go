package main

import (
	"context"
	application "main/src/app"
	"net/http"
)

func main() {

	ctx := context.Background()

	err, app := application.NewApplication(ctx)

	if err != nil {
		return
	}

	//entry := entrypoints.FinanceEntrypoint()

	http.ListenAndServe(":4000", app.AppMux)

}
