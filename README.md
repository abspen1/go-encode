# Encoder & Decoder built with Go
This is an Encoder and eventually decoder that I built in Go. Originally built in C++ for CSE school project but migrated into Go. Interested to see how the run times compare for large files. So far the encoding is working nicely and I think close to finished.

## Running
```bash
go run main.go <input.txt >output.txt
```

## Examples
Input:
```
Mississippi
```
* Create each shift of character and append to str2 slice
* [ississippiM ssissippiMi sissippiMis issippiMiss ssippiMissi sippiMissis ippiMississ ppiMississi piMississip iMississipp Mississippi]
* Sort the strings in the slice lexicographically
* Take the last letter of each string in order
* Ouput the index of original string
* Next line output frequency of cluster followed by character
Output:
```
0
1 i 1 p 2 s 1 M 1 p 1 i 2 s 2 i
```
* So the actual last letters are i p s s M p i s s i
