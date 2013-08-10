go-judy
=======

Go docs can be found here: [http://godoc.org/github.com/gnoso/go-judy](http://godoc.org/github.com/gnoso/go-judy)

**go-judy** is a Go language wrapper of the Judy array implementation at [http://judy.sourceforge.net](http://judy.sourceforge.net).

Judy arrays are a fast and memory efficient dynamic array structure. Judy arrays were invented by Doug Baskins and implemented by Hewlett-Packard.

Judy is designed to avoid cache-line fills wherever possible. There are several different variants of Judy arrays. This package implements the Judy1 bitvector and the JudyL integer map currently. Adding other variants should be relatively simple, however.

Counting and range counting operations are particularly fast, and do not require a scan of the array.

**NOTE:** The Judy array is implemented in C and allocates memory directly from the operating system. It is NOT garbage collected by the Go runtime. **It is very important that you call Free() on a Judy array after using it to prevent memory leaks.** The "defer" pattern is a great way to accomplish this.

Here are some examples of set/unset/test bit operations.

    j := Judy1{}   // declare empty Judy1 bit vector array
    defer j.Free() // make sure the array is freed when finished
    
    j.Set(11235)   // returns true
    j.Set(11235)   // returns false (already set)
    j.Test(11235)  // returns true
    j.Unset(11235) // return true
    j.Unset(11235) // return false (already unset)

Integer Map

    j := JudyL{}   // declare empty JudyL integer map array
    defer j.Free() // make sure the array is freed when finished
    
    j.Insert(11235, 1123)   // returns true
    val, ok := j.Get(11235) // val == 1123, ok == true
    _, ok = j.Get(1)        // ok == false (not found)

    j.Delete(11235)         // returns true
    j.Delete(11235)         // returns false (doesn't exist)

Count of an empty array 

    j := Judy1{}   // declare empty Judy1 bit vector array
    defer j.Free() // make sure the array is freed when finished

    j.CountAll()  // return 0

Check memory used by the array

    j := Judy1{}   // declare empty Judy1 bit vector array
    defer j.Free() // make sure the array is freed when finished

    for i := 0; i < 10; i++ {
      j.Set(uint64(i))
    }
    j.CountAll()   // return 10
    j.MemoryUsed() // returns memory usage info
