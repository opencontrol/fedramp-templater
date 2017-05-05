# FedRAMP Templater [![Build Status](https://travis-ci.org/opencontrol/fedramp-templater.svg?branch=master)](https://travis-ci.org/opencontrol/fedramp-templater)

This is a command-line tool to take the [FedRAMP](http://www.fedramp.gov/) System Security Plan template and fill it with OpenControls data.

## Installation

Requires [Go](https://golang.org/) 1.6+. 


1. [Install `gokogiri` dependencies.](https://github.com/moovweb/gokogiri/pull/95/files)
1. Install the templater:

    For Ubuntu/Debian:
    ```bash
    sudo apt-get install libxml2-dev 
    go get github.com/opencontrol/fedramp-templater
    ```
    
    For OsX (assuming you have [HomeBrew installed](http://brew.sh/))
    ```
    brew install libxml2
    brew install pkg-config
    go get github.com/moovweb/gokogiri
    ```
   
    Note that installation issues are usually caused by the install of Gokogiri, and if you run into issues you may find some help at [this issue](https://github.com/moovweb/gokogiri/issues/14) and with the [update to the GokoGiri README](https://github.com/moovweb/gokogiri/pull/95)


## Usage

1. Follow [the Compliance Masonry instructions](https://github.com/opencontrol/compliance-masonry#readme) to:
    1. Install Compliance Masonry
    1. Create an OpenControl project
    1. Collect the OpenControl dependencies
1. [Download the `System Security Plan (SSP)` template.](https://www.fedramp.gov/resources/templates-2016/) (Tested with [v2.1](https://www.fedramp.gov/files/2015/03/FedRAMP-System-Security-Plan-Template-v2.1.docx).)
1. Run

    ```bash
    # To Fill the SSP.
    fedramp-templater fill <openControlsDir> <inputDoc> <outputDoc>
    # i.e.
    fedramp-templater fill opencontrols/ FedRAMP-System-Security-Plan-Template-v2.1.docx FedRAMP-Masonry-Template-v2.1.docx
    
    # To diff the SSP with the YAML
    fedramp-templater diff <openControlsDir> <inputDoc>
    fedramp-templater diff opencontrols/ FedRAMP-System-Security-Plan-Template-v2.1.docx
    ```

The output document will be the same as the input one, albeit filled in with the data from your OpenControls files.
