#!/usr/bin/env python3
"""Run lint checks on the codebase."""

import subprocess

def run_lint():
    """Run lint checks."""
    print("Running lint checks...")
    
    # Run ruff check
    print("\n=== Running ruff check ===")
    result = subprocess.run(["ruff", "check", "."], capture_output=True, text=True)
    print(result.stdout)
    if result.stderr:
        print("Errors:", result.stderr)
    
    # Run ruff format check
    print("\n=== Running ruff format check ===")
    result = subprocess.run(["ruff", "format", ".", "--check"], capture_output=True, text=True)
    print(result.stdout)
    if result.stderr:
        print("Errors:", result.stderr)
    
    # Run mypy
    print("\n=== Running mypy ===")
    result = subprocess.run(["mypy", "."], capture_output=True, text=True)
    print(result.stdout)
    if result.stderr:
        print("Errors:", result.stderr)

if __name__ == "__main__":
    run_lint()