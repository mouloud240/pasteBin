package initializers

import (
    "github.com/joho/godotenv"

)
func InitEnv(path string) error {
		err := godotenv.Load(path)
		return err
		

}
