package main

import (
	"fmt"
	"os"

	"github.com/aldinokemal/go-whatsapp-web-multidevice/cmd"
)

// @title			Go WhatsApp Web Multi Device API
// @version		2.0
// @description	This is a WhatsApp Web Multi Device API server.
// @termsOfService	http://swagger.io/terms/
// @contact.name	API Support
// @contact.url	https://github.com/aldinokemal/go-whatsapp-web-multidevice
// @license.name	MIT
// @license.url	https://opensource.org/licenses/MIT
// @host			localhost:3000
// @BasePath		/
func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
