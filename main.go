package main

import "cloudreveMirror/routers"

func main() {
	r := routers.InitRouter()

	r.Run(":5001")
}
