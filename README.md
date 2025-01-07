# CLI Calculator in Go 🧮

A simple command-line calculator built with Go (Golang) that supports basic arithmetic operations: addition, subtraction, multiplication, and division.

## Features ✨

- Perform basic operations:
  - Addition (`+`)
  - Subtraction (`-`)
  - Multiplication (`*`)
  - Division (`/`)
- Handles division by zero gracefully.
- Allows continuous calculations in a loop until the user chooses to exit.
- User-friendly prompts and error handling.

## How to Use 🚀

1. Clone this repository:
   ```bash
   git clone https://github.com/your-username/cli-calculator.git
   cd cli-calculator

2.	Build and run the program:
    `go run main.go`

3.	Follow the prompts:
	•	Enter two numbers.
	•	Choose an operation (+, -, *, /).
	•	View the result.
	•	Decide whether to perform another calculation or exit.

Example Interaction 🖥️
Welcome to the calculator!
Enter the first number:
10
Enter the second number:
5
Choose an operation (+, -, *, /):
+
Result: 10.00 + 5.00 = 15.00
Would you like to perform another operation? (yes to continue, no to exit):
yes
Enter the first number:
20
Enter the second number:
0
Choose an operation (+, -, *, /):
/
Error: Division by zero is not allowed.
Would you like to perform another operation? (yes to continue, no to exit):
no
Thank you for using the calculator. Goodbye!


Project Structure 🗂️
	•	main.go: The main file containing the CLI calculator logic.

Requirements 📦
	•	Go version 1.19 or higher.

How to Contribute 🤝
	1.	Fork this repository.
	2.	Create a new branch (git checkout -b feature/your-feature-name).
	3.	Commit your changes (git commit -m 'Add some feature').
	4.	Push to the branch (git push origin feature/your-feature-name).
	5.	Open a pull request.
