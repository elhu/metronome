# Metronome

Metronome is a simple command line metronome written in Go.
Tested on MacOS only.

## Installation

```bash
go build -o metronome main.go
```

## Usage

```bash
./metronome
```

You can set the tempo (beats per minute) and how many beats per bar (for the accent) like this:

```bash
./metronome -bpm 70 -beats 4
```

## Contributing

Pull requests are welcome!

## License

[MIT](https://choosealicense.com/licenses/mit/)
