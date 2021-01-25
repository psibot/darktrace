package examples

import (
	"fmt"
	"github.com/rfizzle/darktrace"
	"log"
	"time"
)

func GetModelBreaches() {
	client, err := darktrace.NewClient(
		"https://darktrace.example.com",
		"publicToken",
		"privateToken",
		darktrace.ClientDisableTLSValidation,
	)

	if err != nil {
		log.Fatal(err)
	}

	results, err := client.ModelBreaches(
		darktrace.Param("starttime", fmt.Sprintf("%d", time.Now().Add(time.Duration(12)*time.Hour*-1).UTC().Unix())),
		darktrace.Param("endtime", fmt.Sprintf("%d", time.Now().UTC().Unix())),
	)

	if err != nil {
		log.Fatal(err)
	}

	if len(results) > 0 {
		return
	}

	return
}
