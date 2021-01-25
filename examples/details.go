package examples

import (
	"github.com/rfizzle/darktrace"
	"log"
)

func GetEventDetails() {
	client, err := darktrace.NewClient(
		"https://darktrace.example.com",
		"publicToken",
		"privateToken",
		darktrace.ClientDisableTLSValidation,
	)

	if err != nil {
		log.Fatal(err)
	}

	results, err := client.EventList(
		darktrace.Param("pbid", "6000000053951"),
		darktrace.Param("includetotalbytes", "true"),
	)

	if err != nil {
		log.Fatal(err)
	}

	if len(results) > 0 {
		return
	}

	return
}
