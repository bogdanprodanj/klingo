# klingo
Small tool to translate a name written in English to Klingon and find out its species.

## Building
Download dependencies (make sure that modules are enabled by setting `GO111MODULE=on`)

```
$ go mod vendor
$ go build -mod=vendor -o ./bin/translator ./cmd/translator/...
```

## Usage

`./bin/translator [input]`

Note, that sometimes one or more results will be shown, as [Star Track API](http://stapi.co​.)
might return several names that contains the input letters.
Output of the program will contain full name, followed by translation and species name separated by new line. 

Example output:
```
$ ./bin/translator Uhura
0xF8DB 0xF8E8 0xF8DD 0xF8E3 0xF8D0 0x0020 0xF8E5 0xF8D6 0xF8E5 0xF8E1 0xF8D0
Nyota Uhura
Human

$ ./bin/translator Kaa
0xF8E9 0xF8D6 0xF8DD 0x0 0xF8DF 0xF8D0 0xF8D0 0xF8DB
Zho'Kaan
Arkonian
0xF8D0 0xF8DF 0xF8D0 0xF8D0 0xF8E1
Akaar
Capellan
0xF8D9 0xF8D4 0xF8DD 0xF8DB 0xF8D0 0xF8E1 0xF8D3 0x0020 0xF8D8 0xF8D0 0xF8DA 0xF8D4 0xF8E2 0x0020 0xF8D0 0xF8DF 0xF8D0 0xF8D0 0xF8E1
Leonard James Akaar
Capellan
0xF8E2 0xF8D6 0xF8D0 0xF8DF 0xF8D0 0xF8D0 0xF8E1 0x0020 0xF8D4 0xF8D3 0xF8DD 0xF8DB
Shakaar Edon
Bajoran
```
