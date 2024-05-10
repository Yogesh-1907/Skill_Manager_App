package service

import (
	"testing"

	"SkillManager_API/model"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// MockSkillRepository is a mock implementation of SkillRepository for testing purposes
type MockSkillRepository struct {
	mock.Mock
}

func (m *MockSkillRepository) FindBySkill(skillName string) (model.Skill, error) {
	args := m.Called(skillName)
	return args.Get(0).(model.Skill), args.Error(1)
}

func (m *MockSkillRepository) AddEmployeeToSkill(skill model.Skill, employee model.Employee) (int, error) {
	args := m.Called(skill, employee)
	return args.Int(0), args.Error(1)
}
func (m *MockSkillRepository) CreateSkill(skill *model.Skill) error {
	args := m.Called(skill)
	return args.Error(0)
}

func (m *MockSkillRepository) DeleteEmployeeAssociation(skill model.Skill) error {
	args := m.Called(skill)
	return args.Error(0)
}

func (m *MockSkillRepository) DeleteSkill(skill model.Skill) error {
	args := m.Called(skill)
	return args.Error(0)
}

func TestAddEmployeeToSkill(t *testing.T) {
	// Create a new instance of the MockSkillRepository
	repo := new(MockSkillRepository)

	// Create a new instance of SkillService and inject the mock repository
	skillService := SkillService{skillRepo: repo}

	// Mock the behavior of FindBySkill method in the repository
	repo.On("FindBySkill", "Java").Return(model.Skill{}, nil)

	// Mock the behavior of AddEmployeeToSkill method in the repository
	repo.On("AddEmployeeToSkill", mock.Anything, mock.Anything).Return(1, nil)

	// Create an EmployeeDTO object for testing
	employeeDTO := model.EmployeeDTO{
		EmpId:      12,
		Name:       "Yogesh choudhar",
		Age:        30,
		SkillLevel: "Intermediate",
		SkillId:    1,
	}

	// Call the AddEmployeeToSkill method
	empId, err := skillService.AddEmployeeToSkill("Java", employeeDTO)

	// Check if the returned empId is as expected
	assert.Equal(t, 1, empId)

	// Check if the error is nil
	assert.Nil(t, err)

	// Assert that the FindBySkill method was called with the correct parameter
	repo.AssertCalled(t, "FindBySkill", "Java")

	// Assert that the AddEmployeeToSkill method was called with the correct parameters
	repo.AssertCalled(t, "AddEmployeeToSkill", mock.Anything, mock.Anything)
}
