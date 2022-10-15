package aggregate

import "github.com/google/uuid"

type Product struct {
	ProductId  uuid.UUID
	SkuNo      int
	UnitWeight int
}
