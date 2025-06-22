package app

import (
	"fmt"
	"os"
	"testing"

	pkgerrors "github.com/pkg/errors"
	"github.com/stretchr/testify/require"
)

func TestConfig_IsValid(t *testing.T) {
	type arg struct {
		givenCfg Config
		expErr   error
	}
	tcs := map[string]arg{
		"all good": {
			givenCfg: Config{
				Env:       "dev",
				Version:   "1.0.0",
				Server:    "server",
				Project:   "project",
				Component: "api",
				Team:      "team",
				AppName:   "project-app",
			},
		},
		"no team": {
			givenCfg: Config{
				Env:       "dev",
				Version:   "1.0.0",
				Server:    "server",
				Project:   "project",
				Component: "api",
				AppName:   "project-app",
			},
		},
		"no version": {
			givenCfg: Config{
				Env:       "dev",
				Server:    "server",
				Project:   "project",
				Component: "api",
				Team:      "team",
				AppName:   "project-app",
			},
			expErr: fmt.Errorf("version missing: %w", ErrInvalidAppConfig),
		},
		"no server": {
			givenCfg: Config{
				Env:       "dev",
				Version:   "1.0.0",
				Project:   "project",
				Component: "api",
				Team:      "team",
				AppName:   "project-app",
			},
		},
		"no env": {
			givenCfg: Config{
				Version:   "1.0.0",
				Server:    "server",
				Project:   "project",
				Component: "api",
				Team:      "team",
				AppName:   "project-app",
			},
			expErr: fmt.Errorf("invalid env: %s given: %w", "", ErrInvalidAppConfig),
		},
		"invalid env": {
			givenCfg: Config{
				Env:       "env",
				Version:   "1.0.0",
				Server:    "server",
				Project:   "project",
				Component: "api",
				Team:      "team",
				AppName:   "project-app",
			},
			expErr: fmt.Errorf("invalid env: %s given: %w", "env", ErrInvalidAppConfig),
		},
		"no project component": {
			givenCfg: Config{
				Env:     "dev",
				Version: "1.0.0",
				Server:  "server",
				Project: "project",
				Team:    "team",
				AppName: "project-app",
			},
			expErr: fmt.Errorf("invalid component: %s given: %w", "", ErrInvalidAppConfig),
		},
		"invalid project component": {
			givenCfg: Config{
				Env:       "dev",
				Version:   "1.0.0",
				Server:    "server",
				Project:   "project",
				Component: "component",
				Team:      "team",
				AppName:   "project-app",
			},
			expErr: fmt.Errorf("invalid component: %s given: %w", "component", ErrInvalidAppConfig),
		},
		"no app": {
			givenCfg: Config{
				Env:       "dev",
				Version:   "1.0.0",
				Server:    "server",
				Project:   "project",
				Component: "api",
				Team:      "team",
			},
			expErr: fmt.Errorf("app name missing: %w", ErrInvalidAppConfig),
		},
		"no project": {
			givenCfg: Config{
				Env:       "dev",
				Version:   "1.0.0",
				Server:    "server",
				Component: "api",
				Team:      "team",
				AppName:   "project-app",
			},
			expErr: fmt.Errorf("project missing: %w", ErrInvalidAppConfig),
		},
		"app name doesn't start with project name": {
			givenCfg: Config{
				Env:       "dev",
				Version:   "1.0.0",
				Server:    "server",
				Project:   "project",
				Component: "api",
				Team:      "team",
				AppName:   "app",
			},
			expErr: fmt.Errorf("app name should start with project name as prefix: %w", ErrInvalidAppConfig),
		},
		"app name and project name doesn't match": {
			givenCfg: Config{
				Env:       "dev",
				Version:   "1.0.0",
				Server:    "server",
				Project:   "project",
				Component: "api",
				Team:      "team",
				AppName:   "pjt-app",
			},
			expErr: fmt.Errorf("app name should start with project name as prefix: %w", ErrInvalidAppConfig),
		},
		"app name format is incorrect": {
			givenCfg: Config{
				Env:       "dev",
				Version:   "1.0.0",
				Server:    "server",
				Project:   "project",
				Component: "api",
				Team:      "team",
				AppName:   "project.app",
			},
			expErr: fmt.Errorf("app name should start with project name as prefix: %w", ErrInvalidAppConfig),
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			require.Equal(t, tc.expErr, pkgerrors.Cause(tc.givenCfg.IsValid()))
		})
	}
}

