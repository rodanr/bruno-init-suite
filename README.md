
# Bruno Init Suite CLI Documentation

## Overview
<img align="right" width="128px" src="https://github.com/user-attachments/assets/11703489-d50a-45bb-8caf-7b5a746454db">

`bruno-init-suite` is a CLI tool designed to automate the creation of Bruno documentation and integration with third-party services, including scripts and configurations. This tool simplifies the initialization of Bruno projects with custom configurations, such as using AWS Cognito for authentication.


# Bruno Init Suite CLI Installation Instructions

## Installation for Linux

### For Linux amd64

```sh
curl -L https://github.com/rodanr/bruno-init-suite/releases/latest/download/bruno-init-suite-linux-amd64.tar.gz -o /tmp/bruno-init-suite.tar.gz
tar -xzf /tmp/bruno-init-suite.tar.gz -C ~/.local/bin
mv ~/.local/bin/bruno-init-suite-linux-amd64 ~/.local/bin/bruis
chmod a+rx ~/.local/bin/bruis
```

### For Linux arm64

```sh
curl -L https://github.com/rodanr/bruno-init-suite/releases/latest/download/bruno-init-suite-linux-arm64.tar.gz -o /tmp/bruno-init-suite.tar.gz
tar -xzf /tmp/bruno-init-suite.tar.gz -C ~/.local/bin
mv ~/.local/bin/bruno-init-suite-linux-arm64 ~/.local/bin/bruis
chmod a+rx ~/.local/bin/bruis
```

## Installation for macOS

### For macOS amd64

```sh
curl -L https://github.com/rodanr/bruno-init-suite/releases/latest/download/bruno-init-suite-darwin-amd64.tar.gz -o /tmp/bruno-init-suite.tar.gz
tar -xzf /tmp/bruno-init-suite.tar.gz -C ~/.local/bin
mv ~/.local/bin/bruno-init-suite-darwin-amd64 ~/.local/bin/bruis
chmod a+rx ~/.local/bin/bruis
```

### For macOS arm64

```sh
curl -L https://github.com/rodanr/bruno-init-suite/releases/latest/download/bruno-init-suite-darwin-arm64.tar.gz -o /tmp/bruno-init-suite.tar.gz
tar -xzf /tmp/bruno-init-suite.tar.gz -C ~/.local/bin
mv ~/.local/bin/bruno-init-suite-darwin-arm64 ~/.local/bin/bruis
chmod a+rx ~/.local/bin/bruis
```

## Installation for Windows

### For Windows amd64

Download the latest release from [GitHub](https://github.com/rodanr/bruno-init-suite/releases/latest/download/bruno-init-suite-windows-amd64.exe.zip), unzip it, and rename the executable to `bruis.exe`. Add the executable to your system's PATH.

### For Windows 386

Download the latest release from [GitHub](https://github.com/rodanr/bruno-init-suite/releases/latest/download/bruno-init-suite-windows-386.exe.zip), unzip it, and rename the executable to `bruis.exe`. Add the executable to your system's PATH.
## Usage

Run the CLI tool with the following commands and options:

### Commands

- **init**: Initialize a new Bruno project with custom configurations.
- **version**: Show the version of Bruno Init Suite.

### Global Flags

- **--output, -o**: Output directory for generated Bruno docs.
- **--version, -v**: Show version of Bruno Init Suite.

### Initialize Command

The `init` command initializes a new Bruno project in the specified output directory with optional configurations.

```sh
bruis init [flags]
```

### Flags

- **--cognito, -c**: Use AWS Cognito for authentication (default: false).
- **--name, -n**: Name of the project (default: "bruno-docs").
- **--baseUrl, -b**: Base URL for the project (default: "http://localhost:8080").

### Example

Initialize a new Bruno project named "my-bruno-docs" with AWS Cognito authentication and a custom base URL:

```sh
mkdir my-bruno-docs
cd my-bruno-docs
bruis init -c -n my-bruno-docs -b https://api.example.com
```
or
```sh
bruis init -c -n my-bruno-docs -b https://api.example.com -o my-bruno-docs
```

This command will generate the necessary Bruno documentation files and configurations in the specified output directory.

## Project Structure

The generated project will have the following structure:

```
<output-dir>/
├── bruno.json
├── collection.bru
├── .env.example
├── environments/
│   └── environment.bru
└── lib/
    └── cognito-auth.js (if --cognito is specified)
```

## Contributing

We welcome contributions to improve this project! If you encounter any issues, have suggestions for features, or want to contribute in any way, please feel free to create a GitHub issue.

## YouTube Channel

For tutorials and more information, check out our [YouTube Channel](https://www.youtube.com/@BrunoInitSuite).
