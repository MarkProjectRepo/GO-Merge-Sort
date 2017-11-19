# GO-Merge-Sort
A short test in multi threaded performance with lessons in more isn't necessarily better.

While attempting to learn a little bit about GO's performance in multithreaded tasks, I tried my hands at a multithreaded solution, and found it to be extemely inefficient. It is clear as to why, especially upon reading this stackoverflow: https://stackoverflow.com/questions/41418210/golang-why-using-goroutines-to-parallelize-calls-ends-up-being-slower 

I took a "shotgun" approach, which leads to overdemanding resources from the CPU and overall slowing down the sorting, this leads to the dispatching mergesorting methods being magnitudes slower than the normal MergeSort, as well as the native package include Sort.
