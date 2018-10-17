### clock
* std::time()
    * returns wall-clock.
    * with precision in seconds
    
* std::clock()
    * cpu time clock (sum of user and system time)
    * with precision in microseconds
    
* clock_gettime(CLOCK_MONOTONIC, ..)
    * is monotonic. (계속 증가)
    * thread safety
    * with precision in nanoseconds
    
* gettimeofday()
    * with precision in microseconds. but this is not guaranteed. because system clock is hardware dependent.
    * obsolescent function. -> Use clock_gettime(..)
    
    
### system_clock VS steady_clock
C++ stdlib is inside GCC source:

* high_resolution_clock is an alias for system_clock or steady_clock

* system_clock forwards to the first of the following that is available:
    * clock_gettime(CLOCK_REALTIME, ...)
    * gettimeofday
    * time
    
* steady_clock forwards to the first of the following that is available:
    * clock_gettime(CLOCK_MONOTONIC, ...)
    * system_clock
    
### CLOCK_REALTIME vs CLOCK_MONOTONIC
* CLOCK_REALTIME
```
This clock represents the clock measuring real time for the system.
For this clock, the values returned by clock_gettime() and specified by clock_settime() represent the amount of time
(in seconds and nanoseconds) since the Epoch.
```

* CLOCK_MONOTONIC
```
For this clock, the value returned by clock_gettime() represents the amount of time (in seconds and nanoseconds)
since an unspecified point in the past (for example, system start-up time, or the Epoch).
This point does not change after system start-up time.
he value of the CLOCK_MONOTONIC clock cannot be set via clock_settime().
```

### 참고
https://stackoverflow.com/questions/12392278/measure-time-in-linux-time-vs-clock-vs-getrusage-vs-clock-gettime-vs-gettimeof
https://stackoverflow.com/questions/13263277/difference-between-stdsystem-clock-and-stdsteady-clock
https://stackoverflow.com/questions/3523442/difference-between-clock-realtime-and-clock-monotonic
