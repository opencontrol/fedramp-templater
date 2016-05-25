# FedRAMP Templater

This is a command-line tool to take the [FedRAMP](http://www.fedramp.gov/) System Security Plan template and transform it to be [Compliance Masonry](https://github.com/opencontrol/compliance-masonry)-compatible.

## Usage

Requires [Go](https://golang.org/) 1.2+.

1. Install the tool:

    ```bash
    go get github.com/opencontrol/fedramp-templater
    ```

1. [Download the `System Security Plan (SSP)` template.](https://www.fedramp.gov/resources/templates-3/) (Tested with [v2.1](https://www.fedramp.gov/files/2015/03/FedRAMP-System-Security-Plan-Template-v2.1.docx).)
1. Run

    ```bash
    fedramp-templater <input> <output>
    # i.e.
    fedramp-templater FedRAMP-System-Security-Plan-Template-v2.1.docx FedRAMP-Masonry-Template-v2.1.docx
    ```

The output document will be the same as the input one, albeit with a bunch of Compliance Masonry tags (e.g. `{{ getControl "NIST-800-53@CM-2"}}`) inserted.
