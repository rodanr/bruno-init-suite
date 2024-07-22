
# Bruno Init Suite CLI Documentation

## Overview

`bruno-init-suite` is a CLI tool designed to automate the creation of Bruno documentation and integration with third-party services, including scripts and configurations. This tool simplifies the initialization of Bruno projects with custom configurations, such as using AWS Cognito for authentication.

## Installation

To install `bruno-init-suite`, download the latest release from GitHub and make the binary executable:

### For Linux and macOS

```sh
curl -L https://github.com/rodanr/bruno-init-suite/releases/latest/download/bruno-init-suite-linux-amd64.tar.gz -o /tmp/bruno-init-suite.tar.gz
tar -xzf /tmp/bruno-init-suite.tar.gz -C ~/.local/bin
chmod a+rx ~/.local/bin/bruno-init-suite
```

### For Windows

Download the latest release from [GitHub](https://github.com/rodanr/bruno-init-suite/releases/latest/download/bruno-init-suite-windows-amd64.exe.zip), unzip it, and add the executable to your system's PATH.

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
bruno-init-suite init [flags]
```

### Flags

- **--cognito, -c**: Use AWS Cognito for authentication (default: false).
- **--name, -n**: Name of the project (default: "bruno-docs").
- **--baseUrl, -b**: Base URL for the project (default: "http://localhost:8080").

### Example

Initialize a new Bruno project named "my-bruno-docs" with AWS Cognito authentication and a custom base URL:

```sh
bruno-init-suite init -c -n my-bruno-docs -b https://api.example.com
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

## Example Workflow

1. **Download and unarchive the `bruno-init-suite` binary**:

    ```sh
    curl -L https://github.com/rodanr/bruno-init-suite/releases/latest/download/bruno-init-suite-linux-amd64.tar.gz -o /tmp/bruno-init-suite.tar.gz
    tar -xzf /tmp/bruno-init-suite.tar.gz -C ~/.local/bin
    chmod a+rx ~/.local/bin/bruno-init-suite
    ```

2. **Initialize a new project**:

    ```sh
    bruno-init-suite init -c -n my-bruno-docs -b https://api.example.com -o ~/my-bruno-project
    ```

3. **Navigate to the project directory**:

    ```sh
    cd ~/my-bruno-project
    ```

4. **View the generated files**:

    ```sh
    tree .
    ```

## Conclusion

`bruno-init-suite` CLI tool simplifies the process of initializing Bruno documentation projects with custom configurations. By following the above documentation, you can easily set up and use this tool to generate Bruno docs tailored to your needs.
