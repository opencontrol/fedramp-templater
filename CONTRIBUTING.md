## Welcome!

We're so glad you're thinking about contributing to an 18F open source project! If you're unsure about anything, just ask -- or submit the issue or pull request anyway. The worst that can happen is you'll be politely asked to change something. We love all friendly contributions.

We want to ensure a welcoming environment for all of our projects. Our staff follow the [18F Code of Conduct](https://github.com/18F/code-of-conduct/blob/master/code-of-conduct.md) and all contributors should do the same.

We encourage you to read this project's CONTRIBUTING policy (you are here), its [LICENSE](LICENSE.md), and its [README](README.md).

If you have any questions or want to read more, check out the [18F Open Source Policy GitHub repository]( https://github.com/18f/open-source-policy), or just [shoot us an email](mailto:18f@gsa.gov).

## Public domain

This project is in the public domain within the United States, and
copyright and related rights in the work worldwide are waived through
the [CC0 1.0 Universal public domain dedication](https://creativecommons.org/publicdomain/zero/1.0/).

All contributions to this project will be released under the CC0
dedication. By submitting a pull request, you are agreeing to comply
with this waiver of copyright interest.

## Development

1. Follow the [installation instructions](README.md#installation).
1. Go to the repository directory.

    ```bash
    cd $GOPATH/src/github.com/opencontrol/fedramp-templater
    ```

1. _code code code_
1. Run the tests.

    ```bash
    go test -v $(go list ./... | grep -v /vendor/)
    ```

1. Run the CLI.

    ```bash
    go run main.go fixtures/FedRAMP_ac-2-1_v2.1.docx tmp/output.docx
    open tmp/output.docx
    ```

[Glide](https://glide.sh/) is used to manage dependencies.

## Resources

### Tools

* [XML Tree Chrome extension](https://chrome.google.com/webstore/detail/xml-tree/gbammbheopgpmaagmckhpjbfgdfkpadb)
* [XML Viewer Chrome extension](https://chrome.google.com/webstore/detail/xv-%E2%80%94-xml-viewer/eeocglpgjdpaefaedpblffpeebgmgddk?hl=en)

### Reference

* [WordprocessingML information](http://officeopenxml.com/anatomyofOOXML.php)
* [Structure of a WordprocessingML document](https://msdn.microsoft.com/en-us/library/office/gg278308.aspx)
