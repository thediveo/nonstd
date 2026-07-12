/*
Package prioerrgroup provides an error group with synchronization, error
propagation as well as error prioritization, and finally Context cancellation
for groups of goroutines working on subtasks of a common task. The new
additional functionality here is error prioritization on top of
[errgroup.Group].
*/
package prioerrgroup
