package views

import (
	"context"

	"app/structs"
)

func GetSessionContext(ctx context.Context) structs.SessionContext {
  return ctx.Value("session").(structs.SessionContext)
}
