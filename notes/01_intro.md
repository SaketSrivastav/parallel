# GO Parallel Computing

## Amdahl's Law

- Thearetical speedup of a program cannot be improved more than the time taken by the sequential part of the program.
- Ex: if a program needs 20 hours to complete using a single thread, but a one-hour portion of the program cannot be parallelized, therefore only the remaining 19 hours (p = 0.95) of execution time can be parallelized, then regardless of how many threads are devoted to a parallelized execution of this program, the minimum execution time cannot be less than one hour. Hence, the theoretical speedup is limited to at most 20 times the single thread performance. (1/1-p) = 20.

## Gustafson's Law

- It is a more positive take on Amdahl's Law. It states that if we increase problem size then we can achieve linear speed up wrt to number of processors. Ex: If you are building a wall with 1 wheel burrow and 10 workers then instead of building 1 wall you can build 3 wall where there are more burrows are available. In the same amount of time that you take to build 1 wall you will be able to complete 3 walls. You may not want to build 3 walls but the idea is to complete more work in same amount of time.
- Ex: there are 3 programs waiting to get executed on a processor, p-1 needs to wait for IO read therefore instead of waiting for the IO and blocking the processor it can be interrupted and pushed to back of the queue while p2 and p3 can execute. This honours Gustafson's Law by doing more work in the given time and not being limited by the sequential part of the program as Amdahl's Law states.
