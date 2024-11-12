# Grimoire

![](/docs/images/grimoire.jpeg)

A prompt engineer's ***Grimoire*** to assist with getting the best results from Generative AI prompts.

## Features

- **Prompt Management**: Easily manage and organize your prompts.
- **Clipboard Integration**: Seamlessly copy and paste prompts using the clipboard service.
- **Command Line Interface**: Interact with Grimoire through a user-friendly CLI.
- **Ollama Integration**: Generate text using AI models with Ollama's local LLM integration.

## Installation

To install Grimoire, follow these steps:

1. Clone the repository and navigate to the directory:

```sh
git clone https://github.com/gphorvath/grimoire.git
cd grimoire
```

1. Build and install using Make:

```sh
# Build the binary
make build

# Install binary and prompts (requires sudo)
make install
```

This will:

- Build the Grimoire binary
- Create the ~/.grimoire directory
- Copy default prompts to ~/.grimoire/prompts
- Install the binary to /usr/local/bin

## Usage

After installation, you can use Grimoire through its CLI:

```sh
# List all available prompts
grimoire list

# Copy a prompt to clipboard
grimoire copy <prompt-name>

# Create a new prompt
grimoire new <prompt-name>

# Edit an existing prompt in default editor
grimoire edit <prompt-name>

# Use echo prompt to verify prompt loading
# Requires Ollama running locally with a configured model (default: llama3)
grimoire generate --prompt echo "Testing generation functionality"

# Modify prompt in default editor before generation (doesn't require arg)
grimoire generate --prompt echo --edit
```

## Environment Variables

GRIMOIRE environment variables control the configuration of the application.

### Ollama Configuration

- `GRIMOIRE_OLLAMA_MODEL` (default: "llama3")
    The model to be used with Ollama for AI inference.

- `GRIMOIRE_OLLAMA_URL` (default: "<http://localhost:11434/api>")
    The URL endpoint for the Ollama API server.

- `GRIMOIRE_OLLAMA_STREAM` (default: true)
    Controls whether to stream responses from Ollama. Set to "true" to enable streaming, "false" to disable.

### Editor Configuration

- `GRIMOIRE_EDITOR` (default: "vim")
    The text editor to use when editing prompts or configuration files.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
