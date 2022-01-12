package bl

import (
	"github.com/khorevaa/r2gitsync/internal/bl/logic"
	"github.com/khorevaa/r2gitsync/internal/di"
)

type BL struct {
	ProjectsLogic logic.IProjectsLogic
	// AuthLogic      *auth.Logic
	// TasksLogic     *tasks.Logic
	// BatchLogic     *batch.Logic
	// CommonUseCases *i_common.CommonUseCases
	// // Validator
	// Utils utils.IUtils
}

func NewBL(di di.IAppDeps) *BL {
	return &BL{
		ProjectsLogic: logic.NewProjectsLogic(di),
		// BatchLogic:     batch.NewBatchLogic(di),
		// TasksLogic:     tasks.NewLogic(di),
		// CommonUseCases: i_common.NewCommonUseCases(di),
		// Utils:          utils.NewUtils(),
	}
}
