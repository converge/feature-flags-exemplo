package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {

	fmt.Println("feature flag value USE_NEW_FEATURE")
	useNewFeature, err := strconv.ParseBool(os.Getenv("USE_NEW_FEATURE"))
	if err != nil {
		panic(err)
	}

	if useNewFeature {
		fmt.Println("using new feature flag")
	} else {
		fmt.Println("not using feature flag")
	}

}
