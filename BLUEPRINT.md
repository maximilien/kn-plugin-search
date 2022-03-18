## Sample Plugin Project

This project is a sample that can be used as a blueprint to kickoff your own,
golang based [kn](https://github.com/knative/client) plugin.

### Init

Just copy over the content of this repository (except maybe this file), and
adapt the following files:

- [hack/global_vars.sh](hack/global_vars.sh) to adjust the name of your plugin.
- [internal/root/root.go](internal/root/root.go) fix the description and name of
  your plugin.
- [pkg/plugin/plugin.go](pkg/plugin/plugin.go) to update the description and the
  name of your plugin.

### Add your code

To add your code, look into [internal/command](internal/command) and check the
commands that are available there. It is recommended to keep `version.go` but
you should probably remove `print.go` and add your own commands.

The commands that you create should then finally be added in
[internal/root/root.go](internal/root/root.go) to your main Cobra command.

_TODO: More integration how to reference to support libraries and connect to the
cluster_
