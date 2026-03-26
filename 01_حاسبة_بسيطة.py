# Enhanced Calculator

class EnhancedCalculator:
    def __init__(self):
        self.history = []

    def add(self, a, b):
        result = a + b
        self.history.append(f"{a} + {b} = {result}")
        return result

    def subtract(self, a, b):
        result = a - b
        self.history.append(f"{a} - {b} = {result}")
        return result

    def multiply(self, a, b):
        result = a * b
        self.history.append(f"{a} * {b} = {result}")
        return result

    def divide(self, a, b):
        if b == 0:
            result = "Error: Division by zero"
            self.history.append(f"{a} / {b} = {result}")
            return result
        result = a / b
        self.history.append(f"{a} / {b} = {result}")
        return result

    def show_history(self):
        return self.history

# Example usage
if __name__ == '__main__':
    calc = EnhancedCalculator()
    print(calc.add(10, 5))       # Should print 15
    print(calc.subtract(10, 5))  # Should print 5
    print(calc.multiply(10, 5))  # Should print 50
    print(calc.divide(10, 0))    # Should print Error
    print(calc.show_history())    # Should show all operations done
