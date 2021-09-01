package command

import "github.com/tinytengu/go-argparse"

type CommandHandler func(cmd Command, result argparse.ParseResult)
