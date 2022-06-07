<<<<<<< HEAD
# templates
Cloud native buildpack templates
=======
# CNB Buildpacks Templates

Each subfolder defines the scaffolding for a buildpack in a particular language:

1. Go - buildpack scaffolding using [libcnb](https://github.com/buildpacks/libcnb) in [Go](https://go.dev/)
2. Bash - buildpack scaffolding targetting [GNU bash](https://www.gnu.org/software/bash/)

## Creating a Buildpack

Follow the [pack install](https://buildpacks.io/docs/tools/pack/) instructions to install the `pack` CLI for your platform.  On successful completion `pack --version` should show a version number of `0.27` or greater.

Creating a buildpack identified as `examples/my_buildpack` results in a scaffolded buildpack in the `my_buildpack` folder:

```bash
$ pack buildpack new examples/my_buildpack
Use the arrow keys to navigate: ↓ ↑ → ←
? Choose an implementation language for your buildpack:
  ▸ Go
    bash
? Choose the buildpack API version (use the default if you are unsure):
  ▸ 0.7
    0.8
✔ Enter the initial buildpack version: 0.0.1
✔ Enter a stack this buildpack will support by default: io.buildpacks.samples.stacks.bionic
✔ Enter a valid Go module name for this buildpack: github.com/user/buildpack
Successfully created examples/my_buildpack
```
>>>>>>> 62e691a (Provide scaffolding for Go and bash buildpacks)
