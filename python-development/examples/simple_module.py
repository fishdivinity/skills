"""Simple module example following PEP8 and best practices."""

from typing import Optional, Sequence


class Calculator:
    """A simple calculator class."""

    def add(self, a: float, b: float) -> float:
        """Add two numbers.

        Args:
            a: First number
            b: Second number

        Returns:
            Sum of the two numbers
        """
        return a + b

    def subtract(self, a: float, b: float) -> float:
        """Subtract two numbers.

        Args:
            a: First number
            b: Second number

        Returns:
            Difference of the two numbers
        """
        return a - b

    def multiply(self, a: float, b: float) -> float:
        """Multiply two numbers.

        Args:
            a: First number
            b: Second number

        Returns:
            Product of the two numbers
        """
        return a * b

    def divide(self, a: float, b: float) -> Optional[float]:
        """Divide two numbers.

        Args:
            a: First number
            b: Second number

        Returns:
            Quotient of the two numbers, or None if division by zero

        Raises:
            ValueError: If b is zero
        """
        if b == 0:
            raise ValueError("Division by zero is not allowed")
        return a / b


def calculate_average(numbers: Sequence[float | int]) -> Optional[float]:
    """Calculate the average of a list of numbers.

    Args:
        numbers: List of numeric values

    Returns:
        Average value or None if list is empty
    """
    if not numbers:
        return None

    return sum(numbers) / len(numbers)


def main() -> None:
    """Main function to demonstrate the module."""
    calculator = Calculator()
    
    print("Calculator operations:")
    print(f"1 + 2 = {calculator.add(1, 2)}")
    print(f"5 - 3 = {calculator.subtract(5, 3)}")
    print(f"2 * 4 = {calculator.multiply(2, 4)}")
    
    try:
        print(f"10 / 2 = {calculator.divide(10, 2)}")
        print(f"10 / 0 = {calculator.divide(10, 0)}")
    except ValueError as e:
        print(f"Error: {e}")
    
    numbers = [1, 2, 3, 4, 5]
    average = calculate_average(numbers)
    print(f"Average of {numbers} is {average}")


if __name__ == "__main__":
    main()