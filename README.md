# GitHub Monitor

GitHub Monitor is a web application designed to monitor pending approvals across GitHub repositories. It provides a simple interface to track pull requests and actions within specified repositories of a GitHub organization.

## Features

- Monitor pending pull requests across multiple repos.
- Track GitHub Actions workflows and their statuses across multiple repos.

## Prerequisites

Before you begin, ensure you have met the following requirements:

- Docker and Docker Compose installed on your machine.
- A GitHub Personal Access Token with appropriate permissions to access the repositories you want to monitor.
  - Navigate to https://github.com/settings/tokens
  - Create a new fine-grained access token
    - Select the Org you want to monitor under `resource owner`
    - Select `All repositories` or manually select each repo
    - Under permissions grant the following:
      - `Actions`: `read/write`
      - `Issues`: `read-only`
      - `Pull requests`: `read/write`

## Getting Started

1. Clone the repository to your local machine:

```sh
git clone https://github.com/steven-harris/github-monitor.git
cd github-monitor
```

2. Copy the `.env-example` file to `.env`:

```sh
cp .env-example .env
```

3. Open the `.env` file in a text editor and replace the placeholder `<>` with your GitHub Personal Access Token:

```plaintext
GITHUB_TOKEN=<your_github_personal_access_token>
```

4. In the [`docker-compose.yml`](./docker-compose.yml) file, you can customize the `GITHUB_ORG`, `PR_REPOS`, and `ACTION_REPOS` environment variables to monitor specific repositories and actions within your GitHub organization.

Example configuration:

```yml
environment:
  GITHUB_TOKEN: ${GITHUB_TOKEN}
  GITHUB_ORG: YourGitHubOrg
  PR_REPOS: |
    repo1,
    repo2+label:important,
  ACTION_REPOS: |
    repo1,
    repo2
  ACTION_FILTER: deploy
```

**Please note and environment variables with the same names will override docker-compose**

## Running the Application

To run GitHub Monitor locally, execute the following command in the root directory of the project:

```sh
docker-compose up --build
```

This command builds the Docker image and starts the application. Once the application is running, you can access it at `http://localhost:8888`.

## License

This project is licensed under the GNU General Public License v3.0. See the [`LICENSE`](.\LICENSE") file for more details.

