## Advanced Slices

### Indexing capacity

```go
slice := make([]string,3,6)
slice[4] // resolves to error (index access is based on length, not cap)
slice[3] = 0 // reoslves to error (cannot insert at index where index >= len)
```

Creating array with capacity only allocates memory for further elements, we cannott index them
### The length is the number of elements slice contain, the capacity is number of elements in backing array

### Slice capacity extension

```go
    when idx >= len, we can append to slice via built-in append
    when inserting element exceed capacity, new array under the slice is created with cap * 2 when cap <= 1024 else cap *= cap*0.25
```

### Changing copied slice which length < capacity

* Results in changing backing array.
    * When changing by append
```go
    s1 := make([]string,3)
    s1[0] = 0
    s1[1] = 1
    s1[2] = 2
    s2 := s1[1:2] // len = 1, cap = 2 -- cap = backingarray cap - amount we cut array (**s1[1**)
    s3 := append(s2,10)
    /* append results in s3 = 0,1,10, s2 = 1, s1 = 0,1,10
       s3 is result of append (where it takes slice, add element to last via len, and increase length of slice which will return)
       s2 didnt changed because we didnt update its len
       s1 changed bcs its len is 3 and (with append we change backing array)
       it worked like this bcs of slices length, appending adds last element to array based on length
    */
```

### How to mitigate append / backing array?
* Via copying using built-in copy(dst,src)
* Via making full-slice expression where we pass slice[1:2:2] whereas last elem is capacity (so append will not change backing array, but will create another one) [be aware of memory leaks, bcs backing array is still referenced by our new slice]

### Slice memory leaks

* Slicing slice and storing it in another variable, but without modifying capacity.
    * e.g.

```go
    s1 := make([]int,3000)
    seedSlice(s1) // seeds slice with 3000 elements
    s2 := s1[:5]
    /*
    westore to s2 new slice with len 5, but capacity is 3000, 
    we assume gc cleaned up s1, because it dont have reference, 
    but it has, reference in s2 slice to the backing array of s1
    so what? memory leak.

        How to mitigate it?
    
        Full slice expression do not works bcs we have ref to backing array (and its all about at),
        we should make slice deep copy,and return it.
    */
```






