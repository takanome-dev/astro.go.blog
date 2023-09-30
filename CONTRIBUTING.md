# Contributing

Thank you so much for your interest in contributing!. All types of contributions are encouraged and valued.

## Bugs and issues

Bug reports help make this project a better experience for everyone. When you report a bug, a template will be created automatically containing information we'd like to know.

Before raising a new issue, please search existing ones to make sure you're not creating a duplicate.

If you see an open issue you would like to work on, just comment `.take` and our Github Action will assign you to the issue.

Issues that are not assigned are assumed open, and to avoid conflicts, please assign yourself before beginning work on any issues.

> **Note**
>
> Assigned issues that have not had any activity in a week will be unassigned by the action. If you think that's too short, please open an issue to discuss it.

Next, you can use `git checkout -b <branch_name>` or `git switch -c <branch_name>` to create a new branch for your work. It's always a good idea to avoid committing changes directly to your default branch (main) - this keeps it clean and avoid some weird issues.

Branch names should be a brief description of your changes, such as `fix-typo` for fixing a typo.

> **Info**
>
> In order to make git commit messages easier to read and faster to reason about, we follow some guidelines on most commits to keep the format predictable.
> Check [Conventional Commits specification](https://www.conventionalcommits.org/) for more information about our guidelines.

## Pull Request

Push your changes to your forked repository by using `git push -u origin <branch_name>`.

- `-u` tells `git` to set the upstream, it's the same as `--set-upstream`
- `origin` tells `git` to push to your fork
- `branch_name` tells `git` to push to a branch - this MUST match the name of the branch you created locally.

Make sure to change the PR title in something like: `fix: correct typo` or `feat: add node latest`

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details
