package env

import (
	"fmt"
	"os"

	"github.com/go-task/task/v3/internal/experiments"
	"github.com/go-task/task/v3/taskfile/ast"
)

const taskVarPrefix = "TASK_"

func Get(t *ast.Task) []string {
	if t.Env == nil {
		return nil
	}

	return GetFromVars(t.Env)
}

func GetFromVars(env *ast.Vars) []string {
	environ := os.Environ()

	for k, v := range env.ToCacheMap() {
		if !isTypeAllowed(v) {
			continue
		}
		if !experiments.EnvPrecedence.Enabled() {
			if _, alreadySet := os.LookupEnv(k); alreadySet {
				continue
			}
		}
		environ = append(environ, fmt.Sprintf("%s=%v", k, v))
	}

	return environ
}

func isTypeAllowed(v any) bool {
	switch v.(type) {
	case string, bool, int, float32, float64:
		return true
	default:
		return false
	}
}

func GetTaskEnv(key string) string {
	return os.Getenv(taskVarPrefix + key)
}
