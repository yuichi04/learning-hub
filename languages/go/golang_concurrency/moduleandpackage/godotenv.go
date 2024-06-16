package moduleandpackage

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GoDotEnv() {
	godotenv.Load()
	fmt.Println(os.Getenv("GO_ENV"))
}
