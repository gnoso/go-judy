go-judy
=======

Go docs can be found here: [http://godoc.org/github.com/gnoso/go-judy](http://godoc.org/github.com/gnoso/go-judy)

**go-judy** is a Go language wrapper of the Judy array implementation at [http://judy.sourceforge.net](http://judy.sourceforge.net).

Judy arrays are a fast and memory efficient dynamic array structure. There are several different variants of Judy arrays, but this package only implements the Judy1 bitvector variety at this time. Adding the other variants would be relatively simple, however.

Here are some examples of set/unset/test bit operations.

    j := Judy1{}   // declare empty Judy1 bit vector array
    defer j.Free() // make sure the array is freed when finished
    
    j.Set(11235)   // returns true
    j.Set(11235)   // returns false (already set)
    j.Test(11235)  // returns true
    j.Unset(11235) // return true
    j.Unset(11235) // return false (already unset)

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
