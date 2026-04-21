# Python 项目结构示例

## 标准项目结构

```
myproject/
├── src/
│   └── myproject/
│       ├── __init__.py
│       ├── core/
│       │   ├── __init__.py
│       │   ├── config.py
│       │   └── utils.py
│       ├── models/
│       │   ├── __init__.py
│       │   └── user.py
│       ├── services/
│       │   ├── __init__.py
│       │   └── user_service.py
│       └── exceptions.py
├── tests/
│   ├── __init__.py
│   ├── unit/
│   │   ├── __init__.py
│   │   ├── test_utils.py
│   │   └── test_user_service.py
│   └── integration/
│       ├── __init__.py
│       └── test_integration.py
├── docs/
├── examples/
├── scripts/
│   └── build.sh
├── requirements.txt
├── requirements-dev.txt
├── pyproject.toml
├── setup.py
└── README.md
```

## 示例代码

### src/myproject/__init__.py
```python
"""MyProject package."""

__version__ = "0.1.0"
```

### src/myproject/core/config.py
```python
"""Configuration module."""

from typing import Dict, Any
import os
from dotenv import load_dotenv

load_dotenv()


class Config:
    """Application configuration."""

    def __init__(self):
        self.api_key = os.getenv("API_KEY", "")
        self.database_url = os.getenv("DATABASE_URL", "")
        self.debug = os.getenv("DEBUG", "False").lower() == "true"

    def to_dict(self) -> Dict[str, Any]:
        """Convert config to dictionary."""
        return {
            "api_key": self.api_key,
            "database_url": self.database_url,
            "debug": self.debug,
        }


config = Config()
```

### src/myproject/core/utils.py
```python
"""Utility functions."""

from typing import List, Any


def validate_email(email: str) -> bool:
    """Validate email address.

    Args:
        email: Email address to validate

    Returns:
        True if email is valid, False otherwise
    """
    import re
    pattern = r"^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$"
    return bool(re.match(pattern, email))


def sanitize_input(input_str: str) -> str:
    """Sanitize user input.

    Args:
        input_str: User input to sanitize

    Returns:
        Sanitized input string
    """
    # Basic sanitization
    return input_str.strip()
```

### src/myproject/models/user.py
```python
"""User model."""

from dataclasses import dataclass
from typing import Optional


@dataclass
class User:
    """User data class."""

    id: int
    username: str
    email: str
    full_name: Optional[str] = None
    active: bool = True

    def __post_init__(self):
        """Validate user data."""
        if not self.username:
            raise ValueError("Username cannot be empty")
        if not self.email:
            raise ValueError("Email cannot be empty")
```

### src/myproject/services/user_service.py
```python
"""User service."""

from typing import List, Optional
from myproject.models.user import User
from myproject.core.utils import validate_email
from myproject.exceptions import UserNotFoundError


class UserService:
    """User service class."""

    def __init__(self):
        self.users: List[User] = []
        self.next_id = 1

    def create_user(self, username: str, email: str, full_name: Optional[str] = None) -> User:
        """Create a new user.

        Args:
            username: Username
            email: Email address
            full_name: Full name (optional)

        Returns:
            Created user

        Raises:
            ValueError: If email is invalid
        """
        if not validate_email(email):
            raise ValueError("Invalid email address")

        user = User(
            id=self.next_id,
            username=username,
            email=email,
            full_name=full_name
        )
        self.users.append(user)
        self.next_id += 1
        return user

    def get_user_by_id(self, user_id: int) -> User:
        """Get user by ID.

        Args:
            user_id: User ID

        Returns:
            User object

        Raises:
            UserNotFoundError: If user not found
        """
        for user in self.users:
            if user.id == user_id:
                return user
        raise UserNotFoundError(f"User with ID {user_id} not found")

    def get_all_users(self) -> List[User]:
        """Get all users.

        Returns:
            List of all users
        """
        return self.users

    def update_user(self, user_id: int, **kwargs) -> User:
        """Update user.

        Args:
            user_id: User ID
            **kwargs: User attributes to update

        Returns:
            Updated user

        Raises:
            UserNotFoundError: If user not found
            ValueError: If email is invalid
        """
        user = self.get_user_by_id(user_id)

        if "email" in kwargs:
            email = kwargs["email"]
            if not validate_email(email):
                raise ValueError("Invalid email address")

        for key, value in kwargs.items():
            if hasattr(user, key):
                setattr(user, key, value)

        return user

    def delete_user(self, user_id: int) -> bool:
        """Delete user.

        Args:
            user_id: User ID

        Returns:
            True if user was deleted

        Raises:
            UserNotFoundError: If user not found
        """
        user = self.get_user_by_id(user_id)
        self.users.remove(user)
        return True
```

