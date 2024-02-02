package errs

import (
	"log"
	"os"
)

func PanicIfErr(err error) {
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
}

func PanicIfErrMsg(err error, msg string) {
	if err != nil {
		log.Println(err)
		log.Println(msg)
		os.Exit(1)
	}
}
