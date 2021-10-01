package gouuid

import (
	"fmt"
	"testing"

	"github.com/google/uuid"
)

func Test_DefaultIfEmpty(t *testing.T) {
	result := DefaultIfEmpty(nil)
	if result != uuid.Nil {
		t.Errorf("Returned = %s; want %s", result, uuid.Nil)
	}

	uid := uuid.New()
	result = DefaultIfEmpty(&uid)
	if result != uid {
		t.Errorf("Returned = %s; want %s", result, uid)
	}
}

func Test_ParseToDefault(t *testing.T) {
	result := ParseToDefault("error")
	if result != uuid.Nil {
		t.Errorf("Returned = %s; want %s", result, uuid.Nil)
	}

	uid := uuid.New().String()
	result = ParseToDefault(uid)
	if result.String() != uid {
		t.Errorf("Returned = %s; want %s", result, uid)
	}
}

func Test_IndexOf(t *testing.T) {
	ids := []uuid.UUID{ParseToDefault("b7729c88-47e9-42a7-92d3-3e6bcc585f73"), ParseToDefault("83adb35a-847a-4962-8e09-8311a45dc2a2"), ParseToDefault("d9d65dfc-4643-44ab-920f-c564259fd96c")}
	result, _ := IndexOf(ids, ParseToDefault("83adb35a-847a-4962-8e09-8311a45dc2a2"))

	if result != 1 {
		t.Errorf("Returned = %d; want %s", result, "1")
	}

	result, _ = IndexOf(ids, ParseToDefault("83adb35a-847a-4962-8e09-8311a45dc2a1"))

	if result != -1 {
		t.Errorf("Returned = %d; want %s", result, "-1")
	}
}

func Test_ExtractIdsFromStructSlice(t *testing.T) {

	type Student struct {
		ID     uuid.UUID `json:"id"`
		Name   string    `json:"name"`
		RoleNo int       `json:"roleno"`
	}

	students := []Student{
		{
			ID:     uuid.New(),
			Name:   "Lucian",
			RoleNo: 7,
		},
		{
			ID:     uuid.New(),
			Name:   "Aman",
			RoleNo: 8,
		},
	}

	// result := ExtractIdsFromStructSlice(students, "ID")

	result := ExtractIdsFromStructSlice(&students, "ID")
	fmt.Println(result)
}
