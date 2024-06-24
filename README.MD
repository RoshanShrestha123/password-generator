# Password Generator

A simple terminal-based password generator written in Go.

## Features
- Generates a random password.
- Option to include numbers in the password.

## Usage

1. Clone the repository:
   ```sh
   git clone https://github.com/RoshanShrestha123/password-generator
   ```

2. Navigate to the project directory:
   ```sh
   cd password-generator
   ```

3. Build the project:
   ```sh
   go build -o password-generator main.go
   ```

4. Run the password generator:
   ```sh
   ./password-generator [OPTIONS]
   ```

## Options

- `--num=true`: Include numbers in the generated password.

## Example

To generate a password with numbers included (optional):
```sh
./password-generator --num=true
```

## Requirements

- Go 1.15 or higher

## License

This project is licensed under the MIT License.

---
Roshan Shrestha | Learning go from scratch