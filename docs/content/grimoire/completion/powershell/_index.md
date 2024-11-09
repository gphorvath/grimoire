---
title: "powershell"
description: "Generate the autocompletion script for powershell"
weight: 5
date: 2024-11-09
draft: false
---

# grimoire completion powershell

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	grimoire completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


## Usage

```bash
grimoire completion powershell [flags]
```

## Flags

| Flag | Description | Default |
|------|-------------|----------|
| `--no-descriptions` | disable completion descriptions | `false` |

