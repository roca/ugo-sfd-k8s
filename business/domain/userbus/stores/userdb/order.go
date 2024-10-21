package userdb

import (
	"fmt"

	"github.com/roca/ugo-sfd-k8s/business/api/order"
	"github.com/roca/ugo-sfd-k8s/business/domain/userbus"
)

var orderByFields = map[string]string{
	userbus.OrderByID:      "user_id",
	userbus.OrderByName:    "name",
	userbus.OrderByEmail:   "email",
	userbus.OrderByRoles:   "roles",
	userbus.OrderByEnabled: "enabled",
}

func orderByClause(orderBy order.By) (string, error) {
	by, exists := orderByFields[orderBy.Field]
	if !exists {
		return "", fmt.Errorf("field %q does not exist", orderBy.Field)
	}

	return " ORDER BY " + by + " " + orderBy.Direction, nil
}
