package main

import (
	"googlemaps.github.io/maps"
	"log"
)






func main()  {
	var clinetGCM *maps.Client
	if clinetGCM == nil {
		// Pre-declare an err variable to avoid shadowing client.
		var err error

		clinetGCM, err = maps.NewClient(maps.WithAPIKey("api-key"))
		if err != nil {
			log.Fatalf("maps.NewClient failed: %v", err)
		}
	}
}

