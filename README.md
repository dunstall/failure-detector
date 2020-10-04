# Phi-Accrual Failure Detector
Implementation of Phi-Accrual Failure Detector with Golang

The protocol samples arrival times of heartbeats and maintains sliding window
of most recent samples. This window is used to estimate the arrival time of
the next heartbeat. The distribution of past samples is used as an
approximation for the probabilistic distribution of future hearbeat messages.
With this information can compute phi with a scale that changes dynamically
to match recent network conditions.

## Resources
* [The φ accrual failure detector](https://dspace.jaist.ac.jp/dspace/bitstream/10119/4784/1/IS-RR-2004-010.pdf)
* [Two-ways Adaptive Failure Detection with the ϕ-Failure Detector](https://pdfs.semanticscholar.org/219b/309d324782ac31fa1e4232a1a51a12ef6af2.pdf)

<a href="https://www.gistgrok.com/">GistGrok: Online C++ Shell and Compiler</a>
