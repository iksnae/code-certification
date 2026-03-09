---
description: Test Drive the plan
argument-hint: [path-to-plan]
---

# Build

Follow the `Workflow` to implement the `PATH_TO_PLAN` then `Report` the completed work.

## Variables

PATH_TO_PLAN: $ARGUMENTS

## Workflow

- If no `PATH_TO_PLAN` is provided, STOP immediately and ask the user to provide it (AskUserQuestion).
- Read and execute the plan at `PATH_TO_PLAN`.
- Check if all tests pass or exist and surface missing or existing issues to user (AskUserQuestion).
- Think hard about the plan and logical tests to verify deliverables.
- If plan includes testing planning, use it as a guide for testing and start `TDD Loop`. Otherwise, create your own test plan.
- Confirm test plan with user (AskUserQuestion) to apply `TDD Loop`.

## TDD Loop

Repeat this 3-step loop until all tests pass: Write Test -> Run Test -> Refactor.


### Write Test

- Create a new test file in the same directory as the plan file.
- Use the plan's description and steps to create a failing test that verifies the expected outcome of each step.
- If the plan includes testing planning, use it as a guide for writing tests. Otherwise, create your own tests.

### Run Test

- Run the tests
- If all tests pass, proceed to refactor. Otherwise, go back to Write Test and refine the tests based on the failing output.

### Refactor

- Review the code and refactor it according to best practices.
- Remove any duplication or unnecessary code.
- Ensure that the code is clean, maintainable, and follows the project's coding standards.
- Once refactoring is complete, go back to Write Test and new tests based on the updated code.
- Commit changes with comprehensive messages.

## Report

Once all tests pass, report the completed work by updating the plan file with the following information:

- The date of completion
- A brief description of what was implemented
- Any issues encountered during implementation and how they were resolved
- Any refactoring that was done to improve the codebase
- Present the `## Report`.
