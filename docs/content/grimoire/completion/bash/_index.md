---
title: "bash"
description: "Generate the autocompletion script for bash"
weight: 3
date: 2024-11-09
draft: false
---

# grimoire completion bash

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(grimoire completion bash)

To load completions for every new session, execute once:

#### Linux:

	grimoire completion bash > /etc/bash_completion.d/grimoire

#### macOS:

	grimoire completion bash > $(brew --prefix)/etc/bash_completion.d/grimoire

You will need to start a new shell for this setup to take effect.


## Usage

```bash
grimoire completion bash
```

## Flags

| Flag | Description | Default |
|------|-------------|----------|
| `--no-descriptions` | disable completion descriptions | `false` |