### src/myproject/exceptions.py
```python
"""Custom exceptions."""


class MyProjectError(Exception):
    """Base exception for MyProject."""
    pass


class UserNotFoundError(MyProjectError):
    """Raised when user is not found."""
    pass


class ValidationError(MyProjectError):
    """Raised when validation fails."""
    pass
```

### tests/unit/test_utils.py
```python
"""Tests for utility functions."""

import pytest
from myproject.core.utils import validate_email, sanitize_input


def test_validate_email():
    """Test email validation."""
    assert validate_email("test@example.com") is True
    assert validate_email("invalid-email") is False
    assert validate_email("") is False


def test_sanitize_input():
    """Test input sanitization."""
    assert sanitize_input("  hello world  ") == "hello world"
    assert sanitize_input("test") == "test"
    assert sanitize_input("") == ""
```

### tests/unit/test_user_service.py
```python
"""Tests for user service."""

import pytest
from myproject.services.user_service import UserService
from myproject.exceptions import UserNotFoundError


def test_create_user():
    """Test user creation."""
    service = UserService()
    user = service.create_user("testuser", "test@example.com", "Test User")
    assert user.id == 1
    assert user.username == "testuser"
    assert user.email == "test@example.com"
    assert user.full_name == "Test User"


def test_get_user_by_id():
    """Test getting user by ID."""
    service = UserService()
    user = service.create_user("testuser", "test@example.com")
    retrieved_user = service.get_user_by_id(user.id)
    assert retrieved_user.id == user.id


def test_get_user_not_found():
    """Test getting non-existent user."""
    service = UserService()
    with pytest.raises(UserNotFoundError):
        service.get_user_by_id(999)


def test_update_user():
    """Test updating user."""
    service = UserService()
    user = service.create_user("testuser", "test@example.com")
    updated_user = service.update_user(user.id, full_name="Updated Name")
    assert updated_user.full_name == "Updated Name"


def test_delete_user():
    """Test deleting user."""
    service = UserService()
    user = service.create_user("testuser", "test@example.com")
    assert service.delete_user(user.id) is True
    with pytest.raises(UserNotFoundError):
        service.get_user_by_id(user.id)
```

### requirements.txt
```
python-dotenv>=1.0.0
```

### requirements-dev.txt
```
-r requirements.txt
pytest>=7.0.0
ruff>=0.4.0
mypy>=1.0.0
```

### pyproject.toml
```toml
[project]
name = "myproject"
version = "0.1.0"
description = "My project description"
authors = [
    { name = "Author Name", email = "author@example.com" },
]
dependencies = [
    "python-dotenv>=1.0.0",
]

[build-system]
requires = ["setuptools", "wheel"]
build-backend = "setuptools.build_meta"

[tool.ruff]
line-length = 79
target-version = "py313"

[tool.ruff.lint]
select = [
    "E",
    "W",
    "F",
    "I",
    "B",
    "C4",
    "UP",
    "SIM",
]

[tool.ruff.format]
quote-style = "double"
indent-style = "space"

[tool.mypy]
python_version = "3.13"
strict = true
warn_return_any = true
warn_unused_ignores = true
disallow_untyped_defs = true
disallow_incomplete_defs = true

[tool.pytest.ini_options]
testpaths = ["tests"]
pythonpath = ["src"]
```

### setup.py
```python
"""Setup script for myproject."""

from setuptools import setup, find_packages

setup(
    name="myproject",
    version="0.1.0",
    packages=find_packages(where="src"),
    package_dir={"": "src"},
    install_requires=[
        "python-dotenv>=1.0.0",
    ],
    extras_require={
        "dev": [
            "pytest>=7.0.0",
            "ruff>=0.4.0",
            "mypy>=1.0.0",
        ],
    },
)
```

### README.md
```markdown
# MyProject

A Python project demonstrating best practices.

## Features

- PEP8 compliant code
- Type hints
- Modular structure
- Comprehensive tests
- Environment management

## Setup

1. Create virtual environment:
   ```
   python -m venv venv
   ```

2. Activate virtual environment:
   ```
   # Windows
   venv\Scripts\activate
   
   # Linux/Mac
   source venv/bin/activate
   ```

3. Install dependencies:
   ```
   pip install -r requirements-dev.txt
   ```

4. Run tests:
   ```
   pytest
   ```

5. Run linting:
   ```
   ruff check .
   ```

6. Run type checking:
   ```
   mypy .
   ```

## Project Structure

- `src/myproject/`: Main source code
- `tests/`: Test files
- `docs/`: Documentation
- `examples/`: Example code
- `scripts/`: Utility scripts

## Usage

```python
from myproject.services.user_service import UserService

# Create user service
service = UserService()

# Create user
user = service.create_user("john", "john@example.com", "John Doe")

# Get user
retrieved_user = service.get_user_by_id(user.id)
print(retrieved_user)

# Update user
service.update_user(user.id, full_name="John Smith")

# Delete user
service.delete_user(user.id)
```

## License

MIT
```