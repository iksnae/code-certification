#!/usr/bin/env python3
"""
Validator: Check that new files follow proper format.

Usage:
  validate_new_file.py --directory <dir> --extension <ext>
"""

import argparse
import sys
from pathlib import Path


def main():
    parser = argparse.ArgumentParser(description="Validate that new files follow proper format")
    parser.add_argument("--directory", required=True, help="Directory to check")
    parser.add_argument("--extension", required=True, help="File extension to check (e.g., .md)")

    args = parser.parse_args()

    directory = Path(args.directory)
    extension = args.extension if args.extension.startswith(".") else f".{args.extension}"

    if not directory.exists():
        print(f"✓ Directory {directory} does not exist, skipping validation")
        sys.exit(0)

    # Find all files with the specified extension
    files = list(directory.glob(f"*{extension}"))

    if not files:
        print(f"✓ No {extension} files found in {directory}")
        sys.exit(0)

    all_valid = True
    for file_path in files:
        try:
            with open(file_path, "r", encoding="utf-8") as f:
                content = f.read()

            # Basic validations for markdown plan files
            issues = []

            # Check for required markdown structure
            if not content.strip():
                issues.append("File is empty")

            # Check for h1 header
            if not any(line.startswith("# ") for line in content.split("\n")):
                issues.append("Missing top-level header (# Title)")

            # Note: Line length check removed—long lines are common in markdown docs

            if issues:
                all_valid = False
                print(f"⚠ {file_path}: Format issues detected:")
                for issue in issues:
                    print(f"  - {issue}")
            else:
                print(f"✓ {file_path}: Format is valid")

        except Exception as e:
            all_valid = False
            print(f"✗ {file_path}: Error reading file - {e}")

    sys.exit(0 if all_valid else 1)


if __name__ == "__main__":
    main()
