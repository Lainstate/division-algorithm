# Custom Division Algorithm

A Go program that divides a given number progressively until the divided values sum up to the original one. The program stops dividing when the divided values become too small (less than `1e-6`).

## Features

- Accepts a number as user input.
- Divides the input by 2 repeatedly.
- Stops when the divided values are smaller than `1e-6`.
- Ensures the sum of all divided values equals the original input value.
- Uses high precision calculations with `math/big` to avoid rounding errors.

## Usage

1. Clone this repository or download the source code.
2. Run the program:

    ```bash
    go run .
    ```

3. Enter a number when prompted.
4. The program will display each divided value and the final sum, which matches the original number.

## Example Output

```bash
Enter a number to divide: 7.51

--- Dividing Values ---
Step  1: 3.7550000000
Step  2: 1.8775000000
Step  3: 0.9387500000
Step  4: 0.4693750000
Step  5: 0.2346875000
...
...
--- Results ---
Final Sum:    7.5100000000
Original Value: 7.5100000000
