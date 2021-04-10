package roles

import (
	"fmt"
	"testing"
)

func TestRoles(t *testing.T) {
	role := Member
	roleType := fmt.Sprintf("%T", role)
	expectedRoleType := "roles.Role"

	if roleType != expectedRoleType {
		t.Fatalf("the role type expected is %s but got %T", expectedRoleType, Member)
	}
}
