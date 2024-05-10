package utility

type SkillzManagerError struct {
	message string
}

func (e SkillzManagerError) Error() string {
	return e.message
}

func NewSkillzManagerError(message string) SkillzManagerError {
	return SkillzManagerError{message: message}
}
