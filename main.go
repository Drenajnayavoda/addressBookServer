package main

import (
	controller "addressBookServer/controllers/stdhttp"
	psg "addressBookServer/gate/psg"
	errorLogger "addressBookServer/pkg/errorLogger"

	"fmt"
)

func main() {

	CONN := "postgres://localhost:5432/address_book"
	psg, err := psg.NewPsg(CONN, "jvs", "123")
	if err != nil {
		fmt.Println(err)
		return
	}

	logger, err := errorLogger.NewErrorLogger(".log")
	if err != nil {
		fmt.Println(err)
		return
	}

	ADDR := "127.0.0.1:8000"
	fmt.Println("server started on ", ADDR)
	ctr := controller.NewController(ADDR, psg, logger)

	ctr.Srv.ListenAndServe()
}
