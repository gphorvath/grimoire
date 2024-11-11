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

2. Build and install using Make:
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

# Edit an existing prompt
grimoire edit <prompt-name>

# Use echo prompt to verify prompt loading
# Requires Ollama running locally with a configured model (default: llama3)
grimoire generate -p echo "Testing generation functionality"
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
