package cmd

import (
	"github.com/genzai-io/sliced/app/api"
	"github.com/genzai-io/sliced/common/resp"
)

func init() { api.Register(&DBCreate{}) }

// Demotes a Voting member to a Non-Voting member.
type DBCreate struct {
	Name_ string
}

func (c *DBCreate) Name() string   { return "+DB" }
func (c *DBCreate) Help() string   { return "" }
func (c *DBCreate) IsError() bool  { return false }
func (c *DBCreate) IsWorker() bool { return true }

func (c *DBCreate) Marshal(buf []byte) []byte {
	buf = resp.AppendArray(buf, 2)
	buf = resp.AppendBulkString(buf, c.Name())
	buf = resp.AppendBulkString(buf, c.Name_)
	return buf
}

func (c *DBCreate) Parse(args [][]byte) Command {
	cmd := &DBCreate{}

	switch len(args) {
	default:
		return Err("ERR invalid params")

	case 2:
		// Set schema and slice to -1 indicating we want the global store raft
		cmd.Name_ = string(args[1])
		return cmd
	}
	return cmd
}

func (c *DBCreate) Handle(ctx *Context) Reply {
	reply := api.Array([]Reply{
		api.Int(10),
		String("hi"),
	})

	return reply
}
