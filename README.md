# FedRAMP Templater [![Build Status](https://travis-ci.org/opencontrol/fedramp-templater.svg?branch=master)](https://travis-ci.org/opencontrol/fedramp-templater)

This is a command-line tool to take the [FedRAMP](http://www.fedramp.gov/) System Security Plan template and transform it to be [Compliance Masonry](https://github.com/opencontrol/compliance-masonry)-compatible.

## Usage

Requires [Go](https://golang.org/) 1.6+. 


1. Install the tool:

    For Ubuntu/Debian:
    ```bash
    sudo apt-get install libxml2-dev 
    go get github.com/opencontrol/fedramp-templater
    ```
    
    For OsX:
    ```
    brew install homebrew/dupes/apple-gcc42
    export CC=/usr/local/Cellar/apple-gcc42/4.2.1-5666.3/bin/gcc-4.2
    brew install libxml2
    ln -s /usr/local/Cellar/libxml2/2.9.4/lib/pkgconfig/libxml-2.0.pc /usr/local/lib/pkgconfig/libxml-2.0.pc
    export PKG_CONFIG_PATH=/usr/local/lib/pkgconfig
    go get github.com/moovweb/gokogiri
    go get github.com/opencontrol/fedramp-templater
    ```
  


1. [Download the `System Security Plan (SSP)` template.](https://www.fedramp.gov/resources/templates-2016/) (Tested with [v2.1](https://www.fedramp.gov/files/2015/03/FedRAMP-System-Security-Plan-Template-v2.1.docx).)
1. Run

    ```bash
    fedramp-templater <input> <output>
    # i.e.
    fedramp-templater FedRAMP-System-Security-Plan-Template-v2.1.docx FedRAMP-Masonry-Template-v2.1.docx
    ```

The output document will be the same as the input one, albeit with a bunch of Compliance Masonry tags (e.g. `{{ getControl "NIST-800-53@CM-2"}}`) inserted. The resulting document can then be run through [the Compliance Masonry .docx formatter](https://github.com/opencontrol/compliance-masonry/#create-docx-template).
