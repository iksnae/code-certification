#!/usr/bin/env python3
"""
Validator: Check that files contain required sections.

Usage:
  validate_file_contains.py --directory <dir> --extension <ext> --contains <text> [--contains <text>...]
"""

import argparse
import sys
from pathlib import Path


def main():
    parser = argparse.ArgumentParser(description="Validate that files contain required text")
    parser.add_argument("--directory", required=True, help="Directory to search")
    parser.add_argument("--extension", required=True, help="File extension to check (e.g., .md)")
    parser.add_argument("--contains", action="append", required=True, help="Required text to find (can be specified multiple times)")

    args = parser.parse_args()

    directory = Path(args.directory)
    extension = args.extension if args.extension.startswith(".") else f".{args.extension}"
    required_texts = args.contains

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
        with open(file_path, "r", encoding="utf-8") as f:
            content = f.read()

        missing = []
        for required_text in required_texts:
            if required_text not in content:
                missing.append(required_text)

        if missing:
            all_valid = False
            print(f"✗ {file_path}: Missing required sections:")
            for text in missing:
                print(f"  - {text}")
        else:
            print(f"✓ {file_path}: All required sections present")

    sys.exit(0 if all_valid else 1)


if __name__ == "__main__":
    main()
