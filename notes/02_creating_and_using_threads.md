# Creating and using Threads

## Processes
- Independent entity.
- Does not share memory with other process.
- Lifecycle of 1 process does not affect lifecycle of other process.
- Spawning a new process takes significant amount of time. They are heavy weight.
- Ex: in C/C++
```
if fork() == 0 {
    // i am child
} else {
    // i am parent
}
```
- Golang does not have a fork() function to spawn a new process as we do in C/C++. There are other ways to do it. https://gobyexample.com/spawning-processes


## Threads
- Spawned inside process.
- All threads inside a process share memory space.
- All threads share process lifecycle, if process terminates then all threads inside process also terminates.
- Threads are lightweight compared to processes.
- All threads need to communicate with each other so that they dont collide with each other.
- OS decides which thread to execute.

## Green Thread (User Level Thread)

- Threads are not kernel level and therefore OS does not decide which thread to run.
- The thread scheduler is present at the user level and that decides which green thread to run inside a kernel thread.
- A kernel thread may have multiple green threads present inside a kernel thread but its the green thread scheduler decides which one to execute.
- This allows for less context switching.
- Problem: If one of the green thread waits for IO then OS preempts all the green thread present inside the kernel thread even if there were other green threads ready to execute.

GOLANG uses a hybrid system where it creates 1 kernel thread per CPU and then manages these green threads inside these kernel threads.

In case of a GREEN thread waiting for IO, GOLANG reshuffles the remaining green threads not waiting for an IO to a different kernel thread therefore improving the performance.
