---
title: "zsh"
description: "Generate the autocompletion script for zsh"
weight: 6
date: 2024-11-09
draft: false
---

# grimoire completion zsh

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(grimoire completion zsh)

To load completions for every new session, execute once:

#### Linux:

	grimoire completion zsh > "${fpath[1]}/_grimoire"

#### macOS:

	grimoire completion zsh > $(brew --prefix)/share/zsh/site-functions/_grimoire

You will need to start a new shell for this setup to take effect.


## Usage

```bash
grimoire completion zsh [flags]
```

## Flags

| Flag | Description | Default |
|------|-------------|----------|
| `--no-descriptions` | disable completion descriptions | `false` |

