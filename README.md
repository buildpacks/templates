# CNB Buildpacks Templates

Each subfolder defines the scaffolding for a buildpack in a particular language:

1. Bash - buildpack scaffolding targetting [GNU bash](https://www.gnu.org/software/bash/)

## Creating a Buildpack

Follow the [pack install](https://buildpacks.io/docs/tools/pack/) instructions to install the `pack` CLI for your platform.  On successful completion `pack --version` should show a version number of `0.28` or greater.

Creating a buildpack identified as `examples/my_buildpack` results in a scaffolded buildpack in the `my_buildpack` folder:

```bash
$ pack buildpack create -s bash
? Enter a directory in which to scaffold the project: bash_buildpack
? Enter an identifier for the buildpack: examples/my_buildpack
? Choose the buildpack API version (use the default if you are unsure):
  ▸ 0.8
    0.7
✔ Enter the initial buildpack version: 0.0.1
✔ Enter a stack this buildpack will support by default: io.buildpacks.samples.stacks.bionic
✔ Enter a valid Go module name for this buildpack: github.com/user/buildpack
Successfully created examples/my_buildpack
```
