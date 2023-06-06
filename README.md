# gh-runmd

This is a hackaton project.

A [gh](https://github.com/cli/cli) extension to run code blocks from markdown files. The goal for this is to make it easier and faster to run setups and examples from README files. The default barebones behaviour should take a repo as an argument, get the README, build one script from all the code blocks and run it. This can also be used by maintainers to test their READMEs.

```bash
gh runmd your/repo
```

Additionally, `gh runmd` should be able to run specific "paths": maintainers should be able to mark certain blocks as part of a path (for example, a certain OS) and users should be able to run only those blocks.

```bash
gh runmd your/repo --path macos
```

Maintainers should be able to mark certain blocks as non runnable, for example, blocks that are only meant to be read.

```markdown
```bash{runmd=false}
```

Users should be able to preview the script and pick which blocks to run.

Users should be able to provide both a URL and a local path.

### Questions

Should maintainers be able to specify wait conditions for each block before on to the next one? This can be useful if there's an async setup happening, but it's also a way of running code that is not in plain sight (unless it's a straight up "sleep").