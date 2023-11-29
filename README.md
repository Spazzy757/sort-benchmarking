# Benchmarking Different Sorting Methods in Golang

This is a simple example of different sorting algorithms sorting a list of
integers

## Supported Agorithms

- Merge Sort 
- Insertion Sort 
- Bubble Sort 

## To Run an Algorithm

```bash
go run . -type {merge|bubble|insertion}
```
> note currently the list that is used is 1000000 and very resource intensive,
> you can change this by setting the flag `-list-length`

## To Run Benchmarking

You will need golang installed and then you can simply run:

```bash
go test -bench=.
```

Currently each sorting algorithm is run with an array with a size of 10000
