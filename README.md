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
    
    For OsX (assuming you have [HomeBrew installed](http://brew.sh/))
    ```
    brew install libxml2
    go get github.com/moovweb/gokogiri
    ```
   
    Note that installation issues are usually caused by installation Gokogiri, and if you run into issues you may find some help at [this issue](https://github.com/moovweb/gokogiri/issues/14) and with the [update to the GokoGiri README](https://github.com/moovweb/gokogiri/pull/95)


1. [Download the `System Security Plan (SSP)` template.](https://www.fedramp.gov/resources/templates-2016/) (Tested with [v2.1](https://www.fedramp.gov/files/2015/03/FedRAMP-System-Security-Plan-Template-v2.1.docx).)
1. Run

    ```bash
    fedramp-templater <input> <output>
    # i.e.
    fedramp-templater FedRAMP-System-Security-Plan-Template-v2.1.docx FedRAMP-Masonry-Template-v2.1.docx
    ```

The output document will be the same as the input one, albeit with a bunch of Compliance Masonry tags (e.g. `{{ getControl "NIST-800-53@CM-2"}}`) inserted. The resulting document can then be run through [the Compliance Masonry .docx formatter](https://github.com/opencontrol/compliance-masonry/#create-docx-template).
