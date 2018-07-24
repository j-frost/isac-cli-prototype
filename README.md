# isac-cli-prototype

This is a prototype concept for how a potential command line interface for our infrastructure management might be structured. 

It is written in Go using [Cobra](https://github.com/spf13/cobra). This prototype tries to implement a natural language command pattern. You will however find it to be trivial to change do domain specific commands. Find more theory in the wiki. 

## Install

If you just want to try it out:

```bash
go get github.com/j-frost/isac-cli-prototype/cmd/isac-cli-prototype 
```

To install the source code for tinkering: 

```bash
go get github.com/j-frost/isac-cli-prototype 
```

## Usage

Should be more or less self documenting. Use 

```bash
isac help
```

for general usage information.

## Contribute

PRs will most likely not be accepted. This is just a little proof of concept. 

## License

see [LICENSE](/LICENSE) 
