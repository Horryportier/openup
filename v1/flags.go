package v1

import (
    "flag"
)

var (
dev = flag.Bool("dev", false, "Sets data.json path to local project for development")
noDefaultEditor = flag.Bool("no-default-editor", false, "If this flag is set the default editor environment variable is ignored")
)
