// Filename: cmd/web/helpers.go
//help us to secure our app to make sure user is authenticated.

//Written by: Abel Blanco, Jovan Alpuche, Lianni Mathews, Cameron Tillet, Talib Marin, Amir Gonzalez
//Tested by: Abel Blanco, Jovan Alpuche, Lianni Mathews, Cameron Tillet, Talib Marin, Amir Gonzalez
//Debbuged by: Abel Blanco, Jovan Alpuche, Lianni Mathews, Cameron Tillet, Talib Marin, Amir Gonzalez

package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
	// deal with the error status
	http.Error(w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError)
}

// func (app *application) clientError(w http.ResponseWriter, status int) {
// 	http.Error(w, http.StatusText(status), status)
// }

// func (app *application) notFound(w http.ResponseWriter) {
// 	app.clientError(w, http.StatusNotFound)
// }

func (app *application) isAuthenticated(r *http.Request) bool {
	return app.sessionManager.Exists(r.Context(), "authenticatedUserID")
} //if its here means we have an authenticated user
