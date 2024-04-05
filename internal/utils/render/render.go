package render

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Struct2Responce() {}

func Request2Struct(r *http.Request, v any) error {
	const op = "utils.render.json"
	data, err := io.ReadAll(r.Body)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}

	err = json.Unmarshal(data, &v)
	if err != nil {
		return fmt.Errorf("%s: %w", op, err)
	}
	return nil
}
