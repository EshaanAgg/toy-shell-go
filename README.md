[![progress-banner](https://backend.codecrafters.io/progress/shell/81b300f3-c367-49d8-a88d-3b72b05cb01b)](https://app.codecrafters.io/users/codecrafters-bot?r=2qF)

This is my Go solution to the ["Build Your Own Shell" Challenge](https://app.codecrafters.io/courses/shell/overview) by Codecrafters, born out of the realization that [Rust is too diffcult for me](https://github.com/EshaanAgg/toy-shell) to sanely code in.  

This is a simple Unix shell built from scratch, which supports:
- Basic commands like `echo`, `type`, `exit`
- Parsing of the `PATH` environment variable to find exectuables to run
- Navigation of the file system with `pwd` and `cd`, along with parsing of the `HOME` environment variable

## Working in Raw Mode

To implement features like autocomplete and tab completion, it was necessary to use the [`term`](https://pkg.go.dev/golang.org/x/term) library to switch the terminal into raw mode. This allows for more control over the input and output of the terminal.

In the context of the raw mode, the following are good to know:
- `\n`: Newline character. It moves the cursor to the next line, at the same column.
- `\r`: Carriage return character. It moves the cursor to the beginning of the current line.
- `\r\n`: Carriage return and newline character. It moves the cursor to the beginning of the next line.