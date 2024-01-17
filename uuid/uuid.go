package uuidUtils

import "github.com/google/uuid"

func New(v ...int) string {
	var u uuid.UUID
	if len(v) == 0 || v[0] == 4 {
		u, _ = uuid.NewRandom()
	} else if v[0] == 1 {
		u, _ = uuid.NewUUID()
	}
	return u.String()
}
