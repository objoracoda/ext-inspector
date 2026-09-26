package cli

import (
	"fmt"

	"github.com/objoracoda/ext-inspector/internal/app"
)

func Run(args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("пустой ID, нужно указать ID через: ext-inspector <extension-id>")
	}
	id := args[0]

	result, err := app.Inspect(id)

	if err != nil {
		return err
	}

	fmt.Println("ID расширения: ", result)
	return nil
}
