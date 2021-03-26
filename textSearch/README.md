# Text Search

## Problem

We need a way of finding all the occurrences of a particular set of characters in a string. It should:

- Accept two strings as input. One called 'textToSearch' and one called 'subtext'
- The solution should match the subtext against the textToSearch, outputting the positions of the beginning of each match for the subtext within the textToSearch.
- Multiple matches are possible
- Matching is case insensitive
- If no matches have been found, no output is generated

## Solution

The main package contains 2 functions outlined below. To execute the program run the following from the commandline:

```
go run main.go -textToSearch="Peter told me that peter the pickle piper piped a pitted pickle before he petered out. Phew!" -subtext="Peter"
```

To run the tests, run:

```
go test ./...
```

### `main`

This function allows the program to run. It uses the `flag` package to read flag arguments passed in to the program. There are 2 flags defined:

- `textToSearch` - the text that you want to search
- `subtext` - the text that you're searching for

The main function calls the `textSearch` function described below passing in the two flags as parameters to the function. It then prints the output from that function.

### `textSearch`

The `textSearch` function which is the one that performs the actual search. It accepts 2 string parameters:

- `textToSearch` - the text that you want to search
- `subtext` - the text that you're searching for

The response from the function is a comma separated string of the start position of each match within the `textToSearch` string.