func TestNewConfigFromEnv(t *testing.T) {
	type arg struct {
		prepareEnvVars func(*testing.T)
		expErr         bool
		expCfg         Config
	}
	tcs := map[string]arg{
		"success: all good": {
			prepareEnvVars: func(t *testing.T) {
				require.NoError(t, os.Setenv("ENVIRONMENT", "qa"))
				require.NoError(t, os.Setenv("VERSION", "1.0.0"))
				require.NoError(t, os.Setenv("SERVER_NAME", "server"))
				require.NoError(t, os.Setenv("PROJECT_NAME", "project"))
				require.NoError(t, os.Setenv("PROJECT_COMPONENT", "api"))
				require.NoError(t, os.Setenv("TEAM_NAME", "team"))
				require.NoError(t, os.Setenv("APP_NAME", "project-app"))
			},
			expCfg: Config{
				Version: "1.0.0", Env: EnvQA, Server: "server", Project: "project", Component: ComponentTypeAPI, Team: "team", AppName: "project-app",
			},
		},
		"success: without server": {
			prepareEnvVars: func(t *testing.T) {
				require.NoError(t, os.Setenv("ENVIRONMENT", "qa"))
				require.NoError(t, os.Setenv("VERSION", "1.0.0"))
				require.NoError(t, os.Setenv("PROJECT_NAME", "project"))
				require.NoError(t, os.Setenv("PROJECT_COMPONENT", "api"))
				require.NoError(t, os.Setenv("TEAM_NAME", "team"))
				require.NoError(t, os.Setenv("APP_NAME", "project-app"))
			},
			expCfg: Config{
				Version: "1.0.0", Env: EnvQA, Project: "project", Component: ComponentTypeAPI, Team: "team", AppName: "project-app",
			},
		},
		"success: without team": {
			prepareEnvVars: func(t *testing.T) {
				require.NoError(t, os.Setenv("ENVIRONMENT", "qa"))
				require.NoError(t, os.Setenv("VERSION", "1.0.0"))
				require.NoError(t, os.Setenv("SERVER_NAME", "server"))
				require.NoError(t, os.Setenv("PROJECT_NAME", "project"))
				require.NoError(t, os.Setenv("PROJECT_COMPONENT", "api"))
				require.NoError(t, os.Setenv("APP_NAME", "project-app"))
			},
			expCfg: Config{
				Version: "1.0.0", Env: EnvQA, Server: "server", Project: "project", Component: ComponentTypeAPI, AppName: "project-app",
			},
		},
		"err: without env": {
			prepareEnvVars: func(t *testing.T) {
				require.NoError(t, os.Setenv("VERSION", "1.0.0"))
				require.NoError(t, os.Setenv("SERVER_NAME", "server"))
				require.NoError(t, os.Setenv("PROJECT_NAME", "project"))
				require.NoError(t, os.Setenv("PROJECT_COMPONENT", "api"))
				require.NoError(t, os.Setenv("TEAM_NAME", "team"))
				require.NoError(t, os.Setenv("APP_NAME", "project-app"))
			},
			expErr: true,
		},
		"err: with invalid env": {
			prepareEnvVars: func(t *testing.T) {
				require.NoError(t, os.Setenv("ENVIRONMENT", "env"))
				require.NoError(t, os.Setenv("VERSION", "1.0.0"))
				require.NoError(t, os.Setenv("SERVER_NAME", "server"))
				require.NoError(t, os.Setenv("PROJECT_NAME", "project"))
				require.NoError(t, os.Setenv("PROJECT_COMPONENT", "api"))
				require.NoError(t, os.Setenv("TEAM_NAME", "team"))
				require.NoError(t, os.Setenv("APP_NAME", "project-app"))
			},
			expErr: true,
		},
		"err: without component": {
			prepareEnvVars: func(t *testing.T) {
				require.NoError(t, os.Setenv("ENVIRONMENT", "qa"))
				require.NoError(t, os.Setenv("VERSION", "1.0.0"))
				require.NoError(t, os.Setenv("SERVER_NAME", "server"))
				require.NoError(t, os.Setenv("PROJECT_NAME", "project"))
				require.NoError(t, os.Setenv("TEAM_NAME", "team"))
				require.NoError(t, os.Setenv("APP_NAME", "project-app"))
			},
			expErr: true,
		},
		"err: with invalid component": {
			prepareEnvVars: func(t *testing.T) {
				require.NoError(t, os.Setenv("ENVIRONMENT", "qa"))
				require.NoError(t, os.Setenv("VERSION", "1.0.0"))
				require.NoError(t, os.Setenv("SERVER_NAME", "server"))
				require.NoError(t, os.Setenv("PROJECT_NAME", "project"))
				require.NoError(t, os.Setenv("PROJECT_COMPONENT", "comp"))
				require.NoError(t, os.Setenv("TEAM_NAME", "team"))
				require.NoError(t, os.Setenv("APP_NAME", "project-app"))
			},
			expErr: true,
		},
		"err: without project": {
			prepareEnvVars: func(t *testing.T) {
				require.NoError(t, os.Setenv("ENVIRONMENT", "qa"))
				require.NoError(t, os.Setenv("VERSION", "1.0.0"))
				require.NoError(t, os.Setenv("SERVER_NAME", "server"))
				require.NoError(t, os.Setenv("PROJECT_COMPONENT", "api"))
				require.NoError(t, os.Setenv("TEAM_NAME", "team"))
				require.NoError(t, os.Setenv("APP_NAME", "project-app"))
			},
			expErr: true,
		},
		"err: without app": {
			prepareEnvVars: func(t *testing.T) {
				require.NoError(t, os.Setenv("ENVIRONMENT", "qa"))
				require.NoError(t, os.Setenv("VERSION", "1.0.0"))
				require.NoError(t, os.Setenv("SERVER_NAME", "server"))
				require.NoError(t, os.Setenv("PROJECT_NAME", "project"))
				require.NoError(t, os.Setenv("PROJECT_COMPONENT", "api"))
				require.NoError(t, os.Setenv("TEAM_NAME", "team"))
			},
			expErr: true,
		},
		"err: without version": {
			prepareEnvVars: func(t *testing.T) {
				require.NoError(t, os.Setenv("ENVIRONMENT", "qa"))
				require.NoError(t, os.Setenv("SERVER_NAME", "server"))
				require.NoError(t, os.Setenv("PROJECT_NAME", "project"))
				require.NoError(t, os.Setenv("PROJECT_COMPONENT", "api"))
				require.NoError(t, os.Setenv("TEAM_NAME", "team"))
				require.NoError(t, os.Setenv("APP_NAME", "project-app"))
			},
			expErr: true,
		},
	}
	for desc, tc := range tcs {
		t.Run(desc, func(t *testing.T) {
			// Given:
			defer func() {
				require.NoError(t, os.Unsetenv("ENVIRONMENT"))
				require.NoError(t, os.Unsetenv("VERSION"))
				require.NoError(t, os.Unsetenv("SERVER_NAME"))
				require.NoError(t, os.Unsetenv("PROJECT_NAME"))
				require.NoError(t, os.Unsetenv("PROJECT_COMPONENT"))
				require.NoError(t, os.Unsetenv("TEAM_NAME"))
				require.NoError(t, os.Unsetenv("APP_NAME"))
			}()

			tc.prepareEnvVars(t)

			// When:
			cfg := NewConfigFromEnv()
			err := cfg.IsValid()

			// Then:
			if tc.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
				require.Equal(t, tc.expCfg, cfg)
			}
		})
	}
}
