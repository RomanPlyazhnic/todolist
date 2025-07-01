package unit

import (
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"

	"github.com/RomanPlyazhnic/todolist/internal/app/contracts"
)

func TestCreateTodoList_Success(t *testing.T) {
	request := contracts.TodoList{
		UserId: rand.Intn(1000000),
		Text:   faker.Sentence(),
		Checkboxes: []*contracts.Checkbox{
			{
				Checked: true,
				Text:    faker.Sentence(),
			},
			{
				Checked: false,
				Text:    faker.Sentence(),
			},
		},
	}

	_, err := request.Validate()

	assert.Equal(t, nil, err)
}
