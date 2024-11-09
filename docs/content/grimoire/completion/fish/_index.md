---
title: "fish"
description: "Generate the autocompletion script for fish"
weight: 4
date: 2024-11-09
draft: false
---

# grimoire completion fish

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	grimoire completion fish | source

To load completions for every new session, execute once:

	grimoire completion fish > ~/.config/fish/completions/grimoire.fish

You will need to start a new shell for this setup to take effect.


## Usage

```bash
grimoire completion fish [flags]
```

## Flags

| Flag | Description | Default |
|------|-------------|----------|
| `--no-descriptions` | disable completion descriptions | `false` |

