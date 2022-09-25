#chapter 6 pointers

## pointers are a last resort

The only time you should use pointer parameters to modify a variable is when the function expects an interface. You see this pattern when working whith JSON (we'll talk more about the JSON support in Go's standared library in "encodking/json"

## pointer passing performance

The behavior for returning a pointer versus returning a value is more interesting. Fordata structures that are smaller than a megabyte, it is actually slower to return apointer type than a value type. For example, a 100-byte data structure takes around 10nanoseconds to be returned, but a pointer to that data structure takes about 30 nano‐seconds. Once your data structures are larger than a megabyte, the performanceadvantage flips. It takes nearly 2 milliseconds to return 10 megabytes of data, but alittle more than half a millisecond to return a pointer to it.You should be aware that these are very short times. For the vast majority of cases, thedifference between using a pointer and a value won’t affect your program’s perfor‐mance. But if you are passing megabytes of data between functions, consider using apointer even if the data is meant to be immutable.All of these numbers are from an i7-8700 computer with 32GB of RAM. You can runyour own performance tests by using the code on GitHub.
https://oreil.ly/uVEin


