package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/richdon/bookings/pkg/config"
	"github.com/richdon/bookings/pkg/handlers"
	"github.com/richdon/bookings/pkg/render"
)

const port = ":8080"

// sessions shouuld be available package wide
var session *scs.SessionManager

// this will be used to pass around data for packages to use via repo pattern
var app config.AppConfig

func main() {
	// create a session for app config, which will be used by handlers
	session = scs.New()
	session.Lifetime = 24 * time.Hour              // session live for 24 hrs
	session.Cookie.Persist = true                  // session cookie will exist after user exits browser
	session.Cookie.SameSite = http.SameSiteLaxMode // csrf allows cookie to follow across links
	session.Cookie.Secure = app.InProd

	// add our session to the app config
	app.Session = session

	// create a template cache to put into app config, so it can be load templates from memory not disk
	tc, err := render.CreateTemplateCache()

	if err != nil {
		log.Fatalf("cannot load template cache: err: %s", err)
	}
	// add template cache to app config so app can set as a repo in packages
	app.TemplateCache = tc
	app.UseCache = true

	// set app repo in render package, which now has template cache
	render.NewRepo(&app)

	// set app repo in handlers package so it can be used there
	repo := handlers.NewRepo(&app)
	handlers.NewHandlers(repo)

	// instantiate handlers

	srv := &http.Server{
		Addr:    port,
		Handler: routes(),
	}

	fmt.Fprintf(os.Stdout, "Listening on port %s\n", port)

	if err = srv.ListenAndServe(); err != nil {
		log.Fatal(err)
	}

}
