# Daily Coding Problem: Problem #541 [Easy]

This is also "Daily Coding Problem: Problem #29".
The Daily Coding Problem repeats a few problems.

This problem was asked by Amazon.

Run-length encoding is a fast and simple method of encoding strings. The
basic idea is to represent repeated successive characters as a single
count and character.

For example, the string "AAAABBBCCDAA" would be encoded as "4A3B2C1D2A".

Implement run-length encoding and decoding. You can assume the string to
be encoded have no digits and consists solely of alphabetic characters.
You can assume the string to be decoded is valid.

## Building and running

```sh
$ go build rl.go
$ ./rl AAAABBBCCDAA   # encodes by default
$ ./rl -d 4A3B2C1D2A  # decodes with -d flag
$ ./rl "thequick" "brownfoxjumps" "overthe"  # more than one string
```

## Analysis

Encode and decode are dissimilar enough to merit seperate functions,
but each is a loop over the input.

Encoding has to keep track of the "last character" it saw
so as to be able to output a run length and an encoded character,
but also has to track that the length of the run is less than 10:
the problem statement says "a single count and character".
So a run of 10 'A' characters would encode as "9A1A".
Encoding also has to check for a run that it hasn't output
at the end of the loop over the input characters.

Even though the programming part of this is not complex,
there's enough coding to do for an interviewer to get a good
sense of how the candidate works.

The candidate should talk about what they're doing as they
write the code.
The candidate should offer test inputs:
10 (or more) of a single character
and the zero-length string
should both be test inputs for encoding.
Test inputs for decoding should include things like '9A1A',
'1A1A1A' and maybe even un-encoded strings.
The problem statements says "You can assume the string to be decoded is valid".
But why not try potentially "valid" encoded strings in testing?

I think I'd agree with the "easy" rating on this one.
It doesn't require any flashes of insight,
memorizing obscure O(n) algorithm variants,
or insanely detailed programming language knowledge.
