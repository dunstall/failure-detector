# Phi-Accrual Failure Detector
Implementation of Phi-Accrual Failure Detector with Golang

The protocol samples arrival times of heartbeats and maintains sliding window
of most recent samlples. This window is used to estimate the arrival time of
the next heartbeat. The distribution of past samples is used as an
approximation for the probabilistic distribution of future hearbeat messages.
With this information can compute phi with a scale that changes dynamically
to match recent network conditions.

Configurable:
* Rate of heartbeats sent (eg every 100ms)

Design:
* Keep the calculation monitoring and interpretation separate so its easy to
use own monitoring system and just notify interpretation of the results. Also
provide own monitoring (optional).
* See https://github.com/dgryski/go-failure/blob/master/failure.go



## Resources
* [The φ accrual failure detector](https://dspace.jaist.ac.jp/dspace/bitstream/10119/4784/1/IS-RR-2004-010.pdf)
* [Two-ways Adaptive Failure Detection with the ϕ-Failure Detector](https://pdfs.semanticscholar.org/219b/309d324782ac31fa1e4232a1a51a12ef6af2.pdf)
