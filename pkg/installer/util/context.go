package util

import (
	"github.com/sirupsen/logrus"

	"github.com/kubermatic/kubeone/pkg/config"
	"github.com/kubermatic/kubeone/pkg/ssh"
)

// Context hold together currently test flags and parsed info, along with
// utilities like logger
type Context struct {
	Cluster        *config.Cluster
	Logger         logrus.FieldLogger
	Connector      *ssh.Connector
	Configuration  *Configuration
	WorkDir        string
	JoinCommand    string
	Verbose        bool
	BackupFile     string
	DestroyWorkers bool
}

// Clone returns a shallow copy of the context.
func (c *Context) Clone() *Context {
	newCtx := *c
	return &newCtx
}