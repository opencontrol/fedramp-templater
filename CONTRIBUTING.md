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

### Manual Steps

Use these steps if you just want to get started quickly and you have your Go environment configured as per Step 1 below.

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

### Adding dependencies

```bash
glide get --strip-vcs --strip-vendor <package>

```

### Using the supplied Makefile

The Makefile is adapted from [a standard Go Makefile](https://github.com/vincentbernat/hellogopher/tree/feature/glide). Want to use it? Then follow these steps:

1. _Directory Structure._ The `Makefile` assumes you have a ["standard" structure](https://github.com/golang/go/wiki/GithubCodeLayout) when you pull down the software. For example:

    ```bash
    [top-level-dir]
      -> git
         -> src
            -> github.com
               -> opencontrol
                  -> fedramp-templater
    ```

    The advantage to this structure is that it works very well with any Git project, from any workspace.
1. _Install Go._ Strongly recommend *not* to install Go directly but instead to use [Go Version Manager (gvm)](https://github.com/vincentbernat/hellogopher/tree/feature/glide). As an example, selecting a specific Go version is as easy as:

    ```bash
    gvm use [installed go-version from 'gvm list']
    ```
1. _Install <code>make</code>_. See [GNU Make](https://www.gnu.org/software/make/) for details on what `make` is and can do.

    `yum`-based systems (RHEL / CentOS / etc.):

    ```bash
    sudo yum install make
    ```

    `apt`-based systems (Ubuntu / etc.). The below actually installs lots more than just `make`:

    ```bash
    sudo apt-get install build-essential
    ```

    How about Mac? First, install [Homebrew](https://brew.sh). Then:

    ```bash
    brew install make
    ```
1. _OPTIONAL: Install Debugger._ Unfortunately, software does at times behave oddly. The debugger used by this `Makefile` is [Delve (dlv)](https://github.com/derekparker/delve). It can take a little effort to get the debugger installed and working, but the results are far worth it!
1. _One-Time Setup._ Just like in the Manual Instructions above, the code dependencies must be installed. There is a simple target for this:

    ```bash
    make depend
    ```

    The above will pull down all the required dependencies and will not need to be run again.
1. _Common Targets_. Invoke a target by using:

    ```bash
    make [target-name]
    ```

    The common targets include:
    * `all` - Performs a `build`
    * `build` - Builds the source code and places `fedramp-templater` binary into the `./bin` folder (excluded via `.gitignore`)
    * `debug` - Invoke `dlv` _(did you install it above?)_ and invoke it with any `DEBUG_OPTIONS` defined on the command line (example below).
    * `clean` - Simply removes the `compliance-masonry` binary from the `./bin` folder
    * `rebuild` - Invokes `clean` and `build`.
    * `test` - Runs the tests.
    * `lint` - Checks to see if the Go code is properly formatted. (If you want to contribute to the project, use this target; you will need to make sure your code follows accepted standards.)
1. _Examples_.

    * Build the CLI:

        ```bash
        make build
        ```

    * Debug the CLI using `dlv`, running the `fill` command. For this to work, let's assume a simple layout as shown below:

        ```bash
        ~/myproj/myapp
          -> opencontrols (output from 'compliance-masonry get')
          -> input
             -> Input.docx
          -> output
             -> (empty folder; will be filled by 'fedramp-templater fill')
        ```

        (Of course, the `Input.docx` document must be a [FedRAMP SSP Template](https://www.fedramp.gov/resources/templates-2016/).)

        Given the above, invoke the debugger by using:

        ```bash
        MY_DIR="$HOME/myproj/myapp"
        MY_DEBUG_OPTIONS="fill '$MY_DIR/opencontrols' '$MY_DIR/input/Input.docx' '$MY_DIR/output/Output.docx'"
        make debug DEBUG_OPTIONS="$MY_DEBUG_OPTIONS"
        ```

        If all goes well, you should get a debugger prompt and be able to execute debugger commands:
        
        ```bash
        Type 'help' for list of commands.
        (dlv) b main.go:119
        Breakpoint 1 set at 0x4353793 for main.main() ./main.go:119
        (dlv) c
        > main.main() ./main.go:119 (hits goroutine(1):1 total:1) (PC: 0x4353793)
           114:  }
           115:
           116:  func main() {
           117:    opts := parseArgs()
           118:
        => 119:    openControlData := loadOpenControls(opts.openControlsDir)
           120:    doc, err := ssp.Load(opts.inputPath)
           121:    if err != nil {
           122:      log.Fatalln(err)
           123:    }
           124:    defer doc.Close()
        (dlv) p opts
        main.options {
          openControlsDir: "/Users/l.abruce/myproj/myapp/opencontrols",
          inputPath: "/Users/l.abruce/myproj/myapp/input/Input.docx",
          outputPath: "/Users/l.abruce/myproj/myapp/output/Output.docx",
          cmd: 2,}
        (dlv) q
        ```

## Inspecting Word docx XML

Only tested on Mac.

```bash
./scripts/preview-doc <path/to/word.docx>
```

## Creating Binaries

### One Time Setup for Uploading Binaries

1. Install [goxc](go get github.com/laher/goxc)

    ```bash
    go get github.com/laher/goxc
    ```

1. [Get a GitHub API token](https://github.com/settings/tokens/new). The token should have write access to repos.
1. Add a .goxc.local.json file with a github api key

    ```bash
    goxc -wlc default publish-github -apikey=123456789012
    ```

### Compiling and Uploading Binaries

1. Set version number in:
    * [`.goxc.json`](.goxc.json)
1. Run the release script

    ```bash
    ./release.sh
    ```

## Resources

### Tools

* [XML Tree Chrome extension](https://chrome.google.com/webstore/detail/xml-tree/gbammbheopgpmaagmckhpjbfgdfkpadb)
* [XML Viewer Chrome extension](https://chrome.google.com/webstore/detail/xv-%E2%80%94-xml-viewer/eeocglpgjdpaefaedpblffpeebgmgddk?hl=en)

### Reference

* [WordprocessingML information](http://officeopenxml.com/anatomyofOOXML.php)
* [Structure of a WordprocessingML document](https://msdn.microsoft.com/en-us/library/office/gg278308.aspx)
