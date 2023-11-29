README
======

This `release-notes-publisher` converts HTML content to XHTML, with additional functionality to handle JSON data representing project versions and the option to send the generated XHTML content to a Confluence page via REST API.

Requirements
------------
- Access to the terminal or command line
- Input HTML content
- JSON files with project version information

Installation
------------
To install this program, you need to clone the repository or download the source code to your local machine.

Usage
-----
Run the program using the command line. Here are the available flags:

* `-i` : Input HTML file path (optional, reads from stdin if not provided)
* `-o` : Output XHTML file path or Confluence API URL (required)
```shell
  - File output example: file:///path/to/output.xhtml
  - Confluence output example: confluence://https://your-confluence-instance.com
```
* `--minify` : Minify the XHTML output (optional)
* `--escape-for-json` : Escape XHTML for embedding into JSON (optional)
* `--versions-filepath` : Path of versions JSON file (default: project_versions.json)
* `--mocks-versions-filepath` : Path of mocks versions JSON file (default: project_versions_mocks.json)
* `--confluence-page-title` : Title of the Confluence page (required for Confluence output)
* `--confluence-space-code` : Space code of the Confluence space (required for Confluence output)
* `--confluence-auth-personal-token` : Personal access token for Confluence API (required for Confluence output)

Example
-------

```shell
# Convert and send to a file:
./release-notes-publisher -i input.html -o file:///path/to/output.xhtml --minify

# Convert and send to Confluence:
./release-notes-publisher -i input.html -o confluence://https://your-confluence-instance.com --confluence-page-title "My Page" --confluence-space-code "MYSPACE" --confluence-auth-personal-token "token_goes_here"

# Full example
./release-notes-publisher -i releasenotes.html --versions-filepath ./project_versions.json --mocks-versions-filepath ./project_versions_mocks.json --minify --escape-for-json -o confluence://https://some.url.com --confluence-ancestor-page-id 86572322 --confluence-page-title "new page 1" --confluence-space-code SOMESPACE --confluence-auth-personal-token "SOMETOKEN"
```
Note: Replace the placeholders with actual file paths and Confluence details.

Contributing
------------
Contributions to this project are welcome. Please fork the repository and submit a pull request with your changes.

License
-------
The software is open for unrestricted use.

Support
-------
For support, please open an issue on the project's issue tracker.

