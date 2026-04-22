package job

import (
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

type Handler struct {
	DB        *gorm.DB
	CloneRepo func(token, branch, repoName, destination string) error
}

func (h *Handler) CreateJobHandler(w http.ResponseWriter, r *http.Request) {

}

func RunJob(commands []string) error {
	return fmt.Errorf("No implementation error")
}
