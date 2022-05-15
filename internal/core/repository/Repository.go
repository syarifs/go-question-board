package repository

type Repository struct {
	Auth AuthRepository
	User UserRepository
	Tag TagRepository
	Subject SubjectRepository
	Major MajorRepository
	Questionnaire QuestionnaireRepository
	EvaluateTeacher EvaluateRepository
}

