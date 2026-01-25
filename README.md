# File Deduplication Tool

`dugo`, a fast and efficient command-line tool written in Go to find and remove duplicate files in a directory. It supports concurrency for improved performance compared to other tools.

---

## Features

- **Fast Duplicate Detection**: Uses file size and MD5 hashing to quickly identify potential duplicates.
- **Accurate Comparison**: Performs byte-by-byte comparison to confirm duplicates, avoiding false positives due to hash collisions.
- **Concurrency Support**: Leverages Go's goroutines to process files in parallel, speeding up the deduplication process.
- **Interactive Deletion**: Optionally prompts the user to delete selected duplicate files interactively.
- **Flexible Ignore Options**: Allows ignoring files or directories by name or regex pattern.
- **Customizable Workers**: Lets you control the number of concurrent workers for optimal performance.

---

## Installation

### Prerequisites
- Go (for building from source).

### Build from Source
1. Clone the repository:
   
2. Build the tool:
   ```bash
   go build -o dogu
   ```
3. Move the binary to a directory in your `PATH` (optional):
   ```bash
   sudo mv dogu /usr/local/bin/
   ```

---

## Usage

### Basic Usage
To find duplicates in a directory:
```bash
./dugo /path/to/directory
```

### Enable Interactive Deletion
To interactively delete duplicates:
```bash
./dugo -it /path/to/directory
```

### Ignore Files or Directories
- Ignore specific files or directories by name:
  ```bash
  ./dugo -ignore-names=".git,temp,backup" /path/to/directory
  ```
- Ignore files or directories using a regex pattern:
  ```bash
  ./dugo -ignore-regex=".*\.tmp$" /path/to/directory
  ```

### Control Concurrency
Set the number of concurrent workers (default: 4):
```bash
./dugo -workers=8 /path/to/directory
```

### Full Example
Find duplicates, ignore `.tmp` files, enable interactive deletion, and use 8 workers:
```bash
./dugo -ignore-regex=".*\.tmp$" -it -workers=8 /path/to/directory
```

---

## Options

| Flag            | Description                                                                 |
|-----------------|-----------------------------------------------------------------------------|
| `-ignore-names` | Comma-separated list of file/directory names to ignore (exact match).       |
| `-ignore-regex` | Regex pattern to ignore files/directories by path.                          |
| `-workers`      | Number of concurrent workers (default: 4).                                  |
| `-it`           | Enable interactive deletion of duplicate files.                             |

---

## How It Works

1. **Scan Directory**: The tool scans the specified directory and groups files by size.
2. **Hash Files**: Files with the same size are hashed using MD5.
3. **Compare Files**: Files with the same hash are compared byte-by-byte to confirm duplicates.
4. **Report or Delete**: Duplicates are either reported to the user or deleted interactively.

---

## Contributing

Contributions are welcome! Here’s how you can help:
1. Fork the repository.
2. Create a new branch for your feature or bugfix.
3. Submit a pull request with a detailed description of your changes.

---

## License

This project is licensed under the MIT License. See the [LICENSE](https://github.com/aladdin-io/dugo/blob/main/LICENSE) file for details.

---

## Acknowledgments

- Inspired by the need for a fast and accurate file deduplication tool.
- Built with Go’s powerful concurrency model for high performance.
